package helpers

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

func ScheduleWakaDataFetch() {
	defer func() {
		if r := recover(); r != nil {
			Logger.Error("ScheduleWakaDataFetch Crashed: ", zap.Any("error", r))
		}
	}()
	for {
		go WakaDataFetch()
		time.Sleep(24 * time.Hour)
	}
}

func WakaDataFetch() {
	defer func() {
		if r := recover(); r != nil {
			Logger.Error("WakaDataFetch Crashed: ", zap.Any("error", r))
		}
	}()
	date := time.Now().Add(-24 * time.Hour)
	formated_date := date.Format("2006-01-02")
	data, err := FetchWakaData(formated_date)
	if err != nil {
		Logger.Error(fmt.Sprintf("Failed to fetch waka data for date: %s", formated_date), zap.Error(err))
		return
	}
	datainsert := []interface{}{}
	for _, v := range data.Data {
		datainsert = append(datainsert, v)
	}
	if !MongoAddManyDoc("daily_logs", datainsert) {
		Logger.Error(fmt.Sprintf("Failed to insert waka data for date: %s", date))
	}
}
