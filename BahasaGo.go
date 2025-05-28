package main

import "fmt"

// --- Tipe Bentukan (Struct) ---
type Aktivitas struct {
	ID           int
	Kategori     string  // Contoh: Transportasi, Makanan, Energi, Sampah
	Deskripsi    string  // Lebih umum dari Nama, bisa "Naik motor 10km"
	DampakKarbon float64 // Dalam kg CO2e
	Frekuensi    int     // Berapa kali dilakukan dalam periode tertentu (misal: per bulan)
}

// --- Variabel Global (Array Statis) ---
const MAX_AKTIVITAS = 100 // Ukuran maksimum array
var daftarAktivitas [MAX_AKTIVITAS]Aktivitas
var jumlahAktivitas int // Jumlah aktivitas yang saat ini tersimpan dalam array

// --- Fungsi Bantu (untuk String, karena tanpa "strings" package) ---
// Mengonversi string ke huruf kecil
func toLower(s string) string {
	var result []rune
	// Iterasi melalui setiap karakter (rune) dalam string
	for _, r := range s {
		// Jika karakter adalah huruf kapital (A-Z)
		if r >= 'A' && r <= 'Z' {
			// Ubah menjadi huruf kecil dengan menambahkan selisih 'a' dan 'A'
			result = append(result, r+'a'-'A')
		} else {
			// Jika bukan huruf kapital, biarkan apa adanya
			result = append(result, r)
		}
	}
	return string(result)
}

// Mengecek apakah string s mengandung substring sub (case-insensitive)
// Menggunakan implementasi sederhana pengganti strings.Contains
func stringContains(s, sub string) bool {
	sLower := toLower(s)
	subLower := toLower(sub)

	// Jika substring yang dicari kosong, anggap selalu terkandung
	if len(subLower) == 0 {
		return true
	}
	// Jika panjang string utama lebih pendek dari substring, tidak mungkin terkandung
	if len(sLower) < len(subLower) {
		return false
	}

	// Iterasi melalui string utama untuk mencari substring
	for i := 0; i <= len(sLower)-len(subLower); i++ {
		match := true
		// Bandingkan karakter per karakter
		for j := 0; j < len(subLower); j++ {
			if sLower[i+j] != subLower[j] {
				match = false
				// Tidak menggunakan break, inner loop akan selesai
				// dan 'match' akan tetap false jika ada ketidakcocokan
			}
		}
		if match {
			return true // Substring ditemukan
		}
	}
	return false // Substring tidak ditemukan
}

// --- Fungsi untuk mendapatkan input dari pengguna ---
// Catatan: fmt.Scan hanya membaca hingga spasi. Untuk deskripsi atau kategori multi-kata,
// ini akan membaca hanya kata pertama.

// Fungsi untuk mendapatkan input string dari pengguna
func getInputString(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input) // Membaca satu kata
	return input
}

// Fungsi untuk mendapatkan input integer dari pengguna
func getInputInt(prompt string) int {
	var input int
	fmt.Print(prompt)
	for { // Loop tak terbatas hingga input valid
		_, err := fmt.Scan(&input)
		if err == nil {
			return input // Input valid, keluar dari loop
		}
		fmt.Println("Input tidak valid. Masukkan angka.")
		// Membersihkan buffer input jika ada karakter yang tidak sesuai
		// Ini penting karena fmt.Scan tidak mengonsumsi newline setelah error
		var dummy string
		fmt.Scanln(&dummy) // Baca sisa baris untuk membersihkan buffer
	}
}

// Fungsi untuk mendapatkan input float64 dari pengguna
func getInputFloat(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	for { // Loop tak terbatas hingga input valid
		_, err := fmt.Scan(&input)
		if err == nil {
			return input // Input valid, keluar dari loop
		}
		fmt.Println("Input tidak valid. Masukkan angka desimal.")
		// Membersihkan buffer input
		var dummy string
		fmt.Scanln(&dummy) // Baca sisa baris untuk membersihkan buffer
	}
}

