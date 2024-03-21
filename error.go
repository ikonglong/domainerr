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
func (e *Error) AugmentMessage(moreContext string) {
	e.status.AugmentMessage(moreContext)
}

// AugmentMessagef sees comment for AugmentMessage
func (e *Error) AugmentMessagef(moreContextFmt string, args ...any) {
	e.status.AugmentMessage(fmt.Sprintf(moreContextFmt, args...))
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
