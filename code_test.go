package domainerr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusCode_OK(t *testing.T) {
	codeOK := CodeOK
	assert.Equal(t, "OK", codeOK.Name())
	assert.Equal(t, 0, codeOK.Value())
	assert.Equal(t, "OK(0)", codeOK.String())
	assert.Equal(t, HTTPStatusOK, codeOK.ToHTTPStatus())
}

func TestStatusCode_Cancelled(t *testing.T) {
	codeCancelled := CodeCancelled
	assert.Equal(t, "OperationCancelled", codeCancelled.Name())
	assert.Equal(t, 1, codeCancelled.Value())
	assert.Equal(t, "OperationCancelled(1)", codeCancelled.String())
	assert.Equal(t, HTTPStatusClientClosedRequest, codeCancelled.ToHTTPStatus())
}

func TestStatusCode_Unknown(t *testing.T) {
	codeUnknown := CodeUnknown
	assert.Equal(t, "UnknownError", codeUnknown.Name())
	assert.Equal(t, 2, codeUnknown.Value())
	assert.Equal(t, "UnknownError(2)", codeUnknown.String())
	assert.Equal(t, HTTPStatusInternalServerError, codeUnknown.ToHTTPStatus())
}

func TestStatusCode_InvalidArgument(t *testing.T) {
	codeInvalidArgument := CodeInvalidArgument
	assert.Equal(t, "InvalidArgument", codeInvalidArgument.Name())
	assert.Equal(t, 3, codeInvalidArgument.Value())
	assert.Equal(t, "InvalidArgument(3)", codeInvalidArgument.String())
	assert.Equal(t, HTTPStatusBadRequest, codeInvalidArgument.ToHTTPStatus())
}

func TestStatusCode_DeadlineExceeded(t *testing.T) {
	codeDeadlineExceeded := CodeDeadlineExceeded
	assert.Equal(t, "DeadlineExceeded", codeDeadlineExceeded.Name())
	assert.Equal(t, 4, codeDeadlineExceeded.Value())
	assert.Equal(t, "DeadlineExceeded(4)", codeDeadlineExceeded.String())
	assert.Equal(t, HTTPStatusTimeout, codeDeadlineExceeded.ToHTTPStatus())
}

func TestStatusCode_NotFound(t *testing.T) {
	codeNotFound := CodeNotFound
	assert.Equal(t, "NotFound", codeNotFound.Name())
	assert.Equal(t, 5, codeNotFound.Value())
	assert.Equal(t, "NotFound(5)", codeNotFound.String())
	assert.Equal(t, HTTPStatusNotFound, codeNotFound.ToHTTPStatus())
}

func TestStatusCode_AlreadyExists(t *testing.T) {
	codeAlreadyExists := CodeAlreadyExists
	assert.Equal(t, "AlreadyExists", codeAlreadyExists.Name())
	assert.Equal(t, 6, codeAlreadyExists.Value())
	assert.Equal(t, "AlreadyExists(6)", codeAlreadyExists.String())
	assert.Equal(t, HTTPStatusConflict, codeAlreadyExists.ToHTTPStatus())
}

func TestStatusCode_PermissionDenied(t *testing.T) {
	codePermissionDenied := CodePermissionDenied
	assert.Equal(t, "PermissionDenied", codePermissionDenied.Name())
	assert.Equal(t, 7, codePermissionDenied.Value())
	assert.Equal(t, "PermissionDenied(7)", codePermissionDenied.String())
	assert.Equal(t, HTTPStatusForbidden, codePermissionDenied.ToHTTPStatus())
}

func TestStatusCode_Unauthenticated(t *testing.T) {
	codeUnauthenticated := CodeUnauthenticated
	assert.Equal(t, "Unauthenticated", codeUnauthenticated.Name())
	assert.Equal(t, 16, codeUnauthenticated.Value())
	assert.Equal(t, "Unauthenticated(16)", codeUnauthenticated.String())
	assert.Equal(t, HTTPStatusUnauthorized, codeUnauthenticated.ToHTTPStatus())
}

