package db_redis

import (
	"context"
	"time"
	"webServer/config"
	inredis "webServer/pkg/redis"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var DB *redis.Client

func Start() {
	redisip := config.Config.GetString("redis.ip")
	if redisip != "" {
		redisip := config.Config.GetString("redis.ip")
		redislink := inredis.Link{
			Password: config.Config.GetString("redis.password"),
			Db:       config.Config.GetInt("redis.db"),
			Ip:       redisip,
			Port:     config.Config.GetString("redis.port"),
		}
		DB = inredis.Start(redislink)
	}
}

func keyUserLoginCode(account string) string {
	return "UserLoginCode|" + account
}

func keyUseridByAccount(account string) string {
	return "UseridByAccount|" + account
}

func SetUserLoginCode(account, code string) (bool, error) {
	key := keyUserLoginCode(account)
	return DB.SetNX(ctx, key, code, 300*time.Second).Result()
}

func GetUserLoginCode(account string) string {
	key := keyUserLoginCode(account)
	return DB.Get(ctx, key).Val()
}

func DelUserLoginCode(account string) bool {
	key := keyUserLoginCode(account)
	_, err := DB.Del(ctx, key).Result()
	return err == nil
}

func GetUseridByAccount(account string) (string, error) {
	key := keyUseridByAccount(account)
	return DB.Get(ctx, key).Result()
}

func SetUseridByAccount(account, userid string) error {
	key := keyUseridByAccount(account)
	return DB.Set(ctx, key, userid, 0).Err()
}
