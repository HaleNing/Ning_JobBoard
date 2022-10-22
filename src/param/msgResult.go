package param

import "github.com/gofiber/fiber/v2"

type Result struct {
	Code int64  `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func CreateSuccessRes(data any) Result {
	return Result{
		Code: fiber.StatusOK,
		Data: data,
		Msg:  "success",
	}
}
