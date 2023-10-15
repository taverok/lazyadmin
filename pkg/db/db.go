package database

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type DataSource struct {
	Host    string
	Port    int
	User    string
	Pass    string
	Name    string
	SSLMode string
}

func NewMysqlDb(ds DataSource) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s", ds.User, ds.Pass, ds.Host, ds.Port, ds.Name)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(10 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
