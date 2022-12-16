package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type User struct {
	Name, Avatar, Password string
}

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // password set
		DB:       0,        // use default DB
	})

	status := rdb.Ping(ctx)
	fmt.Println(status)

	err := rdb.Set(ctx, "Tifa", 16, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "Tifa").Result()
	if err == redis.Nil {
		fmt.Println("The key[Tifa] does not exist!!!")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("Tifa", val)
	}

	val2, err := rdb.Get(ctx, "Cloud").Result()
	if err == redis.Nil {
		fmt.Println("The key[Cloud] does not exist!!!")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("Cloud", val2)
	}
	// Output: key value
	// key2 does not exist

	user1 := &User{
		Name:     "Cloud",
		Avatar:   "/0x2345/0x4562",
		Password: "123456",
	}
	err = rdb.Set(ctx, "user:1", user1, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	ret, err := rdb.Get(ctx, "user:1").Result()
	if err == redis.Nil {
		fmt.Println("The key[user:1] does not exist!!!")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("user:1", ret)
	}
}
