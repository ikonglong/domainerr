package domainerr

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestError_WithoutCause(t *testing.T) {
	err := NewError(StatusInternal.WithMessage("internal error"))
	want := `
error occurred: InternalError: internal error
github.com/ikonglong/domainerr.TestError_WithoutCause
	.+/github.com/ikonglong/domainerr/error_test.go:\d+
testing.tRunner
	.+/testing/testing.go:\d+
runtime.goexit
	.+/runtime/asm_.+.s:\d+`
	testTextRegexp(t, want, fmt.Sprintf("%+v", err))
}

func TestError_WithCauseMissingStack(t *testing.T) {
	cause := fmt.Errorf("causal error")
	err := NewError(StatusInternal.WithMessage("internal error"), WithCause(cause))

	want := `
error occurred: InternalError: internal error
github.com/ikonglong/domainerr.TestError_WithCauseMissingStack
	.+/github.com/ikonglong/domainerr/error_test.go:\d+
testing.tRunner
	.+/testing/testing.go:\d+
runtime.goexit
	.+/runtime/asm_.+.s:\d+
caused by: causal error`
	testTextRegexp(t, want, fmt.Sprintf("%+v", err))
}

func TestError_WithCauseHavingStack(t *testing.T) {
	app := application{
		s: service{
			repo: repository{
				db: database{},
			},
		},
	}
	err := app.exec("test")

	want := `
error occurred: InternalError: failed to save exec record
github.com/ikonglong/domainerr.service.exec
	.+/github.com/ikonglong/domainerr/error_test.go:\d+
github.com/ikonglong/domainerr.application.exec
	.+/github.com/ikonglong/domainerr/error_test.go:\d+
github.com/ikonglong/domainerr.TestError_WithCauseHavingStack
	.+/github.com/ikonglong/domainerr/error_test.go:\d+
testing.tRunner
	.+/testing/testing.go:\d+
runtime.goexit
	.+/runtime/asm_.+.s:\d+
caused by: network error
github.com/ikonglong/domainerr.database.insert
	.+/github.com/ikonglong/domainerr/error_test.go:\d+
github.com/ikonglong/domainerr.repository.save
	.+/github.com/ikonglong/domainerr/error_test.go:\d+
github.com/ikonglong/domainerr.service.exec
	.+/github.com/ikonglong/domainerr/error_test.go:\d+`
	got := fmt.Sprintf("%+v", err)
	testTextRegexp(t, want, got)
}

func TestError_WithCauseMissingStack_WithCauseHavingStack(t *testing.T) {
	app := application{
		s: service{
			repo: repository{
				db: database{},
			},
		},
	}
	err := app.exec("just wrap error")

	want := `
error occurred: InternalError: failed to save exec record
github.com/ikonglong/domainerr.service.exec
	.+/github.com/ikonglong/domainerr/error_test.go:\d+
github.com/ikonglong/domainerr.application.exec
	.+/github.com/ikonglong/domainerr/error_test.go:\d+
github.com/ikonglong/domainerr.TestError_WithCauseMissingStack_WithCauseHavingStack
	.+/github.com/ikonglong/domainerr/error_test.go:\d+
testing.tRunner
	.+/testing/testing.go:\d+
runtime.goexit
	.+/runtime/asm_.+.s:\d+
caused by: failed to persist exec record, network error
caused by: network error
github.com/ikonglong/domainerr.database.insert
	.+/github.com/ikonglong/domainerr/error_test.go:\d+
github.com/ikonglong/domainerr.repository.save
	.+/github.com/ikonglong/domainerr/error_test.go:\d+
github.com/ikonglong/domainerr.service.exec
	.+/github.com/ikonglong/domainerr/error_test.go:\d+`
	got := fmt.Sprintf("%+v", err)
	testTextRegexp(t, want, got)
}

func testTextRegexp(t *testing.T, want, got string) {
	gotLines := strings.SplitN(got, "\n", -1)
	wantLines := strings.SplitN(want, "\n", -1)

	if len(wantLines) > len(gotLines) {
		t.Errorf("wantLines(%d) > gotLines(%d):\n got: %q\nwant: %q", len(wantLines), len(gotLines), got, want)
		return
	}

	for i, w := range wantLines {
		match, err := regexp.MatchString(w, gotLines[i])
		if err != nil {
			t.Fatal(err)
		}
		if !match {
			t.Errorf("line %d: got: %q\nwant: %q", i+1, gotLines[i], wantLines[i])
		}
	}
}

type application struct {
	s service
}
type service struct {
	repo repository
}
type execRecord struct {
	cmd string
}
type repository struct {
	db database
}
type database struct{}

func (a application) exec(command string) error {
	return a.s.exec(command)
}

func (s service) exec(command string) error {
	e := s.repo.save(execRecord{cmd: command})
	if e != nil {
		return NewError(StatusInternal.WithMessage("failed to save exec record"), WithCause(e))
	}
	return nil
}

func (r repository) save(rec execRecord) error {
	err := r.db.insert(fmt.Sprintf("%v", rec))
	if err != nil {
		if rec.cmd == "just wrap error" {
			return fmt.Errorf("failed to persist exec record, %w", err)
		}
		return err
	}
	return nil
}

func (d database) insert(record string) error {
	return errors.Errorf("network error")
}

func TestChainMsg(t *testing.T) {
	e1 := NewInternalError().WithMessage("error a").Build()
	assert.Equal(t, "error a", e1.ChainMsg())

	e2 := NewInternalError().WithMessage("error a").
		WithCause(fmt.Errorf("error b")).
		Build()
	assert.Equal(t, "error a -> error b", e2.ChainMsg())

	e3 := NewInternalError().WithMessage("error a").
		WithCause(NewUnknownError().WithMessage("error b").WithCause(fmt.Errorf("error c")).Build()).
		Build()
	assert.Equal(t, "error a -> error b -> error c", e3.ChainMsg())

	var nilDomainErr *Error
	e4 := NewInternalError().WithMessage("error a").
		WithCause(nilDomainErr).Build()
	assert.Equal(t, "error a", e4.ChainMsg())

	e5 := NewInternalError().WithMessage("error a").
		WithCause(nil).Build()
	assert.Equal(t, "error a", e5.ChainMsg())

	var nilErr error = nil
	e6 := NewInternalError().WithMessage("error a").
		WithCause(nilErr).Build()
	assert.Equal(t, "error a", e6.ChainMsg())

	e7 := NewInternalError().WithMessage("error a").
		WithCause(errors.WithStack(fmt.Errorf("error b"))).Build()
	assert.Equal(t, "error a -> error b -> error b", e7.ChainMsg())

	e8 := NewInternalError().WithMessage("error a").
		WithCause(fmt.Errorf("wrap error: %w", fmt.Errorf("error c"))).Build()
	assert.Equal(t, "error a -> wrap error: error c", e8.ChainMsg())
}
