package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"questionsandanswers.com/pkg/question/endpoint"
	"questionsandanswers.com/pkg/question/persistence/pg"
	"questionsandanswers.com/pkg/question/service"
	transport "questionsandanswers.com/pkg/question/transport/http"
)

func main() {
	addr := os.Getenv("DB_ADDR")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	logger := log.NewLogfmtLogger(os.Stderr)

	db := pg.Connection(addr, user, password, dbName)
	defer db.Close()

	repository := new(pg.QuestionRepositoryPgImpl).NewRepository(db)
	service := new(service.QuestionServiceImpl).NewService(repository)
	// Wire logging middleware
	service = transport.LoggingMiddleware{logger, service}

	endpoints := endpoint.NewEndpoints(service)
	handler := transport.NewHTTPTransport(endpoints)
	handler = transport.ContentTypeMiddleware(handler)
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		logger.Log(err.Error())
	}
}
