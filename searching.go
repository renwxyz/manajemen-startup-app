package main

import "fmt"

func seqSearchByNamaStartup(nama string, db *[100]Startup) {
	var (
		ditemukan bool
	)

	ditemukan = false
	for i := 0; i < len(db); i++ {
		if db[i].NamaStartup != "" && nama == db[i].NamaStartup {
			detailStartup(db, i)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Startup tidak ditemukan")
	}
}

func selectionSortByBidangUsaha(db *[100]Startup, n int) {
	var indeksMin int = n
	for i := 0; i < n-1; i++ {
		indeksMin = i
		for j := i + 1; j < n; j++ {
			if db[j].BidangUsaha < db[indeksMin].BidangUsaha {
				indeksMin = j
			}
		}
		if indeksMin != i {
			db[i], db[indeksMin] = db[indeksMin], db[i]
		}
	}
}

func binarySearchByBidangUsaha(cari string, db *[100]Startup, n int) {
	var (
		kiri, kanan, tengah, indeks int
		ditemukan                   bool
		count                       int
	)

	selectionSortByBidangUsaha(db, n)

	if n <= 0 {
		fmt.Println("Tidak ada data startup.")
		return
	}

	kiri = 0
	kanan = n - 1
	ditemukan = false

	for kiri <= kanan && !ditemukan {
		tengah = (kiri + kanan) / 2
		if db[tengah].BidangUsaha == cari {
			ditemukan = true
			indeks = tengah
		} else if cari < db[tengah].BidangUsaha {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if !ditemukan {
		fmt.Println("Bidang usaha tidak ditemukan.")
		return
	}

	// Cari semua yang cocok ke kiri
	i := indeks
	for i >= 0 && db[i].BidangUsaha == cari {
		i--
	}
	i++

	// Cetak semua yang cocok ke kanan
	count = 1

	uiHeaderTableStartup()
	for i < n && db[i].BidangUsaha == cari {
		tampilkanStartup(db, i, count)
		i++
		count++
	}
	uiFooterTable()
}

func cariStartup(db *[100]Startup) {
	var pilihan int

	for {

		uiCariStartup(&pilihan)

		switch pilihan {
		case 1:
			cariByNama(db)
		case 2:
			cariByBidangUsaha(db)
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}

	}

}

func cariByNama(db *[100]Startup) {
	var (
		nama    string
		pilihan int
	)

	pilihan = 1
	for pilihan == 1 {
		uiInsertTeks("nama startup", &nama)
		seqSearchByNamaStartup(nama, db)
		uiOpsiLanjutan("Cari lagi", &pilihan)
	}
}

func cariByBidangUsaha(db *[100]Startup) {
	var (
		bidangUsaha      string
		pilihan          int
		validData, count = getDataStartupValid(db)
	)

	pilihan = 1
	for pilihan == 1 {
		uiInsertTeks("bidang usaha", &bidangUsaha)

		binarySearchByBidangUsaha(bidangUsaha, &validData, count)

		uiOpsiLanjutan("Cari lagi", &pilihan)
	}
}
