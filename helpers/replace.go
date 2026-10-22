package helpers

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// ProjectReplaceKey is the Redis list the dashboard pushes project rename
// rules to. Each item is JSON: {"from":"old-name","to":"new-name"}.
const ProjectReplaceKey = "waka:project_replace"

type projectRule struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// LoadProjectRules reads the rename rules; on failure it logs and returns none
// so a Redis outage never blocks storing data.
func LoadProjectRules() map[string]string {
	rules := map[string]string{}
	url := os.Getenv("REDIS_URL")
	if url == "" {
		return rules
	}
	opts, err := redis.ParseURL(url)
	if err != nil {
		Logger.Error("Invalid REDIS_URL", zap.Error(err))
		return rules
	}
	client := redis.NewClient(opts)
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	items, err := client.LRange(ctx, ProjectReplaceKey, 0, -1).Result()
	if err != nil {
		Logger.Error("Failed to load project replace rules", zap.Error(err))
		return rules
	}
	// Rules are applied in list order, so a later rule wins and chains
	// (a->b, then b->c) resolve through ResolveProject.
	for _, item := range items {
		var r projectRule
		if json.Unmarshal([]byte(item), &r) == nil && r.From != "" && r.To != "" {
			rules[r.From] = r.To
		}
	}
	return rules
}

// ResolveProject follows rename rules (including chains) for a project name.
func ResolveProject(rules map[string]string, name string) string {
	for i := 0; i < len(rules)+1; i++ { // bounded to survive cycles
		next, ok := rules[name]
		if !ok || next == name {
			break
		}
		name = next
	}
	return name
}
