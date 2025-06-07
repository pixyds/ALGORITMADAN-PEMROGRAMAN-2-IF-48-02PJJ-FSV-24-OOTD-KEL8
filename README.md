# 👗 OOTD Planner - Aplikasi Perencanaan Outfit Harian

**Kelas:** IF-48-02PJJ  
**Anggota Kelompok 8:**
- Dorkas Santania — 103042310078  
- Zhafir Rasyid — 103042310092  

---

## 🧰 Deskripsi Proyek

OOTD Planner adalah aplikasi berbasis terminal yang dibuat menggunakan bahasa Go.
Aplikasi ini dapat menyimpan sementara data outfit, melakukan pencarian dan pengurutan, serta merekomendasikan pakaian berdasarkan tingkat formalitas dan waktu terakhir digunakan.

---

## 🎯 Fitur Utama

- Menambah, mengedit, dan menghapus outfit
- Menampilkan daftar outfit dengan opsi pengurutan:
  - Nama (Selection Sort)
  - Kategori (Insertion Sort)
  - Formalitas (Selection Sort)
  - Terakhir Dipakai (Insertion Sort)
- Pencarian outfit berdasarkan:
  - Nama (Sequential Search)
  - Kategori (Binary Search)
  - Warna (Sequential Search)
- Fitur perencanaan outfit harian (Plan OOTD) berdasarkan tingkat formalitas
- Data disimpan dalam array statik berukuran maksimal 100 (tidak ada penyimpanan permanen)

---

## 🏗️ Cara Menjalankan

### 📌 Prasyarat

- Sudah menginstal Go (versi 1.17 atau lebih tinggi)  
  [Download Go](https://golang.org/dl/)

### ▶️ Menjalankan Langsung (tanpa build)

Jalankan perintah berikut di terminal:

```bash
go run outfit_manager.go
```

## 📦 Build ke Executable

### 🔹 Untuk Windows (.exe)

```bash
go build -o outfit_manager.exe outfit_manager.go
./outfit_manager.exe
```

### 🔹 Untuk macOS/Linux

```bash
go build -o outfit_manager outfit_manager.go
./outfit_manager
```

###  Struktur File
```
├── outfit_manager.go   # Source code utama aplikasi
└── README.md           # Dokumentasi penggunaan
```
