package main

import "fmt"

func dashboardPendiri(p *Pengguna) {
	for {
		var pilihan int
		fmt.Println("=== DASHBOARD PENDIRI ===")
		fmt.Printf("Halo, %s! Anda masuk sebagai pendiri.\n", p.Nama)
		fmt.Println("1. Menu A")
		fmt.Println("2. Menu B")
		fmt.Println("3. Logout")
		fmt.Print("Pilih menu:")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("Fitur A dalam Pengembangan")
			// Masukkan suatu fungsi
		case 2:
			fmt.Println("Fitur B dalam Pengembangan")
			// Masukkan suatu fungsi
		case 3:
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
		fmt.Println("1. Menu A")
		fmt.Println("2. Menu B")
		fmt.Println("3. Logout")
		fmt.Print("Pilih menu:")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("Fitur A dalam Pengembangan")
			// Masukkan suatu fungsi
		case 2:
			fmt.Println("Fitur B dalam Pengembangan")
			// Masukkan suatu fungsi
		case 3:
			fmt.Println("Keluar dari Dashboard Investor")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
