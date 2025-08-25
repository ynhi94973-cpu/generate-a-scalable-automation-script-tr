package main

import (
	"time"
)

type AutomationScript struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Status      string    `json:"status"` // "active" or "inactive"
}

type Tracker struct {
	ID        string `json:"id"`
	ScriptID  string `json:"script_id"`
	RunCount  int    `json:"run_count"`
	LastError string `json:"last_error"`
	LastRunAt time.Time `json:"last_run_at"`
}

type AutomationScriptTracker struct {
	Scripts      []AutomationScript `json:"scripts"`
	TrackerLogs []Tracker          `json:"tracker_logs"`
}

func main() {}