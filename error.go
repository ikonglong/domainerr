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

// NewError Deprecated
func NewError(status *Status, opts ...ErrorOpt) *Error {
	e := &Error{
		status: status,
	}
	for _, setOpt := range opts {
		setOpt(e)
	}
	e.stack = errors.Callers(1)
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

func (e *Error) Details() any {
	return e.status.Details()
}

func (e *Error) Error() string {
	return e.status.ToObjStyleStr()
}

// ChainMsg strings each error's message in this chain together with separator '->'.
//
// The next error in the chain is got by `interface{ Cause() error }`. If an error implements
// `interface{ Unwrap() error }`, this error wrapper and the wrapped error are together handled
// as an error on the chain.
func (e *Error) ChainMsg() string {
	return ChainMsg(e)
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

func (b *ErrorBuilder) WithDetails(v any) *ErrorBuilder {
	b.status = b.status.WithDetails(v)
	return b
}

func (b *ErrorBuilder) WithCause(cause error) *ErrorBuilder {
	b.cause = cause
	return b
}

func (b *ErrorBuilder) Build() *Error {
	return &Error{
		status: b.status,
		cause:  b.cause,
		stack:  errors.Callers(1),
	}
}

func NewWithStatus(s *Status) *ErrorBuilder {
	return &ErrorBuilder{
		status: s,
	}
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
