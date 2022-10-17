package model

import "database/sql"

type Menu struct {
	Id        string `form:"id"`
	Nama      string `form:"nama" binding:"required"`
	Jenis     string `form:"jenis" binding:"required"`
	Harga     int    `form:"harga" binding:"required"`
	Stock     int    `form:"stock" binding:"required"`
	Img       string
	CreatedAt string       `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
