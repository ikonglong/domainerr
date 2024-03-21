package numcase

import (
	"testing"

	"github.com/ikonglong/domainerr"

	"github.com/stretchr/testify/assert"
)

func TestNewCodeMapper(t *testing.T) {
	cm := NewCodeMapper(&DefaultCodeMapper{})
	if cm == nil {
		t.Error("NewCodeMapper() should not be nil")
	}
}

var defaultCodeMapper = NewCodeMapper(&DefaultCodeMapper{})

func TestHasMappingFor(t *testing.T) {
	for _, code := range domainerr.CodeList {
		if notNeedToMap(code) {
			continue
		}
		assert.True(t, defaultCodeMapper.HasMappingFor(code),
			"defaultCodeMapper should have mapping for %s", code.String())
	}
}

func TestCaseCodeSegmentFor(t *testing.T) {
	for _, code := range domainerr.CodeList {
		if notNeedToMap(code) {
			continue
		}
		assert.NotNil(t, defaultCodeMapper.CaseCodeSegmentFor(code),
			"defaultCodeMapper should map %s to a CaseCodeSegment", code.String())
	}
}

var codesNotNeedToMap = []domainerr.Code{
	domainerr.CodeOK, domainerr.CodeCancelled,
	domainerr.CodeUnknown, domainerr.CodeUnauthenticated,
	domainerr.CodeUnimplemented, domainerr.CodeUnavailable,
	domainerr.CodeAuthorizationExpired, domainerr.CodeUndefined,
}

func notNeedToMap(code domainerr.Code) bool {
	for _, e := range codesNotNeedToMap {
		if code == e {
			return true
		}
	}
	return false
}

func numCodeNotNeedToMap() int {
	return 8
}

func TestCaseCodeSegments(t *testing.T) {
	assert.Equal(t, len(domainerr.CodeList)-numCodeNotNeedToMap(), len(defaultCodeMapper.CaseCodeSegments()))
}

func TestMappings(t *testing.T) {
	assert.Equal(t, len(domainerr.CodeList)-numCodeNotNeedToMap(), len(defaultCodeMapper.Mappings()))
}

func TestInvalidArgument(t *testing.T) {
	r, _ := NewNumRange(1, 50)
	assert.Equal(t, r, defaultCodeMapper.InvalidArgument(),
		"CaseCodeSegment for code %s should be %s", domainerr.CodeInvalidArgument.String(), r.String())
}

func TestDeadlineExceeded(t *testing.T) {
	r, _ := NewNumRange(51, 100)
	assert.Equal(t, r, defaultCodeMapper.DeadlineExceeded(),
		"CaseCodeSegment for code %s should be %s", domainerr.CodeDeadlineExceeded.String(), r.String())
}

func TestNotFound(t *testing.T) {
	r, _ := NewNumRange(101, 150)
	assert.Equal(t, r, defaultCodeMapper.NotFound(),
		"CaseCodeSegment for code %s should be %s", domainerr.CodeNotFound.String(), r.String())
}

func TestAlreadyExists(t *testing.T) {
	r, _ := NewNumRange(151, 200)
	assert.Equal(t, r, defaultCodeMapper.AlreadyExists(),
		"CaseCodeSegment for code %s should be %s", domainerr.CodeAlreadyExists.String(), r.String())
}

func TestPermissionDenied(t *testing.T) {
	r, _ := NewNumRange(201, 250)
	assert.Equal(t, r, defaultCodeMapper.PermissionDenied(),
		"CaseCodeSegment for code %s should be %s", domainerr.CodePermissionDenied.String(), r.String())
}

func TestResourceExhausted(t *testing.T) {
	r, _ := NewNumRange(251, 300)
	assert.Equal(t, r, defaultCodeMapper.ResourceExhausted(),
		"CaseCodeSegment for code %s should be %s", domainerr.CodeResourceExhausted.String(), r.String())
}

func TestFailedPrecondition(t *testing.T) {
	r, _ := NewNumRange(301, 350)
	assert.Equal(t, r, defaultCodeMapper.FailedPrecondition(),
		"CaseCodeSegment for code %s should be %s", domainerr.CodeFailedPrecondition.String(), r.String())
}

func TestAborted(t *testing.T) {
	r, _ := NewNumRange(351, 400)
	assert.Equal(t, r, defaultCodeMapper.Aborted(),
		"CaseCodeSegment for code %s should be %s", domainerr.CodeAborted.String(), r.String())
}

func TestOutOfRange(t *testing.T) {
	r, _ := NewNumRange(401, 450)
	assert.Equal(t, r, defaultCodeMapper.OutOfRange(),
		"CaseCodeSegment for code %s should be %s", domainerr.CodeOutOfRange.String(), r.String())
}

func TestInternalError(t *testing.T) {
	r, _ := NewNumRange(451, 500)
	assert.Equal(t, r, defaultCodeMapper.InternalError(),
		"CaseCodeSegment for code %s should be %s", domainerr.CodeInternalError.String(), r.String())
}

func TestDataLoss(t *testing.T) {
	r, _ := NewNumRange(501, 550)
	assert.Equal(t, r, defaultCodeMapper.DataLoss(),
		"CaseCodeSegment for code %s should be %s", domainerr.CodeDataLoss.String(), r.String())
}
