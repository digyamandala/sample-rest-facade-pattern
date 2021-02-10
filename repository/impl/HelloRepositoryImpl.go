package repository_impl

import "github.com/personal/sample-rest-facade-pattern/entity"

type HelloRepositoryImpl struct {
}

func (repositoryImpl *HelloRepositoryImpl) FindByName(name string) entity.Hello {
	return entity.Hello{
		Name:   name,
		Status: "ACTIVE",
	}
}
