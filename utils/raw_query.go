package utils

const (
	// menu
	InsertMenu     = `insert into mst_menu(id, nama, jenis, harga, stock, img) values (:id, :nama, :jenis, :harga, :stock, :img)`
	SelectAllMenu  = `select id,  nama, jenis, harga, stock, img, created_at, updated_at  from mst_menu limit $1 offset $2`
	SelectMenuById = `select id, nama, jenis, harga, stock, img, created_at, updated_at from mst_menu where id=$1`
	UpdateMenu     = `update mst_menu set nama=:nama, jenis=:jenis, harga=:harga, stock=:stock where id=:id`
	DeleteMenu     = `delete from mst_menu where id=$1`

	// karyawan
	InsertKaryawan     = `insert into mst_karyawan(id, nama, alamat, gaji, username, password) values (:id, :nama, :alamat, :gaji, :username, :password)`
	SelectAllKaryawan  = ` select id, nama, alamat, gaji, username, password, created_at, updated_at  from mst_karyawan limit $1 offset $2`
	SelectKaryawanById = ` select id, nama, alamat, gaji, username, password, created_at, updated_at  from mst_karyawan where id=$1`
	UpdateKaryawan     = `update mst_karyawan set nama=:nama, alamat=:alamat, gaji=:gaji, username=:username, password=:password where id=:id`
	DeleteKaryawan     = `delete from mst_karyawan where id=$1`
)
