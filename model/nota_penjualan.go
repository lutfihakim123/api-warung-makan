package model

import (
	"database/sql"
	"time"
)

type Nota struct {
	Id          string       `json:"id"`
	PelangganId string       `json:"pelanggan_id" binding:"required"`
	KaryawanId  string       `json:"karyawan_id" binding:"required"`
	WaktuPesan  time.Time    `json:"waktu_pesan" binding:"required"`
	Total       int          `json:"total" binding:"required"`
	CreatedAt   string       `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}
