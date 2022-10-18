package model

import "database/sql"

type Meja struct {
	Id        string       `json:"id"`
	NoMeja    string       `json:"nomeja" binding:"required"`
	Status    string       `json:"status" binding:"required"`
	CreatedAt string       `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
