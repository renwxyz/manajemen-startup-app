package main

import "fmt"

func tambahAnggota(t *Team) {
	var jumlahAnggota int
	fmt.Print("Jumlah anggota yang ingin ditambahkan ke tim ", t.NamaTeam, ": ")
	fmt.Scanln(&jumlahAnggota)

	for i := 0; i < jumlahAnggota; i++ {
		var nama string
		var posisi string

		fmt.Printf("Nama anggota ke-%d: ", i+1)
		fmt.Scanln(&nama)
		fmt.Printf("Posisi anggota ke-%d: ", i+1)
		fmt.Scanln(&posisi)

		// Masukkan ke slot kosong
		for j := 0; j < MaxAnggota; j++ {
			if t.AnggotaTeam[j].NamaAnggota == "" {
				t.AnggotaTeam[j] = Anggota{NamaAnggota: nama, Posisi: posisi}
				break
			}
		}
	}
}

func tampilAnggota() {}
func hapusAnggota()  {}
func ubahAnggota()   {}
