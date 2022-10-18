package model

import (
	"database/sql"
	"time"
)

type Nota struct {
	Id            string    `json:"id"`
	PelangganId   string    `json:"pelangganid" binding:"required"`
	KaryawanId    string    `json:"karyawanid" binding:"required"`
	WaktuPesan    time.Time `json:"waktupesan"`
	MejaId        string    `json:"mejaid" binding:"required"`
	Harga         int
	Total         int
	NamaPelanggan string
	CreatedAt     string       `db:"created_at"`
	UpdatedAt     sql.NullTime `db:"updated_at"`
}
