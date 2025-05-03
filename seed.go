package main

func masukanDataDummy() {
	// Dummy User
	DatabasePengguna[0] = Pengguna{
		IdPengguna:    1,
		Nama:          "Rendhi",
		Email:         "rendhi19@gmail.com",
		KataSandi:     "rendhi123",
		JenisPengguna: "pendiri",
	}

	DatabasePengguna[1] = Pengguna{
		IdPengguna:    2,
		Nama:          "Geugeut",
		Email:         "geugeut2@gmail.com",
		KataSandi:     "geugeut123",
		JenisPengguna: "pendiri",
	}

	DatabasePengguna[2] = Pengguna{
		IdPengguna:    3,
		Nama:          "Bob",
		Email:         "bob@gmail.com",
		KataSandi:     "bob123",
		JenisPengguna: "investor",
	}

	DatabasePengguna[3] = Pengguna{
		IdPengguna:    4,
		Nama:          "Alice",
		Email:         "alice@gmail.com",
		KataSandi:     "alice123",
		JenisPengguna: "pendiri",
	}

	DatabasePengguna[4] = Pengguna{
		IdPengguna:    5,
		Nama:          "Charlie",
		Email:         "charlie@gmail.com",
		KataSandi:     "charlie123",
		JenisPengguna: "pendiri",
	}

	DatabasePengguna[5] = Pengguna{
		IdPengguna:    6,
		Nama:          "Dina",
		Email:         "dina@gmail.com",
		KataSandi:     "dina123",
		JenisPengguna: "investor",
	}

	DatabasePengguna[6] = Pengguna{
		IdPengguna:    7,
		Nama:          "Eka",
		Email:         "eka@gmail.com",
		KataSandi:     "eka123",
		JenisPengguna: "pendiri",
	}

	DatabasePengguna[7] = Pengguna{
		IdPengguna:    8,
		Nama:          "Fajar",
		Email:         "fajar@gmail.com",
		KataSandi:     "fajar123",
		JenisPengguna: "pendiri",
	}

	DatabasePengguna[8] = Pengguna{
		IdPengguna:    9,
		Nama:          "Gina",
		Email:         "gina@gmail.com",
		KataSandi:     "gina123",
		JenisPengguna: "investor",
	}

	DatabasePengguna[9] = Pengguna{
		IdPengguna:    10,
		Nama:          "Hendra",
		Email:         "hendra@gmail.com",
		KataSandi:     "hendra123",
		JenisPengguna: "pendiri",
	}

	DatabasePengguna[10] = Pengguna{
		IdPengguna:    11,
		Nama:          "Intan",
		Email:         "intan@gmail.com",
		KataSandi:     "intan123",
		JenisPengguna: "investor",
	}

	DatabasePengguna[11] = Pengguna{
		IdPengguna:    12,
		Nama:          "Joko",
		Email:         "joko@gmail.com",
		KataSandi:     "joko123",
		JenisPengguna: "investor",
	}

	DatabasePengguna[12] = Pengguna{
		IdPengguna:    13,
		Nama:          "Kiki",
		Email:         "kiki@gmail.com",
		KataSandi:     "kiki123",
		JenisPengguna: "pendiri",
	}

	DatabasePengguna[13] = Pengguna{
		IdPengguna:    14,
		Nama:          "Luna",
		Email:         "luna@gmail.com",
		KataSandi:     "luna123",
		JenisPengguna: "investor",
	}

	DatabasePengguna[14] = Pengguna{
		IdPengguna:    15,
		Nama:          "Miko",
		Email:         "miko@gmail.com",
		KataSandi:     "miko123",
		JenisPengguna: "investor",
	}

	JumlahTerdaftar = 15

	// [END] Dummy User

	// Dummy Startup
	DatabaseStartup[0] = Startup{
		IdStartup:      1,
		IdPendiri:      1,
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
		IdStartup:      2,
		IdPendiri:      1,
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
		IdStartup:      3,
		IdPendiri:      1,
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
		IdStartup:      4,
		IdPendiri:      2,
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
		IdStartup:      5,
		IdPendiri:      2,
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
		IdStartup:      6,
		IdPendiri:      3,
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
		IdStartup:      7,
		IdPendiri:      3,
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
		IdStartup:      8,
		IdPendiri:      4,
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
		IdStartup:      9,
		IdPendiri:      5,
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
		IdStartup:      10,
		IdPendiri:      5,
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
