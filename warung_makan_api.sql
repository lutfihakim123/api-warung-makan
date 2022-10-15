-- DDL

CREATE TABLE mst_karyawan (
id_karyawan varchar(255) PRIMARY KEY,
nama VARCHAR(100),
alamat VARCHAR(100),
gaji integer,
username VARCHAR(100),
password VARCHAR(100)
);

CREATE TABLE mst_pelanggan (
id_pelanggan varchar(255) PRIMARY KEY,
nama VARCHAR(100),
nohp integer,
alamat VARCHAR(100)
);

CREATE TABLE mst_menu (
id_menu varchar(255) PRIMARY KEY,
nama VARCHAR(100),
jenis VARCHAR(100),
harga integer,
stock integer
);

CREATE TABLE rincian (
id_rincian varchar(255) PRIMARY KEY,
nota_id varchar(255),
menu_id varchar(255),
kuantitas integer
);

CREATE TABLE nota_penjualan (
id_nota varchar(255) PRIMARY KEY,
pelanggan_id VARCHAR(255),
karyawan_id varchar(255),
waktu_pesan date,
total integer
);

-- foreign KEY

ALTER TABLE rincian
ADD CONSTRAINT fk_rincian_nota FOREIGN KEY (nota_id) REFERENCES nota_penjualan (id_nota);

ALTER TABLE rincian
ADD CONSTRAINT fk_rincian_menu FOREIGN KEY (menu_id) REFERENCES mst_menu (id_menu);

ALTER TABLE nota_penjualan
ADD CONSTRAINT fk_nota_pelanggan FOREIGN KEY (pelanggan_id) REFERENCES mst_pelanggan (id_pelanggan);

ALTER TABLE nota_penjualan
ADD CONSTRAINT fk_nota_karyawan FOREIGN KEY (karyawan_id) REFERENCES mst_karyawan (id_karyawan);

-- end foreign key



