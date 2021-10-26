package lib

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var ctx = context.Background()
var rdb *redis.Client

func CreateRedisClient(ip string, port string, passwd string) *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", ip, port),
		Password: passwd,
		DB:       0,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Errorf("create redis client error: %s", err.Error())
		panic(err)
	}
	log.Infof("create redis client success (ip=%s,prot=%s)", ip, port)
	return rdb
}

func HashSet(key string, filed string, value string) int64 {
	result, err := rdb.HSet(ctx, key, map[string]string{filed: value}).Result()
	if err != nil {
		log.Errorf("hash set error (key=%s):%s", key, err.Error())
	}
	return result
}

func HashGet(key string, field string) string {
	result, err := rdb.HGet(ctx, key, field).Result()
	if err != nil {
		log.Errorf("hash get error (key=%s,field=%s):%s", key, field, err.Error())
	}
	return result

}
