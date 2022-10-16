package model

import "database/sql"

type Karyawan struct {
	Id        string       `json:"id"`
	Nama      string       `json:"nama" binding:"required"`
	Alamat    string       `json:"alamat" binding:"required"`
	Gaji      int          `json:"gaji" binding:"required"`
	Username  string       `json:"username" binding:"required"`
	Password  string       `json:"password" binding:"required"`
	CreatedAt string       `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
