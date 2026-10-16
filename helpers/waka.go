package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"waka-storage/models"
)

var (
	accessToken = ""
	wakaBaseURL = "https://wakatime.com/api/v1/users/current"
	httpClient  = &http.Client{Timeout: 60 * time.Second}
)

// WakaInit sets the OAuth access token used for reading data.
// WakaTime API keys are write-only (plugins only), so reads need an OAuth token
// with scopes read_heartbeats and read_stats.machines.
func WakaInit(token string) {
	accessToken = token
}

func wakaGet(path string, out interface{}) error {
	req, err := http.NewRequest("GET", wakaBaseURL+path, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		return fmt.Errorf("wakatime %s returned %d: %s", path, resp.StatusCode, body)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func FetchWakaDataActivity(date string) (*models.ActivityResponse, error) {
	var result models.ActivityResponse
	if err := wakaGet(fmt.Sprintf("/durations?date=%s", date), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func FetchWakaDataHeartbeat(date string) (*models.HeartbeatResponse, error) {
	var result models.HeartbeatResponse
	if err := wakaGet(fmt.Sprintf("/heartbeats?date=%s", date), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func FetchWakaDataMachineIds(page int) (*models.MachineResponse, error) {
	var result models.MachineResponse
	if err := wakaGet(fmt.Sprintf("/machine_names?page=%d", page), &result); err != nil {
		return nil, err
	}
	return &result, nil
}
