-- DDL

CREATE TABLE mst_karyawan (
id varchar(255) PRIMARY KEY,
nama VARCHAR(100),
alamat VARCHAR(100),
gaji integer,
username VARCHAR(100),
password VARCHAR(100),
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
updated_at timestamp
);

CREATE TABLE mst_pelanggan (
id varchar(255) PRIMARY KEY,
nama VARCHAR(100),
nohp integer,
alamat VARCHAR(100),
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
updated_at timestamp
);

CREATE TABLE mst_menu (
id varchar(255) PRIMARY KEY,
nama VARCHAR(100),
jenis VARCHAR(100),
harga integer,
stock integer,
img    varchar(100),
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
updated_at timestamp
);

CREATE TABLE mst_meja (
id varchar(255) PRIMARY KEY,
no_meja VARCHAR(100),
status VARCHAR(100),
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
updated_at timestamp
);


CREATE TABLE rincian (
id varchar(255) PRIMARY KEY,
nota_id varchar(255),
menu_id varchar(255),
kuantitas integer,
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
updated_at timestamp
);

CREATE TABLE nota_penjualan (
id varchar(255) PRIMARY KEY,
pelanggan_id VARCHAR(255),
karyawan_id varchar(255),
meja_id varchar(255),
waktu_pesan date,
total integer,
created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
updated_at timestamp
);

-- foreign KEY

ALTER TABLE rincian
ADD CONSTRAINT fk_rincian_nota FOREIGN KEY (nota_id) REFERENCES nota_penjualan (id);

ALTER TABLE rincian
ADD CONSTRAINT fk_rincian_menu FOREIGN KEY (menu_id) REFERENCES mst_menu (id);

ALTER TABLE nota_penjualan
ADD CONSTRAINT fk_nota_pelanggan FOREIGN KEY (pelanggan_id) REFERENCES mst_pelanggan (id);

ALTER TABLE nota_penjualan
ADD CONSTRAINT fk_nota_karyawan FOREIGN KEY (karyawan_id) REFERENCES mst_karyawan (id);

ALTER TABLE nota_penjualan
ADD CONSTRAINT fk_nota_meja FOREIGN KEY (meja_id) REFERENCES mst_meja (id);

-- end foreign key



