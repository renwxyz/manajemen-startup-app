package main

import "fmt"

// CRUD TEAM
func tambahTeam(s *Startup) {
	var (
		namaTeam string
	)
	fmt.Print("Nama Team: ")
	fmt.Scanln(&namaTeam)

	for i := 0; i < MaxTeam; i++ {
		// Temukan slot kosong pada DaftarTeam
		if s.DaftarTeam[i].NamaTeam == "" {
			s.DaftarTeam[i].NamaTeam = namaTeam
			return
		}
	}

	fmt.Println("Maksimum team tercapai.")

}

func tampilkanTeam() {}
func hapusTeam()     {}
func ubahTeam()      {}

//[END] CRUD TEAM
