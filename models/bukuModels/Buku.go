package BukuModels

import ""

type Buku struct {struct.Buku}

//get all buku
func (b *Buku) GetAllBuku() ([]struct.Buku, error) {
	db, err := sql.Open("sqlite3", "config/db/DL.db")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM buku")
	if err != nil {
		panic(err)
	}

	var buku []struct.Buku
	for rows.Next() {

		var b struct.Buku
		err = rows.Scan(&b.ID, &b.ISBN, &b.PDF, &b.Cover, &b.Bahasa, &b.JumlahHalaman, &b.Tahun, &b.Judul, &b.Penulis, &b.Penerbit, &b.Kategori, &b.Deskripsi, &b.CreatedAt, &b.UpdatedAt)
		if err != nil {
			panic(err)
		}
		buku = append(buku, b)
	}
	return buku, nil
}

//get buku by id
func (b *Buku) GetBukuById(id int) (struct.Buku, error) {
	db, err := sql.Open("sqlite3", "config/db/DL.db")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM buku WHERE id = ?", id)
	if err != nil {
		panic(err)
	}

	var buku struct.Buku
	for rows.Next() {

		err = rows.Scan(&buku.ID, &buku.ISBN, &buku.PDF, &buku.Cover, &buku.Bahasa, &buku.JumlahHalaman, &buku.Tahun, &buku.Judul, &buku.Penulis, &buku.Penerbit, &buku.Kategori, &buku.Deskripsi, &buku.CreatedAt, &buku.UpdatedAt)
		if err != nil {
			panic(err)
		}
	}
	return buku, nil
}

//create buku
func (b *Buku) CreateBuku(buku struct.Buku) error {
	db, err := sql.Open("sqlite3", "config/db/DL.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO buku (isbn, pdf, cover, bahasa, jumlah_halaman, tahun, judul, penulis, penerbit, kategori, deskripsi, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", buku.ISBN, buku.PDF, buku.Cover, buku.Bahasa, buku.JumlahHalaman, buku.Tahun, buku.Judul, buku.Penulis, buku.Penerbit, buku.Kategori, buku.Deskripsi, buku.CreatedAt, buku.UpdatedAt)
	if err != nil {
		panic(err)
	}
	return nil
}

//upddate buku
func (b *Buku) UpdateBuku(buku struct.Buku) error {
	db, err := sql.Open("sqlite3", "config/db/DL.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("UPDATE buku SET isbn = ?, pdf = ?, cover = ?, bahasa = ?, jumlah_halaman = ?, tahun = ?, judul = ?, penulis = ?, penerbit = ?, kategori = ?, deskripsi = ?, created_at = ?, updated_at = ? WHERE id = ?", buku.ISBN, buku.PDF, buku.Cover, buku.Bahasa, buku.JumlahHalaman, buku.Tahun, buku.Judul, buku.Penulis, buku.Penerbit, buku.Kategori, buku.Deskripsi, buku.CreatedAt, buku.UpdatedAt, buku.ID)
	if err != nil {
		panic(err)
	}
	return nil
}

//delete buku
func (b *Buku) DeleteBuku(id int) error {
	db, err := sql.Open("sqlite3", "config/db/DL.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DELETE FROM buku WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	return nil
}

