package utils

const (
	// menu
	InsertMenu     = `insert into mst_menu(id, nama, jenis, harga, stock, img) values (:id, :nama, :jenis, :harga, :stock, :img)`
	SelectAllMenu  = `select id,  nama, jenis, harga, stock, img, created_at, updated_at  from mst_menu limit $1 offset $2`
	SelectMenuById = `select id, nama, jenis, harga, stock, img, created_at, updated_at from mst_menu where id=$1`
	UpdateMenu     = `update mst_menu set nama=:nama, jenis=:jenis, harga=:harga, img=:img,  stock=:stock where id=:id`
	DeleteMenu     = `delete from mst_menu where id=$1`

	// karyawan
	InsertKaryawan     = `insert into mst_karyawan(id, nama, alamat, gaji, username, password) values (:id, :nama, :alamat, :gaji, :username, :password)`
	SelectAllKaryawan  = ` select id, nama, alamat, gaji, username, password, created_at, updated_at  from mst_karyawan limit $1 offset $2`
	SelectKaryawanById = ` select id, nama, alamat, gaji, username, password, created_at, updated_at  from mst_karyawan where id=$1`
	UpdateKaryawan     = `update mst_karyawan set nama=:nama, alamat=:alamat, gaji=:gaji, username=:username, password=:password where id=:id`
	DeleteKaryawan     = `delete from mst_karyawan where id=$1`

	// pelanggan
	InsertPelanggan     = `insert into mst_pelanggan(id, nama, nohp, alamat) values (:id, :nama, :nohp, :alamat)`
	SelectAllPelanggan  = `select id, nama, nohp, alamat, created_at, updated_at from mst_pelanggan limit $1 offset $2`
	SelectPelangganById = `select id, nama, nohp, alamat, created_at, updated_at from mst_pelanggan where id=$1`
	UpdatePelanggan     = `update mst_pelanggan set id=:id, nama=:nama, nohp=:nohp, alamat=:alamat where id=:id`
	DeletePelanggan     = `delete from mst_pelanggan where id=$1`

	// meja
	InsertMeja     = `insert into mst_meja(id, no_meja, status) values (:id, :nomeja, :status)`
	SelectAllMeja  = `select id, no_meja as nomeja, status, created_at, updated_at from mst_meja limit $1 offset $2`
	SelectMejaById = `select id, no_meja as nomeja, status, created_at, updated_at from mst_meja where id=$1`
	UpdateMeja     = `update mst_meja set id=:id, no_meja=:nomeja, status=:status where id=:id`
	DeleteMeja     = `delete from mst_meja where id=$1`

	//nota penjualan
	InsertNota     = `insert into nota_penjualan(id, pelanggan_id, karyawan_id, meja_id, waktu_pesan) values (:id, :pelangganid, :karyawanid, :mejaid, :waktupesan)`
	SelectAllNota  = `select nota_penjualan.id, nota_penjualan.pelanggan_id, nota_penjualan.karyawan_id, nota_penjualan.waktu_pesan, menu.harga, rincian.total  created_at, updated_at from nota_penjualan inner join rincian on nota_penjualan.id=rincian.nota_id inner join menu on rincian.menu_id=menu.id limit $1 offset $2`
	SelectNotaById = `select nota_penjualan.id, nota_penjualan.pelanggan_id, nota_penjualan.karyawan_id, nota_penjualan.waktu_pesan, menu.harga, rincian.total  created_at, updated_at from nota_penjualan inner join rincian on nota_penjualan.id=rincian.nota_id inner join menu on rincian.menu_id=menu.id where nota_penjualan.id=$1`
	UpdateNota     = `update nota_penjualan set pelanggan_id=:pelangganid,  karyawan_id=:karyawanid, meja_id=:mejaid where id=:id`
	DeleteNota     = `delete from nota_penjualan where id=$1`
)
