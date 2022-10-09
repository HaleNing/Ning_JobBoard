package main

import (
	"context"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent"
	"github.com/HaleNing/Ning_JobBoard/src/database"
	"github.com/HaleNing/Ning_JobBoard/src/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

// world start
func main() {
	app := fiber.New()
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalf("failed to read env file :[%v]", err)
	}

	client, _ := database.NewConnection()
	database.NewRedisConnection()
	ctx := context.Background()
	if err != nil {
		log.Println("redis connect fail")
		log.Fatal(err)
	} else {
		log.Println("redis connect success")
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	handler.BaseAPi(app, ctx)
	baseAPi := app.Group("/api")
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
