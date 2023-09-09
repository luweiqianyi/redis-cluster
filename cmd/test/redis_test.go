package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

const (
	Host = "localhost"
	Port = "7002"
)

func TestRedisNodeGet(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", Host, Port),
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(client.Context()).Result()
	if err != nil {
		fmt.Println("连接 Redis 失败:", err)
		return
	}
	fmt.Println("成功连接到 Redis:", pong)

	resp := client.Get(client.Context(), "name")
	fmt.Println(resp)
}

func TestRedisClusterGet(t *testing.T) {
	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"127.0.0.1:7001",
			"127.0.0.1:7002",
			"127.0.0.1:7003",
			"127.0.0.1:7004",
			"127.0.0.1:7005",
			"127.0.0.1:7006",
		},
		Password: "",
	})

	ctx := context.Background()
	val, err := clusterClient.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key:", val)

	err = clusterClient.Close()
	if err != nil {
		panic(err)
	}
}
