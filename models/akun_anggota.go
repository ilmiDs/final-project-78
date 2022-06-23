package models

import (
	"strings"
	"time"
	"database/sql"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config/helper"
	"github.com/badoux/checkmail"
	"github.com/mattn/go-sqlite3"
)

	type Anggota struct {
		IDAnggota int `db:"id" json:"id"`
		NIM       int    `db:"nim" json:"nim"`
		Username string `db:"username" json:"username"`
		email string `db:"email" json:"email"`
		gambar string `db:"gambar" json:"gambar"`
		Fakultas string `db:"fakultas" json:"fakultas"`
		no_hp    string    `db:"no_hp" json:"no_hp"`
		kelamin string `db:"kelamin" json:"kelamin"`
		password string `db:"password" json:"password"`
		role string `db:"role" json:"role"`
		logged_in bool `db:"logged_in" json:"logged_in"`
		created_at time.Time `db:"created_at" json:"created_at"`
		updated_at time.Time `db:"updated_at" json:"updated_at"`
	}


	func (u *Anggota) Persiapan(action string) {
		u.IDAnggota = 0
		u.Username = helper.EscapeStrings(u.Username)
		u.kelamin = helper.EscapeStrings(u.kelamin)
		if strings.ToLower(action) == "seed" || strings.ToLower(action) == "tambah" || strings.ToLower(action) == "update" {
			u.password = helper.Hash(u.password)
		}
	}

	func (u *Anggota) Validate(aksi string) error {
		switch strings.ToLower(aksi) {
		case "tambah":
			if u.NIM == 0 {
				return helper.ErrorMessage("NIM tidak boleh kosong")
			}
			if u.Username == "" {
				return helper.ErrorMessage("Username tidak boleh kosong")
			}
			if u.email == "" {
				return helper.ErrorMessage("Email tidak boleh kosong")
			}
			if u.Fakultas == "" {
				return helper.ErrorMessage("Fakultas tidak boleh kosong")
			}
			if u.no_hp == "" {
				return helper.ErrorMessage("No HP tidak boleh kosong")
			}
			if u.kelamin == "" {
				return helper.ErrorMessage("Kelamin tidak boleh kosong")
			}
			if u.password == "" {
				return helper.ErrorMessage("Password tidak boleh kosong")
			}
			if u.gambar == "" {
				return helper.ErrorMessage("Gambar tidak boleh kosong")
			}
			if u.role == "" {
				return helper.ErrorMessage("Role tidak boleh kosong")
			}
			if u.logged_in == false {
				return helper.ErrorMessage("Logged in tidak boleh kosong")
			}
	
		case "update":
			if u.IDAnggota == 0 {
				return helper.ErrorMessage("ID Anggota tidak boleh kosong")
			}
			if u.NIM == 0 {
				return helper.ErrorMessage("NIM tidak boleh kosong")
			}
			if u.Username == "" {
				return helper.ErrorMessage("Username tidak boleh kosong")
			}
			if u.email == "" { 
				return helper.ErrorMessage("Email tidak boleh kosong")
			}
			if u.Fakultas == "" {
				return helper.ErrorMessage("Fakultas tidak boleh kosong")
			}
			if u.no_hp == "" {
				return helper.ErrorMessage("No HP tidak boleh kosong")
			}
			if u.kelamin == "" {
				return helper.ErrorMessage("Kelamin tidak boleh kosong")
			}
			if u.password == "" {
				return helper.ErrorMessage("Password tidak boleh kosong")
			}
			if u.role == "" {
				return helper.ErrorMessage("Role tidak boleh kosong")
			}
			if u.logged_in == false {
				return helper.ErrorMessage("Logged in tidak boleh kosong")
			}
		case "login":
			if u.Username == "" {
				return helper.ErrorMessage("Username tidak boleh kosong")
			}
			if u.password == "" {
				return helper.ErrorMessage("Password tidak boleh kosong")
			}
		}
		return nil
	}

 // fungsi untuk menambah data anggota
	func (u *Anggota) Register() error {
		db, err := sql.Open("sqlite3", "./db/perpus.db")
		if err != nil {
			return err
		}
		defer db.Close()
		err = u.Validate("tambah")
		if err != nil {
			return err
		}
		u.Persiapan("tambah")
		statement, err := db.Prepare("INSERT INTO anggota (nim, username, email, fakultas, no_hp, kelamin, password, gambar, role, logged_in, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			return err
		}
		defer statement.Close()
		_, err = statement.Exec(u.NIM, u.Username, u.email, u.Fakultas, u.no_hp, u.kelamin, u.password, u.gambar, u.role, u.logged_in, u.created_at, u.updated_at)
		if err != nil {
			return err
		}
		return nil
	}

	//fungsi untuk logout
	func (u *Anggota) Logout() error {
		db, err := sql.Open("sqlite3", "./db/perpus.db")
		if err != nil {
			return err
		}
		defer db.Close()
		err = u.Validate("login")
		if err != nil {
			return err
		}
		u.Persiapan("login")
		statement, err := db.Prepare("UPDATE anggota SET logged_in = ? WHERE username = ?")
		if err != nil {
			return err
		}
		defer statement.Close()
		_, err = statement.Exec(u.logged_in, u.Username)
		if err != nil {
			return err
		}
		return nil
	}
	
