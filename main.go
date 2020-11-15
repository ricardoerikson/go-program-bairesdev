package main

import (
	"log"
	"net/http"

	"questionsandanswers.com/pkg/question/endpoint"
	"questionsandanswers.com/pkg/question/persistence/mem"
	"questionsandanswers.com/pkg/question/service"
	transport "questionsandanswers.com/pkg/question/transport/http"
)

func main() {
	repository := new(mem.QuestionRepositoryImpl)
	repository.Init()
	service := service.QuestionServiceImpl{Repository: repository}

	endpoints := endpoint.MakeEndpoints(service)
	handler := transport.ConfigureRoutes(endpoints)
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err.Error())
	}
}
