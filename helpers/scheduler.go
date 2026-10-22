package helpers

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func ScheduleWakaDataFetch() {
	defer func() {
		if r := recover(); r != nil {
			Logger.Error("ScheduleWakaDataFetch Crashed: ", zap.Any("error", r))
		}
	}()

	c := cron.New()
	_, err := c.AddFunc("5 0 * * *", WakaDataFetch)
	if err != nil {
		Logger.Error("Failed to schedual cron job", zap.Any("error", err.Error()))
		return
	}
	c.Start()
	Logger.Info("Cron scheduler started", zap.String("time", time.Now().Format("2006-01-02 15:04:05")))
}

func WakaDataFetch() {
	defer func() {
		if r := recover(); r != nil {
			Logger.Error("WakaDataFetch Crashed: ", zap.Any("error", r))
		}
	}()
	formated_date := time.Now().Add(-1 * time.Hour).Format("2006-01-02")
	Logger.Info("WakaDataFetch started", zap.String("date", formated_date), zap.String("time", time.Now().Format("2006-01-02 15:04:05")))
	WakaDataFetchActivity(formated_date)
	WakaDataFetchHeartbeat(formated_date)
	WakaDataFetchMachineIds(1)
	Logger.Info("WakaDataFetch done", zap.String("date", formated_date), zap.String("time", time.Now().Format("2006-01-02 15:04:05")))
}

func WakaDataFetchActivity(formated_date string) {
	defer func() {
		if r := recover(); r != nil {
			Logger.Error("WakaDataFetchActivity Crashed: ", zap.Any("error", r))
		}
	}()

	data, err := FetchWakaDataActivity(formated_date)
	if err != nil {
		Logger.Error(fmt.Sprintf("Failed to fetch waka data for date: %s", formated_date), zap.Error(err))
		return
	}
	rules := LoadProjectRules()
	datainsert := []interface{}{}
	for _, v := range data.Data {
		v.Project = ResolveProject(rules, v.Project)
		datainsert = append(datainsert, v)
	}
	if !MongoAddManyDoc("daily_logs", datainsert) {
		Logger.Error(fmt.Sprintf("Failed to insert waka data for date: %s", formated_date))
	}
}

func WakaDataFetchHeartbeat(formated_date string) {
	defer func() {
		if r := recover(); r != nil {
			Logger.Error("WakaDataFetchHeartbeat Crashed: ", zap.Any("error", r))
		}
	}()

	data, err := FetchWakaDataHeartbeat(formated_date)
	if err != nil {
		Logger.Error(fmt.Sprintf("Failed to fetch waka data for date: %s", formated_date), zap.Error(err))
		return
	}
	rules := LoadProjectRules()
	datainsert := []interface{}{}
	for _, v := range data.Data {
		v.Project = ResolveProject(rules, v.Project)
		datainsert = append(datainsert, v)
	}
	if !MongoAddManyDoc("heartbeats", datainsert) {
		Logger.Error(fmt.Sprintf("Failed to insert waka data for date: %s", formated_date))
	}
}

func WakaDataFetchMachineIds(page int) {
	defer func() {
		if r := recover(); r != nil {
			Logger.Error("WakaDataFetchMachineIds Crashed: ", zap.Any("error", r))
		}
	}()

	data, err := FetchWakaDataMachineIds(page)
	if err != nil {
		Logger.Error("Failed to fetch waka data for machine ids", zap.Error(err))
		return
	}
	datainsert := []interface{}{}
	for _, v := range data.Data {
		datainsert = append(datainsert, v)
	}
	if !MongoAddManyDoc("machine_names", datainsert) {
		Logger.Error("Failed to insert waka data for machine ids")
	}
}
