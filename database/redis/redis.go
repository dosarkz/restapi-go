package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

type Cli struct {
	client *redis.Client
}

var redisCli *Cli = nil

func ConnectToRedis(conf Redis) *Cli {
	redisCli = new(Cli)
	redisCli.client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Get().Host, conf.Get().Port),
		Password: conf.Get().Password,
		DB:       0,
	})
	_, err := redisCli.client.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}

	return redisCli
}

func (r *Cli) SetValue(key string, value string, expiration time.Duration) (bool, error) {
	err := r.client.Set(key, value, expiration).Err()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *Cli) GetValue(key string) (string, error) {
	val, err := r.client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
