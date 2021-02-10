package controller

import (
	"github.com/personal/sample-rest-facade-pattern/service/interface"
	"github.com/personal/sample-rest-facade-pattern/web/model/response"
)

type HomeController struct {
	HelloService service_interface.HelloService
}

func (controller *HomeController) Hello(name string) response.HelloWebResponse {
	return controller.HelloService.Greet(name)
}
