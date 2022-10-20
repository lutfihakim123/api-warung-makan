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
	InsertNota       = `insert into nota_penjualan(id, pelanggan_id, karyawan_id, meja_id, waktu_pesan) values (:id, :pelanggan_id, :karyawan_id, :meja_id, :waktu_pesan)`
	SelectAllNota    = `select id, pelanggan_id, karyawan_id, meja_id, waktu_pesan , created_at, updated_at from nota_penjualan limit $1 offset $2`
	SelectNotaById   = `select id, pelanggan_id, karyawan_id, meja_id, waktu_pesan , created_at, updated_at from nota_penjualan where id=$1`
	UpdateNota       = `update nota_penjualan set pelanggan_id=:pelanggan_id,  karyawan_id=:karyawan_id, meja_id=:meja_id where id=:id`
	DeleteNota       = `delete from nota_penjualan where id=$1`
	UpdateStatusMeja = `update mst_meja set status='used' where id=:meja_id`

	//rincian
	UpdateStockMenu  = `update mst_menu set stock=:temp_stock where id=:menu_id`
	InsertRincian    = `insert into rincian(id, nota_id, menu_id, kuantitas, total) values (:id, :nota_id, :menu_id, :kuantitas, :total)`
	UpdateRincian    = `update rincian set nota_id=:nota_id,  menu_id=:menu_id, kuantitas=:kuantitas, total=:total where id=:id`
	DeleteRincian    = `delete from rincian where id=$1`
	SelectAllRincian = `SELECT rincian.Id as id, rincian.nota_id as nota_id, mst_meja.no_meja as no_meja, rincian.menu_id as menu_id, mst_pelanggan.nama as nama_pelanggan, mst_karyawan.nama as nama_karyawan, mst_menu.nama as nama_menu, mst_menu.harga as harga_menu, rincian.kuantitas as kuantitas, nota_penjualan.waktu_pesan as waktu_pesan,
	rincian.total as total, mst_menu.img as image_menu, rincian.created_at as created_at, rincian.updated_at from mst_menu inner join rincian on mst_menu.id=rincian.menu_id inner join nota_penjualan ON rincian.nota_id=nota_penjualan.id  inner join mst_pelanggan ON mst_pelanggan.id = nota_penjualan.pelanggan_id 
	inner join mst_karyawan ON mst_karyawan.id = nota_penjualan.karyawan_id inner join mst_meja on mst_meja.id=nota_penjualan.meja_id limit $1 offset $2`
	SelectRincianById = `SELECT rincian.Id as id, rincian.nota_id as nota_id, mst_meja.no_meja as no_meja, rincian.menu_id as menu_id, mst_pelanggan.nama as nama_pelanggan, mst_karyawan.nama as nama_karyawan, mst_menu.nama as nama_menu, mst_menu.harga as harga_menu, rincian.kuantitas as kuantitas, nota_penjualan.waktu_pesan as waktu_pesan,
	rincian.total as total, mst_menu.img as image_menu, rincian.created_at as created_at, rincian.updated_at from mst_menu inner join rincian on mst_menu.id=rincian.menu_id inner join nota_penjualan ON rincian.nota_id=nota_penjualan.id  inner join mst_pelanggan ON mst_pelanggan.id = nota_penjualan.pelanggan_id 
	inner join mst_karyawan ON mst_karyawan.id = nota_penjualan.karyawan_id inner join mst_meja on mst_meja.id=nota_penjualan.meja_id where rincian.id=$1`

	//report
	SelectAllReport  = `SELECT nota_penjualan.Id as id, mst_pelanggan.id as pelanggan_id, mst_karyawan.id as karyawan_id, mst_meja.id as meja_id,  mst_meja.no_meja as no_meja, rincian.menu_id as menu_id, mst_pelanggan.nama as nama_pelanggan, mst_karyawan.nama as nama_karyawan, mst_menu.nama as nama_menu, mst_menu.harga as harga_menu, rincian.kuantitas as kuantitas, nota_penjualan.waktu_pesan as waktu_pesan, rincian.total as total, mst_menu.img as image_menu, nota_penjualan.created_at, nota_penjualan.updated_at from mst_menu inner join rincian on mst_menu.id=rincian.menu_id right join nota_penjualan ON rincian.nota_id=nota_penjualan.id  inner join mst_pelanggan ON mst_pelanggan.id = nota_penjualan.pelanggan_id inner join mst_karyawan ON mst_karyawan.id = nota_penjualan.karyawan_id inner join mst_meja on mst_meja.id=nota_penjualan.meja_id limit $1 offset $2`
	SelectReportById = `SELECT nota_penjualan.Id as id, mst_pelanggan.id as pelanggan_id, mst_karyawan.id as karyawan_id, mst_meja.id as meja_id,  mst_meja.no_meja as no_meja, rincian.menu_id as menu_id, mst_pelanggan.nama as nama_pelanggan, mst_karyawan.nama as nama_karyawan, mst_menu.nama as nama_menu, mst_menu.harga as harga_menu, rincian.kuantitas as kuantitas, nota_penjualan.waktu_pesan as waktu_pesan, rincian.total as total, mst_menu.img as image_menu, nota_penjualan.created_at, nota_penjualan.updated_at from mst_menu inner join rincian on mst_menu.id=rincian.menu_id right join nota_penjualan ON rincian.nota_id=nota_penjualan.id  inner join mst_pelanggan ON mst_pelanggan.id = nota_penjualan.pelanggan_id inner join mst_karyawan ON mst_karyawan.id = nota_penjualan.karyawan_id inner join mst_meja on mst_meja.id=nota_penjualan.meja_id where nota_penjualan.id=$1`
)
