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

	// line changes since last duration
	AIAdditions    int `bson:"ai_additions" json:"ai_additions"`
	AIDeletions    int `bson:"ai_deletions" json:"ai_deletions"`
	HumanAdditions int `bson:"human_additions" json:"human_additions"`
	HumanDeletions int `bson:"human_deletions" json:"human_deletions"`

	// AI usage
	AIModelCosts        map[string]float64 `bson:"ai_model_costs,omitempty" json:"ai_model_costs,omitempty"`               // estimated USD cost per model
	AIModelLineChanges  map[string]int     `bson:"ai_model_line_changes,omitempty" json:"ai_model_line_changes,omitempty"` // lines changed per model
	AIModelBreakdown    []AIModelBreakdown `bson:"ai_model_breakdown,omitempty" json:"ai_model_breakdown,omitempty"`
	AIInputTokens       int                `bson:"ai_input_tokens" json:"ai_input_tokens"`
	AICachedInputTokens int                `bson:"ai_cached_input_tokens" json:"ai_cached_input_tokens"`
	AIOutputTokens      int                `bson:"ai_output_tokens" json:"ai_output_tokens"`
	AISessions          int                `bson:"ai_sessions" json:"ai_sessions"`

	// AI prompt stats
	AIPromptLengthAvg              int `bson:"ai_prompt_length_avg" json:"ai_prompt_length_avg"`
	AIPromptLengthAvgPerSession    int `bson:"ai_prompt_length_avg_per_session" json:"ai_prompt_length_avg_per_session"`
	AIPromptLengthMedianPerSession int `bson:"ai_prompt_length_median_per_session" json:"ai_prompt_length_median_per_session"`
	AIPromptLengthSum              int `bson:"ai_prompt_length_sum" json:"ai_prompt_length_sum"`
	AIPromptEventsTotal            int `bson:"ai_prompt_events_total" json:"ai_prompt_events_total"`
	AIPromptEventsAvgPerSession    int `bson:"ai_prompt_events_avg_per_session" json:"ai_prompt_events_avg_per_session"`
	AIPromptEventsMedianPerSession int `bson:"ai_prompt_events_median_per_session" json:"ai_prompt_events_median_per_session"`
}

type AIModelBreakdown struct {
	Name              string  `bson:"name" json:"name"`
	Cost              float64 `bson:"cost" json:"cost"` // estimated USD
	Lines             int     `bson:"lines" json:"lines"`
	InputTokens       int     `bson:"input_tokens" json:"input_tokens"`
	CachedInputTokens int     `bson:"cached_input_tokens" json:"cached_input_tokens"`
	OutputTokens      int     `bson:"output_tokens" json:"output_tokens"`
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

	// AI usage since last heartbeat
	AISession           *string `bson:"ai_session,omitempty" json:"ai_session,omitempty"`
	AISubscriptionPlan  *string `bson:"ai_subscription_plan,omitempty" json:"ai_subscription_plan,omitempty"`
	AIInputTokens       int     `bson:"ai_input_tokens" json:"ai_input_tokens"`
	AICachedInputTokens int     `bson:"ai_cached_input_tokens" json:"ai_cached_input_tokens"`
	AIOutputTokens      int     `bson:"ai_output_tokens" json:"ai_output_tokens"`
	AIPromptLength      int     `bson:"ai_prompt_length" json:"ai_prompt_length"`
}

type MachineResponse struct {
	Data []Machine `bson:"data" json:"data"`
}

type Machine struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Value      string    `json:"value"`
	IP         string    `json:"ip"`
	LastSeenAt time.Time `json:"last_seen_at"`
	Timezone   string    `json:"timezone"`
	CreatedAt  time.Time `json:"created_at"`
}
