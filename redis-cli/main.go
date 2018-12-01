package main

import (
	"fmt"
	"os"
	"time"

	"reflect"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	fmt.Fprintln(os.Stdout, pong)

	if err := client.Set("key1", "value1", time.Second*3).Err(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	val1, err := client.Get("key1").Result()
	if err == redis.Nil {
		fmt.Fprintln(os.Stderr, "no value")
		return
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	fmt.Fprintln(os.Stdout, val1)

	if exists, err := client.Exists("key1").Result(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	} else if exists == 1 {
		fmt.Fprintln(os.Stdout, "key already saved")
	}

	// rewrite
	if err := client.Set("key1", "value2", time.Second*3).Err(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	// reget
	val1, err = client.Get("key1").Result()
	if err == redis.Nil {
		fmt.Fprintln(os.Stderr, "no value")
		return
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	fmt.Fprintln(os.Stdout, val1)

	// set hash
	res, err := client.HSet("hkey1", "user_id", 1).Result()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	fmt.Fprintln(os.Stdout, res)

	// get hash
	userID, err := client.HGet("hkey1", "user_id").Result()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	fmt.Fprintln(os.Stdout, userID)

	// set multiple hash FIXME expiration
	s, err := client.HMSet("hkey2", map[string]interface{}{"user_id": 1, "scope": "scope1"}).Result()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	} else if s != "OK" {
		fmt.Fprintln(os.Stderr, "failed")
		return
	}

	// get multiple hash
	tkr, err := client.HMGet("hkey2", []string{"user_id", "scope"}...).Result()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	for i, v := range tkr {
		fmt.Fprintf(os.Stdout, "key: %#v, val: %#v\n", i, v)
		fmt.Fprintf(os.Stdout, "type of value '%#v' is %s\n", v, reflect.TypeOf(v).String())
	}

	// no record
	tkr, err = client.HMGet("nokey", []string{"user_id", "scope"}...).Result()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	// error返ってこないのか
	fmt.Fprintf(os.Stdout, "len %d\n", len(tkr))
	for i, v := range tkr {
		fmt.Fprintf(os.Stdout, "key: %#v, val: %#v\n", i, v)
		if v != nil {
			fmt.Fprintf(os.Stdout, "type of value '%#v' is %s\n", v, reflect.TypeOf(v).String())
		} else {
			fmt.Fprintf(os.Stdout, "type of value '%#v' is nil\n", v)
		}
	}
}

type Token struct {
	UserID int
	Scope  string
}
