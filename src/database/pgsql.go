package database

import (
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent"
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
	client, err := ent.Open("postgres", "postgres://user:passwd@host:5432/ning_jobboard")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	DBConn = client
	return client, nil
}
