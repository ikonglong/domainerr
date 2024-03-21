package numcase

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ikonglong/domainerr"
)

type CaseCodeSegment = NumRange

// CodeMapper maps an operation status Code to a CaseCodeSegment.
type CodeMapper interface {
	// HasMappingFor tells whether there is a mapping from some CaseCodeSegment to given statusCode.
	HasMappingFor(statusCode domainerr.Code) bool
	// CaseCodeSegmentFor returns the CaseCodeSegment mapped to the given operation status Code if
	// such a mapping exists. Otherwise, it returns nil.
	CaseCodeSegmentFor(statusCode domainerr.Code) *CaseCodeSegment
	// CaseCodeSegments returns all the CaseCodeSegment list contained in this CodeMapper.
	CaseCodeSegments() []*CaseCodeSegment
	// Mappings returns all the mappings that are from operation status Code to CaseCodeSegment.
	Mappings() map[domainerr.Code]*CaseCodeSegment
	// InvalidArgument declares a mapping from CodeInvalidArgument to the returned CaseCodeSegment.
	InvalidArgument() *CaseCodeSegment
	// DeadlineExceeded declares a mapping from CodeDeadlineExceeded to the returned CaseCodeSegment.
	DeadlineExceeded() *CaseCodeSegment
	// NotFound declares a mapping from CodeNotFound to the returned CaseCodeSegment.
	NotFound() *CaseCodeSegment
	// AlreadyExists declares a mapping from CodeAlreadyExists to the returned CaseCodeSegment.
	AlreadyExists() *CaseCodeSegment
	// PermissionDenied declares a mapping from CodePermissionDenied to the returned CaseCodeSegment.
	PermissionDenied() *CaseCodeSegment
	// ResourceExhausted declares a mapping from CodeResourceExhausted to the returned CaseCodeSegment.
	ResourceExhausted() *CaseCodeSegment
	// FailedPrecondition declares a mapping from CodeFailedPrecondition to the returned CaseCodeSegment.
	FailedPrecondition() *CaseCodeSegment
	// Aborted declares a mapping from CodeAborted to the returned CaseCodeSegment.
	Aborted() *CaseCodeSegment
	// OutOfRange declares a mapping from CodeOutOfRange to the returned CaseCodeSegment.
	OutOfRange() *CaseCodeSegment
	// InternalError declares a mapping from CodeInternalError to the returned CaseCodeSegment.
	InternalError() *CaseCodeSegment
	// DataLoss declares a mapping from CodeDataLoss to the returned CaseCodeSegment.
	DataLoss() *CaseCodeSegment
}

type CodeMapperBase struct {
	CodeMapper
	statusCodeToCaseCodeSeg map[domainerr.Code]*CaseCodeSegment
}

func NewCodeMapper(concreteMapper CodeMapper) CodeMapper {
	mappings := make(map[domainerr.Code]*CaseCodeSegment, len(domainerr.CodeList))
	mappings[domainerr.CodeInvalidArgument] = concreteMapper.InvalidArgument()
	mappings[domainerr.CodeDeadlineExceeded] = concreteMapper.DeadlineExceeded()
	mappings[domainerr.CodeNotFound] = concreteMapper.NotFound()
	mappings[domainerr.CodeAlreadyExists] = concreteMapper.AlreadyExists()
	mappings[domainerr.CodePermissionDenied] = concreteMapper.PermissionDenied()
	mappings[domainerr.CodeResourceExhausted] = concreteMapper.ResourceExhausted()
	mappings[domainerr.CodeFailedPrecondition] = concreteMapper.FailedPrecondition()
	mappings[domainerr.CodeAborted] = concreteMapper.Aborted()
	mappings[domainerr.CodeOutOfRange] = concreteMapper.OutOfRange()
	mappings[domainerr.CodeInternalError] = concreteMapper.InternalError()
	mappings[domainerr.CodeDataLoss] = concreteMapper.DataLoss()
	return &CodeMapperBase{
		CodeMapper:              concreteMapper,
		statusCodeToCaseCodeSeg: mappings,
	}
}

