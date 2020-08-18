package factory_method

import (
	"github.com/stretchr/testify/assert"
	"study/designpattern/common"
	"testing"
)

func TestCreate(t *testing.T) {
	factory := &ModelPredictorFactoryImpl{}
	predictor := factory.CreatePredictor(common.TF_ESTIMATOR)
	_, ok := predictor.(*TensorFlowPredictor)
	assert.True(t, ok)

	predictor = factory.CreatePredictor(common.TF_KERAS)
	_, ok = predictor.(*TensorFlowPredictor)
	assert.True(t, ok)

	predictor = factory.CreatePredictor(common.WORKAROUND)
	_, ok = predictor.(*WorkaroundPredictor)
	assert.True(t, ok)

	predictor = factory.CreatePredictor(common.GetModelBuildType(2))
	_, ok = predictor.(*WorkaroundPredictor)
	assert.True(t, ok)

	predictor = factory.CreatePredictor(2)
	_, ok = predictor.(*WorkaroundPredictor)
	assert.True(t, ok)
}

func TestCreateMethod(t *testing.T) {
	predictor := CreatePredictor(common.TF_ESTIMATOR)
	_, ok := predictor.(*TensorFlowPredictor)
	assert.True(t, ok)

	predictor = CreatePredictor(common.TF_KERAS)
	_, ok = predictor.(*TensorFlowPredictor)
	assert.True(t, ok)

	predictor = CreatePredictor(common.WORKAROUND)
	_, ok = predictor.(*WorkaroundPredictor)
	assert.True(t, ok)

	predictor = CreatePredictor(common.GetModelBuildType(2))
	_, ok = predictor.(*WorkaroundPredictor)
	assert.True(t, ok)

	predictor = CreatePredictor(2)
	_, ok = predictor.(*WorkaroundPredictor)
	assert.True(t, ok)
}
