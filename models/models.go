package models

type ActivityResponse struct {
	Data     []Activity `json:"data"`
	End      string     `json:"end"`
	Start    string     `json:"start"`
	Timezone string     `json:"timezone"`
}

type Activity struct {
	Color    *string `json:"color"`    // Using pointer to handle null
	Duration float64 `json:"duration"` // in seconds
	Project  string  `json:"project"`  // name of the project
	Time     float64 `json:"time"`     // Unix timestamp with fractional seconds
}
