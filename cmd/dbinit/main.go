package main

import (
	"os"

	"github.com/go-kit/kit/log"
	gopg "github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"questionsandanswers.com/pkg/question/entity"
	"questionsandanswers.com/pkg/question/persistence/pg"
)

func createSchema(db *gopg.DB) error {
	models := []interface{}{
		(*entity.Question)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	addr := os.Getenv("DB_ADDR")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	logger := log.NewLogfmtLogger(os.Stderr)

	db := pg.Connection(addr, user, password, dbName)
	defer db.Close()
	err := createSchema(db)
	if err != nil {
		logger.Log(err)
	} else {
		logger.Log("Schema created!")
	}
}
