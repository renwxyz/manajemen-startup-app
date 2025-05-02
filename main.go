package main

import "fmt"

func main() {
	masukanDataDummy()
	fmt.Println("=== Menampilkan User Terdaftar ===")
	for i := 0; i < JumlahTerdaftar; i++ {
		fmt.Println("id:", DatabasePengguna[i].UserId)
		fmt.Println("nama:", DatabasePengguna[i].Nama)
		fmt.Println("email:", DatabasePengguna[i].Email)
		fmt.Println("")
	}

	fmt.Println("=== Menampilkan Informasi Startup ===")
	for i := 0; i < JumlahTerdaftar; i++ {

	}
}
