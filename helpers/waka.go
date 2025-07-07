package helpers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"waka-storage/models"
)

var (
	apiKey      = ""
	wakaBaseURL = "https://wakatime.com/api/v1/users/current"
)

func WakaInit(str string) {
	apiKey = str
}

func FetchWakaData(date string) (*models.ActivityResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/durations?date=%s", wakaBaseURL, date)
	req, _ := http.NewRequest("GET", url, nil)

	auth := base64.StdEncoding.EncodeToString([]byte(apiKey))
	req.Header.Set("Authorization", "Basic "+auth)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result models.ActivityResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
