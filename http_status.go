package domainerr

import (
	"fmt"
)

var (
	HTTPStatusOK                  = newHTTPStatus("OK", 200)
	HTTPStatusBadRequest          = newHTTPStatus("BadRequest", 400)
	HTTPStatusUnauthorized        = newHTTPStatus("Unauthorized", 401)
	HTTPStatusForbidden           = newHTTPStatus("Forbidden", 403)
	HTTPStatusNotFound            = newHTTPStatus("NotFound", 404)
	HTTPStatusMethodNotAllowed    = newHTTPStatus("MethodNotAllowed", 405)
	HTTPStatusConflict            = newHTTPStatus("Conflict", 409)
	HTTPStatusTooManyRequests     = newHTTPStatus("TooManyRequests", 429)
	HTTPStatusClientClosedRequest = newHTTPStatus("ClientClosedRequest", 499)
	HTTPStatusInternalServerError = newHTTPStatus("InternalServerError", 500)
	HTTPStatusNotImplemented      = newHTTPStatus("NotImplemented", 501)
	HTTPStatusServiceUnavailable  = newHTTPStatus("ServiceUnavailable", 503)
	HTTPStatusTimeout             = newHTTPStatus("Timeout", 504)

	httpStatusList = func() []*HTTPStatus {
		list := make([]*HTTPStatus, 0, 12)
		list = append(list, HTTPStatusOK)
		list = append(list, HTTPStatusBadRequest)
		list = append(list, HTTPStatusUnauthorized)
		list = append(list, HTTPStatusForbidden)
		list = append(list, HTTPStatusNotFound)
		list = append(list, HTTPStatusMethodNotAllowed)
		list = append(list, HTTPStatusConflict)
		list = append(list, HTTPStatusTooManyRequests)
		list = append(list, HTTPStatusClientClosedRequest)
		list = append(list, HTTPStatusInternalServerError)
		list = append(list, HTTPStatusNotImplemented)
		list = append(list, HTTPStatusServiceUnavailable)
		list = append(list, HTTPStatusTimeout)
		return list
	}()

	codeToStatus = func() map[int]*HTTPStatus {
		aMap := make(map[int]*HTTPStatus, 12)
		for _, status := range httpStatusList {
			aMap[status.Code()] = status
		}
		return aMap
	}()
)

// StatusWithCode returns the HTTPStatus with the given code and true if the code is defined.
// Otherwise, it returns (nil, false).
func StatusWithCode(statusCode int) (*HTTPStatus, bool) {
	s, found := codeToStatus[statusCode]
	return s, found
}

type HTTPStatus struct {
	name string
	code int
}

func newHTTPStatus(name string, code int) *HTTPStatus {
	return &HTTPStatus{
		name: name,
		code: code,
	}
}

func (s *HTTPStatus) Code() int {
	return s.code
}

func (s *HTTPStatus) Name() string {
	return s.name
}

func (s *HTTPStatus) String() string {
	return fmt.Sprintf("%s(%v)", s.Name(), s.Code())
}
