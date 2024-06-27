package domainerr

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatusToObjStyleStr(t *testing.T) {
	s := StatusFailedPrecondition
	assert.Equal(t, fmt.Sprintf("%s{code:9,message:\"\"}", s.code.Name()), s.ToObjStyleStr())

	s2 := StatusFailedPrecondition.WithMessage("err msg")
	assert.Equal(t, fmt.Sprintf("%s{code:9,message:\"err msg\"}", s2.code.Name()), s2.ToObjStyleStr())

	s3 := StatusFailedPrecondition.WithMessage("err msg").WithDetails(map[string]interface{}{"k": "v"})
	assert.Equal(t, fmt.Sprintf("%s{code:9,message:\"err msg\",details:map[k:v]}", s3.code.Name()), s3.ToObjStyleStr())

	s4 := StatusFailedPrecondition.WithMessage("err msg").WithDetails(map[string]interface{}{"k": "v"}).WithCase(
		&case4Test{moduleCode: 1, caseCode: 1})
	assert.Equal(t, fmt.Sprintf("%s{code:9,specificCase:\"1_1\",message:\"err msg\",details:map[k:v]}", s4.code.Name()),
		s4.ToObjStyleStr())
}

type case4Test struct {
	moduleCode int
	caseCode   int
}

func (c *case4Test) Identifier() string {
	return fmt.Sprintf("%d_%d", c.moduleCode, c.caseCode)
}

func (c *case4Test) StatusCode() Code {
	return CodeFailedPrecondition
}