// --- Pencarian: Cari ID Aktivitas (untuk internal) ---
// Mengembalikan indeks aktivitas berdasarkan ID, atau -1 jika tidak ditemukan
func findAktivitasIndexByID(id int) int {
	for i := 0; i < jumlahAktivitas; i++ {
		if daftarAktivitas[i].ID == id {
			return i // ID ditemukan pada indeks ini
		}
	}
	return -1 // ID tidak ditemukan
}

// --- Menu: 1. Tambah Aktivitas ---
// Mengambil parameter sesuai permintaan di main
func tambahAktivitas(id int, kategori, deskripsi string, dampak float64, frekuensi int) {
	if jumlahAktivitas >= MAX_AKTIVITAS {
		fmt.Println("Maaf, kapasitas aktivitas sudah penuh (maksimal 100).")
		return // Keluar dari fungsi
	}

	// Cek apakah ID sudah ada
	for i := 0; i < jumlahAktivitas; i++ {
		if daftarAktivitas[i].ID == id {
			fmt.Println("Gagal menambahkan: ID sudah ada. Gunakan ID lain atau edit aktivitas yang sudah ada.")
			return // Keluar dari fungsi
		}
	}

	// Buat objek Aktivitas baru
	aktivitasBaru := Aktivitas{
		ID:           id,
		Kategori:     kategori,
		Deskripsi:    deskripsi,
		DampakKarbon: dampak,
		Frekuensi:    frekuensi,
	}

	// Tambahkan aktivitas baru ke array dan perbarui jumlahAktivitas
	daftarAktivitas[jumlahAktivitas] = aktivitasBaru
	jumlahAktivitas++
	fmt.Println("Aktivitas berhasil ditambahkan!")
}

// --- Menu: 2. Cari Aktivitas (Sequential & Binary Search) ---

// Implementasi Sequential Search
// Mencari aktivitas berdasarkan Kategori (case-insensitive)
// Mengembalikan indeks pertama yang ditemukan, atau -1 jika tidak ditemukan
func cariSequential(kategoriCari string) int {
	kategoriCariLower := toLower(kategoriCari)
	for i := 0; i < jumlahAktivitas; i++ {
		if toLower(daftarAktivitas[i].Kategori) == kategoriCariLower {
			return i // Mengembalikan indeks pertama yang cocok
		}
	}
	return -1 // Tidak ditemukan
}

// Fungsi bantu untuk mengurutkan daftarAktivitas berdasarkan Kategori untuk Binary Search
// Menggunakan Selection Sort untuk mengurutkan
func selectionSortKategori() {
	for i := 0; i < jumlahAktivitas-1; i++ {
		minIndex := i
		for j := i + 1; j < jumlahAktivitas; j++ {
			// Membandingkan string kategori secara alfabetis (case-insensitive)
			if toLower(daftarAktivitas[j].Kategori) < toLower(daftarAktivitas[minIndex].Kategori) {
				minIndex = j
			}
		}
		// Tukar elemen jika ditemukan yang lebih kecil
		daftarAktivitas[i], daftarAktivitas[minIndex] = daftarAktivitas[minIndex], daftarAktivitas[i]
	}
}

// Implementasi Binary Search (membutuhkan data terurut berdasarkan Kategori)
// Mengembalikan indeks dalam array, atau -1 jika tidak ditemukan
func cariBinary(kategoriCari string) int {
	kategoriCariLower := toLower(kategoriCari)
	low := 0
	high := jumlahAktivitas - 1

	for low <= high { // Loop selama batas bawah tidak melebihi batas atas
		mid := low + (high-low)/2 // Hitung indeks tengah
		currentKategoriLower := toLower(daftarAktivitas[mid].Kategori)

		if currentKategoriLower == kategoriCariLower {
			return mid // Kategori ditemukan pada indeks tengah
		} else if currentKategoriLower < kategoriCariLower {
			low = mid + 1 // Kategori yang dicari ada di paruh kanan
		} else { // currentKategoriLower > kategoriCariLower
			high = mid - 1 // Kategori yang dicari ada di paruh kiri
		}
	}
	return -1 // Kategori tidak ditemukan
}

