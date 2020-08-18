package singleton

var instance ParamRepository = &RepositoryImpl{}

func GetInstance() ParamRepository {
	return instance
	//return &RepositoryImpl{}
}

func (repository *RepositoryImpl) GetParam(id int) Param {
	return Param{Id: id}
}
