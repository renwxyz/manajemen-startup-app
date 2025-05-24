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
