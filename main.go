package main

import (
	"context"
	"fmt"
	"time"

	"gitlab.id.vin/gami/ps2-gami-common/adapters/cache"
	"gitlab.id.vin/gami/ps2-gami-common/logger"
)

type TestStruct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Count       int64  `json:"count"`
}

func main() {
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       0,
		Password: "",
	})

	redisAdapter := cache.NewCacheV2Adapter(redisClient)
	ans, _ := redisAdapter.RGetBits(context.Background(), []string{"group:6", "group:5", "group:7", "group:8"}, 131)
	logger.Infof("%v", ans)

	resp := redisClient.Ping(ctx)
	fmt.Println(resp.String())

	data := CheckInProgressionData{}

	cli := cache.NewCacheV2Adapter(redisClient)

	err := cli.HGet(ctx, "gms_checkin_progression:DailyCheckinT07_Levelup", "11121169", &data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
}

type CheckInProgressionData struct {
	CheckInData uint64    `json:"check_in_data"`
	LastDate    time.Time `json:"last_date"`
	StartDate   time.Time `json:"start_date"`
}
