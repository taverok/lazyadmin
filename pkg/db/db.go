package db

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type DataSource struct {
	Driver  string `yaml:"driver"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	User    string `yaml:"user"`
	Pass    string `yaml:"pass"`
	Name    string `yaml:"name"`
	SSLMode string `yaml:"SSLMode"`
}

func (it DataSource) dsn() string {
	switch it.Driver {
	case "mysql":
		return fmt.Sprintf("%s:%s@(%s:%d)/%s", it.User, it.Pass, it.Host, it.Port, it.Name)
	}

	return fmt.Sprintf("%s:%s@(%s:%d)/%s", it.User, it.Pass, it.Host, it.Port, it.Name)
}

func NewDb(ds DataSource) (*sqlx.DB, error) {
	db, err := sqlx.Connect(ds.Driver, ds.dsn())
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
