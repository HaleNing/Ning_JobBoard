package database

import (
	"crypto/tls"
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func NewRedisConnection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "singapore-re****",
		Password: "rjcsW***", // no password set
		DB:       0,          // use default DB
		Username: "red****",
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12},
	})
	RDB = rdb
	return rdb
}
