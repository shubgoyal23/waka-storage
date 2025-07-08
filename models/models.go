package models

import "time"

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

type HeartbeatResponse struct {
	Data     []WakaHeartbeat `json:"data"`
	End      string          `json:"end"`
	Start    string          `json:"start"`
	Timezone string          `json:"timezone"`
}

type WakaHeartbeat struct {
	ID               string    `bson:"_id" json:"id"`
	Branch           *string   `json:"branch,omitempty"`
	Category         string    `json:"category"`
	CreatedAt        time.Time `json:"created_at"`
	CursorPos        int       `json:"cursorpos"`
	Dependencies     []string  `json:"dependencies"`
	Entity           string    `json:"entity"`
	IsWrite          bool      `json:"is_write"`
	Language         string    `json:"language"`
	LineAdditions    *int      `json:"line_additions,omitempty"`
	LineDeletions    *int      `json:"line_deletions,omitempty"`
	LineNo           int       `json:"lineno"`
	Lines            int       `json:"lines"`
	MachineNameID    string    `json:"machine_name_id"`
	Project          string    `json:"project"`
	ProjectRootCount *int      `json:"project_root_count,omitempty"`
	Time             float64   `json:"time"` // Unix timestamp
	Type             string    `json:"type"`
	UserAgentID      string    `json:"user_agent_id"`
}
