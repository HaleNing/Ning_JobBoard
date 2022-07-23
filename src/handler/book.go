package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/HaleNing/bustrack/src/Model/ent/book"
	"github.com/HaleNing/bustrack/src/database"
	"github.com/HaleNing/bustrack/src/param"
	"github.com/gofiber/fiber/v2"
	"log"
)

var ctx2 context.Context

func BookApi(api fiber.Router, ctx context.Context) {
	ctx2 = ctx
	api.Get("/book/create", bookCrateHandler)
	api.Get("/book/getBid", bookSelectHandler)
}

func bookCrateHandler(ctx *fiber.Ctx) error {
	b := new(param.BookParam)

	if err := ctx.BodyParser(b); err != nil {
		return err
	}
	log.Println(b.BookName) // john
	log.Println(b.Author)   // doe
	_, err := database.DBConn.Book.Create().SetBookName(b.BookName).SetAuthor(b.Author).Save(ctx2)
	if err != nil {
		log.Printf("err:[%v]", err)
		return fiber.NewError(fiber.StatusBadRequest, "insert error!")
	} else {
		return ctx.SendString("insert success")
	}
}

func bookSelectHandler(ctx *fiber.Ctx) error {
	bn := ctx.Query("name")
	all, err := database.DBConn.Book.Query().Where(book.BookNameEQ(bn)).All(ctx2)
	if err != nil {
		fmt.Errorf("bookSelectHandler have error :[%v]", err)
	}
	fmt.Println(all)
	res, err := json.Marshal(all)
	fmt.Println(res)
	return ctx.Send(res)
}
