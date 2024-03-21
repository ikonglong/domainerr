package domainerr

import (
	"fmt"
	"sort"
)

// The set of canonical operation status codes.
//
// Sometimes multiple error codes may apply. Services should return the most specific error
// code that applies. For example, prefer `CodeOutOfRange` over `CodeFailedPrecondition` if both codes
// apply. Similarly prefer `CodeNotFound` or `CodeAlreadyExists` over `CodeFailedPrecondition`.
var (
	// CodeOK not an error; returned on success.
	//
	// HTTP Mapping: 200 OK
	CodeOK = newCode("OK", 0)

	// CodeCancelled means the operation was cancelled, typically by the caller.
	//
	// HTTP Mapping: 499 Client Closed Request
	CodeCancelled = newCode("OperationCancelled", 1)

	// CodeUnknown error.  For example, this error may be returned when
	// a `Status` value received from another address space belongs to
	// an error space that is not known in this address space. Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	//
	// HTTP Mapping: 500 Internal Server Error
	CodeUnknown = newCode("UnknownError", 2)

	// CodeInvalidArgument means that the client specified an invalid argument.
	// Note that this differs from `CodeFailedPrecondition`. `CodeInvalidArgument` indicates
	// arguments that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	//
	// HTTP Mapping: 400 Bad Request
	CodeInvalidArgument = newCode("InvalidArgument", 3)

	// CodeDeadlineExceeded means the deadline expired before the operation could complete.
	// For operations that change the state of the system, this error may be returned
	// even if the operation has completed successfully. For example, a successful
	// response from a server could have been delayed long enough for the deadline
	// to expire.
	//
	// HTTP Mapping: 504 Gateway Timeout
	CodeDeadlineExceeded = newCode("DeadlineExceeded", 4)

	// CodeNotFound means that some requested entity (e.g., file or directory) was not found.
	//
	// Note to server developers: if a request is denied for an entire class
	// of users, such as gradual feature rollout or undocumented allowlist,
	// `CodeNotFound` may be used. If a request is denied for some users within
	// a class of users, such as user-based access control, `CodePermissionDenied`
	// must be used.
	//
	// HTTP Mapping: 404 Not Found
	CodeNotFound = newCode("NotFound", 5)

	// CodeAlreadyExists means that the entity that a client attempted to create
	// (e.g., file or directory) already exists.
	//
	// HTTP Mapping: 409 Conflict
	CodeAlreadyExists = newCode("AlreadyExists", 6)

	// CodePermissionDenied The caller does not have permission to execute the specified
	// operation. `CodePermissionDenied` must not be used for rejections caused by
	// exhausting some resource (use `CodeResourceExhausted` instead for those errors).
	// `CodePermissionDenied` must not be used if the caller can not be identified
	// (use `CodeUnauthenticated` instead for those errors). This error code does not
	// imply the request is valid or the requested entity exists or satisfies
	// other pre-conditions.
	//
	// HTTP Mapping: 403 Forbidden
	CodePermissionDenied = newCode("PermissionDenied", 7)

	// CodeUnauthenticated means that the request does not have valid authentication
	// credentials for the operation.
	//
	// HTTP Mapping: 401 Unauthorized
	CodeUnauthenticated = newCode("Unauthenticated", 16)

	// CodeResourceExhausted means that some resource has been exhausted, perhaps
	// a per-user quota, or perhaps the entire file system is out of space.
	//
	// HTTP Mapping: 429 Too Many Requests
	CodeResourceExhausted = newCode("ResourceExhausted", 8)

	// CodeFailedPrecondition means that the operation was rejected because the system
	// is not in a state required for the operation's execution.  For example,
	// the directory to be deleted is non-empty, a rmdir operation is applied to
	// a non-directory, etc.
	//
	// Service implementors can use the following guidelines to decide
	// between `CodeFailedPrecondition`, `CodeAborted`, and `CodeUnavailable`:
	//  (a) Use `CodeUnavailable` if the client can retry just the failing call.
	//  (b) Use `CodeAborted` if the client should retry at a higher level. For
	//      example, when a client-specified test-and-set fails, indicating the
	//      client should restart a read-modify-write sequence.
	//  (c) Use `CodeFailedPrecondition` if the client should not retry until
	//      the system state has been explicitly fixed. For example, if a "rmdir"
	//      fails because the directory is non-empty, `CodeFailedPrecondition`
	//      should be returned since the client should not retry unless
	//      the files are deleted from the directory.
	//
	// HTTP Mapping: 400 Bad Request
	CodeFailedPrecondition = newCode("FailedPrecondition", 9)

	// CodeAborted means that the operation was aborted, typically due to a concurrency
	// issue such as a sequencer check failure or transaction abort.
	//
	// See the guidelines above for deciding between `CodeFailedPrecondition`,
	// `CodeAborted`, and `CodeUnavailable`.
	//
	// HTTP Mapping: 409 Conflict
	CodeAborted = newCode("OperationAborted", 10)

	// CodeOutOfRange means that the operation was attempted past the valid range.
	// E.g., seeking or reading past end-of-file.
	//
	// Unlike `CodeInvalidArgument`, this error indicates a problem that may
	// be fixed if the system state changes. For example, a 32-bit file
	// system will generate `CodeInvalidArgument` if asked to read at an
	// offset that is not in the range [0,2^32-1], but it will generate
	// `CodeOutOfRange` if asked to read from an offset past the current
	// file size.
	//
	// There is a fair bit of overlap between `CodeFailedPrecondition` and
	// `CodeOutOfRange`.  We recommend using `CodeOutOfRange` (the more specific
	// error) when it applies so that callers who are iterating through
	// a space can easily look for an `CodeOutOfRange` error to detect when
	// they are done.
	//
	// HTTP Mapping: 400 Bad Request
	CodeOutOfRange = newCode("OutOfRange", 11)

	// CodeUnimplemented means that the operation is defined, but not implemented
	// or not supported/enabled in this service.
	//
	// HTTP Mapping: 501 Not Implemented
	CodeUnimplemented = newCode("OperationUnimplemented", 12)

	// CodeInternalError errors. This means that some invariants expected by the
	// underlying system have been broken. This error code is reserved
	// for serious errors.
	//
	// HTTP Mapping: 500 Internal Server Error
	CodeInternalError = newCode("InternalError", 13)

	// CodeUnavailable means that the service is currently unavailable. This is
	// most likely a transient condition, which can be corrected by retrying
	// with a backoff. Note that it is not always safe to retry
	// non-idempotent operations.
	//
	// See the guidelines above for deciding between `CodeFailedPrecondition`,
	// `CodeAborted`, and `CodeUnavailable`.
	//
	// HTTP Mapping: 503 Service Unavailable
	CodeUnavailable = newCode("ServiceUnavailable", 14)

	// CodeDataLoss means that unrecoverable data loss or corruption.
	//
	// HTTP Mapping: 500 Internal Server Error
	CodeDataLoss = newCode("DataLoss", 15)

	// CodeUndefined means that the API operation/method is not defined on the target resource.
	//
	// HTTP Mapping: 405 Method Not Allowed
	CodeUndefined = newCode("OperationNotDefined", 29)

	// CodeAuthorizationExpired means a user's authorization expired, and it is
	// needed to log-in again and reauthorize.
	//
	// HTTP Mapping: 401 Unauthorized
	CodeAuthorizationExpired = newCode("AuthorizationExpired", 30)
)

