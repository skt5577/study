package abstract_factory

type sparkMLModel struct {
}

type sparkMLPredictor struct {
}

type sparkMLConverter struct {
}

func (model *sparkMLModel) Name() string {
	return "sparkMLModel"
}

func (predictor *sparkMLPredictor) Predict() float32 {
	return 2.0
}

func (converter *sparkMLConverter) Convert(vo ModelVO) MLModel {
	return &sparkMLModel{}
}