// --- Menu: 3. Edit Aktivitas & Hapus Aktivitas ---
// Mengambil parameter sesuai permintaan di main, termasuk isDelete
func editAktivitas(id int, kategori, deskripsi string, dampak float64, frekuensi int, isDelete bool) {
	foundIndex := findAktivitasIndexByID(id) // Cari indeks aktivitas berdasarkan ID

	if foundIndex != -1 { // Jika aktivitas ditemukan
		if isDelete {
			// Logika Hapus Aktivitas
			// Geser semua elemen setelah yang dihapus ke depan untuk mengisi "lubang"
			for i := foundIndex; i < jumlahAktivitas-1; i++ {
				daftarAktivitas[i] = daftarAktivitas[i+1]
			}
			// Kosongkan elemen terakhir (opsional, tapi baik untuk kebersihan)
			daftarAktivitas[jumlahAktivitas-1] = Aktivitas{}
			jumlahAktivitas-- // Kurangi jumlah aktivitas yang valid
			fmt.Println("Aktivitas berhasil dihapus!")
		} else {
			// Logika Edit Aktivitas
			daftarAktivitas[foundIndex].Kategori = kategori
			daftarAktivitas[foundIndex].Deskripsi = deskripsi
			daftarAktivitas[foundIndex].DampakKarbon = dampak
			daftarAktivitas[foundIndex].Frekuensi = frekuensi
			fmt.Println("Aktivitas berhasil diupdate!")
		}
	} else {
		fmt.Println("Aktivitas dengan ID tersebut tidak ditemukan.")
	}
}

// --- Menu: 4. Urutkan (Selection Sort & Insertion Sort) ---

// Implementasi Selection Sort (berdasarkan Dampak Karbon)
// Urutan ditentukan oleh parameter 'ascending'
func selectionSortDampak(ascending bool) {
	for i := 0; i < jumlahAktivitas-1; i++ {
		extremeIndex := i // Indeks elemen terkecil (jika ascending) atau terbesar (jika descending)
		for j := i + 1; j < jumlahAktivitas; j++ {
			if ascending { // Urutkan dari dampak terkecil ke terbesar
				if daftarAktivitas[j].DampakKarbon < daftarAktivitas[extremeIndex].DampakKarbon {
					extremeIndex = j
				}
			} else { // Urutkan dari dampak terbesar ke terkecil
				if daftarAktivitas[j].DampakKarbon > daftarAktivitas[extremeIndex].DampakKarbon {
					extremeIndex = j
				}
			}
		}
		// Tukar elemen saat ini dengan elemen ekstrem yang ditemukan
		daftarAktivitas[i], daftarAktivitas[extremeIndex] = daftarAktivitas[extremeIndex], daftarAktivitas[i]
	}
}

// Implementasi Insertion Sort (berdasarkan Frekuensi)
// Urutan ditentukan oleh parameter 'ascending'
func insertionSortFrekuensi(ascending bool) {
	for i := 1; i < jumlahAktivitas; i++ {
		key := daftarAktivitas[i] // Elemen yang akan disisipkan
		j := i - 1
		// Pindahkan elemen daftarAktivitas[0..i-1] yang lebih besar/kecil dari key
		// ke satu posisi di depan posisi mereka saat ini
		if ascending { // Urutkan dari frekuensi terkecil ke terbesar
			for j >= 0 && daftarAktivitas[j].Frekuensi > key.Frekuensi {
				daftarAktivitas[j+1] = daftarAktivitas[j]
				j = j - 1
			}
		} else { // Urutkan dari frekuensi terbesar ke terkecil
			for j >= 0 && daftarAktivitas[j].Frekuensi < key.Frekuensi {
				daftarAktivitas[j+1] = daftarAktivitas[j]
				j = j - 1
			}
		}
		daftarAktivitas[j+1] = key // Sisipkan key di posisi yang benar
	}
}

