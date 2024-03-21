package numcase

import (
	"fmt"
	"testing"

	"github.com/ikonglong/domainerr"

	"github.com/stretchr/testify/assert"
)

var (
	csWith1DigitAppCodeAndModuleCode, _ = NewCodingStrategyBuilder().
						NumDigitsOfAppCode(1).
						NumDigitsOfModuleCode(1).
						NumDigitsOfCaseCode(3).
						StatusCodeMapper(NewCodeMapper(&DefaultCodeMapper{})).Build()
	csWithoutAppCodeAndModuleCode, _ = NewCodingStrategyBuilder().
						NumDigitsOfAppCode(0).
						NumDigitsOfModuleCode(0).
						NumDigitsOfCaseCode(3).
						StatusCodeMapper(NewCodeMapper(&DefaultCodeMapper{})).Build()
)

// Tests for CaseFactory start

func TestWithAppCode_IllegalAppCode(t *testing.T) {
	f := &CaseFactory{}
	err := WithAppCode(-1)(f)
	assert.NotNil(t, err)
	assert.Equal(t, "illegal argument: appCode < 0", err.Error())
}

func TestWithAppCode_LegalAppCode(t *testing.T) {
	f := &CaseFactory{}
	err := WithAppCode(1)(f)
	assert.Nil(t, err)
	assert.Equal(t, 1, f.appCode)
}

func TestWithModuleCode_IllegalModuleCode(t *testing.T) {
	f := &CaseFactory{}
	err := WithModuleCode(-1)(f)
	assert.NotNil(t, err)
	assert.Equal(t, "illegal argument: moduleCode < 0", err.Error())
}

func TestWithModuleCode_LegalModuleCode(t *testing.T) {
	f := &CaseFactory{}
	err := WithModuleCode(1)(f)
	assert.Nil(t, err)
	assert.Equal(t, 1, f.moduleCode)
}

func TestNewFactory_NilCodingStrategy(t *testing.T) {
	_, err := NewFactory(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "illegal argument: codingStrategy is nil", err.Error())
}

func TestNewFactory_IllegalArgument(t *testing.T) {
	f, err := NewFactory(csWith1DigitAppCodeAndModuleCode, WithAppCode(-1), WithModuleCode(1))
	assert.Nil(t, f)
	assert.Equal(t, "illegal argument: appCode < 0", err.Error())
}

func TestNewFactory_AppCodeRangeNotIncludeAppCode(t *testing.T) {
	f, err := NewFactory(csWith1DigitAppCodeAndModuleCode, WithAppCode(10), WithModuleCode(1))
	assert.Nil(t, f)
	assert.Equal(t, "illegal argument: appCodeRange [0, 9] not include appCode 10", err.Error())
}

func TestNewFactory_ModuleCodeRangeNotIncludeModuleCode(t *testing.T) {
	f, err := NewFactory(csWith1DigitAppCodeAndModuleCode, WithAppCode(1), WithModuleCode(10))
	assert.Nil(t, f)
	assert.Equal(t, "illegal argument: moduleCodeRange [0, 9] not include moduleCode 10", err.Error())
}

func TestNewFactory_OK(t *testing.T) {
	f, err := NewFactory(csWith1DigitAppCodeAndModuleCode, WithAppCode(1), WithModuleCode(1))
	assert.Nil(t, err)
	assert.NotNil(t, f)
}

func TestCaseFactory_Create_CodeMapperNotDefineMappingForGivenStatusCode(t *testing.T) {
	f, _ := NewFactory(csWith1DigitAppCodeAndModuleCode, WithAppCode(1), WithModuleCode(1))
	_, err := f.create(domainerr.CodeUnknown, 551) // 100 is start of case code segment for CodeUnknown
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: statusCodeMapper doesn't define a CaseCodeSegment for status code %s",
		domainerr.CodeUnknown.String()), err.Error())
}

