package main

//global-global gunain PascalCase
//lokal-lokal gunain camelCase

//database pengguna

const MaxPengguna int = 255

type Pengguna struct {
	UserId        int //auto increment by urutan daftar
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
	StartupId      int //auto increment by urutan daftar
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

func masukanDataDummy() {
	// Dummy User
	DatabasePengguna[0] = Pengguna{
		UserId:        1,
		Nama:          "Rendhi",
		Email:         "rendhi19@gmail.com",
		KataSandi:     "rendhi123",
		JenisPengguna: "admin",
	}

	DatabasePengguna[1] = Pengguna{
		UserId:        2,
		Nama:          "Geugeut",
		Email:         "geugeut2@gmail.com",
		KataSandi:     "geugeut123",
		JenisPengguna: "pendiri",
	}

	DatabasePengguna[2] = Pengguna{
		UserId:        3,
		Nama:          "Bob",
		Email:         "bob@gmail.com",
		KataSandi:     "bob123",
		JenisPengguna: "investor",
	}

	DatabasePengguna[3] = Pengguna{
		UserId:        4,
		Nama:          "Alice",
		Email:         "alice@gmail.com",
		KataSandi:     "alice123",
		JenisPengguna: "admin",
	}

	DatabasePengguna[4] = Pengguna{
		UserId:        5,
		Nama:          "Charlie",
		Email:         "charlie@gmail.com",
		KataSandi:     "charlie123",
		JenisPengguna: "pendiri",
	}

	DatabasePengguna[5] = Pengguna{
		UserId:        6,
		Nama:          "Dina",
		Email:         "dina@gmail.com",
		KataSandi:     "dina123",
		JenisPengguna: "investor",
	}

	DatabasePengguna[6] = Pengguna{
		UserId:        7,
		Nama:          "Eka",
		Email:         "eka@gmail.com",
		KataSandi:     "eka123",
		JenisPengguna: "pendiri",
	}

	DatabasePengguna[7] = Pengguna{
		UserId:        8,
		Nama:          "Fajar",
		Email:         "fajar@gmail.com",
		KataSandi:     "fajar123",
		JenisPengguna: "admin",
	}

	DatabasePengguna[8] = Pengguna{
		UserId:        9,
		Nama:          "Gina",
		Email:         "gina@gmail.com",
		KataSandi:     "gina123",
		JenisPengguna: "investor",
	}

	DatabasePengguna[9] = Pengguna{
		UserId:        10,
		Nama:          "Hendra",
		Email:         "hendra@gmail.com",
		KataSandi:     "hendra123",
		JenisPengguna: "pendiri",
	}

	DatabasePengguna[10] = Pengguna{
		UserId:        11,
		Nama:          "Intan",
		Email:         "intan@gmail.com",
		KataSandi:     "intan123",
		JenisPengguna: "admin",
	}

	DatabasePengguna[11] = Pengguna{
		UserId:        12,
		Nama:          "Joko",
		Email:         "joko@gmail.com",
		KataSandi:     "joko123",
		JenisPengguna: "investor",
	}

	DatabasePengguna[12] = Pengguna{
		UserId:        13,
		Nama:          "Kiki",
		Email:         "kiki@gmail.com",
		KataSandi:     "kiki123",
		JenisPengguna: "pendiri",
	}

	DatabasePengguna[13] = Pengguna{
		UserId:        14,
		Nama:          "Luna",
		Email:         "luna@gmail.com",
		KataSandi:     "luna123",
		JenisPengguna: "admin",
	}

	DatabasePengguna[14] = Pengguna{
		UserId:        15,
		Nama:          "Miko",
		Email:         "miko@gmail.com",
		KataSandi:     "miko123",
		JenisPengguna: "investor",
	}

	JumlahTerdaftar = 15

	// [END] Dummy User

	// Dummy Startup
	DatabaseStartup[0] = Startup{
		StartupId:      1,
		NamaStartup:    "krisan",
		BidangUsaha:    "Traveling",
		TahunBerdiri:   2024,
		TotalPendanaan: 5000000000,
		DaftarTeam: [MaxTeam]Team{
			//Blok Data Team 1
			{
				NamaTeam: "pendiri",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Geugeut", Posisi: "CEO"},
					{NamaAnggota: "Rendhi", Posisi: "COO"},
				},
			},
			//[END] Blok Data Team 1

			//Blok Data Team 2
			{
				NamaTeam: "pengembangan produk",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Mark", Posisi: "Product Manager"},
					{NamaAnggota: "john", Posisi: "UI/UX Designer"},
				},
			},
			//[END] Blok Data Team 2
		},
	}

	DatabaseStartup[1] = Startup{
		StartupId:      2,
		NamaStartup:    "Cerah",
		BidangUsaha:    "Agritech",
		TahunBerdiri:   2023,
		TotalPendanaan: 2500000000,
		DaftarTeam: [MaxTeam]Team{
			{
				NamaTeam: "founders",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Dita", Posisi: "CEO"},
					{NamaAnggota: "Raka", Posisi: "CTO"},
				},
			},

			{
				NamaTeam: "engineering",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Andre", Posisi: "Backend Developer"},
					{NamaAnggota: "Sari", Posisi: "Frontend Developer"},
				},
			},
		},
	}

	DatabaseStartup[2] = Startup{
		StartupId:      3,
		NamaStartup:    "Lentera",
		BidangUsaha:    "EdTech",
		TahunBerdiri:   2021,
		TotalPendanaan: 3000000000,
		DaftarTeam: [MaxTeam]Team{
			{
				NamaTeam: "tim kreatif",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Putri", Posisi: "Creative Director"},
					{NamaAnggota: "Agus", Posisi: "Content Developer"},
				},
			},
			{
				NamaTeam: "ops",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Vina", Posisi: "COO"},
					{NamaAnggota: "Bayu", Posisi: "HR Manager"},
				},
			},
		},
	}

	DatabaseStartup[3] = Startup{
		StartupId:      4,
		NamaStartup:    "DigiFood",
		BidangUsaha:    "FoodTech",
		TahunBerdiri:   2022,
		TotalPendanaan: 4500000000,
		DaftarTeam: [MaxTeam]Team{
			{
				NamaTeam: "kepemimpinan",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Anita", Posisi: "CEO"},
					{NamaAnggota: "Dion", Posisi: "CFO"},
				},
			},
			{
				NamaTeam: "tim dapur",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Lukman", Posisi: "Head Chef"},
					{NamaAnggota: "Tina", Posisi: "Product Tester"},
				},
			},
		},
	}

	DatabaseStartup[4] = Startup{
		StartupId:      5,
		NamaStartup:    "Vibi",
		BidangUsaha:    "HealthTech",
		TahunBerdiri:   2020,
		TotalPendanaan: 6000000000,
		DaftarTeam: [MaxTeam]Team{
			{
				NamaTeam: "tim medis",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Dr. Sarah", Posisi: "Medical Advisor"},
					{NamaAnggota: "Dina", Posisi: "Health Consultant"},
				},
			},
			{
				NamaTeam: "dev tim",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Wawan", Posisi: "Mobile Developer"},
					{NamaAnggota: "Cindy", Posisi: "Backend Engineer"},
				},
			},
		},
	}

	DatabaseStartup[5] = Startup{
		StartupId:      6,
		NamaStartup:    "QuickSpace",
		BidangUsaha:    "PropertyTech",
		TahunBerdiri:   2023,
		TotalPendanaan: 3200000000,
		DaftarTeam: [MaxTeam]Team{
			{
				NamaTeam: "pengembang",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Riki", Posisi: "Fullstack Developer"},
					{NamaAnggota: "Mita", Posisi: "Project Manager"},
				},
			},
			{
				NamaTeam: "marketing",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Ayu", Posisi: "CMO"},
					{NamaAnggota: "Hendra", Posisi: "Content Strategist"},
				},
			},
		},
	}

	DatabaseStartup[6] = Startup{
		StartupId:      7,
		NamaStartup:    "RuangCerah",
		BidangUsaha:    "Mental Health",
		TahunBerdiri:   2021,
		TotalPendanaan: 2000000000,
		DaftarTeam: [MaxTeam]Team{
			{
				NamaTeam: "tim psikolog",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Dr. Rani", Posisi: "Psikolog Klinis"},
					{NamaAnggota: "Bagus", Posisi: "Psikolog Anak"},
				},
			},
			{
				NamaTeam: "tim teknologi",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Kevin", Posisi: "CTO"},
					{NamaAnggota: "Rizki", Posisi: "App Developer"},
				},
			},
		},
	}

	DatabaseStartup[7] = Startup{
		StartupId:      8,
		NamaStartup:    "SakuMini",
		BidangUsaha:    "FinTech",
		TahunBerdiri:   2019,
		TotalPendanaan: 7000000000,
		DaftarTeam: [MaxTeam]Team{
			{
				NamaTeam: "founders",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Farhan", Posisi: "CEO"},
					{NamaAnggota: "Nisa", Posisi: "CTO"},
				},
			},
			{
				NamaTeam: "keuangan",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Linda", Posisi: "Financial Analyst"},
					{NamaAnggota: "Ari", Posisi: "Compliance Officer"},
				},
			},
		},
	}

	DatabaseStartup[8] = Startup{
		StartupId:      9,
		NamaStartup:    "EcoRasa",
		BidangUsaha:    "Sustainable Food",
		TahunBerdiri:   2022,
		TotalPendanaan: 3500000000,
		DaftarTeam: [MaxTeam]Team{
			{
				NamaTeam: "tim inovasi",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Alam", Posisi: "Product Innovator"},
					{NamaAnggota: "Yani", Posisi: "Food Scientist"},
				},
			},
			{
				NamaTeam: "produksi",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Niko", Posisi: "Manufacturing Lead"},
					{NamaAnggota: "Ira", Posisi: "Quality Control"},
				},
			},
		},
	}

	DatabaseStartup[9] = Startup{
		StartupId:      10,
		NamaStartup:    "LangitPay",
		BidangUsaha:    "Digital Payment",
		TahunBerdiri:   2020,
		TotalPendanaan: 8000000000,
		DaftarTeam: [MaxTeam]Team{
			{
				NamaTeam: "kepemimpinan",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Tegar", Posisi: "CEO"},
					{NamaAnggota: "Lina", Posisi: "COO"},
				},
			},
			{
				NamaTeam: "pengembangan sistem",
				AnggotaTeam: [MaxAnggota]Anggota{
					{NamaAnggota: "Damar", Posisi: "DevOps Engineer"},
					{NamaAnggota: "Fitria", Posisi: "QA Engineer"},
				},
			},
		},
	}
	StartupTerdaftar = 10
	// [END] Dummy Startup
}
