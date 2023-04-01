package database

import (
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func NewRedisConnection() *redis.Client {

	//rdb := redis.NewClient(&redis.Options{
	//	Addr:     os.Getenv("REDIS_ADDR"),
	//	Password: os.Getenv("REDIS_PASS"), // no password set
	//	DB:       0,                       // use default DB
	//	Username: os.Getenv("REDIS_USERNAME"),
	//	TLSConfig: &tls.Config{
	//		MinVersion: tls.VersionTLS12},
	//})
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis://red-cgju5fu4dad69r7718cg:6379",
	})

	RDB = rdb
	return rdb
}
