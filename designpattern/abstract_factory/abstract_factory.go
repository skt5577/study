package abstract_factory

const (
	TENSOR_FLOW MLType = 1
	SPARK_ML    MLType = 2
)

type MLType int

type ModuleFactory interface {
	CreateModelConverter() ModelConverter
	CreatePredictor() Predictor
}

type tensorFlowFactory struct {
}

func (factory *tensorFlowFactory) CreateModelConverter() ModelConverter {
	return &tensorFlowConverter{}
}
func (factory *tensorFlowFactory) CreatePredictor() Predictor {
	return &tensorFlowPredictor{}
}

type sparkMLFactory struct {
}

func (factory *sparkMLFactory) CreateModelConverter() ModelConverter {
	return &sparkMLConverter{}
}
func (factory *sparkMLFactory) CreatePredictor() Predictor {
	return &sparkMLPredictor{}
}

func GetFactory(mlType MLType) ModuleFactory {
	switch mlType {
	case TENSOR_FLOW:
		return &tensorFlowFactory{}
	case SPARK_ML:
		return &sparkMLFactory{}
	}
	return nil
}