func TestCaseFactory_Create_InvalidArgument(t *testing.T) {
	f, _ := NewFactory(csWithoutAppCodeAndModuleCode)

	_, err := f.NewInvalidArgument(0)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [1, 50] for status code %s doesn't include code 0",
		domainerr.CodeInvalidArgument.String()), err.Error())

	_, err = f.NewInvalidArgument(51)
	assert.Equal(t, fmt.Sprintf("illegal argument: CaseCodeSegment [1, 50] for status code %s doesn't include code 51",
		domainerr.CodeInvalidArgument.String()), err.Error())

	case1, err := f.NewInvalidArgument(1)
	assert.Nil(t, err)
	assert.Equal(t, "001", case1.Identifier())

	case25, err := f.NewInvalidArgument(25)
	assert.Nil(t, err)
	assert.Equal(t, "025", case25.Identifier())

	case50, err := f.NewInvalidArgument(50)
	assert.Nil(t, err)
	assert.Equal(t, "050", case50.Identifier())
}

func TestCaseFactory_Create_DeadlineExceeded(t *testing.T) {
	f, _ := NewFactory(csWithoutAppCodeAndModuleCode)

	_, err := f.NewDeadlineExceeded(50)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [51, 100] for status code %s doesn't include code 50",
		domainerr.CodeDeadlineExceeded.String()), err.Error())

	_, err = f.NewDeadlineExceeded(101)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [51, 100] for status code %s doesn't include code 101",
		domainerr.CodeDeadlineExceeded.String()), err.Error())

	case51, err := f.NewDeadlineExceeded(51)
	assert.Nil(t, err)
	assert.Equal(t, "051", case51.Identifier())

	case75, err := f.NewDeadlineExceeded(75)
	assert.Nil(t, err)
	assert.Equal(t, "075", case75.Identifier())

	case100, err := f.NewDeadlineExceeded(100)
	assert.Nil(t, err)
	assert.Equal(t, "100", case100.Identifier())
}

func TestCaseFactory_Create_NotFound(t *testing.T) {
	f, _ := NewFactory(csWithoutAppCodeAndModuleCode)

	_, err := f.NewNotFound(100)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [101, 150] for status code %s doesn't include code 100",
		domainerr.CodeNotFound.String()), err.Error())

	_, err = f.NewNotFound(151)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [101, 150] for status code %s doesn't include code 151",
		domainerr.CodeNotFound.String()), err.Error())

	case101, err := f.NewNotFound(101)
	assert.Nil(t, err)
	assert.Equal(t, "101", case101.Identifier())

	case125, err := f.NewNotFound(125)
	assert.Nil(t, err)
	assert.Equal(t, "125", case125.Identifier())

	case150, err := f.NewNotFound(150)
	assert.Nil(t, err)
	assert.Equal(t, "150", case150.Identifier())
}

func TestCaseFactory_NewAlreadyExists(t *testing.T) {
	f, _ := NewFactory(csWithoutAppCodeAndModuleCode)

	_, err := f.NewAlreadyExists(150)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [151, 200] for status code %s doesn't include code 150",
		domainerr.CodeAlreadyExists.String()), err.Error())

	_, err = f.NewAlreadyExists(201)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [151, 200] for status code %s doesn't include code 201",
		domainerr.CodeAlreadyExists.String()), err.Error())

	case151, err := f.NewAlreadyExists(151)
	assert.Nil(t, err)
	assert.Equal(t, "151", case151.Identifier())

	case175, err := f.NewAlreadyExists(175)
	assert.Nil(t, err)
	assert.Equal(t, "175", case175.Identifier())

	case200, err := f.NewAlreadyExists(200)
	assert.Nil(t, err)
	assert.Equal(t, "200", case200.Identifier())
}

func TestCaseFactory_NewPermissionDenied(t *testing.T) {
	f, _ := NewFactory(csWithoutAppCodeAndModuleCode)

	_, err := f.NewPermissionDenied(200)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [201, 250] for status code %s doesn't include code 200",
		domainerr.CodePermissionDenied.String()), err.Error())

	_, err = f.NewPermissionDenied(251)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [201, 250] for status code %s doesn't include code 251",
		domainerr.CodePermissionDenied.String()), err.Error())

	case201, err := f.NewPermissionDenied(201)
	assert.Nil(t, err)
	assert.Equal(t, "201", case201.Identifier())

	case225, err := f.NewPermissionDenied(225)
	assert.Nil(t, err)
	assert.Equal(t, "225", case225.Identifier())

	case250, err := f.NewPermissionDenied(250)
	assert.Nil(t, err)
	assert.Equal(t, "250", case250.Identifier())
}

