package handler

import "github.com/gofiber/fiber/v2"

func BusApi(api fiber.Router) {
	api.Get("/bus/type", busTypeHandler)
	api.Get("/bus/name", busNameHandler)
}

func busTypeHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("bus age 1")
}

func busNameHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("bus name is bwm")
}
