package service_interface

import (
	"github.com/personal/sample-rest-facade-pattern/web/model/response"
)

type HelloService interface {
	Greet(name string) response.HelloWebResponse
}
