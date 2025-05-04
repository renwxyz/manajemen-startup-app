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
	fmt.Println("")

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
	fmt.Println("========================================")
	fmt.Printf("Startup #%d: %s\n", s.IdStartup, s.NamaStartup)
	fmt.Printf("Bidang     : %s\n", s.BidangUsaha)
	fmt.Printf("Tahun      : %d\n", s.TahunBerdiri)
	fmt.Printf("Pendanaan  : %d\n", s.TotalPendanaan)
	fmt.Printf("Pendiri ID : %d\n", s.IdPendiri)
	fmt.Println("----------------------------------------")

	for i, team := range s.DaftarTeam {
		if team.NamaTeam == "" {
			continue
		}
		fmt.Printf("[Team %d] %s\n", i+1, team.NamaTeam)
		for _, anggota := range team.AnggotaTeam {
			if anggota.NamaAnggota != "" {
				fmt.Printf("  - %-8s (%s)\n", anggota.NamaAnggota, anggota.Posisi)
			}
		}
		fmt.Println()
	}
	fmt.Println("========================================")
	fmt.Println("")
}

func tampilkanStartupSaya(p *Pengguna) {
	var ada bool = false
	fmt.Println("=== Daftar Startup Saya ===")
	for i := 0; i < StartupTerdaftar; i++ {
		if DatabaseStartup[i].IdPendiri == p.IdPengguna {
			cetakStartup(&DatabaseStartup[i])
			ada = true
		}
	}
	if !ada {
		fmt.Println("Anda belum membuat startup.")
		fmt.Println("")
	}
}

func tampilkanSeluruhStartup() {
	fmt.Println("=== Daftar Seluruh Startup ===")
	if StartupTerdaftar == 0 {
		fmt.Println("Belum ada startup yang terdaftar.")
		fmt.Println("")
		return
	}
	for i := 0; i < StartupTerdaftar; i++ {
		cetakStartup(&DatabaseStartup[i])
	}
}

func tampilkanStartup(p *Pengguna) {
	var pilihan int
	fmt.Println("=== Menu Tampilkan Startup ===")
	fmt.Println("[1] Tampilkan Startup Saya")
	fmt.Println("[2] Tampilkan Seluruh Startup")
	fmt.Print("Pilihan Anda: ")
	fmt.Scanln(&pilihan)
	fmt.Println("")

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
