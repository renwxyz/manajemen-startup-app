package main

import "fmt"

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func uiMainMenu(pilihan *int) {
	fmt.Println("=== SELAMAT DATANG ===")
	fmt.Println("[1] Masuk")
	fmt.Println("[2] Daftar")
	fmt.Println("[0] Batal")
	fmt.Print("Pilihan: ")
	fmt.Scanln(pilihan)
}

func uiDashboardPendiri(penggunaAktif *Pengguna, pilihan *int) {

	fmt.Println("=== DASHBOARD PENDIRI ===")
	fmt.Printf("Halo, %s! Anda masuk sebagai pendiri.\n", penggunaAktif.Nama)
	fmt.Println("[1] Buat Startup")
	fmt.Println("[2] Tampilkan Startup saya")
	fmt.Println("[3] Tampilkan Semua Startup")
	fmt.Println("[4] Cari Startup")
	fmt.Println("[5] Tampilkan laporan jumlah per bidang usaha")
	fmt.Println("[0] Keluar")
	fmt.Print("Pilih menu: ")
	fmt.Scanln(pilihan)

}

func uiCariStartup(pilihan *int) {

	fmt.Println("[1] Cari berdasarkan nama")
	fmt.Println("[2] Cari berdasarkan bidang usaha")
	fmt.Println("[0] Batal")
	fmt.Print("Pilihan: ")
	fmt.Scanln(pilihan)

}

func uiInsertTeks(prompt string, nama *string) {
	fmt.Printf("Masukkan %s: ", prompt)
	fmt.Scanln(nama)
}

func uiInsertInt(prompt string, jumlah *int) {
	fmt.Printf("Masukkan %s: ", prompt)
	fmt.Scanln(jumlah)
}

func uiOpsiLanjutan(prompt string, pilihan *int) {
	fmt.Printf("[1] %s\n", prompt)
	fmt.Println("[0] Selesai")
	fmt.Print("Pilih opsi lanjutan: ")
	fmt.Scanln(pilihan)
}

func uiOpsiSelesai(pilihan *int) {
	fmt.Println("[0] Selesai")
	fmt.Print("Masukkan: ")
	fmt.Scanln(pilihan)
}

func uiInformasiTambah(prompt string, indeks int) {
	fmt.Printf("Tambah %s ke-%d", prompt, indeks)
}

func uiTampilkanStartupSaya(pilihan *int) {
	fmt.Println("[1] Detail semua startup Saya")
	fmt.Println("[2] Pilih detail Startup saya")
	fmt.Println("[3] Ubah data Startup")
	fmt.Println("[4] Hapus startup saya")
	fmt.Println("[0] Kembali")
	fmt.Print("Pilihan: ")
	fmt.Scanln(pilihan)
}

func uiHeaderTableStartup() {
	fmt.Println("============================================== === LIST STARTUP === ===============================================")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("%-15s %-15s %-15s %-25s %-30s %-20s %-20s\n", "No", "Id Startup", "Id Pendiri", "Nama Startup", "Bidang Usaha", "Tahun Berdiri", "Total Pendanaan")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------")

}

func uiFooterTable() {
	fmt.Println("------------------------------------------------------------------------------------------------------------------------")
}

func uiHeaderDetailStartup(i int) {
	fmt.Printf("=========== Detail Startup ke-%d ===========\n", i)
}

func uiUbahDataStartup(pilihan *int) {
	fmt.Println("[1] Ubah id pendiri")
	fmt.Println("[2] Ubah nama startup")
	fmt.Println("[3] Ubah bidang usaha")
	fmt.Println("[4] Ubah total pendanaan")
	fmt.Println("[5] Kelola tim")
	fmt.Println("[0] Selesai")
	fmt.Print("Pilih: ")
	fmt.Scanln(pilihan)
}

func uiKelolaTim(pilihan *int) {
	fmt.Println("[1] Ubah tim")
	fmt.Println("[2] Buat tim")
	fmt.Println("[3] Hapus tim")
	fmt.Println("[0] Kembali")
	fmt.Print("Pilihan: ")
	fmt.Scanln(pilihan)
}

func uiUbahTim(pilihan *int) {
	fmt.Println("[1] Ubah nama tim")
	fmt.Println("[2] Ubah anggota")
	fmt.Println("[3] Buat anggota")
	fmt.Println("[4] Hapus anggota")
	fmt.Println("[0] Kembali")
	fmt.Print("Pilihan: ")
	fmt.Scanln(pilihan)
}

func uiTampilkanSemuaStartup(pilihan *int) {
	fmt.Println("[1] Pilih detail Startup")
	fmt.Println("[2] Urutkan naik total pendanaan")
	fmt.Println("[3] Urutkan naik tahun berdiri")
	fmt.Println("[4] Urutkan turun total pendanaan")
	fmt.Println("[5] Urutkan turun tahun berdiri")
	fmt.Println("[0] Kembali")
	fmt.Print("Pilihan: ")
	fmt.Scanln(pilihan)
}

func uiDashboardInvestor(penggunaAktif *Pengguna, pilihan *int) {
	fmt.Println("=== DASHBOARD Investor ===")
	fmt.Printf("Halo, %s! Anda masuk sebagai Investor.\n", penggunaAktif.Nama)
	fmt.Println("[1] Tampilkan Semua Startup")
	fmt.Println("[2] Cari Startup")
	fmt.Println("[0] Keluar")
	fmt.Print("Pilih menu: ")
	fmt.Scanln(pilihan)
}

//

func pesanEror(pesan string) { fmt.Println("[EROR]" + pesan) }

//func pesanBerhasil(pesan string)  { fmt.Println("[BERHASIL]" + pesan) }
func pesanInformasi(pesan string) { fmt.Println("[INFO]" + pesan) }
