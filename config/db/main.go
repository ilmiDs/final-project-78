package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "config/db/DL.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`

	CREATE TABLE IF NOT EXISTS user (
		id integer not null primary key,
		nim integer not null,
		username varchar(255) not null,
		email varchar(255) not null,
		gambar varchar(255) not null,
		falkultas varchar(255) not null,
		no_hp varchar(255) not null,
		kelamin boolean not null,
		password varchar(255) not null,
		role varchar(255) not null,
		logged_in boolean not null,
		created_at datetime not null,
		updated_at datetime not null
	);

	CREATE TABLE IF NOT EXISTS buku (
		id integer not null primary key,
		ISBN integer not null,
		pdf varchar(255) not null,
		cover varchar(255) not null,
		bahasa varchar(255) not null,
		jumlah halaman integer not null,
		tahun varchar(255) not null,
		judul varchar(255) not null,
		penulis varchar(255) not null,
		penerbit varchar(255) not null,
		kategori varchar(255) not null,
		deskripsi varchar(255) not null,_
		created_at datetime not null,
		updated_at datetime not null
	);

	CREATE TABLE IF NOT EXISTS pinjam (
		id integer not null primary key,
		username varchar(255) not null default '',
		nim integer not null default '',
		penulis varchar(255) not null default '',
		no_buku varchar(255) not null default '',
		forign key no buku references buku(id),
		forign key username references user(username),
		forign key nim references user(nim),
		forign key penulis references buku(penulis),
		created_at datetime not null,
		updated_at datetime not null
	);

	CREATE TABLE IF NOT EXISTS peminjaman (
		id integer not null primary key,
		id_user integer not null,
		id_buku integer not null,
		tanggal_pinjam datetime not null,
		tanggal_kembali datetime not null,
		forign key id_user references user(id),
		forign key id_buku references buku(id),
		created_at datetime not null,
		updated_at datetime not null
	);

	CREATE TABLE IF NOT EXISTS pengembalian (
		id integer not null primary key,
		id_user integer not null,
		id_buku integer not null,
		tanggal_pinjam datetime not null,
		tanggal_kembali datetime not null,
		forign key id_user references user(id),
		forign key id_buku references buku(id),
		created_at datetime not null,
		updated_at datetime not null
	);

	create into user(id, nim, username, email, gambar, falkultas, no_hp, kelamin, password, role, logged_in, created_at, updated_at) values
	(1, 123456789, 'admin', 'ilmi','ilmi@gmail.com', 'admin.jpg', 'Fakultas Teknik', '081234567890', true, 'admin', 'admin', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
	`)
	if err != nil {
		panic(err)
	}

	db.Close()
}
