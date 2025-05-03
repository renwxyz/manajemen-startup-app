package main

import "fmt"

func main() {
	masukanDataDummy()

	for {
		var pilihan int
		fmt.Println("===== Selamat Datang! =====")
		fmt.Println("[1] Masuk")
		fmt.Println("[2] Daftar")
		fmt.Println("[0] Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scanln(&pilihan)
		fmt.Println("")

		switch pilihan {
		case 1:
			masuk()
		case 2:
			daftar()
		case 0:
			fmt.Println("Sampai Jumpa Lagi!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}

	}

}
