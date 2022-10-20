package model

import (
	"database/sql"
	"time"
)

type Nota struct {
	Id          string       `json:"id"`
	PelangganId string       `json:"pelanggan_id" db:"pelanggan_id" binding:"required"`
	KaryawanId  string       `json:"karyawan_id" db:"karyawan_id" binding:"required"`
	MejaId      string       `json:"meja_id" db:"meja_id" binding:"required"`
	WaktuPesan  time.Time    `json:"waktu_pesan" db:"waktu_pesan"`
	CreatedAt   string       `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}

type ReportPenjualan struct {
	Id            string       `json:"id"`
	PelangganId   string       `json:"pelanggan_id" db:"pelanggan_id"`
	KaryawanId    string       `json:"karyawan_id" db:"karyawan_id"`
	MejaId        string       `json:"meja_id" db:"meja_id"`
	MenuId        string       `json:"menu_id" db:"menu_id"`
	NamaMenu      string       `json:"nama_menu" db:"nama_menu"`
	NoMeja        string       `json:"no_meja" db:"no_meja"`
	Harga         int          `json:"harga_menu" db:"harga_menu"`
	Img           string       `json:"image_menu" db:"image_menu"`
	WaktuPesan    time.Time    `json:"waktu_pesan" db:"waktu_pesan"`
	NamaPelanggan string       `json:"nama_pelanggan" db:"nama_pelanggan"`
	NamaKaryawan  string       `json:"nama_karyawan" db:"nama_karyawan"`
	TempStock     int          `json:"temp_stock" db:"temp_stock"`
	Kuantitas     int          `json:"kuantitas"`
	Total         int          `json:"total" db:"total"`
	CreatedAt     string       `db:"created_at"`
	UpdatedAt     sql.NullTime `db:"updated_at"`
}
