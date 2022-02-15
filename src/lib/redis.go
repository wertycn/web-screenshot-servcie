package lib

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
	"web-screenshot-service/lib/util"
)

type ClientConfig struct {
	Mode   string
	Ip     string
	Port   string
	Passwd string
}

var ctx = context.Background()
var rdb *redis.Client

func CreateRedisClient(ip string, port string, passwd string) *redis.Client {
	log.Infof("create redis client start (ip=%s,port=%s,auth_md5=%s)", ip, port, util.GetMd5(passwd))
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
	log.Infof("create redis client success (ip=%s,port=%s,auth_md5=%s)", ip, port, util.GetMd5(passwd))
	return rdb
}

func CreateRedisClientByConfig(config ClientConfig) *redis.Client {
	if config.Mode == "sidecar" {
		return CreateRedisClient(
			os.Getenv(config.Ip),
			os.Getenv(config.Port),
			config.Passwd,
		)
	}
	return CreateRedisClient(config.Ip, config.Port, config.Passwd)
}

func HashSet(key string, filed string, value string) int64 {
	result, err := rdb.HSet(ctx, key, filed, value).Result()
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

func LPush(key string, element string) int64 {
	result, err := rdb.LPush(ctx, key, element).Result()
	if err != nil {
		log.Errorf("list LPush error (key=%s,field=%s):%s", key, element, err.Error())
	}
	return result
}

func RPush(key string, element string) int64 {
	result, err := rdb.RPush(ctx, key, element).Result()
	if err != nil {
		log.Errorf("list RPush error (key=%s,field=%s):%s", key, element, err.Error())
	}
	return result
}

func LPop(key string) string {
	log.Info(key)
	pop := rdb.LPop(ctx, key)
	if pop == nil {
		return ""
	}
	result, err := pop.Result()
	if err != nil {
		log.Errorf("list lPop error (key=%s):%s", key, err.Error())
	}
	return result
}

func BLPop(key string, timeout int) string {
	result, err := rdb.BLPop(ctx, time.Duration(timeout)*time.Second, key).Result()
	if err != nil {
		log.Infof("list BLPop error (key=%s,msg=%s)", key, err.Error())
		return ""
	}
	if len(result) < 2 {
		return ""
	}
	return result[1]
}

func RPop(key string) string {
	result, err := rdb.RPop(ctx, key).Result()
	if err != nil {
		log.Errorf("list rPop error (key=%s,field=%s):%s", key, err.Error())
	}
	return result
}

func Get(key string) string {
	result, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Errorf("get error (key=%s,field=%s):%s", key, err.Error())
	}
	return result
}

func Set(key string, value string, timeout int) string {
	ttl := time.Duration(timeout) * time.Second
	result, err := rdb.Set(ctx, key, value, ttl).Result()
	if err != nil {
		log.Errorf("set error (key=%s,value=%s,timeout=%d):%s", key, value, timeout, err.Error())
	}
	return result
}