// --- Menu: 5. Tampilkan Daftar Aktivitas ---
func tampilkanDaftar() {
	fmt.Println("\n--- Daftar Aktivitas ---")
	if jumlahAktivitas == 0 {
		fmt.Println("Belum ada aktivitas yang ditambahkan.")
		return
	}
	// Header tabel
	fmt.Printf("%-5s | %-15s | %-20s | %-10s | %-10s\n", "ID", "Kategori", "Deskripsi", "Dampak (kg)", "Frekuensi")
	fmt.Println("----------------------------------------------------------------------")
	// Cetak setiap aktivitas yang valid
	for i := 0; i < jumlahAktivitas; i++ {
		a := daftarAktivitas[i]
		fmt.Printf("%-5d | %-15s | %-20s | %-10.2f | %-10d\n", a.ID, a.Kategori, a.Deskripsi, a.DampakKarbon, a.Frekuensi)
	}
}

// --- Menu: 6. Laporan Bulanan & Rekomendasi ---
func laporanBulanan() {
	if jumlahAktivitas == 0 {
		fmt.Println("Belum ada aktivitas untuk membuat laporan.")
		return
	}

	fmt.Println("\n--- Laporan Bulanan Jejak Karbon ---")

	totalJejakKarbon := 0.0
	// Inisialisasi dengan nilai yang memastikan update pertama untuk pencarian nilai ekstrim
	aktivitasTerbanyakDampak := Aktivitas{DampakKarbon: -1.0}
	aktivitasTerdikitDampak := Aktivitas{DampakKarbon: 999999999.0}

	// Hitung total jejak karbon dan cari aktivitas dengan dampak max/min
	for i := 0; i < jumlahAktivitas; i++ {
		a := daftarAktivitas[i]
		// Total jejak karbon dihitung dari dampak per aktivitas dikalikan frekuensi
		totalJejakKarbon += a.DampakKarbon * float64(a.Frekuensi)

		// Pencarian Nilai Ekstrim (Max/Min)
		if a.DampakKarbon > aktivitasTerbanyakDampak.DampakKarbon {
			aktivitasTerbanyakDampak = a
		}
		if a.DampakKarbon < aktivitasTerdikitDampak.DampakKarbon {
			aktivitasTerdikitDampak = a
		}
	}

	fmt.Printf("Total Jejak Karbon Bulan Ini: %.2f kg CO2e\n", totalJejakKarbon)

	// Tampilkan aktivitas dengan dampak ekstrim jika ada data valid
	if aktivitasTerbanyakDampak.DampakKarbon != -1.0 {
		fmt.Printf("Aktivitas dengan Dampak Karbon Terbesar: '%s' (%.2f kg CO2e)\n", aktivitasTerbanyakDampak.Deskripsi, aktivitasTerbanyakDampak.DampakKarbon)
	}
	if aktivitasTerdikitDampak.DampakKarbon != 999999999.0 {
		fmt.Printf("Aktivitas dengan Dampak Karbon Terkecil: '%s' (%.2f kg CO2e)\n", aktivitasTerdikitDampak.Deskripsi, aktivitasTerdikitDampak.DampakKarbon)
	}

	// Skor keberlanjutan sederhana (contoh perhitungan)
	// Asumsi: Semakin rendah jejak karbon, semakin tinggi skor.
	// Jika totalJejakKarbon 0, skor 100. Jika 100 kg CO2e, skor 0.
	skor := 100.0 - (totalJejakKarbon / 100.0 * 100.0 / 100.0) // Disederhanakan menjadi 100 - totalJejakKarbon
	if skor < 0 {
		skor = 0
	} else if skor > 100 {
		skor = 100
	}
	fmt.Printf("Skor Keberlanjutan Anda: %.2f/100\n", skor)

	// Rekomendasi berdasarkan total jejak karbon
	fmt.Println("\nRekomendasi untuk Mengurangi Jejak Karbon:")
	if totalJejakKarbon > 500 {
		fmt.Println("- Pertimbangkan untuk mengurangi perjalanan menggunakan kendaraan pribadi.")
		fmt.Println("- Fokus pada makanan nabati dan lokal.")
	} else if totalJejakKarbon > 200 {
		fmt.Println("- Kurangi konsumsi listrik (matikan lampu yang tidak perlu, cabut charger).")
		fmt.Println("- Pilah sampah dengan lebih baik dan kurangi penggunaan plastik sekali pakai.")
	} else {
		fmt.Println("- Terus pertahankan gaya hidup ramah lingkungan Anda!")
		fmt.Println("- Edukasi orang lain tentang pentingnya keberlanjutan.")
	}

	fmt.Println("\nSistem memberikan skor keberlanjutan berdasarkan pola hidup pengguna. Skor Anda dihitung dari total jejak karbon bulanan. Semakin rendah jejak karbon, semakin tinggi skornya.")
}

