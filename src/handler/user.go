package handler

import (
	"github.com/gofiber/fiber/v2"
)

func UserApi(api fiber.Router) {
	api.Get("/bus/name", nameHandler)
	api.Get("/driver/age", ageHandler)
}

func ageHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("driver age is I'm 23")
}

func nameHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("driver name is  liuxining")

}
