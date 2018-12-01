package kvs

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

const Nil = redis.Nil

func New(dsn string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: "",
		DB:       0,
	})
	if err := client.Ping().Err(); err != nil {
		return nil, errors.Wrapf(err, "failed to ping redis server")
	}
	return client, nil
}

const (
	tokenKey      = "TOKEN_"
	tokenDuration = time.Second * 3600
)

func SetToken(cli *redis.Client, token string, userID int) error {
	if err := cli.Set(tokenKey+token, userID, tokenDuration).Err(); err != nil {
		return err
	}
	return nil
}

func GetIDByToken(cli *redis.Client, token string) (int, error) {
	v, err := cli.Get(tokenKey + token).Result()
	if err != nil {
		return 0, errors.Wrapf(err, "failed to get id from redis by token")
	}
	id, err := strconv.Atoi(v)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to convert string to int")
	}
	return id, nil
}
