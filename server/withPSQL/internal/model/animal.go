package model

import "time"

type Animal struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Type           string    `json:"type"`
	Age            int       `json:"age"`
	Gender         string    `json:"gender"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Region         string    `json:"region"`
	LastTimeSeenAt time.Time `json:"last_time_seen_at"`
	SeenByDevice   int       `json:"seen_by_device"`
}
