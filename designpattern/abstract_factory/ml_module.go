package abstract_factory

type Predictor interface {
	Predict() float32
}

type ModelConverter interface {
	Convert(vo ModelVO) MLModel
}

type ModelVO struct {
	Name string
}

type MLModel interface {
	Name() string
}
