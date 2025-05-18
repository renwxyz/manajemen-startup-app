package main

import "fmt"

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
