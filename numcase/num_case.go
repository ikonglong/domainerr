package numcase

import (
	"github.com/ikonglong/domainerr"
)

type NumCase struct {
	appCode    int
	moduleCode int
	caseCode   int

	identifier string
	statusCode domainerr.Code
}

func newNumCase(appCode int, moduleCode int, caseCode int, identifier string, statusCode domainerr.Code) *NumCase {
	return &NumCase{
		appCode:    appCode,
		moduleCode: moduleCode,
		caseCode:   caseCode,
		identifier: identifier,
		statusCode: statusCode,
	}
}

func (c *NumCase) Identifier() string {
	return c.identifier
}

func (c *NumCase) StatusCode() domainerr.Code {
	return c.statusCode
}
