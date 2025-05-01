package main

import "fmt"

func daftar() {
	fmt.Println("=== Halaman Daftar ===")
	fmt.Print("Nama: ")
	fmt.Scan(&database_pengguna[pengguna_terdaftar].nama)
	fmt.Print("Email: ")
	fmt.Scan(&database_pengguna[pengguna_terdaftar].email)
	fmt.Print("Kata Sandi: ")
	fmt.Scan(&database_pengguna[pengguna_terdaftar].kata_sandi)
}

func masuk() {

}
