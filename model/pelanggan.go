package model

import "database/sql"

type Pelanggan struct {
	Id        string       `json:"id"`
	Nama      string       `json:"nama" binding:"required"`
	NoHp      string       `json:"nohp" binding:"required"`
	Alamat    int          `json:"alamat" binding:"required"`
	CreatedAt string       `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
