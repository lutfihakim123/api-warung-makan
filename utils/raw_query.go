package utils

const (
	InsertMenu     = `insert into mst_menu(id, nama, jenis, harga, stock, img) values (:id, :nama, :jenis, :harga, :stock, :img)`
	SelectAllMenu  = `select id,  nama, jenis, harga, stock, img, created_at, updated_at  from mst_menu limit $1 offset $2`
	SelectMenuById = `select id, nama, jenis, harga, stock, img, created_at, updated_at from mst_menu where id=$1`
	UpdateMenu     = `update mst_menu set nama=:nama, jenis=:jenis, harga=:harga, stock=:stock where id=:id`
	DeleteMenu     = `delete from mst_menu where id=$1`
)
