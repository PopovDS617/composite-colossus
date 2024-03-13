package model

import (
	"database/sql"
	"time"
)

type Animal struct {
	ID        int64        `db:"id"`
	Name      string       `db:"name"`
	Type      string       `db:"type"`
	Age       int          `db:"age"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
