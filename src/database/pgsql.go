package database

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent"
	"time"
)

var DBConn *ent.Client

func NewConnection() (*ent.Client, error) {
	//postgres://board_user:MBuvyvVH6t2Iap4uNArCmVgDsnEscoNy@dpg-cgjt1m64dad69r738be0-a/board

	//datasourceName := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") +
	//	"@" + os.Getenv("DB_HOST") + ":5432/" + os.Getenv("DB_NAME")

	datasourceName := "postgres://" + "board_user" + ":" + "MBuvyvVH6t2Iap4uNArCmVgDsnEscoNy" +
		"@" + "dpg-cgjt1m64dad69r738be0-a" + "/board"
	drv, err := sql.Open(dialect.Postgres, datasourceName)

	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgressql database: %v", err)
	}

	db := drv.DB()
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Hour)
	DBConn = ent.NewClient(ent.Driver(drv))
	return DBConn, nil
}
