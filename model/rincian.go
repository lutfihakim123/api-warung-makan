package model

import (
	"database/sql"
	"time"
)

type Rincian struct {
	Id            string       `json:"id"`
	NotaId        string       `json:"nota_id" db:"nota_id" binding:"required"`
	MenuId        string       `json:"menu_id" db:"menu_id" binding:"required"`
	NamaMenu      string       `json:"nama_menu" db:"nama_menu"`
	NoMeja        string       `json:"no_meja" db:"no_meja"`
	Harga         int          `json:"harga_menu" db:"harga_menu"`
	Img           string       `json:"image_menu" db:"image_menu"`
	WaktuPesan    time.Time    `json:"waktu_pesan" db:"waktu_pesan"`
	NamaPelanggan string       `json:"nama_pelanggan" db:"nama_pelanggan"`
	NamaKaryawan  string       `json:"nama_karyawan" db:"nama_karyawan"`
	TempStock     int          `json:"temp_stock" db:"temp_stock"`
	Kuantitas     int          `json:"kuantitas" binding:"required"`
	Total         int          `json:"total" db:"total"`
	CreatedAt     string       `db:"created_at"`
	UpdatedAt     sql.NullTime `db:"updated_at"`
}
