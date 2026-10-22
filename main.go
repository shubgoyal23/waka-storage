package main

import (
	"os"
	"os/signal"
	"syscall"
	"waka-storage/helpers"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	godotenv.Load()
	var Logger *zap.Logger
	if logger, err := helpers.InitLogger(os.Getenv("UPSTASH_REDIS_REST_URL"), os.Getenv("UPSTASH_REDIS_REST_TOKEN"), helpers.RedisKey("worker-logs")); err != nil {
		panic(err)
	} else {
		Logger = logger
	}

	if f := helpers.MongoInit(os.Getenv("MONGODB_URI"), os.Getenv("MONGO_DB")); !f {
		panic(f)
	}
	if os.Getenv("WAKA_ACCESS_TOKEN") == "" {
		panic("WAKA_ACCESS_TOKEN is not set")
	}
	helpers.WakaInit(os.Getenv("WAKA_ACCESS_TOKEN"))

	Logger.Info("Starting Server", zap.String("version", "2.0.0"))

	if os.Getenv("RUN_NOW") == "true" {
		Logger.Info("RUN_NOW flag set, running WakaDataFetch immediately")
		go helpers.WakaDataFetch()
	}

	go helpers.ScheduleWakaDataFetch()
	<-stop
	Logger.Info("Shutting down Server")
}
