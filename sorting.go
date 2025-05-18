package main

func insertionSortByTahunBerdiriAsc(db *[100]Startup, n int) {
	for i := 1; i < n; i++ {
		key := db[i]
		j := i - 1
		for j >= 0 && db[j].TahunBerdiri > key.TahunBerdiri {
			db[j+1] = db[j]
			j--
		}
		db[j+1] = key
	}
}

func insertionSortByTahunBerdiriDesc(db *[100]Startup, n int) {
	for i := 1; i < n; i++ {
		key := db[i]
		j := i - 1
		for j >= 0 && db[j].TahunBerdiri < key.TahunBerdiri {
			db[j+1] = db[j]
			j--
		}
		db[j+1] = key
	}
}

func selectionSortByTotalPendanaanAsc(db *[100]Startup, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if db[j].TotalPendanaan < db[minIdx].TotalPendanaan {
				minIdx = j
			}
		}
		// Tukar
		db[i], db[minIdx] = db[minIdx], db[i]
	}
}

func selectionSortByTotalPendanaanDesc(db *[100]Startup, n int) {
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if db[j].TotalPendanaan > db[maxIdx].TotalPendanaan {
				maxIdx = j
			}
		}
		// Tukar
		db[i], db[maxIdx] = db[maxIdx], db[i]
	}
}
