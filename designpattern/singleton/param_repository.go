package singleton

type ParamRepository interface {
	GetParam(id int) Param
}

type Param struct {
	Id int
}

type RepositoryImpl struct {
	params []Param
}
