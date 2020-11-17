package main

import (
	"log"
	"net/http"
	"os"

	"questionsandanswers.com/pkg/question/endpoint"
	"questionsandanswers.com/pkg/question/service/pg"
	transport "questionsandanswers.com/pkg/question/transport/http"
)

func main() {
	addr := os.Getenv("DB_ADDR")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db := pg.Connection(addr, user, password, dbName)
	defer db.Close()

	service := pg.NewService(db)
	endpoints := endpoint.NewEndpoints(service)
	handler := transport.NewHTTPTransport(endpoints)
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err.Error())
	}
}
