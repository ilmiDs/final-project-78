package seed

import (
	"log"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config/helper"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/models"
)

var admin = []models.Pegawai{
	models.Pegawai{
		Username: "admin",
		Password: helper.Hash("admin"),
	},
	models.Pegawai{
		Username: "afrizal",
		Password: helper.Hash("admin"),
	},
}

func Seed_admin() {
	db := config.Connect()
	defer db.Close()

	for _, v := range admin {
		_, err := db.Exec("INSERT INTO pegawai(username, password) VALUES(?,?)", v.Username, v.Password)
		if err != nil {
			log.Println(err)
		}
	}
}
