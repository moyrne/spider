package server

import "time"

// Work 任务模板
type Work struct {
	CreateAt time.Time  `json:"create_at"`
	Name     string     `json:"name"`
	URL      string     `json:"url"`
	Adapter  string     `json:"adapter"`
	Args     []string   `json:"args"`
	Regex    [][]string `json:"regex"`
}
