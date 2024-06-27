package domainerr

import (
	"fmt"
	"log"
	"strings"
)

// A pseudo-enum of Status instances mapped 1:1 with the Codes. This simplifies construction
// patterns for derived instances of Status.
var (
	// StatusOK means the operation completed successfully.
	//
	// HTTP Mapping: 200 OK
	StatusOK = statusOK.copy()

	// StatusCancelled means the operation was cancelled (typically by the caller).
	//
	// HTTP Mapping: 499 Client Closed Request
	StatusCancelled = statusCancelled.copy()

	// StatusUnknown may be returned when a `Status` value received
	// from another address space belongs to an error space that is not known
	// in this address space. Also errors raised by APIs that do not return
	// enough error information may be converted to this error.
	//
	// HTTP Mapping: 500 Internal Server Error
	StatusUnknown = statusUnknown.copy()

	// StatusInvalidArgument means the client specified an invalid argument.
	// Note that this differs from `FAILED_PRECONDITION`. `INVALID_ARGUMENT`
	// indicates arguments that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	//
	// HTTP Mapping: 400 Bad Request
	StatusInvalidArgument = statusInvalidArgument.copy()

	// StatusDeadlineExceeded means the deadline expired before the operation
	// could complete. For operations that change the state of the system,
	// this error may be returned, even if the operation has completed successfully.
	// For example, a successful response from a server could have been delayed long
	// enough for the deadline to expire.
	//
	// HTTP Mapping: 504 Gateway Timeout
	StatusDeadlineExceeded = statusDeadlineExceeded.copy()

	// StatusNotFound means some requested entity (e.g., file or directory) was not found.
	//
	// Note to server developers: if a request is denied for an entire class
	// of users, such as gradual feature rollout or undocumented allow list,
	// `NOT_FOUND` may be used. If a request is denied for some users within
	// a class of users, such as user-based access control, `PERMISSION_DENIED`
	// must be used.
	//
	// HTTP Mapping: 404 Not Found
	StatusNotFound = statusNotFound.copy()

	// StatusAlreadyExists means the entity that a client attempted to create
	// (e.g., file or directory) already exists.
	//
	// HTTP Mapping: 409 Conflict
	StatusAlreadyExists = statusAlreadyExists.copy()

	// StatusPermissionDenied means the caller does not have permission to execute the specified
	// operation. `PERMISSION_DENIED` must not be used for rejections
	// caused by exhausting some resource (use `RESOURCE_EXHAUSTED`
	// instead for those errors). `PERMISSION_DENIED` must not be
	// used if the caller can not be identified (use `UNAUTHENTICATED`
	// instead for those errors). This error code does not imply the
	// request is valid or the requested entity exists or satisfies
	// other pre-conditions.
	//
	// HTTP Mapping: 403 Forbidden
	StatusPermissionDenied = statusPermissionDenied.copy()

	// StatusUnauthenticated means the request does not have valid authentication
	// credentials for the operation.
	//
	// HTTP Mapping: 401 Unauthorized
	StatusUnauthenticated = statusUnauthenticated.copy()

	// StatusResourceExhausted means some resource has been exhausted,
	// perhaps a per-user quota, or perhaps the entire file system is out of space.
	//
	// HTTP Mapping: 429 Too Many Requests
	StatusResourceExhausted = statusResourceExhausted.copy()

	// StatusFailedPrecondition means the operation was rejected because the system is not in
	// a state required for the operation's execution.  For example, the directory
	// to be deleted is non-empty, an rmdir operation is applied to
	// a non-directory, etc.
	//
	// Service implementors can use the following guidelines to decide
	// between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`:
	//  (a) Use `UNAVAILABLE` if the client can retry just the failing call.
	//  (b) Use `ABORTED` if the client should retry at a higher level. For
	//      example, when a client-specified test-and-set fails, indicating the
	//      client should restart a read-modify-write sequence.
	//  (c) Use `FAILED_PRECONDITION` if the client should not retry until
	//      the system state has been explicitly fixed. For example, if an "rmdir"
	//      fails because the directory is non-empty, `FAILED_PRECONDITION`
	//      should be returned since the client should not retry unless
	//      the files are deleted from the directory.
	//
	// HTTP Mapping: 400 Bad Request
	StatusFailedPrecondition = statusFailedPrecondition.copy()

	// StatusAborted means the operation was aborted, typically due to
	// a concurrency issue such as a sequencer check failure or transaction abort.
	//
	// See the guidelines above for deciding between `FAILED_PRECONDITION`,
	// `ABORTED`, and `UNAVAILABLE`.
	//
	// HTTP Mapping: 409 Conflict
	StatusAborted = statusAborted.copy()

	// StatusOutOfRange means the operation was attempted past the valid range.
	// E.g., seeking or reading past end-of-file.
	//
	// Unlike `INVALID_ARGUMENT`, this error indicates a problem that may
	// be fixed if the system state changes. For example, a 32-bit file
	// system will generate `INVALID_ARGUMENT` if asked to read at an
	// offset that is not in the range [0,2^32-1], but it will generate
	// `OUT_OF_RANGE` if asked to read from an offset past the current
	// file size.
	//
	// There is a fair bit of overlap between `FAILED_PRECONDITION` and
	// `OUT_OF_RANGE`.  We recommend using `OUT_OF_RANGE` (the more specific
	// error) when it applies so that callers who are iterating through
	// a space can easily look for an `OUT_OF_RANGE` error to detect when
	// they are done.
	//
	// HTTP Mapping: 400 Bad Request
	StatusOutOfRange = statusOutOfRange.copy()

	// StatusUnimplemented means the operation is not implemented or is
	// not supported/enabled in this service.
	//
	// HTTP Mapping: 501 Not Implemented
	StatusUnimplemented = statusUnimplemented.copy()

	// StatusInternal means internal errors. This means that some invariants expected by the
	// underlying system have been broken. This error code is reserved for serious errors.
	//
	// HTTP Mapping: 500 Internal Server Error
	StatusInternal = statusInternal.copy()

	// StatusUnavailable means the service is currently unavailable. This is most likely a
	// transient condition, which can be corrected by retrying with
	// a backoff. Note that it is not always safe to retry
	// non-idempotent operations.
	//
	// See the guidelines above for deciding between `FAILED_PRECONDITION`,
	// `ABORTED`, and `UNAVAILABLE`.
	//
	// HTTP Mapping: 503 Service Unavailable
	StatusUnavailable = statusUnavailable.copy()

	// StatusDataLoss means unrecoverable data loss or corruption.
	//
	// HTTP Mapping: 500 Internal Server Error
	StatusDataLoss = statusDataLoss.copy()

	// StatusUndefined means that the API operation/method is not defined on the target resource.
	//
	// HTTP Mapping: 405 Method Not Allowed
	StatusUndefined = statusUndefined.copy()

	// StatusAuthorizationExpired means a user's authorization expired, and it is
	// needed to log-in again and reauthorize.
	//
	// HTTP Mapping: 401 Unauthorized
	StatusAuthorizationExpired = statusAuthorizationExpired.copy()
)

