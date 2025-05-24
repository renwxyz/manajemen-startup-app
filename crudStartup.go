package main

import "fmt"

func tambahStartup(penggunaAktif *Pengguna, db *[100]Startup) *Startup {
	var (
		namaStartup    string
		bidangUsaha    string
		tahunBerdiri   int
		totalPendanaan int

		maxIdTerisi   int
		indekTersedia int
		idTersedia    int
	)

	indekTersedia = cekIndeksDbStartup(db)

	if indekTersedia == -1 {
		fmt.Println("Database penuh.")
		fmt.Println()
		return nil
	}
	fmt.Println("=== Form Pendaftaran Startup ===")
	fmt.Print("Nama Startup: ")
	fmt.Scanln(&namaStartup)
	fmt.Print("Bidang Usaha: ")
	fmt.Scanln(&bidangUsaha)
	fmt.Print("Tahun Berdiri: ")
	fmt.Scanln(&tahunBerdiri)
	fmt.Print("Total Pendanaan: ")
	fmt.Scanln(&totalPendanaan)

	if namaStartup == "" || bidangUsaha == "" || tahunBerdiri == 0 || totalPendanaan == 0 {
		fmt.Println("Daftar gagal, semua form harus terisi!")
		fmt.Println()
		return nil
	}

	maxIdTerisi = cekIdStartup(db)

	// Tentukan ID baru: Jika belum ada ID > 0 (maxIdTerisi == -1), ID pertama adalah 1.
	// Jika sudah ada, ID baru adalah ID terbesar + 1.
	if maxIdTerisi == -1 {
		idTersedia = 1 // ID pertama biasanya dimulai dari 1, bukan 0
	} else {
		idTersedia = maxIdTerisi + 1
	}

	db[indekTersedia] = Startup{
		IdStartup:      idTersedia,
		IdPendiri:      penggunaAktif.IdPengguna,
		NamaStartup:    namaStartup,
		BidangUsaha:    bidangUsaha,
		TahunBerdiri:   tahunBerdiri,
		TotalPendanaan: totalPendanaan,
	}
	return &db[indekTersedia]
}

func tambahTim(startup *Startup) *Tim {
	var (
		namaTim       string
		indekTersedia int
	)
	indekTersedia = cekIndekTim(&startup.DaftarTim)

	if indekTersedia == -1 {
		fmt.Println("Database penuh, tidak bisa mendaftar.")
		fmt.Println()
		return nil
	}

	fmt.Print("Nama Tim: ")
	fmt.Scanln(&namaTim)

	startup.DaftarTim[indekTersedia].NamaTim = namaTim

	return &startup.DaftarTim[indekTersedia]
}

func tambahAnggota(tim *Tim) {
	var (
		namaAnggota   string
		posisi        string
		indekTersedia int
	)

	indekTersedia = cekIndekAnggota(&tim.DaftarAnggota)

	if indekTersedia == -1 {
		fmt.Println("Database penuh, tidak bisa mendaftar.")
		fmt.Println()
		return
	}

	fmt.Print("Nama anggota: ")
	fmt.Scanln(&namaAnggota)
	fmt.Print("Posisi: ")
	fmt.Scanln(&posisi)

	tim.DaftarAnggota[indekTersedia].NamaAnggota = namaAnggota
	tim.DaftarAnggota[indekTersedia].Posisi = posisi

}

func tampilkanStartup(db *[100]Startup, i, count int) {
	fmt.Printf(" %-15d %-15d %-25s %-30s %-20d %-20d \n", count, db[i].IdStartup, db[i].NamaStartup, db[i].BidangUsaha, db[i].TahunBerdiri, db[i].TotalPendanaan)
}

func detailStartup(db *[100]Startup, i int) {
	fmt.Printf("%-18s: %s\n", "Nama startup", db[i].NamaStartup)
	fmt.Printf("%-18s: %d\n", "ID startup", db[i].IdStartup)
	fmt.Printf("%-18s: %d\n", "ID pendiri", db[i].IdPendiri)
	fmt.Printf("%-18s: %s\n", "Bidang usaha", db[i].BidangUsaha)
	fmt.Printf("%-18s: %d\n", "Tahun berdiri", db[i].TahunBerdiri)
	fmt.Printf("%-18s: %d\n", "Total pendanaan", db[i].TotalPendanaan)
	fmt.Println()

	fmt.Println("========== ===== Daftar Tim ===== ==========")
	var indekTim int = cekIndekTim(&db[i].DaftarTim)
	for j := 0; j < indekTim; j++ {
		tampilkanTim(&db[i].DaftarTim, j)
	}
	fmt.Println("--------------------------------------------")
}

