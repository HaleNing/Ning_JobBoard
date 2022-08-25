package delegate

import (
	"context"
	"github.com/HaleNing/bustrack/src/Model/ent"
	"github.com/HaleNing/bustrack/src/database"
	"github.com/HaleNing/bustrack/src/param"
)

func UpBusName(ctx context.Context, upParam *param.BusUpdateParam) (*ent.Bus, error) {
	save, err := database.DBConn.Bus.UpdateOneID(upParam.Id).SetBusName(upParam.BusName).Save(ctx)
	return save, err
}
