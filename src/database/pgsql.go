package database

import (
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent"
	"log"
)

var DBConn *ent.Client

func NewConnection() (*ent.Client, error) {
	//postgres://board_user:MBuvyvVH6t2Iap4uNArCmVgDsnEscoNy@dpg-cgjt1m64dad69r738be0-a/board

	//datasourceName := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") +
	//	"@" + os.Getenv("DB_HOST") + ":5432/" + os.Getenv("DB_NAME")

	datasourceName := "postgres://" + "board_user" + ":" + "MBuvyvVH6t2Iap4uNArCmVgDsnEscoNy" +
		"@" + "dpg-cgjt1m64dad69r738be0-a" + "/board"
	client, err := ent.Open("postgres", datasourceName)

	if err != nil {
		log.Fatalf("failed opening connection to postgressql database: %v", err)
	}
	DBConn = client
	return client, nil
}
