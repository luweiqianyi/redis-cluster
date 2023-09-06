package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"host.docker.internal:7001",
			"host.docker.internal:7002",
			"host.docker.internal:7003",
			"host.docker.internal:7004",
			"host.docker.internal:7005",
			"host.docker.internal:7006",
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
