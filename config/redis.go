package config

import (
	"context"
	"demo/global"
	"log"

	"github.com/go-redis/redis"
)

func InitRedis(ctx context.Context) error {
	
	Client := redis.NewClient(&redis.Options{
		Addr:     AppConfig.Redis.Host + ":" + AppConfig.Redis.Port,
		Password: AppConfig.Redis.Password,
		DB:       AppConfig.Redis.DB,
	})
	_,err:=Client.Ping().Result()
	if err!= nil {
		log.Fatal("failed to connect redis")
		return err
	}
	global.RedisDb=Client
	log.Println("redis connected")
	return nil
}