func TestCaseFactory_NewResourceExhausted(t *testing.T) {
	f, _ := NewFactory(csWithoutAppCodeAndModuleCode)

	_, err := f.NewResourceExhausted(250)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [251, 300] for status code %s doesn't include code 250",
		domainerr.CodeResourceExhausted.String()), err.Error())

	_, err = f.NewResourceExhausted(301)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [251, 300] for status code %s doesn't include code 301",
		domainerr.CodeResourceExhausted.String()), err.Error())

	case251, err := f.NewResourceExhausted(251)
	assert.Nil(t, err)
	assert.Equal(t, "251", case251.Identifier())

	case275, err := f.NewResourceExhausted(275)
	assert.Nil(t, err)
	assert.Equal(t, "275", case275.Identifier())

	case300, err := f.NewResourceExhausted(300)
	assert.Nil(t, err)
	assert.Equal(t, "300", case300.Identifier())
}

func TestCaseFactory_NewFailedPrecondition(t *testing.T) {
	f, _ := NewFactory(csWithoutAppCodeAndModuleCode)

	_, err := f.NewFailedPrecondition(300)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [301, 350] for status code %s doesn't include code 300",
		domainerr.CodeFailedPrecondition.String()), err.Error())

	_, err = f.NewFailedPrecondition(351)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [301, 350] for status code %s doesn't include code 351",
		domainerr.CodeFailedPrecondition.String()), err.Error())

	case301, err := f.NewFailedPrecondition(301)
	assert.Nil(t, err)
	assert.Equal(t, "301", case301.Identifier())

	case325, err := f.NewFailedPrecondition(325)
	assert.Nil(t, err)
	assert.Equal(t, "325", case325.Identifier())

	case350, err := f.NewFailedPrecondition(350)
	assert.Nil(t, err)
	assert.Equal(t, "350", case350.Identifier())
}

func TestCaseFactory_NewAborted(t *testing.T) {
	f, _ := NewFactory(csWithoutAppCodeAndModuleCode)

	_, err := f.NewAborted(350)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [351, 400] for status code %s doesn't include code 350",
		domainerr.CodeAborted.String()), err.Error())

	_, err = f.NewAborted(401)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [351, 400] for status code %s doesn't include code 401",
		domainerr.CodeAborted.String()), err.Error())

	case351, err := f.NewAborted(351)
	assert.Nil(t, err)
	assert.Equal(t, "351", case351.Identifier())

	case375, err := f.NewAborted(375)
	assert.Nil(t, err)
	assert.Equal(t, "375", case375.Identifier())

	case400, err := f.NewAborted(400)
	assert.Nil(t, err)
	assert.Equal(t, "400", case400.Identifier())
}

func TestCaseFactory_NewOutOfRange(t *testing.T) {
	f, _ := NewFactory(csWithoutAppCodeAndModuleCode)

	_, err := f.NewOutOfRange(400)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [401, 450] for status code %s doesn't include code 400",
		domainerr.CodeOutOfRange.String()), err.Error())

	_, err = f.NewOutOfRange(451)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [401, 450] for status code %s doesn't include code 451",
		domainerr.CodeOutOfRange.String()), err.Error())

	case401, err := f.NewOutOfRange(401)
	assert.Nil(t, err)
	assert.Equal(t, "401", case401.Identifier())

	case425, err := f.NewOutOfRange(425)
	assert.Nil(t, err)
	assert.Equal(t, "425", case425.Identifier())

	case450, err := f.NewOutOfRange(450)
	assert.Nil(t, err)
	assert.Equal(t, "450", case450.Identifier())
}

