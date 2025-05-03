package main

import "fmt"

// MENU PENDIRI

func buatStartup(p *Pengguna) {
	var (
		s          *Startup
		jumlahTeam int
	)
	tambahStartup(p, &DatabaseStartup, &StartupTerdaftar)
	s = &DatabaseStartup[StartupTerdaftar-1]

	fmt.Print("Jumlah tim yang ingin ditambahkan: ")
	fmt.Scanln(&jumlahTeam)

	for i := 0; i < jumlahTeam; i++ {
		tambahTeam(s)
		// Tambahkan anggota ke tim yang baru saja ditambahkan
		// Temukan indeks tim yang baru ditambahkan (yang paling akhir)
		for j := 0; j < MaxTeam; j++ {
			if s.DaftarTeam[j].NamaTeam != "" {
				// Jika ini tim terakhir (berikutnya kosong atau sudah di ujung)
				if j == MaxTeam-1 || s.DaftarTeam[j+1].NamaTeam == "" {
					tambahAnggota(&s.DaftarTeam[j])
					break
				}
			}
		}
	}
}

func dashboardPendiri(p *Pengguna) {
	for {
		var pilihan int
		fmt.Println("=== DASHBOARD PENDIRI ===")
		fmt.Printf("Halo, %s! Anda masuk sebagai pendiri.\n", p.Nama)
		fmt.Println("[1] Buat Startup")
		fmt.Println("[2] Tampilkan Startup")
		fmt.Println("[3] Ubah Data Startup")
		fmt.Println("[4] Cari Startup")
		fmt.Println("[0] Keluar")
		fmt.Print("Pilih menu:")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			buatStartup(p)
		case 2:
			tampilkanStartup(p)
		case 0:
			fmt.Println("Keluar dari Dashboard Pendiri")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

//[END] MENU PENDIRI

func dashboardInvestor(p *Pengguna) {
	for {
		var pilihan int
		fmt.Println("=== DASHBOARD INVESTOR ===")
		fmt.Printf("Halo, %s! Anda masuk sebagai pendiri.\n", p.Nama)
		fmt.Println("[1] Tampilkan Semua Startup")
		fmt.Println("[2] Cari Startup")
		fmt.Println("[3] Laporan Statistik")
		fmt.Println("[0] Keluar")
		fmt.Print("Pilih menu:")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("Fitur A dalam Pengembangan")
			// Masukkan suatu fungsi
		case 2:
			fmt.Println("Fitur B dalam Pengembangan")
			// Masukkan suatu fungsi
		case 0:
			fmt.Println("Keluar dari Dashboard Investor")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
