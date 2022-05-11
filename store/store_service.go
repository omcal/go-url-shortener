package store

import (
	"context"
	"fmt"
	redis2 "github.com/go-redis/redis/v8"
	"time"
)

type StorageService struct {
	redisClient *redis2.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

/*
It  helps  me to initialize the store service which we can
assign to our variable that declared at above it is empty

*/
func InitializeStore() *StorageService {
	redisClient := redis2.NewClient(&redis2.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil { //nolint:wsl
		panic(fmt.Sprintf("Error init redis: %v\n", err)) //nolint:govet
	}
	fmt.Printf("\n Redis stared successfully : pong message ={%s}", pong)
	storeService.redisClient = redisClient
	return storeService //nolint:wsl

}

/*
Storage API design and Implementation
*/

func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))

	}
}

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}
