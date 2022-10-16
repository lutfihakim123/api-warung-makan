package model

import "database/sql"

type Rincian struct {
	Id        string       `json:"id"`
	NotaId    string       `json:"nota_id" binding:"required"`
	MenuId    string       `json:"menu_id" binding:"required"`
	Kuantitas int          `json:"kuantitas" binding:"required"`
	CreatedAt string       `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
