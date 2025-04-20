package main
import "fmt"

func main() {
var pilihan int	\*mencetak menu pilihan pada aplikasi*\
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
fmt.Scan(&pilihan)		\*input pilihan yang ingin dipilih*\
}
