package pg

import pg "github.com/go-pg/pg/v10"

func Connection(addr, username, password, dbName string) *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     addr,
		User:     username,
		Password: password,
		Database: dbName,
	})
}
