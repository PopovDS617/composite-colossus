package model

import "time"

type Animal struct {
	ID        int64
	Name      string
	Type      string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