// CodeList contains all the well-defined operation status codes indexed by their values
var CodeList = func() []Code {
	list := make([]Code, 0, 20)
	list = append(list, CodeOK)
	list = append(list, CodeCancelled)
	list = append(list, CodeUnknown)
	list = append(list, CodeInvalidArgument)
	list = append(list, CodeDeadlineExceeded)
	list = append(list, CodeNotFound)
	list = append(list, CodeAlreadyExists)
	list = append(list, CodePermissionDenied)
	list = append(list, CodeUnauthenticated)
	list = append(list, CodeResourceExhausted)
	list = append(list, CodeFailedPrecondition)
	list = append(list, CodeAborted)
	list = append(list, CodeOutOfRange)
	list = append(list, CodeUnimplemented)
	list = append(list, CodeInternalError)
	list = append(list, CodeUnavailable)
	list = append(list, CodeDataLoss)
	list = append(list, CodeUndefined)
	list = append(list, CodeAuthorizationExpired)
	sort.Slice(list, func(i, j int) bool { return list[i].value < list[j].value })
	return list
}()

var codeToHTTPStatus = map[Code]*HTTPStatus{
	CodeOK:                   HTTPStatusOK,
	CodeInvalidArgument:      HTTPStatusBadRequest,
	CodeFailedPrecondition:   HTTPStatusBadRequest,
	CodeOutOfRange:           HTTPStatusBadRequest,
	CodeUnauthenticated:      HTTPStatusUnauthorized,
	CodePermissionDenied:     HTTPStatusForbidden,
	CodeNotFound:             HTTPStatusNotFound,
	CodeAborted:              HTTPStatusConflict,
	CodeAlreadyExists:        HTTPStatusConflict,
	CodeResourceExhausted:    HTTPStatusTooManyRequests,
	CodeCancelled:            HTTPStatusClientClosedRequest,
	CodeDataLoss:             HTTPStatusInternalServerError,
	CodeUnknown:              HTTPStatusInternalServerError,
	CodeInternalError:        HTTPStatusInternalServerError,
	CodeUnimplemented:        HTTPStatusNotImplemented,
	CodeUnavailable:          HTTPStatusServiceUnavailable,
	CodeDeadlineExceeded:     HTTPStatusTimeout,
	CodeAuthorizationExpired: HTTPStatusUnauthorized,
}

// Code represents a status code of an operation.
type Code struct {
	name  string
	value int
}

func newCode(name string, value int) Code {
	return Code{
		name:  name,
		value: value,
	}
}

func (c *Code) Name() string {
	return c.name
}

// Value returns the numerical value of this code.
func (c *Code) Value() int {
	return c.value
}

// ToHTTPStatus returns the HTTPStatus corresponding to this status code.
func (c *Code) ToHTTPStatus() *HTTPStatus {
	return codeToHTTPStatus[*c]
}

func (c *Code) String() string {
	return fmt.Sprintf("%s(%d)", c.name, c.value)
}