func TestStatusCode_ResourceExhausted(t *testing.T) {
	codeResourceExhausted := CodeResourceExhausted
	assert.Equal(t, "ResourceExhausted", codeResourceExhausted.Name())
	assert.Equal(t, 8, codeResourceExhausted.Value())
	assert.Equal(t, "ResourceExhausted(8)", codeResourceExhausted.String())
	assert.Equal(t, HTTPStatusTooManyRequests, codeResourceExhausted.ToHTTPStatus())
}

func TestStatusCode_FailedPrecondition(t *testing.T) {
	codeFailedPrecondition := CodeFailedPrecondition
	assert.Equal(t, "FailedPrecondition", codeFailedPrecondition.Name())
	assert.Equal(t, 9, codeFailedPrecondition.Value())
	assert.Equal(t, "FailedPrecondition(9)", codeFailedPrecondition.String())
	assert.Equal(t, HTTPStatusBadRequest, codeFailedPrecondition.ToHTTPStatus())
}

func TestStatusCode_Aborted(t *testing.T) {
	codeAborted := CodeAborted
	assert.Equal(t, "OperationAborted", codeAborted.Name())
	assert.Equal(t, 10, codeAborted.Value())
	assert.Equal(t, "OperationAborted(10)", codeAborted.String())
	assert.Equal(t, HTTPStatusConflict, codeAborted.ToHTTPStatus())
}

func TestStatusCode_OutOfRange(t *testing.T) {
	codeOutOfRange := CodeOutOfRange
	assert.Equal(t, "OutOfRange", codeOutOfRange.Name())
	assert.Equal(t, 11, codeOutOfRange.Value())
	assert.Equal(t, "OutOfRange(11)", codeOutOfRange.String())
	assert.Equal(t, HTTPStatusBadRequest, codeOutOfRange.ToHTTPStatus())
}

func TestStatusCode_Unimplemented(t *testing.T) {
	codeUnimplemented := CodeUnimplemented
	assert.Equal(t, "OperationUnimplemented", codeUnimplemented.Name())
	assert.Equal(t, 12, codeUnimplemented.Value())
	assert.Equal(t, "OperationUnimplemented(12)", codeUnimplemented.String())
	assert.Equal(t, HTTPStatusNotImplemented, codeUnimplemented.ToHTTPStatus())
}

func TestStatusCode_InternalError(t *testing.T) {
	codeInternalError := CodeInternalError
	assert.Equal(t, "InternalError", codeInternalError.Name())
	assert.Equal(t, 13, codeInternalError.Value())
	assert.Equal(t, "InternalError(13)", codeInternalError.String())
	assert.Equal(t, HTTPStatusInternalServerError, codeInternalError.ToHTTPStatus())
}

func TestStatusCode_Unavailable(t *testing.T) {
	codeUnavailable := CodeUnavailable
	assert.Equal(t, "ServiceUnavailable", codeUnavailable.Name())
	assert.Equal(t, 14, codeUnavailable.Value())
	assert.Equal(t, "ServiceUnavailable(14)", codeUnavailable.String())
	assert.Equal(t, HTTPStatusServiceUnavailable, codeUnavailable.ToHTTPStatus())
}

func TestStatusCode_DataLoss(t *testing.T) {
	codeDataLoss := CodeDataLoss
	assert.Equal(t, "DataLoss", codeDataLoss.Name())
	assert.Equal(t, 15, codeDataLoss.Value())
	assert.Equal(t, "DataLoss(15)", codeDataLoss.String())
	assert.Equal(t, HTTPStatusInternalServerError, codeDataLoss.ToHTTPStatus())
}

func TestStatusCode_AuthorizationExpired(t *testing.T) {
	codeAuthorizationExpired := CodeAuthorizationExpired
	assert.Equal(t, "AuthorizationExpired", codeAuthorizationExpired.Name())
	assert.Equal(t, 30, codeAuthorizationExpired.Value())
	assert.Equal(t, "AuthorizationExpired(30)", codeAuthorizationExpired.String())
	assert.Equal(t, HTTPStatusUnauthorized, codeAuthorizationExpired.ToHTTPStatus())
}
