package utils

import (
	"context"
	"github.com/HaleNing/Ning_JobBoard/src/database"
	"github.com/gofiber/fiber/v2"
)

// CheckLogin check user token is still working?
func CheckLogin(fiberCtx *fiber.Ctx) bool {
	userToken := fiberCtx.Cookies("token")
	result, _ := database.RDB.Exists(context.Background(), userToken).Result()
	if result > 0 {
		return true
	}
	return false
}
