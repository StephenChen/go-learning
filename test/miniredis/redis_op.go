package miniredis_demo

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

const (
	KeyValidWebsite = "app:valid:website:list"
)

func DoSomethingWithRedis(rdb *redis.Client, key string) bool {
	ctx := context.TODO()
	if !rdb.SIsMember(ctx, KeyValidWebsite, key).Val() {
		return false
	}
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return false
	}
	if !strings.HasPrefix(val, "https://") {
		val = "https://" + val
	}
	// 设置 blog key 五秒过期
	if err := rdb.Set(ctx, "blog", val, 5*time.Second).Err(); err != nil {
		return false
	}
	return true
}
