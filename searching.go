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
