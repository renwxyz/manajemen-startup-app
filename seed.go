package main

func seedPengguna() {
	dbPengguna = [100]Pengguna{
		{IdPengguna: 1, Nama: "Rendhi", Email: "rendhi@gmail.com", KataSandi: "123", JenisPengguna: "pendiri"},
		{IdPengguna: 2, Nama: "Geugeut", Email: "geugeut@gmail.com", KataSandi: "456", JenisPengguna: "pendiri"},
		{IdPengguna: 3, Nama: "John", Email: "john@gmail.com", KataSandi: "789", JenisPengguna: "investor"},
		{IdPengguna: 4, Nama: "Mark", Email: "mark@gmail.com", KataSandi: "987", JenisPengguna: "investor"},
		{IdPengguna: 5, Nama: "Alice", Email: "alice@example.com", KataSandi: "abc", JenisPengguna: "pengguna"},
		{IdPengguna: 6, Nama: "Bob", Email: "bob@example.com", KataSandi: "def", JenisPengguna: "pengguna"},
		{IdPengguna: 7, Nama: "Charlie", Email: "charlie@example.com", KataSandi: "ghi", JenisPengguna: "investor"},
		{IdPengguna: 8, Nama: "Diana", Email: "diana@example.com", KataSandi: "jkl", JenisPengguna: "pengguna"},
		{IdPengguna: 9, Nama: "Ethan", Email: "ethan@example.com", KataSandi: "mno", JenisPengguna: "pendiri"},
		{IdPengguna: 10, Nama: "Fiona", Email: "fiona@example.com", KataSandi: "pqr", JenisPengguna: "pengguna"},
	}
}

