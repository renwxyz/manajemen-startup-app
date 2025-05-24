package main

import "fmt"

func masuk(db *[100]Pengguna, penggunaAktif *Pengguna) {
	var (
		nEmail         string
		nKataSandi     string
		maxIndekTerisi int
	)

	//Kredensial pengguna akan ditangkap dan dimasukkan kedalam variabel
	fmt.Println("=== MASUK ===")
	fmt.Print("Email: ")
	fmt.Scanln(&nEmail)
	fmt.Print("Kata sandi: ")
	fmt.Scanln(&nKataSandi)

	maxIndekTerisi = cekIndeksDbPengguna(db) // dicek terlebih dahulu maksimal indeks terisi terkini
	if maxIndekTerisi == -1 {
		maxIndekTerisi = len(db)
	}

	// maxIndekTerisi digunakan sebagai batas iterasi pengecekan kredensial supaya tidak perlu dicek hingga maksimal ukuran array
	for i := 0; i < maxIndekTerisi; i++ {
		if nEmail == db[i].Email && nKataSandi == db[i].KataSandi {
			*penggunaAktif = db[i] //jika ditemukan maka akan mengisi alamat dari pointer pengguna aktif untuk digunakan data nya nanti
			return                 // berhenti tepat setelah ketemu
		}
	}

	*penggunaAktif = Pengguna{} //jika tidak ada kredensial yang cocok maka akan mengisi alamat pointer dengan himpunan kosong
}
