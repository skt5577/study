package singleton

import "sync"

var repository ParamRepository
var once sync.Once

func GetInstanceLazy() ParamRepository {
	once.Do(func() {
		repository = &RepositoryImpl{}
	})

	return repository
	//return &RepositoryImpl{}
}
