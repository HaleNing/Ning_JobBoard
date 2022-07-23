package database

import (
	"github.com/HaleNing/bustrack/src/Model/ent"
	"log"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

var DBConn *ent.Client

func NewConnection() (*ent.Client, error) {
	client, err := ent.Open("postgres", "postgres://ning:@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	DBConn = client
	return client, nil
}
