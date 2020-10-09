package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	// err := rdb.Set(ctx, "key", "value", 0).Err()
	// if err != nil {
	// panic(err)
	// }

	// rdb.Del(ctx, "key")

	val, err := rdb.Keys(ctx, "dex_descomplica_auth_*").Result()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(len(val))
	for _, v := range val {
		fmt.Println(v)
		val2, _ := rdb.Get(ctx, v).Result()
		if len(val2) < 1024 {
			fmt.Println(val2)
		}
	}
}
