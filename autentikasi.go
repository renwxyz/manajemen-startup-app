package main

import "fmt"

func daftar() {
	var (
		nama          string
		email         string
		kataSandi     string
		jenisPengguna string
	)

	fmt.Println("===== DAFTAR =====")
	fmt.Print("Nama :")
	fmt.Scanln(&nama)
	fmt.Print("Email :")
	fmt.Scanln(&email)
	fmt.Print("Kata sandi: ")
	fmt.Scanln(&kataSandi)
	fmt.Print("Jenis Pengguna [pendiri/investor] :")
	fmt.Scanln(&jenisPengguna)
	fmt.Println("")

	DatabasePengguna[JumlahTerdaftar].UserId = JumlahTerdaftar + 1
	DatabasePengguna[JumlahTerdaftar].Nama = nama
	DatabasePengguna[JumlahTerdaftar].Email = email
	DatabasePengguna[JumlahTerdaftar].KataSandi = kataSandi
	DatabasePengguna[JumlahTerdaftar].JenisPengguna = jenisPengguna

	JumlahTerdaftar += 1
}

func masuk() {
	var (
		email          string
		kataSandi      string
		jenisPengguna  string
		statusMasuk    bool
		indeksPengguna int
	)

	fmt.Println("===== Masuk =====")
	fmt.Print("Email :")
	fmt.Scanln(&email)
	fmt.Print("Kata sandi: ")
	fmt.Scanln(&kataSandi)
	fmt.Println("")

	for i := 0; i < JumlahTerdaftar; i++ {
		if email == DatabasePengguna[i].Email && kataSandi == DatabasePengguna[i].KataSandi {
			jenisPengguna = DatabasePengguna[i].JenisPengguna
			statusMasuk = true
			indeksPengguna = i
			break
		}
	}

	if statusMasuk {
		switch jenisPengguna {
		case "pendiri":
			dashboardPendiri(&DatabasePengguna[indeksPengguna])
		case "investor":
			dashboardInvestor(&DatabasePengguna[indeksPengguna])
		}
	} else {
		fmt.Println("Email atau kata sandi salah.")
	}
}