func seedStartup() {
	dbStartup[0] = Startup{
		IdStartup:      1,
		IdPendiri:      1, // Rendhi
		NamaStartup:    "Krisan",
		BidangUsaha:    "Traveling",
		TahunBerdiri:   2024,
		TotalPendanaan: 1000000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Geugeut", Posisi: "CEO"},
					{NamaAnggota: "Rendhi", Posisi: "CEO"},
				},
			},
		},
	}

	dbStartup[1] = Startup{
		IdStartup:      2,
		IdPendiri:      1, // Rendhi (bisa jadi co-founder di startup lain)
		NamaStartup:    "Cerah",
		BidangUsaha:    "Agritech",
		TahunBerdiri:   2023,
		TotalPendanaan: 2500000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Dita", Posisi: "CEO"}, // Nama anggota bisa di luar daftar pendiri sebelumnya
					{NamaAnggota: "Raka", Posisi: "CTO"}, // Nama anggota bisa di luar daftar pendiri sebelumnya
				},
			},
			{
				NamaTim: "engineering",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Andre", Posisi: "BackendDeveloper"},
					{NamaAnggota: "Sari", Posisi: "FrontendDeveloper"},
				},
			},
		},
	}

	dbStartup[2] = Startup{
		IdStartup:      3,
		IdPendiri:      2, // Geugeut
		NamaStartup:    "Lentera",
		BidangUsaha:    "EdTech",
		TahunBerdiri:   2021,
		TotalPendanaan: 3000000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "timkreatif",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Putri", Posisi: "CreativeDirector"},
					{NamaAnggota: "Agus", Posisi: "ContentDeveloper"},
				},
			},
			{
				NamaTim: "ops",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Vina", Posisi: "COO"},
					{NamaAnggota: "Bayu", Posisi: "HRManager"},
				},
			},
		},
	}

	// Generate data Startup tambahan (dari index 3 hingga 24)
	dbStartup[3] = Startup{
		IdStartup:      4,
		IdPendiri:      9, // Ethan
		NamaStartup:    "SyntaxSolutions",
		BidangUsaha:    "SoftwareDevelopment",
		TahunBerdiri:   2022,
		TotalPendanaan: 1500000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Ethan", Posisi: "CEO"},
				},
			},
			{
				NamaTim: "development",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Budi", Posisi: "LeadDeveloper"},
					{NamaAnggota: "Citra", Posisi: "Developer"},
					{NamaAnggota: "Eko", Posisi: "QAAnalyst"},
				},
			},
		},
	}

	dbStartup[4] = Startup{
		IdStartup:      5,
		IdPendiri:      15, // Kevin
		NamaStartup:    "HealthBridge",
		BidangUsaha:    "HealthTech",
		TahunBerdiri:   2023,
		TotalPendanaan: 4000000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Kevin", Posisi: "CEO"},
					{NamaAnggota: "Lina", Posisi: "MedicalAdvisor"},
				},
			},
			{
				NamaTim: "product",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Maya", Posisi: "ProductManager"},
					{NamaAnggota: "Nico", Posisi: "UX/UIDesigner"},
				},
			},
		},
	}

	dbStartup[5] = Startup{
		IdStartup:      6,
		IdPendiri:      20, // Penelope
		NamaStartup:    "EcoHarvest",
		BidangUsaha:    "AgriTech",
		TahunBerdiri:   2024,
		TotalPendanaan: 1800000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Penelope", Posisi: "CEO"},
				},
			},
			{
				NamaTim: "operations",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Agus", Posisi: "FarmManager"},
					{NamaAnggota: "Bambang", Posisi: "SupplyChain"},
				},
			},
			{
				NamaTim: "sales",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Dewi", Posisi: "SalesLead"},
				},
			},
		},
	}

	dbStartup[6] = Startup{
		IdStartup:      7,
		IdPendiri:      26, // Victor
		NamaStartup:    "FinFlow",
		BidangUsaha:    "FinTech",
		TahunBerdiri:   2022,
		TotalPendanaan: 5000000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Victor", Posisi: "CEO"},
					{NamaAnggota: "Santi", Posisi: "CFO"},
				},
			},
			{
				NamaTim: "engineering",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Joko", Posisi: "BackendEngineer"},
					{NamaAnggota: "Rina", Posisi: "FrontendEngineer"},
					{NamaAnggota: "Tomi", Posisi: "DevOpsEngineer"},
				},
			},
			{
				NamaTim: "compliance",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Wahyu", Posisi: "ComplianceOfficer"},
				},
			},
		},
	}

	dbStartup[7] = Startup{
		IdStartup:      8,
		IdPendiri:      31, // Adrian
		NamaStartup:    "LearnSphere",
		BidangUsaha:    "EdTech",
		TahunBerdiri:   2023,
		TotalPendanaan: 2200000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Adrian", Posisi: "CEO"},
					{NamaAnggota: "Bella", Posisi: "HeadContent"},
				},
			},
			{
				NamaTim: "content",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Candra", Posisi: "ContentCreator"},
					{NamaAnggota: "Dian", Posisi: "CurriculumDeveloper"},
				},
			},
		},
	}

	dbStartup[8] = Startup{
		IdStartup:      9,
		IdPendiri:      37, // Gabriel
		NamaStartup:    "ArtisanMarket",
		BidangUsaha:    "E-commerce",
		TahunBerdiri:   2021,
		TotalPendanaan: 1200000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Gabriel", Posisi: "CEO"},
					{NamaAnggota: "Hailey", Posisi: "HeadMarketing"},
				},
			},
			{
				NamaTim: "marketing",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Ivan", Posisi: "DigitalMarketer"},
					{NamaAnggota: "Jessica", Posisi: "SocialMediaSpecialist"},
				},
			},
			{
				NamaTim: "operations",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Koko", Posisi: "LogisticsManager"},
				},
			},
		},
	}

	dbStartup[9] = Startup{
		IdStartup:      10,
		IdPendiri:      42, // Lily
		NamaStartup:    "UrbanGreen",
		BidangUsaha:    "SmartCity",
		TahunBerdiri:   2024,
		TotalPendanaan: 3500000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Lily", Posisi: "CEO"},
					{NamaAnggota: "Mason", Posisi: "CTO"},
				},
			},
			{
				NamaTim: "engineering",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Nina", Posisi: "HardwareEngineer"},
					{NamaAnggota: "Omar", Posisi: "SoftwareEngineer"},
					{NamaAnggota: "Putri", Posisi: "DataScientist"},
				},
			},
			{
				NamaTim: "partnerships",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Quentin", Posisi: "PartnershipManager"},
				},
			},
		},
	}

	dbStartup[10] = Startup{
		IdStartup:      11,
		IdPendiri:      48, // Riley
		NamaStartup:    "FoodieConnect",
		BidangUsaha:    "FoodTech",
		TahunBerdiri:   2022,
		TotalPendanaan: 2800000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Riley", Posisi: "CEO"},
					{NamaAnggota: "Samuel", Posisi: "HeadOperations"},
				},
			},
			{
				NamaTim: "operations",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Tania", Posisi: "Logistics"},
					{NamaAnggota: "Umar", Posisi: "Partnerships"},
				},
			},
			{
				NamaTim: "marketing",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Vina", Posisi: "MarketingSpecialist"},
					{NamaAnggota: "Wayan", Posisi: "CommunityManager"},
				},
			},
		},
	}

	dbStartup[11] = Startup{
		IdStartup:      12,
		IdPendiri:      53, // William
		NamaStartup:    "TravelWander",
		BidangUsaha:    "TravelTech",
		TahunBerdiri:   2023,
		TotalPendanaan: 1900000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "William", Posisi: "CEO"},
				},
			},
			{
				NamaTim: "product",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Ximena", Posisi: "ProductLead"},
					{NamaAnggota: "Yusuf", Posisi: "UXDesigner"},
				},
			},
			{
				NamaTim: "customersupport",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Zoe", Posisi: "SupportManager"},
				},
			},
		},
	}

	dbStartup[12] = Startup{
		IdStartup:      13,
		IdPendiri:      59, // Christopher
		NamaStartup:    "FitZone",
		BidangUsaha:    "FitnessTech",
		TahunBerdiri:   2024,
		TotalPendanaan: 1100000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Christopher", Posisi: "CEO"},
					{NamaAnggota: "Destiny", Posisi: "HeadFitness"},
				},
			},
			{
				NamaTim: "appdevelopment",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Eric", Posisi: "MobileDeveloper"},
					{NamaAnggota: "Gabriella", Posisi: "BackendDeveloper"},
				},
			},
		},
	}

	dbStartup[13] = Startup{
		IdStartup:      14,
		IdPendiri:      64, // Isabelle
		NamaStartup:    "HomeEase",
		BidangUsaha:    "PropertyTech",
		TahunBerdiri:   2023,
		TotalPendanaan: 3200000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Isabelle", Posisi: "CEO"},
					{NamaAnggota: "Jackson", Posisi: "COO"},
				},
			},
			{
				NamaTim: "sales",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Kylie", Posisi: "SalesManager"},
					{NamaAnggota: "Liam", Posisi: "SalesExecutive"},
				},
			},
			{
				NamaTim: "marketing",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Maya", Posisi: "MarketingSpecialist"},
				},
			},
		},
	}

	dbStartup[14] = Startup{
		IdStartup:      15,
		IdPendiri:      70, // Olivia
		NamaStartup:    "FashionNova",
		BidangUsaha:    "Ecommerce(Fashion)", // Changed E-commerce (Fashion) to Ecommerce(Fashion)
		TahunBerdiri:   2021,
		TotalPendanaan: 2600000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Olivia", Posisi: "CEO"},
				},
			},
			{
				NamaTim: "merchandising",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Peyton", Posisi: "Buyer"},
					{NamaAnggota: "Quincy", Posisi: "InventoryManager"},
				},
			},
			{
				NamaTim: "marketing",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Reagan", Posisi: "SocialMediaManager"},
					{NamaAnggota: "Skylar", Posisi: "InfluencerRelations"},
				},
			},
		},
	}

	dbStartup[15] = Startup{
		IdStartup:      16,
		IdPendiri:      75, // Tyler
		NamaStartup:    "GameCraft",
		BidangUsaha:    "GameDevelopment",
		TahunBerdiri:   2024,
		TotalPendanaan: 4500000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Tyler", Posisi: "CEO"},
					{NamaAnggota: "Violet", Posisi: "LeadGameDesigner"},
				},
			},
			{
				NamaTim: "development",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Wyatt", Posisi: "UnityDeveloper"},
					{NamaAnggota: "Zelda", Posisi: "Artist"},
					{NamaAnggota: "Aaron", Posisi: "SoundDesigner"},
				},
			},
		},
	}

	dbStartup[16] = Startup{
		IdStartup:      17,
		IdPendiri:      81, // Cameron
		NamaStartup:    "CleanEnergySolutions",
		BidangUsaha:    "RenewableEnergy",
		TahunBerdiri:   2023,
		TotalPendanaan: 6000000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Cameron", Posisi: "CEO"},
				},
			},
			{
				NamaTim: "engineering",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Delilah", Posisi: "ElectricalEngineer"},
					{NamaAnggota: "Evan", Posisi: "MechanicalEngineer"},
				},
			},
			{
				NamaTim: "partnerships",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Genesis", Posisi: "BusinessDevelopment"},
				},
			},
		},
	}

	dbStartup[17] = Startup{
		IdStartup:      18,
		IdPendiri:      86, // Ivy
		NamaStartup:    "PetPal",
		BidangUsaha:    "PetCareTech",
		TahunBerdiri:   2022,
		TotalPendanaan: 1700000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Ivy", Posisi: "CEO"},
					{NamaAnggota: "Jaxon", Posisi: "VeterinaryAdvisor"},
				},
			},
			{
				NamaTim: "platform",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Kennedy", Posisi: "PlatformManager"},
					{NamaAnggota: "Leo", Posisi: "CommunityManager"},
				},
			},
		},
	}

	dbStartup[18] = Startup{
		IdStartup:      19,
		IdPendiri:      1, // Rendhi (example of multiple startups by one founder)
		NamaStartup:    "WorkFlow",
		BidangUsaha:    "ProductivitySoftware",
		TahunBerdiri:   2024,
		TotalPendanaan: 2100000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Rendhi", Posisi: "CEO"},
				},
			},
			{
				NamaTim: "development",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Madison", Posisi: "FullstackDeveloper"},
					{NamaAnggota: "Nathan", Posisi: "UI/UXDesigner"},
				},
			},
		},
	}

	dbStartup[19] = Startup{
		IdStartup:      20,
		IdPendiri:      2, // Geugeut (example of multiple startups by one founder)
		NamaStartup:    "CultureConnect",
		BidangUsaha:    "CulturalExchangePlatform",
		TahunBerdiri:   2023,
		TotalPendanaan: 1400000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Geugeut", Posisi: "CEO"},
					{NamaAnggota: "Oscar", Posisi: "HeadPartnerships"},
				},
			},
			{
				NamaTim: "community",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Penelope", Posisi: "CommunityManager"},
					{NamaAnggota: "Quinn", Posisi: "EventCoordinator"},
				},
			},
			{
				NamaTim: "content",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Ryan", Posisi: "ContentCurator"},
				},
			},
		},
	}

	dbStartup[20] = Startup{
		IdStartup:      21,
		IdPendiri:      9, // Ethan
		NamaStartup:    "AquaTechInnovations",
		BidangUsaha:    "WaterTechnology",
		TahunBerdiri:   2024,
		TotalPendanaan: 3800000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Ethan", Posisi: "CEO"},
					{NamaAnggota: "Sophia", Posisi: "LeadResearcher"},
				},
			},
			{
				NamaTim: "R&D",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Thomas", Posisi: "ResearchScientist"},
					{NamaAnggota: "Uma", Posisi: "LabTechnician"},
				},
			},
		},
	}

	dbStartup[21] = Startup{
		IdStartup:      22,
		IdPendiri:      15, // Kevin
		NamaStartup:    "EduSpark",
		BidangUsaha:    "EdTech",
		TahunBerdiri:   2023,
		TotalPendanaan: 2400000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Kevin", Posisi: "CEO"},
					{NamaAnggota: "Victor", Posisi: "HeadEducation"}, // Victor is also a founder of another startup
				},
			},
			{
				NamaTim: "curriculum",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Wendy", Posisi: "CurriculumDesigner"},
					{NamaAnggota: "Xavier", Posisi: "InstructionalDesigner"},
				},
			},
			{
				NamaTim: "platform",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Yvonne", Posisi: "PlatformManager"},
					{NamaAnggota: "Zachary", Posisi: "UIDesigner"},
				},
			},
		},
	}

	dbStartup[22] = Startup{
		IdStartup:      23,
		IdPendiri:      20, // Penelope
		NamaStartup:    "GreenThumbApp",
		BidangUsaha:    "AgriTech",
		TahunBerdiri:   2024,
		TotalPendanaan: 1600000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Penelope", Posisi: "CEO"},
					{NamaAnggota: "Adrian", Posisi: "AgronomyAdvisor"}, // Adrian is also a founder
				},
			},
			{
				NamaTim: "development",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Bella", Posisi: "MobileDeveloper"},
					{NamaAnggota: "Caleb", Posisi: "BackendDeveloper"},
				},
			},
		},
	}

	dbStartup[23] = Startup{
		IdStartup:      24,
		IdPendiri:      26, // Victor
		NamaStartup:    "SecureChain",
		BidangUsaha:    "Cybersecurity",
		TahunBerdiri:   2023,
		TotalPendanaan: 5500000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Victor", Posisi: "CEO"},
				},
			},
			{
				NamaTim: "securityresearch",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Daisy", Posisi: "SecurityAnalyst"},
					{NamaAnggota: "Elijah", Posisi: "ThreatHunter"},
					{NamaAnggota: "Faith", Posisi: "Cryptographer"},
				},
			},
			{
				NamaTim: "consulting",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Gabriel", Posisi: "SecurityConsultant"}, // Gabriel is also a founder
					{NamaAnggota: "Hailey", Posisi: "PenetrationTester"},   // Hailey is also in another startup
				},
			},
		},
	}

	dbStartup[24] = Startup{
		IdStartup:      25,
		IdPendiri:      31, // Adrian
		NamaStartup:    "CreativeHub",
		BidangUsaha:    "CreativeServicesPlatform",
		TahunBerdiri:   2024,
		TotalPendanaan: 1300000000,
		DaftarTim: [100]Tim{
			{
				NamaTim: "pendiri",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Adrian", Posisi: "CEO"},
				},
			},
			{
				NamaTim: "platformteam",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Isaac", Posisi: "PlatformManager"},
					{NamaAnggota: "Jasmine", Posisi: "UX/UIDesigner"},
				},
			},
			{
				NamaTim: "communitymanagement",
				DaftarAnggota: [100]Anggota{
					{NamaAnggota: "Kayden", Posisi: "CommunityLead"},
				},
			},
		},
	}
}
