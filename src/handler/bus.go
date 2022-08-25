package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/HaleNing/bustrack/src/Model/ent/bus"
	"github.com/HaleNing/bustrack/src/database"
	"github.com/HaleNing/bustrack/src/delegate"
	"github.com/HaleNing/bustrack/src/param"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

var (
	busCtx context.Context
)

func BusApi(api fiber.Router, ctx context.Context) {
	busCtx = ctx
	api.Post("/bus/create", createBusHandler)
	api.Get("/bus/getById", getBusByNameHandler)
	api.Post("/bus/upById", updateBusHandler)

}

func createBusHandler(ctx *fiber.Ctx) error {
	bus := new(param.BusCreateParam)
	if err := ctx.BodyParser(bus); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "param is nil")
	}
	log.Println(bus) // john
	_, err := database.DBConn.Bus.Create().SetBusName(bus.BusName).SetUserID(0).Save(busCtx)
	if err != nil {
		log.Printf("err:[%v]", err)
		return fiber.NewError(fiber.StatusBadRequest, "create bus error!")
	} else {
		return ctx.SendString("create bus success")
	}
}

func updateBusHandler(ctx *fiber.Ctx) error {
	param := new(param.BusUpdateParam)
	if err := ctx.BodyParser(param); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "param is nil")
	}
	res, err := delegate.UpBusName(busCtx, param)
	log.Println(res)
	log.Println(err)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "update have error")
	} else {
		return ctx.SendString("update bus success")
	}
}

func getBusByNameHandler(ctx *fiber.Ctx) error {
	bn := ctx.Query("id")
	id, _ := strconv.ParseInt(bn, 10, 64)
	all, err := database.DBConn.Bus.Query().Where(bus.IDEQ(id)).All(busCtx)
	if err != nil {
		_ = fmt.Errorf("getBusByNameHandler have error :[%v]", err)
	}
	log.Println(all)
	res, err := json.Marshal(all)
	log.Println(res)
	return ctx.Send(res)
}
