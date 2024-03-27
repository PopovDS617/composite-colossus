package model

type Region struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