func (m *CodeMapperBase) HasMappingFor(statusCode domainerr.Code) bool {
	_, present := m.statusCodeToCaseCodeSeg[statusCode]
	return present
}

func (m *CodeMapperBase) CaseCodeSegmentFor(statusCode domainerr.Code) *CaseCodeSegment {
	return m.statusCodeToCaseCodeSeg[statusCode]
}

func (m *CodeMapperBase) CaseCodeSegments() []*CaseCodeSegment {
	segments := make([]*CaseCodeSegment, 0, len(m.statusCodeToCaseCodeSeg))
	for _, segment := range m.statusCodeToCaseCodeSeg {
		segments = append(segments, segment)
	}
	return segments
}

func (m *CodeMapperBase) Mappings() map[domainerr.Code]*CaseCodeSegment {
	_copy := make(map[domainerr.Code]*CaseCodeSegment, len(m.statusCodeToCaseCodeSeg))
	for statusCode, segment := range m.statusCodeToCaseCodeSeg {
		_copy[statusCode] = segment
	}
	return _copy
}

func (m *CodeMapperBase) String() string {
	// invert the map, and sort the keys
	segmentToOpStatusCode := make(map[*CaseCodeSegment]domainerr.Code, len(m.statusCodeToCaseCodeSeg))
	segments := make([]*CaseCodeSegment, 0, len(segmentToOpStatusCode))
	for statusCode, segment := range m.statusCodeToCaseCodeSeg {
		segmentToOpStatusCode[segment] = statusCode
		segments = append(segments, segment)
	}
	sort.Slice(segments, func(i, j int) bool {
		return segments[i].End() <= segments[j].Start()
	})

	// build the string
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("+%s+\n", strings.Repeat("-", 88)))
	sb.WriteString(fmt.Sprintf("| %-20s | %-30s | %-30s |\n",
		"Case Code Segment", "Operation Status (Name:Code)", "HTTP Status (Name:Code)"))
	for _, segment := range segments {
		opStatusCode := segmentToOpStatusCode[segment]
		sb.WriteString(fmt.Sprintf("| %-20s | %-20s:%-10v | %-20s:%-10v |\n",
			segment, opStatusCode.Name(), opStatusCode.Value(),
			opStatusCode.ToHTTPStatus().Name(), opStatusCode.ToHTTPStatus().Code()))
	}
	return sb.String()
}

type DefaultCodeMapper struct {
	CodeMapperBase
}

func (m *DefaultCodeMapper) InvalidArgument() *CaseCodeSegment {
	r, _ := NewNumRange(1, 50)
	return r
}

func (m *DefaultCodeMapper) DeadlineExceeded() *CaseCodeSegment {
	r, _ := NewNumRange(51, 100)
	return r
}

func (m *DefaultCodeMapper) NotFound() *CaseCodeSegment {
	r, _ := NewNumRange(101, 150)
	return r
}

func (m *DefaultCodeMapper) AlreadyExists() *CaseCodeSegment {
	r, _ := NewNumRange(151, 200)
	return r
}

func (m *DefaultCodeMapper) PermissionDenied() *CaseCodeSegment {
	r, _ := NewNumRange(201, 250)
	return r
}

func (m *DefaultCodeMapper) ResourceExhausted() *CaseCodeSegment {
	r, _ := NewNumRange(251, 300)
	return r
}

func (m *DefaultCodeMapper) FailedPrecondition() *CaseCodeSegment {
	r, _ := NewNumRange(301, 350)
	return r
}

func (m *DefaultCodeMapper) Aborted() *CaseCodeSegment {
	r, _ := NewNumRange(351, 400)
	return r
}

func (m *DefaultCodeMapper) OutOfRange() *CaseCodeSegment {
	r, _ := NewNumRange(401, 450)
	return r
}

func (m *DefaultCodeMapper) InternalError() *CaseCodeSegment {
	r, _ := NewNumRange(451, 500)
	return r
}

func (m *DefaultCodeMapper) DataLoss() *CaseCodeSegment {
	r, _ := NewNumRange(501, 550)
	return r
}
