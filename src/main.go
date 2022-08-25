package main

import (
	"context"
	"github.com/HaleNing/bustrack/src/Model/ent"
	"github.com/HaleNing/bustrack/src/database"
	"github.com/HaleNing/bustrack/src/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

// world start
func main() {
	app := fiber.New()

	baseAPi := app.Group("/api")
	// file in handler will add interface to the fiber app.

	//读取环境变量文件
	err := godotenv.Load("src/config/config.env")
	if err != nil {
		log.Fatal(err)
	}

	//host=localhost port=5432 user=ning  dbname=postgres sslmode=disable
	client, _ := database.NewConnection()

	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	handler.BaseAPi(app)
	handler.DriverApi(baseAPi)
	handler.BusApi(baseAPi, ctx)
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			// notify me
		}
	}(client)

	log.Fatal(app.Listen(":6099"))
}
