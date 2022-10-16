package model

import "database/sql"

type Menu struct {
	Id        string       `json:"id"`
	Nama      string       `json:"nama" binding:"required"`
	Jenis     string       `json:"jenis" binding:"required"`
	Harga     int          `json:"harga" binding:"required"`
	Stock     int          `json:"stock" binding:"required"`
	CreatedAt string       `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