var (
	statusOK                   = newStatus(CodeOK)
	statusCancelled            = newStatus(CodeCancelled)
	statusUnknown              = newStatus(CodeUnknown)
	statusInvalidArgument      = newStatus(CodeInvalidArgument)
	statusDeadlineExceeded     = newStatus(CodeDeadlineExceeded)
	statusNotFound             = newStatus(CodeNotFound)
	statusAlreadyExists        = newStatus(CodeAlreadyExists)
	statusPermissionDenied     = newStatus(CodePermissionDenied)
	statusUnauthenticated      = newStatus(CodeUnauthenticated)
	statusResourceExhausted    = newStatus(CodeResourceExhausted)
	statusFailedPrecondition   = newStatus(CodeFailedPrecondition)
	statusAborted              = newStatus(CodeAborted)
	statusOutOfRange           = newStatus(CodeOutOfRange)
	statusUnimplemented        = newStatus(CodeUnimplemented)
	statusInternal             = newStatus(CodeInternalError)
	statusUnavailable          = newStatus(CodeUnavailable)
	statusDataLoss             = newStatus(CodeDataLoss)
	statusUndefined            = newStatus(CodeUndefined)
	statusAuthorizationExpired = newStatus(CodeAuthorizationExpired)
)

// statusList contains all the well-defined operation statuses indexed by their code values
var statusList = func() []Status {
	list := make([]Status, 0, len(CodeList))
	for _, code := range CodeList {
		list = append(list, newStatus(code))
	}
	return list
}()

