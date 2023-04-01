package main

import (
	"context"
	"github.com/HaleNing/Ning_JobBoard/src/Model/ent"
	"github.com/HaleNing/Ning_JobBoard/src/database"
	"github.com/HaleNing/Ning_JobBoard/src/handler"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

// world start
func main() {
	app := fiber.New()
	err := godotenv.Load("src/config.env")
	if err != nil {
		log.Fatalf("failed to read env file :[%v]", err)
	}

	client, dbErr := database.NewConnection()
	if dbErr != nil {
		log.Println("pgsql connect fail")
		log.Fatal(dbErr)
	} else {
		log.Println("pgsql connect success")
	}
	redisConnect := database.NewRedisConnection()
	if redisConnect == nil {
		log.Println("redis connect fail")
		log.Fatal(err)
	} else {
		log.Println("redis connect success")
	}

	ctx := context.Background()

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
			// todo  notify me
		}
	}(client)

	defer func(redisClient *redis.Client) {
		err2 := redisClient.Close()
		if err2 != nil {
			// todo notify me
		}
	}(redisConnect)

	log.Fatal(app.Listen(":6099"))
}
