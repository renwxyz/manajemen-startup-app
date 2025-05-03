package main

//global-global gunain PascalCase
//lokal-lokal gunain camelCase

//database pengguna

const MaxPengguna int = 255

type Pengguna struct {
	IdPengguna    int //auto increment by urutan daftar
	Nama          string
	Email         string
	KataSandi     string
	JenisPengguna string
}

var JumlahTerdaftar int
var DatabasePengguna [MaxPengguna]Pengguna

//[END] database Pengguna

//database Startup

const maxStartup int = 24
const MaxTeam int = 6
const MaxAnggota int = 12

type Startup struct {
	IdStartup      int //auto increment by urutan daftar
	IdPendiri      int
	NamaStartup    string
	BidangUsaha    string
	TahunBerdiri   int
	TotalPendanaan int
	DaftarTeam     [MaxTeam]Team
}

type Team struct {
	NamaTeam    string
	AnggotaTeam [MaxAnggota]Anggota
}

type Anggota struct {
	NamaAnggota string
	Posisi      string
}

var StartupTerdaftar int
var DatabaseStartup [maxStartup]Startup

//[END] database Startup