var httpStatusToStatus = map[*HTTPStatus]*Status{
	HTTPStatusOK:                  StatusOK,
	HTTPStatusBadRequest:          StatusInvalidArgument,
	HTTPStatusUnauthorized:        StatusUnauthenticated,
	HTTPStatusForbidden:           StatusPermissionDenied,
	HTTPStatusNotFound:            StatusNotFound,
	HTTPStatusMethodNotAllowed:    StatusUndefined,
	HTTPStatusConflict:            StatusAlreadyExists,
	HTTPStatusTooManyRequests:     StatusResourceExhausted,
	HTTPStatusClientClosedRequest: StatusCancelled,
	HTTPStatusInternalServerError: StatusInternal,
	HTTPStatusNotImplemented:      StatusUnimplemented,
	HTTPStatusServiceUnavailable:  StatusUnavailable,
	HTTPStatusTimeout:             StatusDeadlineExceeded,
}

// NewByHTTPStatus returns a copy of the status prototype mapped to given http status code.
func NewByHTTPStatus(statusCode int) *Status {
	httpStatus, isDefined := StatusWithCode(statusCode)
	if !isDefined {
		return StatusUnknown
	}

	// Internally assure that there must be a unique operation status mapped to any defined https status
	// in order that the caller can take the fluid coding style.
	opStatus, found := httpStatusToStatus[httpStatus]
	if found {
		log.Printf("[Error] not found op-status mapped to given defined http status %v\n", statusCode)
	}
	return opStatus
}

// NewWithCodeValue returns a copy of the status prototype mapped to given op status code.
func NewWithCodeValue(codeValue int) *Status {
	if codeValue < 0 || codeValue > CodeList[len(CodeList)-1].Value() {
		return StatusUnknown.WithMessagef("Unknown op status code: %v", codeValue)
	}
	return &statusList[codeValue]
}

// NewWithCode returns a copy of the status prototype mapped to given op status code.
func NewWithCode(code Code) *Status {
	return &statusList[code.value]
}

// Status defines the status of an operation by providing a standard Code in conjunction with an
// optional Case and an optional message. Instances of Status are created by starting with the
// template for the appropriate Code and supplementing it with additional information:
//
//	StatusNotFound.WithMessage("Could not find 'important_file.txt'")
type Status struct {
	code         Code
	specificCase Case
	// A developer-facing error message, which should be in English. Any
	// user-facing error message should be localized and sent in the
	// details field, or localized by the client.
	message string
	details any
}

func newStatus(code Code) Status {
	return Status{
		code: code,
	}
}

// WithMessage returns a derived instance of this Status with the given message. Leading and
// trailing whitespace is removed.
func (s *Status) WithMessage(msg string) *Status {
	msg = strings.TrimSpace(msg)
	if s.message == msg {
		// variable 'copy' collides with the 'builtin' function, so name it _copy
		_copy := *s
		return &_copy // return a copy of this Status
	}
	return &Status{
		code:         s.code,
		specificCase: s.specificCase,
		message:      msg,
		details:      s.details,
	}
}

