package main

import (
	"encoding/json"
	repositoryimpl "github.com/personal/sample-rest-facade-pattern/repository/impl"
	serviceimpl "github.com/personal/sample-rest-facade-pattern/service/impl"
	"github.com/personal/sample-rest-facade-pattern/web/controller"
	"io"
	"log"
	"net/http"
)

func main() {
	helloRepository := repositoryimpl.HelloRepositoryImpl{}

	helloService := serviceimpl.HelloServiceImpl{
		Repository: &helloRepository,
	}

	homeController := controller.HomeController{
		HelloService: &helloService,
	}

	h1 := func(w http.ResponseWriter, r *http.Request) {
		res, err := json.Marshal(homeController.Hello(r.URL.Query().Get("name")))
		if err != nil {
			panic(err)
		}
		_, _ = io.WriteString(w, string(res))
	}

	http.HandleFunc("/hello", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
