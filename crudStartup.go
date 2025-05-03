package main

import "fmt"

func tambahStartup(p *Pengguna, db *[maxStartup]Startup, index *int) {
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

	db[*index] = Startup{
		IdStartup:      *index + 1,
		IdPendiri:      p.IdPengguna,
		NamaStartup:    namaStartup,
		BidangUsaha:    bidangUsaha,
		TahunBerdiri:   tahunBerdiri,
		TotalPendanaan: totalPendanaan,
	}

	*index++

}

func cetakStartup(s *Startup) {
	fmt.Println("-------------------------------")
	fmt.Printf("Nama Startup    : %s\n", s.NamaStartup)
	fmt.Printf("Bidang Usaha    : %s\n", s.BidangUsaha)
	fmt.Printf("Tahun Berdiri   : %d\n", s.TahunBerdiri)
	fmt.Printf("Total Pendanaan : %d\n", s.TotalPendanaan)
	fmt.Println("Daftar Tim:")
	for i := 0; i < MaxTeam; i++ {
		if s.DaftarTeam[i].NamaTeam != "" {
			fmt.Printf("  - %s:\n", s.DaftarTeam[i].NamaTeam)
			for j := 0; j < MaxAnggota; j++ {
				if s.DaftarTeam[i].AnggotaTeam[j].NamaAnggota != "" {
					fmt.Printf("    > %s (%s)\n", s.DaftarTeam[i].AnggotaTeam[j].NamaAnggota, s.DaftarTeam[i].AnggotaTeam[j].Posisi)
				}
			}
		}
	}
	fmt.Println("-------------------------------")
}

func tampilkanStartupSaya(p *Pengguna) {
	fmt.Println("=== Daftar Startup Saya ===")
	ada := false
	for i := 0; i < StartupTerdaftar; i++ {
		if DatabaseStartup[i].IdPendiri == p.IdPengguna {
			cetakStartup(&DatabaseStartup[i])
			ada = true
		}
	}
	if !ada {
		fmt.Println("Anda belum membuat startup.")
	}
}

func tampilkanSeluruhStartup() {
	fmt.Println("=== Daftar Seluruh Startup ===")
	if StartupTerdaftar == 0 {
		fmt.Println("Belum ada startup yang terdaftar.")
		return
	}
	for i := 0; i < StartupTerdaftar; i++ {
		cetakStartup(&DatabaseStartup[i])
	}
}

func tampilkanStartup(p *Pengguna) {
	var pilihan int
	fmt.Println("=== Menu Tampilkan Startup ===")
	fmt.Println("1. Tampilkan Startup Saya")
	fmt.Println("2. Tampilkan Seluruh Startup")
	fmt.Print("Pilihan Anda: ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		tampilkanStartupSaya(p)
	case 2:
		tampilkanSeluruhStartup()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func ubahStartup()  {}
func hapusStartup() {}
