# Aplikasi Pelacak Gaya Hidup Ramah Lingkungan

Aplikasi Pelacak Gaya Hidup Ramah Lingkungan adalah program berbasis konsol sederhana yang dirancang untuk membantu pengguna memantau dan mengelola jejak karbon mereka. Aplikasi ini memungkinkan pencatatan aktivitas sehari-hari yang berdampak pada lingkungan, menghitung total jejak karbon bulanan, dan memberikan rekomendasi untuk gaya hidup yang lebih hijau.

## Fitur

Aplikasi ini mengimplementasikan berbagai konsep dasar algoritma dan pemrograman yang diajarkan, meliputi:

##### - Tipe Bentukan (Struct): Menggunakan struktur data Aktivitas untuk menyimpan detail setiap kegiatan, seperti ID, kategori, deskripsi, dampak karbon, dan frekuensi.
##### - Array: Data aktivitas disimpan dalam array berukuran tetap ([100]Aktivitas), dengan pengelolaan jumlah elemen aktif secara manual.
##### - Subprogram (Fungsi/Prosedur): Setiap fungsionalitas (tambah, cari, edit, urutkan, laporan) diimplementasikan sebagai fungsi terpisah untuk modularitas kode.
##### - Pencarian Nilai Ekstrem: Sistem secara otomatis mengidentifikasi aktivitas dengan dampak karbon terbesar dan terkecil dalam laporan bulanan.
##### - Sequential Search: Digunakan untuk mencari aktivitas berdasarkan kategori secara berurutan.
##### - Binary Search: Digunakan untuk mencari aktivitas berdasarkan kategori. Penting untuk diingat bahwa data akan diurutkan terlebih dahulu berdasarkan kategori sebelum pencarian biner dilakukan.
##### - Selection Sort: Diimplementasikan untuk mengurutkan daftar aktivitas berdasarkan Dampak Karbon.
##### - Insertion Sort: Diimplementasikan untuk mengurutkan daftar aktivitas berdasarkan Frekuensi.
##### - Manajemen Aktivitas: Pengguna dapat menambahkan, mengedit, dan menghapus aktivitas.
##### - Laporan Bulanan & Skor Keberlanjutan: Sistem menyediakan laporan total jejak karbon bulanan, skor keberlanjutan, dan rekomendasi sederhana untuk mengurangi dampak lingkungan.
##### - Interaksi Menu: Antarmuka berbasis teks yang mudah digunakan dengan pilihan menu numerik.

## Cara Menjalankan Program
Untuk menjalankan aplikasi ini, kamu memerlukan lingkungan Go (Golang) terinstal di komputermu.

##### - Simpan Kode: Salin seluruh kode program dan simpan dalam sebuah file dengan ekstensi .go, misalnya main.go.
##### - Buka Terminal/Command Prompt: Navigasikan ke direktori tempat kamu menyimpan file main.go.
##### - Jalankan Program: Ketik perintah berikut di terminal: 'go run main.go'
##### - Program akan mulai berjalan dan menampilkan menu utama di konsol.
