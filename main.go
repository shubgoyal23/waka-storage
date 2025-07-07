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
	if err := helpers.InitRediGo(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PWD")); err != nil {
		panic(err)
	}

	if logger, err := helpers.InitLogger(); err != nil {
		panic(err)
	} else {
		Logger = logger
	}

	if f := helpers.MongoInit(os.Getenv("MONGODB_URI"), os.Getenv("MONGO_DB")); !f {
		panic(f)
	}
	helpers.WakaInit(os.Getenv("WAKA_API_KEY"))

	Logger.Info("Starting Server", zap.String("version", "2.0.0"))
	go helpers.ScheduleWakaDataFetch()
	<-stop
	Logger.Info("Shutting down Server")
}
