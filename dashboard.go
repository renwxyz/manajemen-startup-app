package main

import "fmt"

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
			fmt.Println("Fitur A dalam Pengembangan")
			// Masukkan suatu fungsi
		case 2:
			fmt.Println("Fitur B dalam Pengembangan")
			// Masukkan suatu fungsi
		case 0:
			fmt.Println("Keluar dari Dashboard Pendiri")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

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
