package models

import "time"

type Reminder struct {
	ID string `json:"id"`
	Title      string        `json:"title"`
	Message    string        `json:"message"`
	Duration   time.Duration `json:"duration"`
	CreatedAt  time.Time     `json:"created_at"`
	ModifiedAt time.Time     `json:"modified_at"`
	
}