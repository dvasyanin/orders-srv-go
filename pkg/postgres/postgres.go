package postgres

import "github.com/go-pg/pg/v10"

type Postgres struct{}

func NewPostgres(addr string, database string, user string, password string) *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     addr,
		Database: database,
		User:     user,
		Password: password,
	})
}