// WithMessagef returns a derived instance of this Status with the formatted message. Leading and
// trailing whitespace is removed.
func (s *Status) WithMessagef(msgFmt string, fmtArgs ...any) *Status {
	return s.WithMessage(fmt.Sprintf(msgFmt, fmtArgs...))
}

// WithCase returns a derived instance of this Status with the given case.
func (s *Status) WithCase(c Case) *Status {
	if s.specificCase == c { // todo 深度比较 case
		_copy := *s
		return &_copy // return a copy of this Status
	}
	return &Status{
		code:         s.code,
		specificCase: c,
		message:      s.message,
		details:      s.details,
	}
}

// WithCaseAndMsg returns a derived instance of this Status with the given case and message.
func (s *Status) WithCaseAndMsg(theCase Case, message string) *Status {
	message = strings.TrimSpace(message)
	if s.specificCase == theCase && s.message == message { // todo 深度比较 case
		_copy := *s
		return &_copy
	}
	return &Status{
		code:         s.code,
		specificCase: theCase,
		message:      message,
		details:      s.details,
	}
}

// WithCaseAndMsgf returns a derived instance of this Status with the given case and formatted message.
func (s *Status) WithCaseAndMsgf(theCase Case, msgFmt string, fmtArgs ...any) *Status {
	msg := fmt.Sprintf(msgFmt, fmtArgs...)
	return s.WithCaseAndMsg(theCase, msg)
}

// AugmentMessage augments this Status's message with more contextual information of current use case scenario.
func (s *Status) AugmentMessage(moreContext string) {
	if moreContext == "" {
		return
	}

	newMsg := ""
	if s.message == "" {
		newMsg = moreContext
	} else {
		newMsg = s.message + "\n" + moreContext
	}
	s.message = newMsg
}

func (s *Status) WithDetails(v any) *Status {
	return &Status{
		code:         s.code,
		specificCase: s.specificCase,
		message:      s.message,
		details:      v,
	}
}

func (s *Status) Code() Code {
	return s.code
}

func (s *Status) Message() string {
	return s.message
}

func (s *Status) SpecificCase() Case {
	return s.specificCase
}

func (s *Status) Details() any {
	return s.details
}

// IsOK tells if this status is OK, i.e., not an error
func (s *Status) IsOK() bool {
	return s.code == CodeOK
}

// String returns the message prefixed with "{status name}: ".
func (s *Status) String() string {
	if s.message == "" {
		return s.code.Name()
	}
	return s.code.Name() + ": " + s.message
}

func (s *Status) ToObjStyleStr() string {
	var b strings.Builder
	b.Grow(200)
	fmt.Fprintf(&b, "%s{", s.code.Name())
	fmt.Fprintf(&b, "code:%d", s.code.Value())
	if s.specificCase != nil {
		fmt.Fprintf(&b, `,specificCase:"%s"`, s.specificCase.Identifier())
	}
	fmt.Fprintf(&b, `,message:"%s"`, s.Message())
	if s.details != nil {
		fmt.Fprintf(&b, ",details:%+v", s.Details())
	}
	fmt.Fprintf(&b, "}")
	return b.String()
}

// RetryAdvice provides advice on retry for this status.
func (s *Status) RetryAdvice() RetryAdvice {
	advice := NoAdvice
	if s.code == CodeUnavailable {
		advice = JustRetryFailingCall
	} else if s.code == CodeFailedPrecondition {
		advice = NotRetryUntilStateFixed
	} else if s.code == CodeAborted || s.code == CodeResourceExhausted {
		advice = RetryAtHigherLevel
	}
	return advice
}

func (s *Status) Equal(s2 *Status) bool {
	if s.specificCase == nil {
		return s.code == s2.code
	}
	return s.code == s2.code && s.specificCase.Identifier() == s.SpecificCase().Identifier()
}

func (s *Status) copy() *Status {
	copy := *s
	return &copy
}
