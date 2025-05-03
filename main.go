package main

import "fmt"

func main() {
	masukanDataDummy()
	for {
		fmt.Println("1. Pendiri")
		fmt.Println("2. Investor")
		fmt.Println("3. Batal")
		fmt.Print("piliih menu:")
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			dashboardPendiri(&DatabasePengguna[0])
		case 2:
			dashboardInvestor(&DatabasePengguna[1])
		case 3:
			fmt.Println("batal")
			return
		default:
			fmt.Println("tidak ada opsi")
		}
	}
}