func hapusStartupSaya(penggunaAktif *Pengguna, db *[100]Startup) {
	var namaStartup string
	fmt.Print("Nama startup yang ingin dihapus: ")
	fmt.Scanln(&namaStartup)

	indeksTersedia := cekIndeksDbStartup(db)
	for i := 0; i < indeksTersedia; i++ {
		if db[i].NamaStartup == namaStartup && db[i].IdPendiri == penggunaAktif.IdPengguna {
			for j := i; j < indeksTersedia-1; j++ {
				db[j] = db[j+1]
			}
			db[indeksTersedia-1] = Startup{}
			fmt.Println("Startup berhasil dihapus.")
			return
		}
	}
	fmt.Println("Startup tidak ditemukan.")
}

func tampilkanTim(db *[100]Tim, j int) {
	fmt.Printf(">>> Tim: %s\n", db[j].NamaTim)

	var indekAnggota int = cekIndekAnggota(&db[j].DaftarAnggota)
	for k := 0; k < indekAnggota; k++ {
		tampilkanAnggota(&db[j].DaftarAnggota, k)
	}
}

func tampilkanAnggota(db *[100]Anggota, k int) {
	fmt.Printf("	- %-10s | %s\n", db[k].NamaAnggota, db[k].Posisi)
}

func ubahIdPendiri(db *[100]Startup, indeks int) {
	var nIdPendiri int
	fmt.Print("Masukan id pendiri baru: ")
	fmt.Scanln(&nIdPendiri)
	db[indeks].IdPendiri = nIdPendiri
	fmt.Println("ID Pendiri berhasil diubah.")
}

func ubahNamaStartup(db *[100]Startup, indeks int) {
	var nNamaStartup string
	fmt.Print("Masukan nama startup baru: ")
	fmt.Scanln(&nNamaStartup)
	db[indeks].NamaStartup = nNamaStartup
	fmt.Println("Nama Startup berhasil diubah.")
}

func ubahBidangUsaha(db *[100]Startup, indeks int) {
	var nBidangUsaha string
	fmt.Print("Masukan kategori bidang usaha baru: ")
	fmt.Scanln(&nBidangUsaha)
	db[indeks].BidangUsaha = nBidangUsaha
	fmt.Println("Bidang Usaha berhasil diubah.")
}

func ubahTotalPendanaan(db *[100]Startup, indeks int) {
	var nTotalPendanaan int
	fmt.Print("Masukan total pendanaan baru: ")
	fmt.Scanln(&nTotalPendanaan)
	db[indeks].TotalPendanaan = nTotalPendanaan
	fmt.Println("Total Pendanaan berhasil diubah.")
}

func ubahNamaTim(tim *Tim) {
	var nBaruNamaTim string
	fmt.Print("Nama tim baru: ")
	fmt.Scanln(&nBaruNamaTim)
	tim.NamaTim = nBaruNamaTim
}

func ubahAnggota(tim *Tim) {
	var pilihNamaAnggota string
	var indekAnggota int = cekIndekAnggota(&tim.DaftarAnggota)
	var indeksDiubah int

	fmt.Print("Pilih anggota: ")
	fmt.Scanln(&pilihNamaAnggota)
	for i := 0; i < indekAnggota; i++ {
		if pilihNamaAnggota == tim.DaftarAnggota[i].NamaAnggota {
			indeksDiubah = i
		}
	}

	var nNamaAnggota string

	fmt.Print("Anggota baru: ")
	fmt.Scanln(&nNamaAnggota)
	tim.DaftarAnggota[indeksDiubah].NamaAnggota = nNamaAnggota

	var nPosisi string
	fmt.Print("Posisi baru: ")
	fmt.Scanln(&nNamaAnggota)
	tim.DaftarAnggota[indeksDiubah].Posisi = nPosisi

}

func hapusAnggota(tim *Tim) {
	var namaAnggota string
	fmt.Print("Nama anggota yang ingin dihapus: ")
	fmt.Scanln(&namaAnggota)

	jumlahAnggota := cekIndekAnggota(&tim.DaftarAnggota)
	if jumlahAnggota == -1 {
		jumlahAnggota = len(tim.DaftarAnggota) // penuh
	}

	for i := 0; i < jumlahAnggota; i++ {
		if tim.DaftarAnggota[i].NamaAnggota == namaAnggota {
			tim.DaftarAnggota[i] = Anggota{} // kosongkan elemen
			fmt.Println("Anggota berhasil dihapus.")
			return
		}
	}

	fmt.Println("Anggota tidak ditemukan.")
}

func hapusTim(startup *Startup) {
	var namaTim string
	fmt.Print("Nama tim yang ingin dihapus: ")
	fmt.Scanln(&namaTim)

	jumlahTim := cekIndekTim(&startup.DaftarTim)
	if jumlahTim == -1 {
		jumlahTim = len(startup.DaftarTim) // array penuh
	}

	for i := 0; i < jumlahTim; i++ {
		if startup.DaftarTim[i].NamaTim == namaTim {
			startup.DaftarTim[i] = Tim{} // kosongkan tim
			fmt.Println("Tim berhasil dihapus.")
			return
		}
	}

	fmt.Println("Tim tidak ditemukan.")
}
