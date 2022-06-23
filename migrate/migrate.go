package migration

import (
	"database/sql"
	
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/models"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/structs"
)

func MigratekeDB(db *sql.DB) {
	models.MigratekeDB(db)
}
