package main

import "fmt"

func masuk(db *[100]Pengguna, penggunaAktif *Pengguna) {
	var (
		nEmail         string
		nKataSandi     string
		maxIndekTerisi int
	)

	fmt.Println("=== MASUK ===")
	fmt.Print("Email: ")
	fmt.Scanln(&nEmail)
	fmt.Print("Kata sandi: ")
	fmt.Scanln(&nKataSandi)

	maxIndekTerisi = cekIndeksDbPengguna(db)
	if maxIndekTerisi == -1 {
		maxIndekTerisi = len(db)
	}

	for i := 0; i < maxIndekTerisi; i++ {
		if nEmail == db[i].Email && nKataSandi == db[i].KataSandi {
			*penggunaAktif = db[i]
			return
		}
	}

	*penggunaAktif = Pengguna{}
}

func daftar(db *[100]Pengguna) {
	var (
		nNama          string
		nEmail         string
		nKataSandi     string
		nJenisPengguna string

		maxIdTerisi   int
		indekTersedia int
		idTersedia    int
	)

	indekTersedia = cekIndeksDbPengguna(db)

	if indekTersedia == -1 {
		fmt.Println("Database penuh, tidak bisa mendaftar.")
		fmt.Println()
		return
	}

	fmt.Println("=== DAFTAR ===")
	fmt.Print("Nama :")
	fmt.Scanln(&nNama)
	fmt.Print("Email :")
	fmt.Scanln(&nEmail)
	fmt.Print("Kata sandi: ")
	fmt.Scanln(&nKataSandi)
	fmt.Print("Jenis Pengguna [pendiri/investor] :")
	fmt.Scanln(&nJenisPengguna)

	if nNama == "" || nEmail == "" || nKataSandi == "" || nJenisPengguna == "" {
		fmt.Println("[Eror] Daftar gagal, semua form harus terisi!")
		fmt.Println()
		return
	}

	maxIdTerisi = cekIdPengguna(db)

	// Tentukan ID baru: Jika belum ada ID > 0 (maxIdTerisi == -1), ID pertama adalah 1.
	// Jika sudah ada, ID baru adalah ID terbesar + 1.
	if maxIdTerisi == -1 {
		idTersedia = 1 // ID pertama biasanya dimulai dari 1, bukan 0
	} else {
		idTersedia = maxIdTerisi + 1
	}

	db[indekTersedia] = Pengguna{
		IdPengguna:    idTersedia,
		Nama:          nNama,
		Email:         nEmail,
		KataSandi:     nKataSandi,
		JenisPengguna: nJenisPengguna,
	}

	// 7. Beri pesan sukses
	fmt.Println("Pendaftaran berhasil!")
	fmt.Printf("ID Pengguna Anda: %d\n", maxIdTerisi+1)
	fmt.Println()
}
