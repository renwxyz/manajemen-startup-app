package main

//global-global gunain PascalCase
//lokal-lokal gunain camelCase

//database pengguna

const MaxPengguna int = 255

type Pengguna struct {
	Nama          string
	Email         string
	KataSandi     string
	JenisPengguna string
}

var PenggunaTerdaftar int
var DatabasePengguna [MaxPengguna]Pengguna

//[END] database Pengguna

//database Startup

const max_startup int = 24
const MaxTeam int = 6
const MaxAnggota int = 12

type Startup struct {
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
	Nama   string
	Posisi string
}

var StartupTerdaftar int = 2
var DatabaseStartup [10]Startup

//[END] database Startup
