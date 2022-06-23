package struct

import "time"

type Buku struct {
	ID          int       `db:"id" json:"id"`
	ISBN        int       `db:"isbn" json:"isbn"`
	PDF         string    `db:"pdf" json:"pdf"`
	Cover       string    `db:"cover" json:"cover"`
	Bahasa      string    `db:"bahasa" json:"bahasa"`
	JumlahHalaman int       `db:"jumlah_halaman" json:"jumlah_halaman"`
	Tahun       int       `db:"tahun" json:"tahun"`
	Judul       string    `db:"judul" json:"judul"`
	Penulis     string    `db:"penulis" json:"penulis"`
	Penerbit    string    `db:"penerbit" json:"penerbit"`
	Kategori    string    `db:"kategori" json:"kategori"`
	Deskripsi   string    `db:"deskripsi" json:"deskripsi"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}