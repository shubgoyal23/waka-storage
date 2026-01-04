package models

import "time"

type ActivityResponse struct {
	Data     []Activity `bson:"data" json:"data"`
	End      string     `bson:"end" json:"end"`
	Start    string     `bson:"start" json:"start"`
	Timezone string     `bson:"timezone" json:"timezone"`
}

type Activity struct {
	Color    *string `bson:"color,omitempty" json:"color,omitempty"` // Using pointer to handle null
	Duration float64 `bson:"duration" json:"duration"`               // in seconds
	Project  string  `bson:"project" json:"project"`                 // name of the project
	Time     float64 `bson:"time" json:"time"`                       // Unix timestamp with fractional seconds
}

type HeartbeatResponse struct {
	Data     []WakaHeartbeat `bson:"data" json:"data"`
	End      string          `bson:"end" json:"end"`
	Start    string          `bson:"start" json:"start"`
	Timezone string          `bson:"timezone" json:"timezone"`
}

type WakaHeartbeat struct {
	ID               string    `bson:"_id" json:"id"`
	Branch           *string   `bson:"branch,omitempty" json:"branch,omitempty"`
	Category         string    `bson:"category" json:"category"`
	CreatedAt        time.Time `bson:"created_at" json:"created_at"`
	CursorPos        int       `bson:"cursorpos" json:"cursorpos"`
	Dependencies     []string  `bson:"dependencies" json:"dependencies"`
	Entity           string    `bson:"entity" json:"entity"`
	IsWrite          bool      `bson:"is_write" json:"is_write"`
	Language         string    `bson:"language" json:"language"`
	LineNo           int       `bson:"lineno" json:"lineno"`
	Lines            int       `bson:"lines" json:"lines"`
	MachineNameID    string    `bson:"machine_name_id" json:"machine_name_id"`
	Project          string    `bson:"project" json:"project"`
	ProjectRootCount *int      `bson:"project_root_count,omitempty" json:"project_root_count,omitempty"`
	Time             float64   `bson:"time" json:"time"` // Unix timestamp
	Type             string    `bson:"type" json:"type"`
	UserAgentID      string    `bson:"user_agent_id" json:"user_agent_id"`
	AILineChanges    *int      `bson:"ai_line_changes,omitempty" json:"ai_line_changes,omitempty"`
	HumanLineChanges *int      `bson:"human_line_changes,omitempty" json:"human_line_changes,omitempty"`
	UserID           string    `bson:"user_id" json:"user_id"`
}
