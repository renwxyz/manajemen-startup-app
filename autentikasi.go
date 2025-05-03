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

	DatabasePengguna[JumlahTerdaftar].Nama = nama
	DatabasePengguna[JumlahTerdaftar].Email = email
	DatabasePengguna[JumlahTerdaftar].KataSandi = kataSandi
	DatabasePengguna[JumlahTerdaftar].JenisPengguna = jenisPengguna
}

func masuk() {
	fmt.Println("on development")
}