func TestCaseFactory_NewInternalError(t *testing.T) {
	f, _ := NewFactory(csWithoutAppCodeAndModuleCode)

	_, err := f.NewInternalError(450)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [451, 500] for status code %s doesn't include code 450",
		domainerr.CodeInternalError.String()), err.Error())

	_, err = f.NewInternalError(501)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [451, 500] for status code %s doesn't include code 501",
		domainerr.CodeInternalError.String()), err.Error())

	case451, err := f.NewInternalError(451)
	assert.Nil(t, err)
	assert.Equal(t, "451", case451.Identifier())

	case475, err := f.NewInternalError(475)
	assert.Nil(t, err)
	assert.Equal(t, "475", case475.Identifier())

	case500, err := f.NewInternalError(500)
	assert.Nil(t, err)
	assert.Equal(t, "500", case500.Identifier())
}

func TestCaseFactory_NewDataLoss(t *testing.T) {
	f, _ := NewFactory(csWithoutAppCodeAndModuleCode)

	_, err := f.NewDataLoss(500)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [501, 550] for status code %s doesn't include code 500",
		domainerr.CodeDataLoss.String()), err.Error())

	_, err = f.NewDataLoss(551)
	assert.Equal(t, fmt.Sprintf(
		"illegal argument: CaseCodeSegment [501, 550] for status code %s doesn't include code 551",
		domainerr.CodeDataLoss.String()), err.Error())

	case501, err := f.NewDataLoss(501)
	assert.Nil(t, err)
	assert.Equal(t, "501", case501.Identifier())

	case525, err := f.NewDataLoss(525)
	assert.Nil(t, err)
	assert.Equal(t, "525", case525.Identifier())

	case550, err := f.NewDataLoss(550)
	assert.Nil(t, err)
	assert.Equal(t, "550", case550.Identifier())
}

func TestCaseFactory_Create(t *testing.T) {
	strategy, _ := NewCodingStrategyBuilder().
		NumDigitsOfAppCode(2).
		NumDigitsOfModuleCode(2).
		NumDigitsOfCaseCode(3).
		StatusCodeMapper(NewCodeMapper(&DefaultCodeMapper{})).Build()
	f, _ := NewFactory(strategy, WithAppCode(1), WithModuleCode(1))
	c, err := f.create(domainerr.CodeInvalidArgument, 1)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, 1, c.appCode)
	assert.Equal(t, 1, c.moduleCode)
	assert.Equal(t, 1, c.caseCode)
	assert.Equal(t, domainerr.CodeInvalidArgument, c.statusCode)
	assert.Equal(t, "01_01_001", c.Identifier())
}

func TestCaseFactory_PadLeftZeros(t *testing.T) {
	f := &CaseFactory{}
	assert.Equal(t, "001", f.padLeftZeros(1, 3))
	assert.Equal(t, "010", f.padLeftZeros(10, 3))
	assert.Equal(t, "100", f.padLeftZeros(100, 3))
	assert.Equal(t, "1000", f.padLeftZeros(1000, 3))
}

// Tests for CaseFactory end

// Tests for CodingStrategyBuilder start

func TestNewCodingStrategyBuilder(t *testing.T) {
	b := NewCodingStrategyBuilder()
	assert.NotNil(t, b)
}

func TestCodingStrategyBuilder_NumDigitsOfAppCode(t *testing.T) {
	b := NewCodingStrategyBuilder()
	b.NumDigitsOfAppCode(1)
	assert.Equal(t, 1, b.s.numDigitsOfAppCode)
}

func TestCodingStrategyBuilder_NumDigitsOfModuleCode(t *testing.T) {
	b := NewCodingStrategyBuilder()
	b.NumDigitsOfModuleCode(1)
	assert.Equal(t, 1, b.s.numDigitsOfModuleCode)
}

func TestCodingStrategyBuilder_NumDigitsOfCaseCode(t *testing.T) {
	b := NewCodingStrategyBuilder()
	b.NumDigitsOfCaseCode(3)
	assert.Equal(t, 3, b.s.numDigitsOfCaseCode)
}

