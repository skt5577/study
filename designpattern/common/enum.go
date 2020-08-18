package common

import "fmt"

type ModelBuildType int

const (
	TF_ESTIMATOR ModelBuildType = 1 + iota
	WORKAROUND
	TF_KERAS
)

func GetModelBuildType(id int) ModelBuildType {
	switch id {
	case int(TF_ESTIMATOR):
		return TF_ESTIMATOR
	case int(WORKAROUND):
		return WORKAROUND
	case int(TF_KERAS):
		return TF_KERAS
	default:
		return TF_ESTIMATOR
	}
}
func Test() {
	fmt.Println(GetModelBuildType(1))
}
