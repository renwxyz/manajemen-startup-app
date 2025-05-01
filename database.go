package main

//database pengguna

const max_pengguna int = 255

var pengguna_terdaftar int = 3

type pengguna struct {
	nama       string
	email      string
	kata_sandi string
}

var database_pengguna = [max_pengguna]pengguna{
	{nama: "Rendhi", email: "rendhi@gmail.com", kata_sandi: "rendhi123"},
	{nama: "Bob", email: "bob@gmail.com", kata_sandi: "bob123"},
	{nama: "Alice", email: "alice@gmail.com", kata_sandi: "alice123"},
}

//database pengguna

const max_startup int = 24
const max_anggota int = 12
const max_team int = 6

var startup_terdaftar int = 2

type anggota struct {
	nama   string
	posisi string
}

type team struct {
	nama_team    string
	anggota_team [max_anggota]anggota
}

type startup struct {
	nama_startup    string
	bidang_usaha    string
	tahun_berdiri   int
	total_pendanaan int
	daftar_team     [max_team]team
}

var database_startup = [10]startup{
	// Blok data 1
	{
		nama_startup:    "microtrans",
		bidang_usaha:    "banking",
		tahun_berdiri:   2025,
		total_pendanaan: 1000000000,
		daftar_team: [max_team]team{
			{
				nama_team: "Tim Pendiri",
				anggota_team: [max_anggota]anggota{
					{nama: "Rendhi", posisi: "CEO"},
					{nama: "Mark", posisi: "COO"},
					{nama: "Dio", posisi: "CTO"},
				}},
			{
				nama_team: "Tim Pengembangan Produk",
				anggota_team: [max_anggota]anggota{
					{nama: "Alice", posisi: "Product Manager"},
					{nama: "Bob", posisi: "UI/UX Designer"},
				}},
		}},
	// Blok data 2
	{
		nama_startup:    "EduSmart",
		bidang_usaha:    "edtech",
		tahun_berdiri:   2023,
		total_pendanaan: 250000000,
		daftar_team: [max_team]team{
			{
				nama_team: "Tim Pendiri",
				anggota_team: [max_anggota]anggota{
					{nama: "Tania", posisi: "CEO"},
					{nama: "Yusuf", posisi: "CTO"},
				}},
			{
				nama_team: "Tim Akademik",
				anggota_team: [max_anggota]anggota{
					{nama: "Rani", posisi: "Curriculum Specialist"},
					{nama: "Hadi", posisi: "Learning Designer"},
				}},
		}},
	// Blok data 3
	{
		nama_startup:    "AgroLink",
		bidang_usaha:    "agritech",
		tahun_berdiri:   2022,
		total_pendanaan: 750000000,
		daftar_team: [max_team]team{
			{
				nama_team: "Tim Pendiri",
				anggota_team: [max_anggota]anggota{
					{nama: "Andi", posisi: "CEO"},
					{nama: "Budi", posisi: "CFO"},
				}},
			{
				nama_team: "Tim Teknologi",
				anggota_team: [max_anggota]anggota{
					{nama: "Rika", posisi: "Backend Engineer"},
					{nama: "Kevin", posisi: "Data Analyst"},
				}},
		}},
	// Blok data 4
	{
		nama_startup:    "MediCareX",
		bidang_usaha:    "healthtech",
		tahun_berdiri:   2021,
		total_pendanaan: 500000000,
		daftar_team: [max_team]team{
			{
				nama_team: "Tim Pendiri",
				anggota_team: [max_anggota]anggota{
					{nama: "Sari", posisi: "CEO"},
					{nama: "Rino", posisi: "COO"},
				}},
			{
				nama_team: "Tim Medis",
				anggota_team: [max_anggota]anggota{
					{nama: "Dr. Nadia", posisi: "Chief Medical Officer"},
					{nama: "Erwin", posisi: "Health Advisor"},
				}},
		}},
	// Blok data 5
	{
		nama_startup:    "LogiTrack",
		bidang_usaha:    "logistics",
		tahun_berdiri:   2024,
		total_pendanaan: 300000000,
		daftar_team: [max_team]team{
			{
				nama_team: "Tim Pendiri",
				anggota_team: [max_anggota]anggota{
					{nama: "Bayu", posisi: "CEO"},
					{nama: "Fina", posisi: "CTO"},
				}},
			{
				nama_team: "Tim Operasional",
				anggota_team: [max_anggota]anggota{
					{nama: "Nina", posisi: "Logistics Coordinator"},
					{nama: "Udin", posisi: "Fleet Manager"},
				}},
		}},
	// Blok data 6
	{
		nama_startup:    "FinSavr",
		bidang_usaha:    "fintech",
		tahun_berdiri:   2020,
		total_pendanaan: 900000000,
		daftar_team: [max_team]team{
			{
				nama_team: "Tim Pendiri",
				anggota_team: [max_anggota]anggota{
					{nama: "Dewi", posisi: "CEO"},
					{nama: "Joko", posisi: "COO"},
				}},
			{
				nama_team: "Tim Teknologi",
				anggota_team: [max_anggota]anggota{
					{nama: "Linda", posisi: "Mobile Developer"},
					{nama: "Surya", posisi: "Security Engineer"},
				}},
		}},
	// Blok data 7
	{
		nama_startup:    "TourNest",
		bidang_usaha:    "travel",
		tahun_berdiri:   2023,
		total_pendanaan: 150000000,
		daftar_team: [max_team]team{
			{
				nama_team: "Tim Pendiri",
				anggota_team: [max_anggota]anggota{
					{nama: "Rama", posisi: "CEO"},
					{nama: "Yanti", posisi: "CMO"},
				}},
			{
				nama_team: "Tim Layanan",
				anggota_team: [max_anggota]anggota{
					{nama: "Gina", posisi: "Customer Support"},
					{nama: "Eka", posisi: "Tour Planner"},
				}},
		}},
	// Blok data 8
	{
		nama_startup:    "Shopora",
		bidang_usaha:    "e-commerce",
		tahun_berdiri:   2025,
		total_pendanaan: 1200000000,
		daftar_team: [max_team]team{
			{
				nama_team: "Tim Pendiri",
				anggota_team: [max_anggota]anggota{
					{nama: "Hana", posisi: "CEO"},
					{nama: "Tomi", posisi: "CTO"},
				}},
			{
				nama_team: "Tim Marketplace",
				anggota_team: [max_anggota]anggota{
					{nama: "Ari", posisi: "Seller Manager"},
					{nama: "Vina", posisi: "Buyer Support"},
				}},
		}},
	// Blok data 9
	{
		nama_startup:    "SkillBoost",
		bidang_usaha:    "education",
		tahun_berdiri:   2019,
		total_pendanaan: 800000000,
		daftar_team: [max_team]team{
			{
				nama_team: "Tim Pendiri",
				anggota_team: [max_anggota]anggota{
					{nama: "Taufik", posisi: "CEO"},
					{nama: "Lisa", posisi: "COO"},
				}},
			{
				nama_team: "Tim Instruktur",
				anggota_team: [max_anggota]anggota{
					{nama: "Bambang", posisi: "Instructor Lead"},
					{nama: "Mira", posisi: "Content Developer"},
				}},
		}},
	// Blok data 10
	{
		nama_startup:    "CityCharge",
		bidang_usaha:    "energy",
		tahun_berdiri:   2024,
		total_pendanaan: 950000000,
		daftar_team: [max_team]team{
			{
				nama_team: "Tim Pendiri",
				anggota_team: [max_anggota]anggota{
					{nama: "Rio", posisi: "CEO"},
					{nama: "Zara", posisi: "CFO"},
				}},
			{
				nama_team: "Tim Infrastruktur",
				anggota_team: [max_anggota]anggota{
					{nama: "Fajar", posisi: "Engineer"},
					{nama: "Desi", posisi: "Project Manager"},
				}},
		}},
}
