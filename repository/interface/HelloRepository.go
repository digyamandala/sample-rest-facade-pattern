package repository_interface

import "github.com/personal/sample-rest-facade-pattern/entity"

type HelloRepository interface {
	FindByName(name string) entity.Hello
}
