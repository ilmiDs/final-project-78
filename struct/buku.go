package struct

import "time"

type Buku struct {
	ID          int       `json:"id"`
	ISBN        int       `json:"isbn"`
	PDF         string    `json:"pdf"`
	Cover       string    `json:"cover"`
	Bahasa      string    `json:"bahasa"`
	JumlahHalaman int       `json:"jumlah_halaman"`
	Tahun        string    `json:"tahun"`
	Judul        string    `json:"judul"`
	Penulis      string    `json:"penulis"`
	Penerbit     string    `json:"penerbit"`
	Kategori     string    `json:"kategori"`
	Deskripsi    string    `json:"deskripsi"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}