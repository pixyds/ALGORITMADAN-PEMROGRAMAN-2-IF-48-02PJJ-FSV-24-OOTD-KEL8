package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// === Deklarasi Struct, Konstanta, dan Variabel Global ===
const MAX_OUTFITS = 100

type Outfit struct {
	ID         int
	Nama       string
	Kategori   string
	Warna      string
	Musim      string
	Deskripsi  string
	Formalitas int 
	LastUsed   time.Time
}

var (
	outfits  [MAX_OUTFITS]Outfit
	nOutfits int = 0
	scanner      = bufio.NewScanner(os.Stdin)
)

// === Fungsi Input Helper ===
func input() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// === Fungsi Utility / Pencarian / Sorting ===
func cariIdxByID(id int) int {
	for i := 0; i < nOutfits; i++ {
		if outfits[i].ID == id {
			return i
		}
	}
	return -1
}

func seqSearchNama(nama string) int {
	for i := 0; i < nOutfits; i++ {
		if strings.EqualFold(outfits[i].Nama, nama) {
			return i
		}
	}
	return -1
}

func seqSearchWarna(warna string) []int {
	var hasil []int
	for i := 0; i < nOutfits; i++ {
		if strings.EqualFold(outfits[i].Warna, warna) {
			hasil = append(hasil, i)
		}
	}
	return hasil
}

