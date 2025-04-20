package main
import "fmt"

func main() {
    var pilihan int
    var id, frekuensi int
    var kategori, deskripsi string
    var dampak float64
    var ascending bool

    for {
        fmt.Println("\n=== Aplikasi Pelacak Gaya Hidup Ramah Lingkungan ===")
        fmt.Println("1. Tambah Aktivitas")
        fmt.Println("2. Cari Aktivitas (Sequential)")
        fmt.Println("3. Cari Aktivitas (Binary)")
        fmt.Println("4. Edit Aktivitas")
        fmt.Println("5. Hapus Aktivitas")
        fmt.Println("6. Urutkan berdasarkan Dampak (Selection Sort)")
        fmt.Println("7. Urutkan berdasarkan Frekuensi (Insertion Sort)")
        fmt.Println("8. Tampilkan Daftar Aktivitas")
        fmt.Println("9. Laporan Bulanan")
        fmt.Println("10. Keluar")
        fmt.Print("Pilih: ")
        fmt.Scan(&pilihan)

        if pilihan == 10 {
            return
        }

        switch pilihan {
        case 1:
            fmt.Print("Masukkan ID: ")
            fmt.Scan(&id)
            fmt.Print("Masukkan Kategori (Transportasi/Energi/Makanan): ")
            fmt.Scan(&kategori)
            fmt.Print("Masukkan Deskripsi: ")
            fmt.Scan(&deskripsi)
            fmt.Print("Masukkan Dampak Karbon (kg CO2 per aktivitas): ")
            fmt.Scan(&dampak)
            fmt.Print("Masukkan Frekuensi (kali per bulan): ")
            fmt.Scan(&frekuensi)
            tambahAktivitas(id, kategori, deskripsi, dampak, frekuensi)
        case 2:
            fmt.Print("Masukkan Kategori yang dicari: ")
            fmt.Scan(&kategori)
            idx := cariSequential(kategori)
            if idx != -1 {
                fmt.Printf("Ditemukan: %s (%s)\n", daftarAktivitas[idx].Deskripsi, daftarAktivitas[idx].Kategori)
            } else {
                fmt.Println("Aktivitas tidak ditemukan!")
            }
        case 3:
            fmt.Print("Masukkan Kategori yang dicari: ")
            fmt.Scan(&kategori)
            // Asumsi array diurutkan terlebih dahulu berdasarkan kategori
            selectionSortDampak(true) // Placeholder, idealnya urutkan berdasarkan kategori
            idx := cariBinary(kategori)
            if idx != -1 {
                fmt.Printf("Ditemukan: %s (%s)\n", daftarAktivitas[idx].Deskripsi, daftarAktivitas[idx].Kategori)
            } else {
                fmt.Println("Aktivitas tidak ditemukan!")
            }
        case 4:
            fmt.Print("Masukkan ID aktivitas yang akan diedit: ")
            fmt.Scan(&id)
            fmt.Print("Masukkan Kategori baru: ")
            fmt.Scan(&kategori)
            fmt.Print("Masukkan Deskripsi baru: ")
            fmt.Scan(&deskripsi)
            fmt.Print("Masukkan Dampak Karbon baru: ")
            fmt.Scan(&dampak)
            fmt.Print("Masukkan Frekuensi baru: ")
            fmt.Scan(&frekuensi)
            editAktivitas(id, kategori, deskripsi, dampak, frekuensi)
        case 5:
            fmt.Print("Masukkan ID aktivitas yang akan dihapus: ")
            fmt.Scan(&id)
            hapusAktivitas(id)
        case 6:
            fmt.Print("Urutkan ascending? (true/false): ")
            fmt.Scan(&ascending)
            selectionSortDampak(ascending)
            tampilkanDaftar()
        case 7:
            fmt.Print("Urutkan ascending? (true/false): ")
            fmt.Scan(&ascending)
            insertionSortFrekuensi(ascending)
            tampilkanDaftar()
        case 8:
            tampilkanDaftar()
        case 9:
            laporanBulanan()
        }
    }
}