func TestCodingStrategyBuilder_StatusCodeMapper(t *testing.T) {
	b := NewCodingStrategyBuilder()
	b.StatusCodeMapper(NewCodeMapper(&DefaultCodeMapper{}))
	assert.NotNil(t, b.s.statusCodeMapper)
}

func TestCodingStrategyBuilder_Build_IllegalNumDigitsOfAppCode(t *testing.T) {
	b := NewCodingStrategyBuilder()
	b.NumDigitsOfAppCode(-1)
	b.NumDigitsOfModuleCode(0)
	b.NumDigitsOfCaseCode(0)
	b.StatusCodeMapper(NewCodeMapper(&DefaultCodeMapper{}))
	_, err := b.Build()
	assert.Equal(t, "illegal argument: numDigitsOfAppCode < 0", err.Error())
}

func TestCodingStrategyBuilder_Build_IllegalNumDigitsOfModuleCode(t *testing.T) {
	b := NewCodingStrategyBuilder()
	b.NumDigitsOfAppCode(0)
	b.NumDigitsOfModuleCode(-1)
	b.NumDigitsOfCaseCode(0)
	b.StatusCodeMapper(NewCodeMapper(&DefaultCodeMapper{}))
	_, err := b.Build()
	assert.Equal(t, "illegal argument: numDigitsOfModuleCode < 0", err.Error())
}

func TestCodingStrategyBuilder_Build_IllegalNumDigitsOfCaseCode(t *testing.T) {
	b := NewCodingStrategyBuilder()
	b.NumDigitsOfAppCode(0)
	b.NumDigitsOfModuleCode(0)
	b.NumDigitsOfCaseCode(-1)
	b.StatusCodeMapper(NewCodeMapper(&DefaultCodeMapper{}))
	_, err := b.Build()
	assert.Equal(t, "illegal argument: numDigitsOfCaseCode < 0", err.Error())
}

func TestCodingStrategyBuilder_Build_NilStatusCodeMapper(t *testing.T) {
	b := NewCodingStrategyBuilder()
	b.NumDigitsOfAppCode(1)
	b.NumDigitsOfModuleCode(1)
	b.NumDigitsOfCaseCode(3)
	b.StatusCodeMapper(nil)
	_, err := b.Build()
	assert.Equal(t, "illegal argument: statusCodeMapper is nil", err.Error())
}

func TestCodingStrategyBuilder_Build(t *testing.T) {
	b := NewCodingStrategyBuilder()
	b.NumDigitsOfAppCode(1)
	b.NumDigitsOfModuleCode(1)
	b.NumDigitsOfCaseCode(3)
	b.StatusCodeMapper(NewCodeMapper(&DefaultCodeMapper{}))
	s, err := b.Build()
	assert.Nil(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, 1, s.numDigitsOfAppCode)
	assert.Equal(t, 1, s.numDigitsOfModuleCode)
	assert.Equal(t, 3, s.numDigitsOfCaseCode)
	appCodeR, _ := NewNumRange(0, 9)
	assert.Equal(t, s.appCodeRange, appCodeR)
	moduleCodeR, _ := NewNumRange(0, 9)
	assert.Equal(t, s.moduleCodeRange, moduleCodeR)
	caseCodeR, _ := NewNumRange(0, 999)
	assert.Equal(t, s.caseCodeRange, caseCodeR)
	assert.Equal(t, s.statusCodeMapper, NewCodeMapper(&DefaultCodeMapper{}))
}

func TestCodingStrategyBuilder_Build_CaseCodeRangeNotIncludeCaseCodeSegmentDefinedByCodeMapper(t *testing.T) {
	b := NewCodingStrategyBuilder()
	b.NumDigitsOfAppCode(1)
	b.NumDigitsOfModuleCode(1)
	b.NumDigitsOfCaseCode(2)
	b.StatusCodeMapper(NewCodeMapper(&DefaultCodeMapper{}))
	_, err := b.Build()
	assert.Equal(t,
		"illegal argument: case code range [0, 99] of CodingStrategy doesn't include caseCodeSegment [51, 100] defined by CodeMapper",
		err.Error())
}

// Tests for CodingStrategyBuilder end
