package main

type Pengguna struct {
	IdPengguna    int
	Nama          string
	Email         string
	KataSandi     string
	JenisPengguna string
}

var dbPengguna [100]Pengguna

type Startup struct {
	IdStartup      int
	IdPendiri      int
	NamaStartup    string
	BidangUsaha    string
	TahunBerdiri   int
	TotalPendanaan int
	DaftarTim      [100]Tim
}

type Tim struct {
	NamaTim       string
	DaftarAnggota [100]Anggota
}

type Anggota struct {
	NamaAnggota string
	Posisi      string
}

var dbStartup [100]Startup
