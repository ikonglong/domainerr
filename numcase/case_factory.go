package numcase

import (
	"bytes"
	"math"
	"sort"
	"strconv"

	"github.com/ikonglong/domainerr"
)

type CaseFactory struct {
	codingStrategy *CodingStrategy
	appCode        int
	moduleCode     int
}

type FactoryOpt func(f *CaseFactory) error

func WithAppCode(appCode int) FactoryOpt {
	return func(f *CaseFactory) error {
		err := domainerr.CheckArgument(appCode >= 0, "appCode < 0")
		if err != nil {
			return err
		}
		f.appCode = appCode
		return nil
	}
}

func WithModuleCode(moduleCode int) FactoryOpt {
	return func(f *CaseFactory) error {
		err := domainerr.CheckArgument(moduleCode >= 0, "moduleCode < 0")
		if err != nil {
			return err
		}
		f.moduleCode = moduleCode
		return nil
	}
}

func NewFactory(codingStrategy *CodingStrategy, opts ...FactoryOpt) (*CaseFactory, error) {
	err := domainerr.CheckArgument(codingStrategy != nil, "codingStrategy is nil")
	if err != nil {
		return nil, err
	}

	f := &CaseFactory{codingStrategy: codingStrategy}
	for _, setOpt := range opts {
		err = setOpt(f)
		if err != nil {
			return nil, err
		}
	}

	err = domainerr.CheckArgument(codingStrategy.appCodeRange.include(f.appCode),
		"appCodeRange %v not include appCode %d", codingStrategy.appCodeRange, f.appCode)
	if err != nil {
		return nil, err
	}

	err = domainerr.CheckArgument(codingStrategy.moduleCodeRange.include(f.moduleCode),
		"moduleCodeRange %v not include moduleCode %d", codingStrategy.moduleCodeRange, f.moduleCode)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// NewInvalidArgument creates a case that represents a more specific InvalidArgument status.
//
// The arg caseCode must be in the code segment corresponding to InvalidArgument status.
//
// If the out-of-box DefaultCodeMapper is used, the code segment corresponding to
// InvalidArgument status is [1, 50].
func (f *CaseFactory) NewInvalidArgument(caseCode int) (*NumCase, error) {
	return f.create(domainerr.CodeInvalidArgument, caseCode)
}

// NewDeadlineExceeded creates a case that represents a more specific DeadlineExceeded status.
//
// The arg caseCode must be in the code segment corresponding to DeadlineExceeded status.
//
// If the out-of-box DefaultCodeMapper is used, the code segment corresponding to
// DeadlineExceeded status is [51, 100].
func (f *CaseFactory) NewDeadlineExceeded(caseCode int) (*NumCase, error) {
	return f.create(domainerr.CodeDeadlineExceeded, caseCode)
}

// NewNotFound creates a case that represents a more specific NotFound status.
//
// The arg caseCode must be in the code segment corresponding to NotFound status.
//
// If the out-of-box DefaultCodeMapper is used, the code segment corresponding to
// NotFound status is [101, 150].
func (f *CaseFactory) NewNotFound(caseCode int) (*NumCase, error) {
	return f.create(domainerr.CodeNotFound, caseCode)
}

// NewAlreadyExists creates a case that represents a more specific AlreadyExists status.
//
// The arg caseCode must be in the code segment corresponding to AlreadyExists status.
//
// If the out-of-box DefaultCodeMapper is used, the code segment corresponding to
// AlreadyExists status is [151, 200].
func (f *CaseFactory) NewAlreadyExists(caseCode int) (*NumCase, error) {
	return f.create(domainerr.CodeAlreadyExists, caseCode)
}

// NewPermissionDenied creates a case that represents a more specific PermissionDenied status.
//
// The arg caseCode must be in the code segment corresponding to PermissionDenied status.
//
// If the out-of-box DefaultCodeMapper is used, the code segment corresponding to
// PermissionDenied status is [201, 250].
func (f *CaseFactory) NewPermissionDenied(caseCode int) (*NumCase, error) {
	return f.create(domainerr.CodePermissionDenied, caseCode)
}

// NewResourceExhausted creates a case that represents a more specific ResourceExhausted status.
//
// The arg caseCode must be in the code segment corresponding to ResourceExhausted status.
//
// If the out-of-box DefaultCodeMapper is used, the code segment corresponding to
// ResourceExhausted status is [251, 300].
func (f *CaseFactory) NewResourceExhausted(caseCode int) (*NumCase, error) {
	return f.create(domainerr.CodeResourceExhausted, caseCode)
}

// NewFailedPrecondition creates a case that represents a more specific FailedPrecondition status.
//
// The arg caseCode must be in the code segment corresponding to FailedPrecondition status.
//
// If the out-of-box DefaultCodeMapper is used, the code segment corresponding to
// FailedPrecondition status is [301, 350].
func (f *CaseFactory) NewFailedPrecondition(caseCode int) (*NumCase, error) {
	return f.create(domainerr.CodeFailedPrecondition, caseCode)
}

// NewAborted creates a case that represents a more specific Aborted status.
//
// The arg caseCode must be in the code segment corresponding to Aborted status.
//
// If the out-of-box DefaultCodeMapper is used, the code segment corresponding to
// Aborted status is [351, 400].
func (f *CaseFactory) NewAborted(caseCode int) (*NumCase, error) {
	return f.create(domainerr.CodeAborted, caseCode)
}

// NewOutOfRange creates a case that represents a more specific OutOfRange status.
//
// The arg caseCode must be in the code segment corresponding to OutOfRange status.
//
// If the out-of-box DefaultCodeMapper is used, the code segment corresponding to
// OutOfRange status is [401, 450].
func (f *CaseFactory) NewOutOfRange(caseCode int) (*NumCase, error) {
	return f.create(domainerr.CodeOutOfRange, caseCode)
}

// NewInternalError creates a case that represents a more specific InternalError status.
//
// The arg caseCode must be in the code segment corresponding to InternalError status.
//
// If the out-of-box DefaultCodeMapper is used, the code segment corresponding to
// InternalError status is [451, 500].
func (f *CaseFactory) NewInternalError(caseCode int) (*NumCase, error) {
	return f.create(domainerr.CodeInternalError, caseCode)
}

// NewDataLoss creates a case that represents a more specific DataLoss status.
//
// The arg caseCode must be in the code segment corresponding to DataLoss status.
//
// If the out-of-box DefaultCodeMapper is used, the code segment corresponding to
// DataLoss status is [501, 550].
func (f *CaseFactory) NewDataLoss(caseCode int) (*NumCase, error) {
	return f.create(domainerr.CodeDataLoss, caseCode)
}

func (f *CaseFactory) create(statusCode domainerr.Code, caseCode int) (*NumCase, error) {
	codeSeg := f.codingStrategy.statusCodeMapper.CaseCodeSegmentFor(statusCode)
	err := domainerr.CheckArgument(codeSeg != nil,
		"statusCodeMapper doesn't define a CaseCodeSegment for status code %s", statusCode.String())
	if err != nil {
		return nil, err
	}
	err = domainerr.CheckArgument(codeSeg.include(caseCode),
		"CaseCodeSegment %s for status code %s doesn't include code %d",
		codeSeg.String(), statusCode.String(), caseCode)
	if err != nil {
		return nil, err
	}

	var caseID bytes.Buffer
	if f.codingStrategy.numDigitsOfAppCode > 0 {
		caseID.WriteString(f.padLeftZeros(f.appCode, f.codingStrategy.numDigitsOfAppCode))
		caseID.WriteByte('_')
	}
	if f.codingStrategy.numDigitsOfModuleCode > 0 {
		caseID.WriteString(f.padLeftZeros(f.moduleCode, f.codingStrategy.numDigitsOfModuleCode))
		caseID.WriteByte('_')
	}
	caseID.WriteString(f.padLeftZeros(caseCode, f.codingStrategy.numDigitsOfCaseCode))
	return newNumCase(f.appCode, f.moduleCode, caseCode, caseID.String(), statusCode), nil
}

func (f *CaseFactory) padLeftZeros(num int, minLen int) string {
	s := strconv.Itoa(num)
	if len(s) >= minLen {
		return s
	}

	var buf bytes.Buffer
	for i := len(s); i < minLen; i++ {
		buf.WriteByte('0')
	}
	buf.WriteString(s)
	return buf.String()
}

type CodingStrategy struct {
	numDigitsOfAppCode int
	appCodeRange       *NumRange

	numDigitsOfModuleCode int
	moduleCodeRange       *NumRange

	numDigitsOfCaseCode int
	caseCodeRange       *NumRange
	statusCodeMapper    CodeMapper
}

type CodingStrategyBuilder struct {
	s *CodingStrategy
}

func NewCodingStrategyBuilder() *CodingStrategyBuilder {
	return &CodingStrategyBuilder{
		s: &CodingStrategy{},
	}
}

func (b *CodingStrategyBuilder) NumDigitsOfAppCode(n int) *CodingStrategyBuilder {
	b.s.numDigitsOfAppCode = n
	return b
}

func (b *CodingStrategyBuilder) NumDigitsOfModuleCode(n int) *CodingStrategyBuilder {
	b.s.numDigitsOfModuleCode = n
	return b
}

func (b *CodingStrategyBuilder) NumDigitsOfCaseCode(n int) *CodingStrategyBuilder {
	b.s.numDigitsOfCaseCode = n
	return b
}

func (b *CodingStrategyBuilder) StatusCodeMapper(m CodeMapper) *CodingStrategyBuilder {
	b.s.statusCodeMapper = m
	return b
}

func (b *CodingStrategyBuilder) Build() (*CodingStrategy, error) {
	err := domainerr.CheckArgument(b.s.numDigitsOfAppCode >= 0, "numDigitsOfAppCode < 0")
	if err != nil {
		return nil, err
	}
	err = domainerr.CheckArgument(b.s.numDigitsOfModuleCode >= 0, "numDigitsOfModuleCode < 0")
	if err != nil {
		return nil, err
	}
	err = domainerr.CheckArgument(b.s.numDigitsOfCaseCode >= 0, "numDigitsOfCaseCode < 0")
	if err != nil {
		return nil, err
	}
	err = domainerr.CheckArgument(b.s.statusCodeMapper != nil, "statusCodeMapper is nil")
	if err != nil {
		return nil, err
	}

	b.s.appCodeRange, _ = NewNumRange(0, domainerr.FloatToInt(math.Pow10(b.s.numDigitsOfAppCode))-1)
	b.s.moduleCodeRange, _ = NewNumRange(0, domainerr.FloatToInt(math.Pow10(b.s.numDigitsOfModuleCode))-1)
	b.s.caseCodeRange, _ = NewNumRange(0, domainerr.FloatToInt(math.Pow10(b.s.numDigitsOfCaseCode))-1)

	segs := b.s.statusCodeMapper.CaseCodeSegments()
	sort.Slice(segs, func(i, j int) bool {
		return segs[i].End() <= segs[j].Start()
	})
	for _, cs := range segs {
		err = domainerr.CheckArgument(b.s.caseCodeRange.includeRange(cs),
			"case code range %v of CodingStrategy doesn't include caseCodeSegment %v defined by CodeMapper",
			b.s.caseCodeRange, cs)
		if err != nil {
			return nil, err
		}
	}
	return b.s, nil
}
