package domainerr

import (
	"fmt"

	"github.com/pkg/errors"
)

type Error struct {
	cause  error
	status *Status
	*stack
}

type stack = errors.Stack

type ErrorOpt func(e *Error)

func WithCause(cause error) ErrorOpt {
	return func(e *Error) {
		e.cause = cause
	}
}

func NewError(status *Status, opts ...ErrorOpt) *Error {
	e := &Error{
		status: status,
	}
	for _, setOpt := range opts {
		setOpt(e)
	}
	e.stack = errors.Callers()
	return e
}

func (e *Error) Status() *Status {
	return e.status
}

func (e *Error) Cause() error {
	return e.cause
}

func (e *Error) Unwrap() error {
	return e.cause
}

func (e *Error) Error() string {
	return e.status.String()
}

// AugmentMessage is a shortcut of err.Status().AugmentMessage(...), and augments the message of
// the status of this error with more contextual information of current use case scenario.
func (e *Error) AugmentMessage(moreCtx string) {
	e.status.AugmentMessage(moreCtx)
}

// AugmentMessagef sees comment for AugmentMessage
func (e *Error) AugmentMessagef(moreCtxFmt string, args ...any) {
	e.status.AugmentMessage(fmt.Sprintf(moreCtxFmt, args...))
}

// Format implements the fmt.Formatter interface.
func (e *Error) Format(s fmt.State, verb rune) {
	errors.FormatError(e, s, verb)
}

// StatusFromErrChain finds the first op Error from the causal chain of given error.
// If one is found, return its status. Otherwise, return nil
func StatusFromErrChain(err error) *Status {
	if IsNil(err) {
		return nil
	}
	cause := err
	for NotNil(cause) {
		if match, opErr := AsOpError(cause); match {
			return opErr.Status()
		}
		cause = errors.Unwrap(cause)
	}
	return nil
}

// AsOpError finds the first error in given error chain that is of type opError,
// and if one is found, sets target to that error value and returns true. Otherwise,
// it returns false.
func AsOpError(err error) (bool, *Error) {
	panic("implement me")
	//var opErr Error
	//return errors.As(err, &opErr), &opErr
}

type ErrorBuilder struct {
	status *Status
	cause  error
}

func (b *ErrorBuilder) WithMessage(msg string) *ErrorBuilder {
	b.status = b.status.WithMessage(msg)
	return b
}

func (b *ErrorBuilder) WithMessagef(msgFmt string, a ...any) *ErrorBuilder {
	b.status = b.status.WithMessagef(msgFmt, a...)
	return b
}

func (b *ErrorBuilder) WithSpecificCase(c Case) *ErrorBuilder {
	b.status = b.status.WithCase(c)
	return b
}

func (b *ErrorBuilder) WithDetails(details map[string]any) *ErrorBuilder {
	b.status.AddDetails(details)
	return b
}

func (b *ErrorBuilder) WithCause(cause error) *ErrorBuilder {
	b.cause = cause
	return b
}

func (b *ErrorBuilder) Build() *Error {
	return NewError(b.status, WithCause(b.cause))
}

func NewCancelled() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusCancelled,
	}
}

func NewUnknownError() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusUnknown,
	}
}

func NewInvalidArgument() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusInvalidArgument,
	}
}

func NewDeadlineExceeded() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusDeadlineExceeded,
	}
}

func NewNotFound() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusNotFound,
	}
}

func NewAlreadyExists() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusAlreadyExists,
	}
}

func NewPermissionDenied() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusPermissionDenied,
	}
}

func NewUnauthenticated() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusUnauthenticated,
	}
}

func NewResourceExhausted() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusResourceExhausted,
	}
}

func NewFailedPrecondition() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusFailedPrecondition,
	}
}

func NewAborted() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusAborted,
	}
}

func NewOutOfRange() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusOutOfRange,
	}
}

func NewUnimplemented() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusUnimplemented,
	}
}

func NewInternalError() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusInternal,
	}
}

func NewUnavailable() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusUnavailable,
	}
}

func NewDataLoss() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusDataLoss,
	}
}

func NewUndefined() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusUndefined,
	}
}

func NewAuthorizationExpired() *ErrorBuilder {
	return &ErrorBuilder{
		status: StatusAuthorizationExpired,
	}
}
