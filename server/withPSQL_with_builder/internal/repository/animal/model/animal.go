package model

import (
	"time"
)

type Animal struct {
	ID             int64     `db:"id"`
	Name           string    `db:"name"`
	Type           string    `db:"type"`
	Age            int       `db:"age"`
	Gender         string    `db:"gender"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	RegionID       int       `db:"region_id"`
	RegionName     string    `db:"region_name"`
	LastTimeSeenAt time.Time `db:"last_time_seen_at"`
	SeenByDevice   int       `db:"seen_by_device"`
}
