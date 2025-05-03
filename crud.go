package main

import "fmt"

// CRUD STARTUP
func tambahStartup() {
	var (
		namaStartup    string
		bidangUsaha    string
		tahunBerdiri   int
		totalPendanaan int
	)

	fmt.Print("Nama Startup: ")
	fmt.Scanln(&namaStartup)
	fmt.Print("Bidang Usaha: ")
	fmt.Scanln(&bidangUsaha)
	fmt.Print("Tahun Berdiri: ")
	fmt.Scanln(&tahunBerdiri)
	fmt.Print("Total Pendanaan: ")
	fmt.Scanln(&totalPendanaan)

	DatabaseStartup[StartupTerdaftar] = Startup{
		StartupId:      StartupTerdaftar + 1,
		NamaStartup:    namaStartup,
		BidangUsaha:    bidangUsaha,
		TahunBerdiri:   tahunBerdiri,
		TotalPendanaan: totalPendanaan,
	}

	StartupTerdaftar++

}

func tampilkanStartup() {}
func hapusStartup()     {}
func ubahStartup()      {}

//[END] CRUD STARTUP

// CRUD TEAM
func tambahTeam() {
	var (
		namaTeam string
	)
	fmt.Print("Nama Team: ")
	fmt.Scanln(&namaTeam)

	for i := 0; i < MaxTeam; i++ {
		// Temukan slot kosong pada DaftarTeam
		if DatabaseStartup[StartupTerdaftar-1].DaftarTeam[i].NamaTeam == "" {
			DatabaseStartup[StartupTerdaftar-1].DaftarTeam[i] = Team{
				NamaTeam: namaTeam,
			}
			break
		}
	}

	fmt.Println("Team berhasil ditambahkan!")

}
func tampilkanTeam() {}
func hapusTeam()     {}
func ubahTeam()      {}

//[END] CRUD TEAM

//CRUD ANGGOTA
func tambahAnggota() {
	var (
		namaAnggota string
		posisi      string
		teamIndex   int
	)

	// Input nama anggota dan posisi
	fmt.Print("Nama Anggota: ")
	fmt.Scanln(&namaAnggota)
	fmt.Print("Posisi: ")
	fmt.Scanln(&posisi)

	// Pilih team yang ingin ditambahkan anggotanya
	fmt.Print("Pilih nomor team (0-5): ")
	fmt.Scanln(&teamIndex)

	// Menambahkan anggota ke team yang dipilih
	for i := 0; i < MaxAnggota; i++ {
		// Temukan slot kosong pada AnggotaTeam
		if DatabaseStartup[StartupTerdaftar-1].DaftarTeam[teamIndex].AnggotaTeam[i].NamaAnggota == "" {
			DatabaseStartup[StartupTerdaftar-1].DaftarTeam[teamIndex].AnggotaTeam[i] = Anggota{
				NamaAnggota: namaAnggota,
				Posisi:      posisi,
			}
			break
		}
	}

	fmt.Println("Anggota berhasil ditambahkan!")
}

func tampilAnggota() {}
func hapusAnggota()  {}
func ubahAnggota()   {}

//[END] CRUD ANGGOTA
