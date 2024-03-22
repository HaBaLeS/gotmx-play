package database

import (
	"fmt"
	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	PostgresUser = "postgres"
	PostgresPass = "postgres"
	PostgresHost = "localhost"
	PostgresPort = 5432
	PostgresDB   = "todoapp"
)

type Db struct {
	DB *sqlx.DB
}

func (d *Db) initConnection() {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", PostgresUser, PostgresPass, PostgresHost, PostgresPort, PostgresDB)
	connConfig, err := pgx.ParseConfig(url)
	if err != nil {
		panic(err)
	}
	dbc := stdlib.OpenDB(*connConfig)
	db := sqlx.NewDb(dbc, "pgx")

	d.DB = db
}

func NewDao() *Db {

	db := &Db{}
	db.initConnection()

	return db
}
