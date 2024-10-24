package model

import "time"

// Record is a model for record in log table
type Record struct {
	ID        int64
	ChatID    int64
	Action    string
	CreatedAt time.Time
}
