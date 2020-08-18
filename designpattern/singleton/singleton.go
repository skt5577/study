package singleton

var instance ParamRepository = &RepositoryImpl{}

type ParamRepository interface {
	GetParam(id int) Param
}

type Param struct {
	Id int
}

type RepositoryImpl struct {
	params []Param
}

func GetInstance() ParamRepository {
	return instance
	//return &RepositoryImpl{}
}

func (repository *RepositoryImpl) GetParam(id int) Param {
	return Param{Id: id}
}
