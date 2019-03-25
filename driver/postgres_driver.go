package driver

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	SQL *sql.DB
}

var Postgres = &PostgresDB{}

func Connect(host, port, user, password, dbname string) (*PostgresDB) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
							host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)	
	if err != nil {
		panic(err)
	}			

	Postgres.SQL = db
	return Postgres
}