package abstract_factory

type tensorFlowModel struct {
}

type tensorFlowPredictor struct {
}

type tensorFlowConverter struct {
}

func (model *tensorFlowModel) Name() string {
	return "tensorFlowModel"
}

func (predictor *tensorFlowPredictor) Predict() float32 {
	return 1.0
}

func (converter *tensorFlowConverter) Convert(vo ModelVO) MLModel {
	return &tensorFlowModel{}
}
