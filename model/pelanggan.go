package model

import "database/sql"

type Pelanggan struct {
	Id        string       `json:"id"`
	Nama      string       `json:"nama" binding:"required"`
	Alamat    string       `json:"alamat" binding:"required"`
	NoHp      int          `json:"nohp" binding:"required"`
	CreatedAt string       `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
