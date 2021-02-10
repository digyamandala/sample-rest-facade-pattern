package service_impl

import (
	repository_impl "github.com/personal/sample-rest-facade-pattern/repository/impl"
	"github.com/personal/sample-rest-facade-pattern/web/model/response"
)

type HelloServiceImpl struct {
	Repository *repository_impl.HelloRepositoryImpl
}

func (helloServiceImpl *HelloServiceImpl) Greet(name string) response.HelloWebResponse {
	entity := helloServiceImpl.Repository.FindByName(name)

	return response.HelloWebResponse{
		Name:   entity.Name,
		Status: entity.Status,
	}
}
