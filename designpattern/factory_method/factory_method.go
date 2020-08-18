package factory_method

import (
	"study/designpattern/common"
)

type ModelPredictor interface {
	Predict() float64
}

type ModelPredictorFactory interface {
	CreatePredictor(buildType common.ModelBuildType) ModelPredictor
}

type ModelPredictorFactoryImpl struct {
}

func (factory *ModelPredictorFactoryImpl) CreatePredictor(buildType common.ModelBuildType) ModelPredictor {
	if buildType == common.WORKAROUND {
		return &WorkaroundPredictor{}
	} else {
		return &TensorFlowPredictor{}
	}
}

type TensorFlowPredictor struct {
}

func (predictor *TensorFlowPredictor) Predict() float64 {
	return 0.1
}

type WorkaroundPredictor struct {
}

func (predictor *WorkaroundPredictor) Predict() float64 {
	return 0.2
}

func CreatePredictor(buildType common.ModelBuildType) ModelPredictor {
	if buildType == common.WORKAROUND {
		return &WorkaroundPredictor{}
	} else {
		return &TensorFlowPredictor{}
	}
}
