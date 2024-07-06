
# Dokumentasi Proyek Gin_Gorm_Api

## Pendahuluan
**Deskripsi Proyek:**
Gin_Gorm_Api adalah aplikasi API berbasis Go yang menggunakan framework Gin dan ORM Gorm untuk mengelola database. Proyek ini bertujuan untuk menyediakan API yang efisien dan mudah digunakan untuk berbagai operasi data.

## Struktur Direktori
**Struktur utama direktori:**
- **bootstrap/**: Berisi file bootstrap untuk aplikasi.
- **configs/**: Berisi file konfigurasi.
- **controllers/**: Berisi logika kontroler untuk menangani permintaan API.
- **costanta/**: Berisi konstanta yang digunakan dalam aplikasi.
- **database/**: Berisi migrasi dan seeder database.
- **middleware/**: Berisi middleware untuk aplikasi.
- **models/**: Berisi definisi model untuk ORM Gorm.
- **public/**: Direktori publik untuk file yang dapat diakses oleh pengguna.
- **requests/**: Berisi validasi permintaan.
- **responses/**: Berisi format respon API.
- **routes/**: Berisi definisi rute API.
- **tmp/**: Direktori sementara.
- **utils/**: Berisi utilitas dan fungsi bantu.
- **.air.toml**: Konfigurasi untuk live reloading dengan air.
- **.env**: File lingkungan untuk konfigurasi aplikasi.
- **go.mod**: File modul Go.
- **go.sum**: File checksum untuk dependensi Go.
- **main.go**: Entry point utama untuk aplikasi.

## Instalasi
**Clone Repository:**
```bash
git clone https://github.com/RahmatRafiq/Gin_Gorm_Api.git
cd Gin_Gorm_Api
```

**Instalasi Dependensi:**
```bash
go mod tidy
```

**Pengaturan Lingkungan:**
Salin file `.env.example` menjadi `.env` dan sesuaikan konfigurasi sesuai kebutuhan Anda.

## Migrasi Database
**Menjalankan Migrasi:**
```bash
go run main.go migrate
```

**Seeder Data (jika ada):**
```bash
go run main.go seed
```

## Penggunaan
**Menjalankan Server Pengembangan:**
```bash
go run main.go
```
Akses API di `http://localhost:8080`.

## Fitur Utama
**Deskripsi Fitur:**
- **User Management:** Sistem manajemen pengguna, termasuk pendaftaran, login, dan otorisasi.
- **CRUD Operations:** Operasi CRUD penuh untuk entitas dalam aplikasi.
- **Error Handling:** Penanganan kesalahan yang konsisten dan informatif.
- **Middleware:** Middleware untuk logging, otorisasi, dan validasi permintaan.
- **Database Migration:** Migrasi dan seeder database otomatis menggunakan Gorm.

## Contributing
**Panduan untuk Berkontribusi:**
1. Fork repositori.
2. Buat branch fitur (`git checkout -b feature/AmazingFeature`).
3. Commit perubahan (`git commit -m 'Add some AmazingFeature'`).
4. Push ke branch (`git push origin feature/AmazingFeature`).
5. Buat Pull Request.

## License
Proyek ini dilisensikan di bawah MIT License - lihat file [LICENSE](https://github.com/RahmatRafiq/Gin_Gorm_Api/blob/main/LICENSE) untuk detail lebih lanjut.
