package helpers

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func ScheduleWakaDataFetch() {
	defer func() {
		if r := recover(); r != nil {
			Logger.Error("ScheduleWakaDataFetch Crashed: ", zap.Any("error", r))
		}
	}()

	// skip if last update is less than 24 hours
	skip := false
	var data bson.M
	MongoGetLastOneDoc("daily_logs", &data)
	if data != nil {
		date, ok := data["time"].(float64)
		if ok {
			lastUpdate := time.Unix(int64(date), 0)
			ct := time.Now().Add(-24 * time.Hour)
			if lastUpdate.Day() >= ct.Day() {
				skip = true
			}
		}
	}

	for {
		if !skip {
			formated_date := time.Now().Add(-24 * time.Hour).Format("2006-01-02")
			WakaDataFetchActivity(formated_date)
			WakaDataFetchHeartbeat(formated_date)
		}
		skip = false
		time.Sleep(24 * time.Hour)
	}
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
	datainsert := []interface{}{}
	for _, v := range data.Data {
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
	datainsert := []interface{}{}
	for _, v := range data.Data {
		datainsert = append(datainsert, v)
	}
	if !MongoAddManyDoc("heartbeats", datainsert) {
		Logger.Error(fmt.Sprintf("Failed to insert waka data for date: %s", formated_date))
	}
}
