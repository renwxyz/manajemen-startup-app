package main

import "fmt"

func laporanJumlahPerBidangUsaha(db *[100]Startup) {
	var pilihan int
	var selesai bool
	for !selesai {
		var bidangUnik [100]string
		var jumlah [100]int
		var jumlahKategori int = 0

		for i := 0; i < 100; i++ {
			if db[i].IdStartup != 0 {
				var ditemukan bool = false
				var indeks int = 0

				for j := 0; j < jumlahKategori; j++ {
					if db[i].BidangUsaha == bidangUnik[j] {
						ditemukan = true
						indeks = j
					}
				}

				if ditemukan {
					jumlah[indeks]++
				} else {
					bidangUnik[jumlahKategori] = db[i].BidangUsaha
					jumlah[jumlahKategori] = 1
					jumlahKategori++
				}
			}
		}

		// Tampilkan hasil
		fmt.Println("Laporan Jumlah Startup per Kategori Bidang Usaha")
		fmt.Println("===============================================")
		fmt.Printf(" %-30s | %-10s\n", "Bidang Usaha", "Jumlah")
		fmt.Println("===============================================")
		for i := 0; i < jumlahKategori; i++ {
			fmt.Printf(" %-30s | %-10d\n", bidangUnik[i], jumlah[i])
		}
		fmt.Println("===============================================")

		uiOpsiSelesai(&pilihan)

		if pilihan == 0 {
			selesai = true
		} else if pilihan != 0 {
			pesanEror("Pilihan tidak valid.")
		}
	}
}