// --- Main Program (Menu Interaktif) ---
func main() {
	// Variabel lokal untuk input menu
	var pilihan, subPilihan, id, frekuensi int
	var kategori, deskripsi string
	var dampak float64
	var ascending bool

	// Contoh data awal untuk demonstrasi
	daftarAktivitas[0] = Aktivitas{ID: 1, Deskripsi: "Motor", Kategori: "Transportasi", DampakKarbon: 2.3, Frekuensi: 20}
	daftarAktivitas[1] = Aktivitas{ID: 2, Deskripsi: "Ayam", Kategori: "Makanan", DampakKarbon: 1.5, Frekuensi: 30}
	daftarAktivitas[2] = Aktivitas{ID: 3, Deskripsi: "Lampu", Kategori: "Energi", DampakKarbon: -0.2, Frekuensi: 60}
	daftarAktivitas[3] = Aktivitas{ID: 4, Deskripsi: "Anorganik", Kategori: "Sampah", DampakKarbon: 0.8, Frekuensi: 15}
	jumlahAktivitas = 4 // Set jumlah aktivitas sesuai data awal yang dimasukkan

	// Loop utama program untuk menampilkan menu
	for {
		fmt.Println("\n=== Aplikasi Pelacak Gaya Hidup Ramah Lingkungan ===")
		fmt.Println("1. Tambah Aktivitas")
		fmt.Println("2. Cari Aktivitas")
		fmt.Println("3. Edit Aktivitas")
		fmt.Println("4. Urutkan")
		fmt.Println("5. Tampilkan Daftar Aktivitas")
		fmt.Println("6. Laporan Bulanan")
		fmt.Println("0. Keluar") // Opsi untuk keluar dari program
		fmt.Print("Pilih menu (berupa angka, contoh: 1): ")
		fmt.Scan(&pilihan) // Membaca pilihan menu

		// Logika untuk keluar dari program
		if pilihan == 0 {
			fmt.Println("Terima kasih telah menggunakan Aplikasi Pelacak Gaya Hidup Ramah Lingkungan!")
			return // Keluar dari fungsi main, mengakhiri program
		}

		// Logika untuk setiap pilihan menu menggunakan if-else if-else
		if pilihan == 1 {
			fmt.Print("Masukkan ID (berupa angka, contoh: 1): ")
			fmt.Scan(&id)
			fmt.Print("Masukkan Kategori (satu kata, contoh: Transportasi): ")
			fmt.Scan(&kategori)
			fmt.Print("Masukkan Deskripsi (satu kata, contoh: Mobil): ")
			fmt.Scan(&deskripsi)
			fmt.Print("Masukkan Dampak Karbon (kg CO2 per aktivitas, contoh: 0.5): ")
			fmt.Scan(&dampak)
			fmt.Print("Masukkan Frekuensi (kali per bulan, contoh: 20): ")
			fmt.Scan(&frekuensi)
			tambahAktivitas(id, kategori, deskripsi, dampak, frekuensi)
		} else if pilihan == 2 {
			fmt.Print("Masukkan Kategori yang dicari (satu kata, contoh: Transportasi): ")
			fmt.Scan(&kategori)
			fmt.Println("Pilih metode pencarian:")
			fmt.Println("1. Sequential Search")
			fmt.Println("2. Binary Search")
			fmt.Print("Pilih metode (berupa angka, contoh: 1): ")
			fmt.Scan(&subPilihan)
			if subPilihan == 1 {
				idx := cariSequential(kategori)
				if idx != -1 {
					fmt.Printf("Ditemukan: ID: %d, Deskripsi: %s, Kategori: %s, Dampak: %.2f kgCO2e, Frekuensi: %d\n",
						daftarAktivitas[idx].ID, daftarAktivitas[idx].Deskripsi, daftarAktivitas[idx].Kategori,
						daftarAktivitas[idx].DampakKarbon, daftarAktivitas[idx].Frekuensi)
				} else {
					fmt.Println("Aktivitas tidak ditemukan!")
				}
			} else if subPilihan == 2 {
				// Penting: Binary Search memerlukan data terurut.
				// Urutkan dulu berdasarkan kategori sebelum mencari.
				selectionSortKategori() // Mengurutkan array global
				idx := cariBinary(kategori)
				if idx != -1 {
					fmt.Printf("Ditemukan: ID: %d, Deskripsi: %s, Kategori: %s, Dampak: %.2f kgCO2e, Frekuensi: %d\n",
						daftarAktivitas[idx].ID, daftarAktivitas[idx].Deskripsi, daftarAktivitas[idx].Kategori,
						daftarAktivitas[idx].DampakKarbon, daftarAktivitas[idx].Frekuensi)
				} else {
					fmt.Println("Aktivitas tidak ditemukan!")
				}
			} else {
				fmt.Println("Pilihan tidak valid!")
			}
		} else if pilihan == 3 {
			fmt.Print("Masukkan ID aktivitas (berupa angka, contoh: 1): ")
			fmt.Scan(&id)
			fmt.Println("Pilih aksi:")
			fmt.Println("1. Edit Aktivitas")
			fmt.Println("2. Hapus Aktivitas")
			fmt.Print("Pilih aksi (berupa angka, contoh: 1): ")
			fmt.Scan(&subPilihan)
			if subPilihan == 1 {
				fmt.Print("Masukkan Kategori baru (satu kata, contoh: Transportasi): ")
				fmt.Scan(&kategori)
				fmt.Print("Masukkan Deskripsi baru (satu kata, contoh: Bus): ")
				fmt.Scan(&deskripsi)
				fmt.Print("Masukkan Dampak Karbon baru (contoh: 0.2): ")
				fmt.Scan(&dampak)
				fmt.Print("Masukkan Frekuensi baru (contoh: 15): ")
				fmt.Scan(&frekuensi)
				editAktivitas(id, kategori, deskripsi, dampak, frekuensi, false) // false berarti edit
			} else if subPilihan == 2 {
				// Untuk hapus, nilai kategori, deskripsi, dampak, frekuensi tidak relevan,
				// jadi bisa dilewatkan nilai default atau kosong.
				editAktivitas(id, "", "", 0, 0, true) // true berarti hapus
			} else {
				fmt.Println("Pilihan tidak valid!")
			}
		} else if pilihan == 4 {
			fmt.Println("Pilih kriteria pengurutan:")
			fmt.Println("1. Dampak Karbon (Selection Sort)")
			fmt.Println("2. Frekuensi (Insertion Sort)")
			fmt.Print("Pilih kriteria (berupa angka, contoh: 1): ")
			fmt.Scan(&subPilihan)
			fmt.Print("Urutkan ascending? (true/false, contoh: true): ")
			fmt.Scan(&ascending) // Membaca boolean true/false
			if subPilihan == 1 {
				selectionSortDampak(ascending)
				tampilkanDaftar()
			} else if subPilihan == 2 {
				insertionSortFrekuensi(ascending)
				tampilkanDaftar()
			} else {
				fmt.Println("Pilihan tidak valid!")
			}
		} else if pilihan == 5 {
			tampilkanDaftar()
		} else if pilihan == 6 {
			laporanBulanan()
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
