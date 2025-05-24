package main

func cekIndeksDbPengguna(db *[100]Pengguna) int {
	for i := 0; i < len(db); i++ {
		if db[i].IdPengguna == 0 {
			return i
		}
	}
	return -1
}

func cekIdPengguna(db *[100]Pengguna) int {
	var maxIndekTerisi = cekIndeksDbPengguna(db)
	var maxIdTerisi = -1

	if maxIndekTerisi != -1 {
		maxIdTerisi = 0
		for i := 0; i < maxIndekTerisi; i++ {
			if db[i].IdPengguna > maxIdTerisi {
				maxIdTerisi = db[i].IdPengguna
			}
		}
	}

	return maxIdTerisi
}

func cekIndeksDbStartup(db *[100]Startup) int {
	for i := 0; i < len(db); i++ {
		if db[i].IdStartup == 0 {
			return i
		}
	}
	return -1
}

func cekIdStartup(db *[100]Startup) int {
	var maxIndekTerisi = cekIndeksDbStartup(db)
	var maxIdTerisi = -1

	if maxIndekTerisi != -1 {
		maxIdTerisi = 0
		for i := 0; i < maxIndekTerisi; i++ {
			if db[i].IdStartup > maxIdTerisi {
				maxIdTerisi = db[i].IdStartup
			}
		}
	}

	return maxIdTerisi
}

func cekIndekTim(db *[100]Tim) int {
	for i := 0; i < len(db); i++ {
		if db[i].NamaTim == "" {
			return i
		}
	}
	return -1
}

func cekIndekAnggota(db *[100]Anggota) int {
	for i := 0; i < len(db); i++ {
		if db[i].NamaAnggota == "" {
			return i
		}
	}
	return -1
}

func getDataStartupValid(db *[100]Startup) ([100]Startup, int) {
	var hasil [100]Startup
	var jumlahDataValid int = 0

	for i := 0; i < len(db); i++ {
		if db[i].IdStartup != 0 {
			hasil[jumlahDataValid] = db[i]
			jumlahDataValid++
		}
	}

	return hasil, jumlahDataValid
}

/*
Dokumentasi cekIndeksDbPengguna(db *[100]Pengguna)int{}
Parameter bertipe data pointer ke suatu array [100]Pengguna

Fungsi Utama: Mengecek apakah masih ada indeks kosong
Program bekerja dengan cara mengiterasi db sebanyak len(db)
kemudian dicek apakah ada db[i].IdPengguna == 0
jika ada return index ke-i untuk kemudian siap dipakai

jika tidak ada db[i].IdPengguna == 0
maka akan return -1 yang artinya tidak ada indeks siap pakai
*/

/*
Dokumentasi cekIdPengguna(db *[100]Pengguna)int{}
Parameter bertipe data pointer ke suatu array [100]Pengguna

Fungsi utama: Mencari id terbesar yang sudah digunakan.
Program bekerja dengan cara mengiterasi sebanyak sebelum indeks
kosong.

jika maxIndeksTerisi bukan -1 (database masih belum penuh)
maka iterasi sepanjang indeks terisi. kemudian akan dicek
diantara 0 hingga indeks terisi terakhir berapa id yang paling besar
yang sudah terpakai. program mereturn id terbesar yang digunakan

jika jika maxIndeksTerisi -1 (database penuh)
maka tidak perlu ada pengecekan untuk id yang tersedia
karena sudah tidak bisa menambahkan pengguna.
program akan mereturn -1 (tidak tersedia)
*/
