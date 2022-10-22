package database

import (
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent"
	"log"
	"os"
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
	datasourceName := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") +
		"@" + os.Getenv("DB_HOST") + ":5432/" + os.Getenv("DB_NAME")
	client, err := ent.Open("postgres", datasourceName)
	if err != nil {
		log.Fatalf("failed opening connection to postgressql database: %v", err)
	}
	DBConn = client
	return client, nil
}
