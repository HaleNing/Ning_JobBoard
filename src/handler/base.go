package handler

import "github.com/gofiber/fiber/v2"

func BaseAPi(api *fiber.App) {
	api.Get("/login", loginHandler)
}

func loginHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("you are accepted")
}