func binarySearchKategori(kat string) int {
	low, high := 0, nOutfits-1
	for low <= high {
		mid := (low + high) / 2
		if strings.EqualFold(outfits[mid].Kategori, kat) {
			return mid
		} else if strings.ToLower(outfits[mid].Kategori) < strings.ToLower(kat) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func selectionSortNama(asc bool) {
	for i := 0; i < nOutfits-1; i++ {
		idx := i
		for j := i + 1; j < nOutfits; j++ {
			cmp := strings.Compare(strings.ToLower(outfits[j].Nama), strings.ToLower(outfits[idx].Nama))
			if (asc && cmp < 0) || (!asc && cmp > 0) {
				idx = j
			}
		}
		if idx != i {
			outfits[i], outfits[idx] = outfits[idx], outfits[i]
		}
	}
}

func insertionSortKategori(asc bool) {
	for i := 1; i < nOutfits; i++ {
		temp := outfits[i]
		j := i - 1
		compare := func(a, b string) bool {
			if asc {
				return strings.ToLower(a) > strings.ToLower(b)
			}
			return strings.ToLower(a) < strings.ToLower(b)
		}
		for j >= 0 && compare(outfits[j].Kategori, temp.Kategori) {
			outfits[j+1] = outfits[j]
			j--
		}
		outfits[j+1] = temp
	}
}

func selectionSortFormalitas(asc bool) {
	for i := 0; i < nOutfits-1; i++ {
		idx := i
		for j := i + 1; j < nOutfits; j++ {
			if (asc && outfits[j].Formalitas < outfits[idx].Formalitas) || (!asc && outfits[j].Formalitas > outfits[idx].Formalitas) {
				idx = j
			}
		}
		if idx != i {
			outfits[i], outfits[idx] = outfits[idx], outfits[i]
		}
	}
}

func insertionSortLastUsed(asc bool) {
	for i := 1; i < nOutfits; i++ {
		temp := outfits[i]
		j := i - 1
		for j >= 0 && ((asc && outfits[j].LastUsed.After(temp.LastUsed)) || (!asc && outfits[j].LastUsed.Before(temp.LastUsed))) {
			outfits[j+1] = outfits[j]
			j--
		}
		outfits[j+1] = temp
	}
}

// === Fungsi Tampil ===
func tampilOutfits() {
	fmt.Println("\nDaftar Outfit:")
	for i := 0; i < nOutfits; i++ {
		fmt.Printf("%d. %s | %s | %s | %s | Formalitas: %d | Terakhir Dipakai: %s\n",
			outfits[i].ID, outfits[i].Nama, outfits[i].Kategori, outfits[i].Warna, outfits[i].Musim, outfits[i].Formalitas,
			outfits[i].LastUsed.Format("2006-01-02"))
		fmt.Printf("   %s\n", outfits[i].Deskripsi)
	}
}

func tampilDetail(idx int) {
	fmt.Println("\nDetail Outfit:")
	fmt.Printf("ID: %d\n", outfits[idx].ID)
	fmt.Printf("Nama: %s\n", outfits[idx].Nama)
	fmt.Printf("Kategori: %s\n", outfits[idx].Kategori)
	fmt.Printf("Warna: %s\n", outfits[idx].Warna)
	fmt.Printf("Musim: %s\n", outfits[idx].Musim)
	fmt.Printf("Deskripsi: %s\n", outfits[idx].Deskripsi)
	fmt.Printf("Formalitas: %d\n", outfits[idx].Formalitas)
	fmt.Printf("Terakhir Dipakai: %s\n", outfits[idx].LastUsed.Format("2006-01-02"))
}

// === Fungsi CRUD ===
func tambahOutfit() {
	if nOutfits >= MAX_OUTFITS {
		fmt.Println("Kapasitas penuh!")
		return
	}
	fmt.Println("\nTambah Outfit Baru")
	fmt.Print("Nama Outfit: ")
	nama := input()
	fmt.Print("Kategori (Atasan/Bawahan/Luar/Aksesoris/Sepatu): ")
	kategori := input()
	fmt.Print("Warna: ")
	warna := input()
	fmt.Print("Musim (Panas/Hujan/Dingin): ")
	musim := input()
	fmt.Print("Deskripsi: ")
	deskripsi := input()
	fmt.Print("Tingkat Formalitas (1=Santai,2=Semi Formal,3=Formal): ")
	formalStr := input()
	formal, err := strconv.Atoi(formalStr)
	if err != nil || formal < 1 || formal > 3 {
		fmt.Println("Input formalitas tidak valid, diset ke 1 (Santai).")
		formal = 1
	}

	outfits[nOutfits] = Outfit{
		ID:         nOutfits + 1,
		Nama:       nama,
		Kategori:   kategori,
		Warna:      warna,
		Musim:      musim,
		Deskripsi:  deskripsi,
		Formalitas: formal,
		LastUsed:   time.Now(),
	}
	nOutfits++
	fmt.Println("Outfit berhasil ditambahkan!")
}

func editOutfit() {
	fmt.Print("Masukkan ID outfit yang ingin diedit: ")
	var id int
	fmt.Scan(&id)
	scanner.Scan()
	idx := cariIdxByID(id)
	if idx == -1 {
		fmt.Println("Outfit tidak ditemukan.")
		return
	}

	fmt.Println("Edit data (kosongkan jika tidak ingin mengubah):")
	fmt.Printf("Nama [%s]: ", outfits[idx].Nama)
	nama := input()
	if nama != "" {
		outfits[idx].Nama = nama
	}
	fmt.Printf("Kategori [%s]: ", outfits[idx].Kategori)
	kat := input()
	if kat != "" {
		outfits[idx].Kategori = kat
	}
	fmt.Printf("Warna [%s]: ", outfits[idx].Warna)
	warna := input()
	if warna != "" {
		outfits[idx].Warna = warna
	}
	fmt.Printf("Musim [%s]: ", outfits[idx].Musim)
	musim := input()
	if musim != "" {
		outfits[idx].Musim = musim
	}
	fmt.Printf("Deskripsi [%s]: ", outfits[idx].Deskripsi)
	desk := input()
	if desk != "" {
		outfits[idx].Deskripsi = desk
	}
	fmt.Printf("Formalitas [%d]: ", outfits[idx].Formalitas)
	formalStr := input()
	if formalStr != "" {
		formal, err := strconv.Atoi(formalStr)
		if err == nil && formal >= 1 && formal <= 3 {
			outfits[idx].Formalitas = formal
		} else {
			fmt.Println("Input formalitas tidak valid, tidak diubah.")
		}
	}
	fmt.Println("Outfit berhasil diubah.")
}

func hapusOutfit() {
	fmt.Print("Masukkan ID outfit yang ingin dihapus: ")
	var id int
	fmt.Scan(&id)
	scanner.Scan()
	idx := cariIdxByID(id)
	if idx == -1 {
		fmt.Println("Outfit tidak ditemukan.")
		return
	}
	for i := idx; i < nOutfits-1; i++ {
		outfits[i] = outfits[i+1]
	}
	nOutfits--
	fmt.Println("Outfit berhasil dihapus.")
}

// === Fungsi Menu Cari Outfit ===
func menuCariOutfit() {
	if nOutfits == 0 {
		fmt.Println("Belum ada outfit.")
		return
	}
	fmt.Println("\nCari berdasarkan:")
	fmt.Println("1. Nama (Sequential Search)")
	fmt.Println("2. Kategori (Binary Search, urutkan dulu)")
	fmt.Println("3. Warna (Sequential Search)")
	fmt.Print("Pilih: ")
	var pil int
	fmt.Scan(&pil)
	scanner.Scan()

	switch pil {
	case 1:
		fmt.Print("Masukkan nama: ")
		nama := input()
		idx := seqSearchNama(nama)
		if idx != -1 {
			tampilDetail(idx)
		} else {
			fmt.Println("Outfit tidak ditemukan.")
		}
	case 2:
		insertionSortKategori(true)
		fmt.Print("Masukkan kategori: ")
		kat := input()
		idx := binarySearchKategori(kat)
		if idx != -1 {
			tampilDetail(idx)
		} else {
			fmt.Println("Outfit tidak ditemukan.")
		}
	case 3:
		fmt.Print("Masukkan warna: ")
		warna := input()
		idxs := seqSearchWarna(warna)
		if len(idxs) > 0 {
			for _, idx := range idxs {
				tampilDetail(idx)
			}
		} else {
			fmt.Println("Outfit dengan warna tersebut tidak ditemukan.")
		}
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// === Fungsi Menu Lihat Outfit ===
func menuLihatOutfit() {
	if nOutfits == 0 {
		fmt.Println("Belum ada outfit.")
		return
	}

	fmt.Println("\nUrutkan berdasarkan:")
	fmt.Println("1. Nama (Selection Sort)")
	fmt.Println("2. Kategori (Insertion Sort)")
	fmt.Println("3. Formalitas (Selection Sort)")
	fmt.Println("4. Terakhir Dipakai (Insertion Sort)")
	fmt.Print("Pilih: ")
	var pil int
	fmt.Scan(&pil)
	scanner.Scan()

	fmt.Println("Urutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Print("Pilih: ")
	var urut int
	fmt.Scan(&urut)
	scanner.Scan()
	asc := urut == 1

	switch pil {
	case 1:
		selectionSortNama(asc)
	case 2:
		insertionSortKategori(asc)
	case 3:
		selectionSortFormalitas(asc)
	case 4:
		insertionSortLastUsed(asc)
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}

	tampilOutfits()
}

// === Fungsi Plan OOTD ===
func planOOTD() {
	if nOutfits == 0 {
		fmt.Println("Belum ada outfit untuk direncanakan.")
		return
	}
	fmt.Println("\nPilih outfit berdasarkan formalitas:")
	fmt.Println("1. Santai (1)")
	fmt.Println("2. Semi Formal (2)")
	fmt.Println("3. Formal (3)")
	fmt.Print("Masukkan pilihan: ")
	var f int
	fmt.Scan(&f)
	scanner.Scan()
	if f < 1 || f > 3 {
		fmt.Println("Pilihan formalitas tidak valid.")
		return
	}

	// Cari outfit yang formalitasnya sesuai dan terbaru (LastUsed paling lama)
	var kandidatIdx = -1
	var oldest time.Time = time.Now()
	for i := 0; i < nOutfits; i++ {
		if outfits[i].Formalitas == f {
			if kandidatIdx == -1 || outfits[i].LastUsed.Before(oldest) {
				kandidatIdx = i
				oldest = outfits[i].LastUsed
			}
		}
	}
	if kandidatIdx == -1 {
		fmt.Println("Tidak ada outfit dengan formalitas tersebut.")
		return
	}

	// Update LastUsed ke sekarang karena akan dipakai
	outfits[kandidatIdx].LastUsed = time.Now()
	fmt.Println("Rekomendasi Outfit OOTD:")
	tampilDetail(kandidatIdx)
}

// === Fungsi Main Menu ===
func mainMenu() {
	for {
		fmt.Println("\n--- Manajemen Outfit ---")
		fmt.Println("1. Tambah Outfit")
		fmt.Println("2. Edit Outfit")
		fmt.Println("3. Hapus Outfit")
		fmt.Println("4. Cari Outfit")
		fmt.Println("5. Lihat Outfit")
		fmt.Println("6. Plan OOTD")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pil int
		inputLine := input()
		pil, err := strconv.Atoi(inputLine)
		if err != nil {
			fmt.Println("Input tidak valid, silakan masukkan angka.")
			continue
		}

		switch pil {
		case 1:
			tambahOutfit()
		case 2:
			editOutfit()
		case 3:
			hapusOutfit()
		case 4:
			menuCariOutfit()
		case 5:
			menuLihatOutfit()
		case 6:
			planOOTD()
		case 0:
			fmt.Println("Terima kasih, sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func main() {
	mainMenu()
}
