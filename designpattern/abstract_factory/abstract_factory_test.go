package abstract_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFactory(t *testing.T) {
	factory := GetFactory(TENSOR_FLOW)
	_, ok := factory.(*tensorFlowFactory)
	assert.True(t, ok)

	factory = GetFactory(SPARK_ML)
	_, ok = factory.(*sparkMLFactory)
	assert.True(t, ok)
}

func TestCreateModule(t *testing.T) {
	factory := GetFactory(TENSOR_FLOW)
	predictor := factory.CreatePredictor()
	converter := factory.CreateModelConverter()

	_, ok := predictor.(*tensorFlowPredictor)
	assert.True(t, ok)
	_, ok = converter.(*tensorFlowConverter)
	assert.True(t, ok)

	factory = GetFactory(SPARK_ML)
	predictor = factory.CreatePredictor()
	converter = factory.CreateModelConverter()

	_, ok = predictor.(*sparkMLPredictor)
	assert.True(t, ok)
	_, ok = converter.(*sparkMLConverter)
	assert.True(t, ok)
}
