package main

import "fmt"

func main() {
	seedPengguna()
	seedStartup()

	for {

		var (
			penggunaAktif Pengguna
			pilihan       int
		)

		uiMainMenu(&pilihan)

		switch pilihan {
		case 1:
			clearScreen()
			masuk(&dbPengguna, &penggunaAktif)
			if penggunaAktif.IdPengguna != 0 {
				arahkanKeDashboard(&penggunaAktif)
			} else {
				fmt.Println("[Eror] Masuk gagal, ulangi.")
				fmt.Println()
			}
		case 2:
			clearScreen()
			daftar(&dbPengguna)
		case 0:
			fmt.Println("Terimakasih, sampai jumpa.")
			return
		default:
			fmt.Println("Opsi yang dipilih tidak valid.")
		}
	}
}

func arahkanKeDashboard(penggunaAktif *Pengguna) {
	switch penggunaAktif.JenisPengguna {
	case "pendiri":
		dashboardPendiri(penggunaAktif)
	case "investor":
		dashboardInvestor(penggunaAktif)
	default:
		fmt.Println("Jenis pengguna tidak valid.")
	}

}
