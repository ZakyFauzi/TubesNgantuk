package main
import "fmt"

type Aktivitas struct {
	ID int
	Kategori string
	Deskripsi string
	DampakKarbon float64
	Frekuensi int
}
const MAX_AKTIVITAS = 100
var daftarAktivitas [MAX_AKTIVITAS]Aktivitas
var jumlahAktivitas int

func tambahAktivitas(id int, kategori, deskripsi string, dampak float64, frekuensi int) {
	if jumlahAktivitas < MAX_AKTIVITAS {
		daftarAktivitas[jumlahAktivitas] = Aktivitas{
			ID: id,
			Kategori: kategori,
			Deskripsi: deskripsi,
			DampakKarbon: dampak,
			Frekuensi: frekuensi,
		}
		jumlahAktivitas++
	} else {
		fmt.Println("Kapasitas penuh!")
	}
}

func cariSequential(kategori string) int {
	for i := 0; i < jumlahAktivitas; i++ {
		if daftarAktivitas[i].Kategori == kategori {
			return i
		}
	}
	return -1
}

func selectionSortKategori() {
	for i := 0; i < jumlahAktivitas-1; i++ {
		idxMin := i
		for j := i + 1; j < jumlahAktivitas; j++ {
			if daftarAktivitas[j].Kategori < daftarAktivitas[idxMin].Kategori {
				idxMin = j
			}
		}
		daftarAktivitas[i], daftarAktivitas[idxMin] = daftarAktivitas[idxMin], daftarAktivitas[i]
	}
}

func cariBinary(kategori string) int {
	kiri := 0
	kanan := jumlahAktivitas - 1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if daftarAktivitas[tengah].Kategori == kategori {
			return tengah
		} else if daftarAktivitas[tengah].Kategori < kategori {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func editAktivitas(id int, kategori, deskripsi string, dampak float64, frekuensi int, hapus bool) {
	for i := 0; i < jumlahAktivitas; i++ {
		if daftarAktivitas[i].ID == id {
			if hapus {
				for j := i; j < jumlahAktivitas-1; j++ {
					daftarAktivitas[j] = daftarAktivitas[j+1]
				}
				jumlahAktivitas--
				fmt.Println("Aktivitas berhasil dihapus!")
			} else {
				daftarAktivitas[i].Kategori = kategori
				daftarAktivitas[i].Deskripsi = deskripsi
				daftarAktivitas[i].DampakKarbon = dampak
				daftarAktivitas[i].Frekuensi = frekuensi
				fmt.Println("Aktivitas berhasil diedit!")
			}
			return
		}
	}
	fmt.Println("Aktivitas tidak ditemukan!")
}

func selectionSortDampak(ascending bool) {
	for i := 0; i < jumlahAktivitas-1; i++ {
		idxEkstrim := i
		for j := i + 1; j < jumlahAktivitas; j++ {
			if ascending {
				if daftarAktivitas[j].DampakKarbon < daftarAktivitas[idxEkstrim].DampakKarbon {
					idxEkstrim = j
				}
			} else {
				if daftarAktivitas[j].DampakKarbon > daftarAktivitas[idxEkstrim].DampakKarbon {
					idxEkstrim = j
				}
			}
		}
		daftarAktivitas[i], daftarAktivitas[idxEkstrim] = daftarAktivitas[idxEkstrim], daftarAktivitas[i]
	}
}

func insertionSortFrekuensi(ascending bool) {
	for i := 1; i < jumlahAktivitas; i++ {
		key := daftarAktivitas[i]
		j := i - 1
		for j >= 0 {
			if ascending {
				if daftarAktivitas[j].Frekuensi > key.Frekuensi {
					daftarAktivitas[j+1] = daftarAktivitas[j]
					j--
				} else {
					j = -1 // Keluar dari loop tanpa break
				}
			} else {
				if daftarAktivitas[j].Frekuensi < key.Frekuensi {
					daftarAktivitas[j+1] = daftarAktivitas[j]
					j--
				} else {
					j = -1 // Keluar dari loop tanpa break
				}
			}
		}
		daftarAktivitas[j+1] = key
	}
}

func hitungSkorKeberlanjutan() float64 {
	totalDampak := 0.0
	for i := 0; i < jumlahAktivitas; i++ {
		totalDampak += daftarAktivitas[i].DampakKarbon * float64(daftarAktivitas[i].Frekuensi)
	}
	skor := 100.0 - (totalDampak / 100.0 * 100.0)
	if skor < 0 {
		skor = 0
	}
	return skor
}

func laporanBulanan() {
	totalDampak := 0.0
	dampakKategori := make(map[string]float64)
	for i := 0; i < jumlahAktivitas; i++ {
		dampak := daftarAktivitas[i].DampakKarbon * float64(daftarAktivitas[i].Frekuensi)
		totalDampak += dampak
		dampakKategori[daftarAktivitas[i].Kategori] += dampak
	}
	fmt.Println("=== Laporan Bulanan ===")
	fmt.Printf("Total Jejak Karbon: %.2f kg CO2\n", totalDampak)
	fmt.Printf("Skor Keberlanjutan: %.2f/100\n", hitungSkorKeberlanjutan())
	fmt.Println("Dampak per Kategori:")
	for kat, dampak := range dampakKategori {
		fmt.Printf("- %s: %.2f kg CO2\n", kat, dampak)
	}
	fmt.Println("\nRekomendasi:")
	for kat, dampak := range dampakKategori {
		if dampak > totalDampak/3.0 {
			fmt.Printf("- Kurangi aktivitas di kategori %s, misalnya gunakan alternatif ramah lingkungan.\n", kat)
		}
	}
}

func tampilkanDaftar() {
	fmt.Println("=== Daftar Aktivitas ===")
	for i := 0; i < jumlahAktivitas; i++ {
		fmt.Printf("ID: %d, Kategori: %s, Deskripsi: %s, Dampak: %.2f kg CO2, Frekuensi: %d\n",
			daftarAktivitas[i].ID, daftarAktivitas[i].Kategori, daftarAktivitas[i].Deskripsi,
			daftarAktivitas[i].DampakKarbon, daftarAktivitas[i].Frekuensi)
	}
}

func main() {
	var pilihan, subPilihan int
	var id, frekuensi int
	var kategori, deskripsi string
	var dampak float64
	var ascending bool

	for {
		fmt.Println("=== Aplikasi Pelacak Gaya Hidup Ramah Lingkungan ===")
		fmt.Println("1. Tambah Aktivitas")
		fmt.Println("2. Cari Aktivitas")
		fmt.Println("3. Edit Aktivitas")
		fmt.Println("4. Urutkan")
		fmt.Println("5. Tampilkan Daftar Aktivitas")
		fmt.Println("6. Laporan Bulanan")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu (berupa angka, contoh: 1): ")
		fmt.Scan(&pilihan)

		if pilihan == 7 {
			return
		}

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
					fmt.Printf("Ditemukan: %s (%s)\n", daftarAktivitas[idx].Deskripsi, daftarAktivitas[idx].Kategori)
				} else {
					fmt.Println("Aktivitas tidak ditemukan!")
				}
			} else if subPilihan == 2 {
				selectionSortKategori()
				idx := cariBinary(kategori)
				if idx != -1 {
					fmt.Printf("Ditemukan: %s (%s)\n", daftarAktivitas[idx].Deskripsi, daftarAktivitas[idx].Kategori)
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
				editAktivitas(id, kategori, deskripsi, dampak, frekuensi, false)
			} else if subPilihan == 2 {
				editAktivitas(id, "", "", 0, 0, true)
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
			fmt.Scan(&ascending)
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
