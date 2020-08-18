package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyParam(t *testing.T) {
	parameter := (&ParameterBuilder{}).Build()
	assert.Equal(t, 0, parameter.id)
	assert.Equal(t, float64(0), parameter.cpi)
	assert.Equal(t, float64(0), parameter.cpc)
	//string can not be nil.
	assert.Equal(t, "", parameter.tag)

	parameter = Parameter{}
	assert.Equal(t, 0, parameter.id)
	assert.Equal(t, float64(0), parameter.cpi)
	assert.Equal(t, float64(0), parameter.cpc)
	assert.Equal(t, "", parameter.tag)
}

func TestBuild(t *testing.T) {
	builder := ParameterBuilder{}
	parameter := builder.SetId(1).SetCpc(0.1).SetCpi(0.2).SetTag("test").Build()
	assert.Equal(t, 1, parameter.id)
	assert.Equal(t, 0.1, parameter.cpc)
	assert.Equal(t, 0.2, parameter.cpi)
	assert.Equal(t, "test", parameter.tag)

	parameter = Parameter{
		id:  1,
		cpc: 0.1,
		cpi: 0.2,
		tag: "test"}
	assert.Equal(t, 1, parameter.id)
	assert.Equal(t, 0.1, parameter.cpc)
	assert.Equal(t, 0.2, parameter.cpi)
	assert.Equal(t, "test", parameter.tag)
}
