package main

import "fmt"

func dashboardPendiri(penggunaAktif *Pengguna) {
	for {
		clearScreen()
		var pilihan int

		uiDashboardPendiri(penggunaAktif, &pilihan)

		switch pilihan {
		case 1:
			buatStartup(penggunaAktif, &dbStartup)
		case 2:
			tampilkanStartupSaya(penggunaAktif, &dbStartup)
		case 3:
			tampilkanSemuaStartup(&dbStartup)
		case 4:
			cariStartup(&dbStartup)
		case 5:
			laporanJumlahPerBidangUsaha(&dbStartup)
		case 0:
			fmt.Println("Keluar dari Dashboard Pendiri")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

//BLOK OPSI 1
func buatStartup(penggunaAktif *Pengguna, db *[100]Startup) {
	var (
		startupBaru *Startup
		jumlahTim   int
		timBaru     *Tim
	)

	clearScreen()

	startupBaru = tambahStartup(penggunaAktif, db)

	uiInsertInt("jumlah tim", &jumlahTim)

	for i := 0; i < jumlahTim; i++ {
		var jumlahAnggota int
		fmt.Printf("=== Tambah tim ke-%d ===\n", i+1)
		timBaru = tambahTim(startupBaru)

		fmt.Print("Jumlah anggota: ")
		fmt.Scanln(&jumlahAnggota)

		for i := 0; i < jumlahAnggota; i++ {
			fmt.Printf("[Tambah anggota ke-%d]\n", i+1)
			tambahAnggota(timBaru)
		}

		fmt.Println()
	}

}

//[END] BLOK OPSI 1

//BLOK OPSI 2
func tampilkanStartupSaya(penggunaAktif *Pengguna, db *[100]Startup) {
	for {
		var indekTersedia int = cekIndeksDbStartup(db)
		var count int = 1

		clearScreen()
		uiHeaderTableStartup()

		for i := 0; i < indekTersedia; i++ {
			if penggunaAktif.IdPengguna == db[i].IdPendiri {
				tampilkanStartup(db, i, count)
				count++
			}
		}

		uiFooterTable()

		var pilihan int

		uiTampilkanStartupSaya(&pilihan)

		switch pilihan {
		case 1:
			detailSemuaStartupSaya(penggunaAktif, db)
		case 2:
			pilihDetailStartupSaya(penggunaAktif, db)
		case 3:
			ubahDataStartup(penggunaAktif, db)
		case 4:
			hapusStartupSaya(penggunaAktif, db)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}

}

func detailSemuaStartupSaya(penggunaAktif *Pengguna, db *[100]Startup) {
	var pilihan int
	var selesai bool = false

	for !selesai {
		var indeksTersedia int = cekIndeksDbStartup(db)
		var counter int = 1
		var ditemukan bool = false

		clearScreen()

		for i := 0; i < indeksTersedia; i++ {
			if penggunaAktif.IdPengguna == db[i].IdPendiri {
				uiHeaderDetailStartup(counter)
				detailStartup(db, i)
				fmt.Println()
				counter++
				ditemukan = true
			}
		}

		if !ditemukan {
			pesanInformasi("Anda belum memiliki starup.")
		}

		fmt.Println()
		uiOpsiLanjutan("Refresh", &pilihan)

		if pilihan == 0 {
			selesai = true
		} else if pilihan != 1 {
			pesanEror("Pilihan tidak valid.")
		}
	}
}

func pilihDetailStartupSaya(penggunaAktif *Pengguna, db *[100]Startup) {
	var nNamaStartup string
	var pilihan int
	var selesai bool = false

	for !selesai {
		clearScreen()
		uiInsertTeks("nama startup", &nNamaStartup)

		var indeksTersedia int = cekIndeksDbStartup(db)
		var ditemukan bool = false

		for i := 0; i < indeksTersedia; i++ {
			if penggunaAktif.IdPengguna == db[i].IdPendiri && nNamaStartup == db[i].NamaStartup {
				detailStartup(db, i)
				ditemukan = true
			}
		}

		if !ditemukan {
			pesanInformasi("Startup tidak ditemukan.")
		}

		fmt.Println()
		uiOpsiLanjutan("Pilih startup lain", &pilihan)

		if pilihan == 0 {
			selesai = true
		} else if pilihan != 1 {
			pesanEror("Pilihan tidak valid.")
		}
	}
}

func ubahDataStartup(penggunaAktif *Pengguna, db *[100]Startup) {
	for {
		var nNamaStartup string
		var idStartup int

		clearScreen()
		uiInsertTeks("nama startup", &nNamaStartup)

		var indekTersedia int = cekIndeksDbStartup(db)
		for i := 0; i < indekTersedia; i++ {
			if penggunaAktif.IdPengguna == db[i].IdPendiri && nNamaStartup == db[i].NamaStartup {
				detailStartup(db, i)
				idStartup = i
			}
		}

		fmt.Println()

		var pilihan int
		uiUbahDataStartup(&pilihan)

		switch pilihan {
		case 1:
			ubahIdPendiri(db, idStartup)
		case 2:
			ubahNamaStartup(db, idStartup)
		case 3:
			ubahBidangUsaha(db, idStartup)
		case 4:
			ubahTotalPendanaan(db, idStartup)
		case 5:
			kelolaTimStartup(db, idStartup)
		case 0:
			return
		default:
			pesanEror("Pilihan tidak valid.")
		}
	}
}

func kelolaTimStartup(db *[100]Startup, indeks int) {
	for {
		var indekTim int = cekIndekTim(&db[indeks].DaftarTim)

		clearScreen()

		for j := 0; j < indekTim; j++ {
			tampilkanTim(&db[indeks].DaftarTim, j)
		}

		var pilihan int
		uiKelolaTim(&pilihan)

		switch pilihan {
		case 1:
			ubahTim(db, indeks)
		case 2:
			buatTim(db, indeks)
		case 3:
			hapusTim(&db[indeks])
		case 0:
			return
		default:
			pesanEror("Pilihan tidak valid.")
		}
	}
}

func ubahTim(db *[100]Startup, indeks int) {
	var timDitemukan bool = false
	var indeksTim int = -1

	for !timDitemukan {
		clearScreen()
		var nNamaTim string

		fmt.Print("Nama Tim: ")
		fmt.Scanln(&nNamaTim)

		var jumlahTim = cekIndekTim(&db[indeks].DaftarTim)

		// Cek apakah nama tim ditemukan
		for j := 0; j < jumlahTim; j++ {
			if db[indeks].DaftarTim[j].NamaTim == nNamaTim {
				indeksTim = j
				timDitemukan = true
			}
		}

		if !timDitemukan {
			pesanEror("Tim tidak ditemukan.")
		}
	}

	// tim ditemukan
	tampilkanTim(&db[indeks].DaftarTim, indeksTim)
	var selectedTim *Tim = &db[indeks].DaftarTim[indeksTim]

	for {
		var pilihan int
		uiUbahTim(&pilihan)

		switch pilihan {
		case 1:
			ubahNamaTim(selectedTim)
		case 2:
			ubahAnggota(selectedTim)
		case 3:
			buatAnggota(selectedTim)
		case 4:
			hapusAnggota(selectedTim)
		case 0:
			return
		default:
			pesanEror("Pilihan tidak valid.")
		}
	}
}

func buatAnggota(tim *Tim) {
	var jumlahAnggota int

	uiInsertInt("jumlah anggota", &jumlahAnggota)

	for i := 0; i < jumlahAnggota; i++ {
		uiInformasiTambah("anggota", i+1)
		tambahAnggota(tim)
	}
}

func buatTim(db *[100]Startup, indeks int) {
	var jumlahTim int
	uiInsertInt("jumlah tim", &jumlahTim)

	for i := 0; i < jumlahTim; i++ {
		var jumlahAnggota int
		var timBaru *Tim

		uiInformasiTambah("tim", i+1)

		timBaru = tambahTim(&db[indeks])

		uiInsertInt("jumlah anggota", &jumlahAnggota)

		for i := 0; i < jumlahAnggota; i++ {
			uiInformasiTambah("anggota", i+1)
			tambahAnggota(timBaru)
		}

		fmt.Println()
	}

}

func tampilkanSemuaStartup(db *[100]Startup) {
	var indekTersedia int = cekIndeksDbStartup(db)

	for {
		var count int = 1
		clearScreen()
		uiHeaderTableStartup()

		for i := 0; i < indekTersedia; i++ {
			tampilkanStartup(db, i, count)
			count++
		}

		uiFooterTable()

		var pilihan int
		uiTampilkanSemuaStartup(&pilihan)

		switch pilihan {
		case 1:
			pilihDetail(db)
		case 2:
			urutNaikTotalPendanaan(&dbStartup)
		case 3:
			urutNaikTahunBerdiri(&dbStartup)
		case 4:
			urutTurunTotalPendanaan(&dbStartup)
		case 5:
			urutTurunTahunBerdiri(&dbStartup)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func pilihDetail(db *[100]Startup) {
	var nNamaStartup string
	var indekTersedia int = cekIndeksDbStartup(db)
	var pilihan int
	var selesai bool

	for !selesai {
		uiInsertTeks("nama startup", &nNamaStartup)
		for i := 0; i < indekTersedia; i++ {
			if nNamaStartup == db[i].NamaStartup {
				detailStartup(db, i)
			}
		}

		uiOpsiLanjutan("Pilih startup lain", &pilihan)

		if pilihan == 0 {
			selesai = true
		} else if pilihan != 1 {
			pesanEror("Pilihan tidak valid.")
		}
	}

}

func urutNaikTotalPendanaan(db *[100]Startup) {
	var validData, count = getDataStartupValid(db)
	var pilihan int
	var selesai bool

	for !selesai {
		clearScreen()
		var numList int = 1
		selectionSortByTotalPendanaanAsc(&validData, count)

		uiHeaderTableStartup()
		for i := 0; i < count; i++ {
			tampilkanStartup(&validData, i, numList)
			numList++
		}
		uiFooterTable()

		uiOpsiSelesai(&pilihan)

		if pilihan == 0 {
			selesai = true
		} else if pilihan != 0 {
			pesanEror("Pilihan tidak valid.")
		}
	}

}

func urutNaikTahunBerdiri(db *[100]Startup) {
	var validData, count = getDataStartupValid(db)
	var pilihan int
	var selesai bool

	for !selesai {
		var numList int = 1
		clearScreen()
		insertionSortByTahunBerdiriAsc(&validData, count)

		uiHeaderTableStartup()
		for i := 0; i < count; i++ {
			tampilkanStartup(&validData, i, numList)
			numList++
		}
		uiFooterTable()

		uiOpsiSelesai(&pilihan)

		if pilihan == 0 {
			selesai = true
		} else if pilihan != 0 {
			pesanEror("Pilihan tidak valid.")
		}
	}
}

func urutTurunTotalPendanaan(db *[100]Startup) {
	var validData, count = getDataStartupValid(db)

	var pilihan int
	var selesai bool
	for !selesai {
		var numList int = 1
		clearScreen()
		selectionSortByTotalPendanaanDesc(&validData, count)

		uiHeaderTableStartup()

		for i := 0; i < count; i++ {
			tampilkanStartup(&validData, i, numList)
			numList++
		}

		uiFooterTable()

		uiOpsiSelesai(&pilihan)

		if pilihan == 0 {
			selesai = true
		} else if pilihan != 0 {
			pesanEror("Pilihan tidak valid.")
		}
	}
}
func urutTurunTahunBerdiri(db *[100]Startup) {
	var validData, count = getDataStartupValid(db)
	var pilihan int
	var selesai bool
	for !selesai {
		var numList int = 1
		clearScreen()
		insertionSortByTahunBerdiriDesc(&validData, count)

		uiHeaderTableStartup()
		for i := 0; i < count; i++ {
			tampilkanStartup(&validData, i, numList)
			numList++
		}
		uiFooterTable()

		uiOpsiSelesai(&pilihan)

		if pilihan == 0 {
			selesai = true
		} else if pilihan != 0 {
			pesanEror("Pilihan tidak valid.")
		}
	}
}

func dashboardInvestor(penggunaAktif *Pengguna) {
	var pilihan int
	for {
		clearScreen()
		uiDashboardInvestor(penggunaAktif, &pilihan)

		switch pilihan {
		case 1:
			tampilkanSemuaStartup(&dbStartup)
		case 2:
			cariStartup(&dbStartup)
		case 0:
			fmt.Println("Keluar dari Dashboard investor")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
