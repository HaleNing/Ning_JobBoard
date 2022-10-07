package main

import (
	"context"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent"
	"github.com/HaleNing/Ning_JobBoard/src/database"
	"github.com/HaleNing/Ning_JobBoard/src/handler"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"log"
)

// world start
func main() {
	app := fiber.New()

	// file in handler will add interface to the fiber app.

	//host=localhost port=5432 user=ning  dbname=postgres sslmode=disable
	client, _ := database.NewConnection()
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	baseAPi := app.Group("/api")

	handler.BaseAPi(app)
	handler.UserApi(baseAPi)
	handler.JobApi(baseAPi, ctx)
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			// notify me
		}
	}(client)

	log.Fatal(app.Listen(":6099"))
}
