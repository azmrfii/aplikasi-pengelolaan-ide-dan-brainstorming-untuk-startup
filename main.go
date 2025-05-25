package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
)

const MAX_IDE = 100

type Ide struct {
	Judul       string
	Kategori    string
	Tanggal     string
	Upvotes     int
}

var daftarIde [MAX_IDE]Ide
var jumlahIde int

func inputTanggal() string {
	now := time.Now()
	return now.Format("2006-01-02")
}

func tambahIde(judul, kategori string) {
	if jumlahIde < MAX_IDE {
		daftarIde[jumlahIde] = Ide{judul, kategori, inputTanggal(), 0}
		jumlahIde++
		fmt.Println("Ide berhasil ditambahkan!")
	} else {
		fmt.Println("Kapasitas penyimpanan penuh!")
	}
}

func cariIdeSequential(keyword string) {
	fmt.Println("\nHasil Pencarian (Sequential Search):")
	ketemu := false
	for i := 0; i < jumlahIde; i++ {
		if strings.Contains(strings.ToLower(daftarIde[i].Judul), strings.ToLower(keyword)) {
			tampilkanIde(daftarIde[i])
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Tidak ditemukan.")
	}
}

func binarySearch(keyword string) int {
	kiri := 0
	kanan := jumlahIde - 1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if daftarIde[tengah].Judul == keyword {
			return tengah
		} else if daftarIde[tengah].Judul < keyword {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func insertionSortByJudul() {
	for i := 1; i < jumlahIde; i++ {
		temp := daftarIde[i]
		j := i
		for j > 0 && daftarIde[j-1].Judul > temp.Judul {
			daftarIde[j] = daftarIde[j-1]
			j--
		}
		daftarIde[j] = temp
	}
}

func hapusIde(judul string) {
	idx := -1
	for i := 0; i < jumlahIde; i++ {
		if daftarIde[i].Judul == judul {
			idx = i
		}
	}
	if idx != -1 {
		for i := idx; i < jumlahIde-1; i++ {
			daftarIde[i] = daftarIde[i+1]
		}
		jumlahIde--
		fmt.Println("Ide berhasil dihapus.")
	} else {
		fmt.Println("Ide tidak ditemukan.")
	}
}

func editIde(judul string, judulBaru string, kategoriBaru string) {
	idx := -1
	for i := 0; i < jumlahIde; i++ {
		if daftarIde[i].Judul == judul {
			idx = i
		}
	}
	if idx != -1 {
		daftarIde[idx].Judul = judulBaru
		daftarIde[idx].Kategori = kategoriBaru
		fmt.Println("Ide berhasil diubah.")
	} else {
		fmt.Println("Ide tidak ditemukan.")
	}
}

func upvoteIde(judul string) {
	idx := -1
	for i := 0; i < jumlahIde; i++ {
		if daftarIde[i].Judul == judul {
			idx = i
		}
	}
	if idx != -1 {
		daftarIde[idx].Upvotes++
		fmt.Println("Upvote berhasil.")
	} else {
		fmt.Println("Ide tidak ditemukan.")
	}
}

func selectionSortByUpvotes(ascending bool) {
	for i := 0; i < jumlahIde-1; i++ {
		idx := i
		for j := i + 1; j < jumlahIde; j++ {
			if (ascending && daftarIde[j].Upvotes < daftarIde[idx].Upvotes) || (!ascending && daftarIde[j].Upvotes > daftarIde[idx].Upvotes) {
				idx = j
			}
		}
		temp := daftarIde[i]
		daftarIde[i] = daftarIde[idx]
		daftarIde[idx] = temp
	}
}

func tampilkanIde(ide Ide) {
	fmt.Printf("Judul: %s\nKategori: %s\nTanggal: %s\nUpvotes: %d\n\n", ide.Judul, ide.Kategori, ide.Tanggal, ide.Upvotes)
}

func tampilkanSemuaIde() {
	fmt.Println("\nDaftar Ide:")
	for i := 0; i < jumlahIde; i++ {
		tampilkanIde(daftarIde[i])
	}
}

func tampilkanIdePopuler(start, end string) {
	fmt.Printf("\nIde Terpopuler (%s s.d %s):\n", start, end)
	for i := 0; i < jumlahIde; i++ {
		if daftarIde[i].Tanggal >= start && daftarIde[i].Tanggal <= end && daftarIde[i].Upvotes > 0 {
			tampilkanIde(daftarIde[i])
		}
	}
}

// func menu() {
// 	var pilihan int
// 	for pilihan != 0 {
// 		fmt.Println("\n===== MENU =====")
// 		fmt.Println("1. Tambah Ide")
// 		fmt.Println("2. Tampilkan Semua Ide")
// 		fmt.Println("3. Cari Ide (Sequential Search)")
// 		fmt.Println("4. Cari Ide (Binary Search)")
// 		fmt.Println("5. Edit Ide")
// 		fmt.Println("6. Hapus Ide")
// 		fmt.Println("7. Upvote Ide")
// 		fmt.Println("8. Urutkan Ide (Upvotes - Selection Sort)")
// 		fmt.Println("9. Urutkan Ide (Judul - Insertion Sort)")
// 		fmt.Println("10. Tampilkan Ide Populer (Tanggal)")
// 		fmt.Println("0. Keluar")
// 		fmt.Print("Pilihan Anda: ")
// 		fmt.Scan(&pilihan)

// 		switch pilihan {
// 		case 1:
// 			var judul, kategori string
// 			fmt.Print("Masukkan Judul Ide: ")
// 			fmt.Scan(&judul)
// 			fmt.Print("Masukkan Kategori: ")
// 			fmt.Scan(&kategori)
// 			tambahIde(judul, kategori)
// 		case 2:
// 			tampilkanSemuaIde()
// 		case 3:
// 			var keyword string
// 			fmt.Print("Masukkan kata kunci: ")
// 			fmt.Scan(&keyword)
// 			cariIdeSequential(keyword)
// 		case 4:
// 			var keyword string
// 			fmt.Print("Masukkan judul lengkap: ")
// 			fmt.Scan(&keyword)
// 			insertionSortByJudul()
// 			idx := binarySearch(keyword)
// 			if idx != -1 {
// 				fmt.Println("✅ Ide ditemukan:")
// 				tampilkanIde(daftarIde[idx])
// 			} else {
// 				fmt.Println("❌ Ide tidak ditemukan.")
// 			}
// 		case 5:
// 			var judul, baru, kategori string
// 			fmt.Print("Judul yang ingin diedit: ")
// 			fmt.Scan(&judul)
// 			fmt.Print("Judul baru: ")
// 			fmt.Scan(&baru)
// 			fmt.Print("Kategori baru: ")
// 			fmt.Scan(&kategori)
// 			editIde(judul, baru, kategori)
// 		case 6:
// 			var judul string
// 			fmt.Print("Masukkan judul yang ingin dihapus: ")
// 			fmt.Scan(&judul)
// 			hapusIde(judul)
// 		case 7:
// 			var judul string
// 			fmt.Print("Masukkan judul untuk di-upvote: ")
// 			fmt.Scan(&judul)
// 			upvoteIde(judul)
// 		case 8:
// 			var asc int
// 			fmt.Print("Urutan (1: ascending, 0: descending): ")
// 			fmt.Scan(&asc)
// 			selectionSortByUpvotes(asc == 1)
// 			tampilkanSemuaIde()
// 		case 9:
// 			insertionSortByJudul()
// 			tampilkanSemuaIde()
// 		case 10:
// 			var start, end string
// 			fmt.Print("Tanggal mulai (yyyy-mm-dd): ")
// 			fmt.Scan(&start)
// 			fmt.Print("Tanggal akhir (yyyy-mm-dd): ")
// 			fmt.Scan(&end)
// 			tampilkanIdePopuler(start, end)
// 		case 0:
// 			fmt.Println("👋 Keluar dari program.")
// 		default:
// 			fmt.Println("❌ Pilihan tidak valid.")
// 		}
// 	}
// }

// func menu() {
// 	for {
// 		var pilihan int
// 		fmt.Println("\n===== MENU =====")
// 		fmt.Println("1. Tambah Ide")
// 		fmt.Println("2. Tampilkan Semua Ide")
// 		fmt.Println("3. Cari Ide (Sequential Search)")
// 		fmt.Println("4. Cari Ide (Binary Search)")
// 		fmt.Println("5. Edit Ide")
// 		fmt.Println("6. Hapus Ide")
// 		fmt.Println("7. Upvote Ide")
// 		fmt.Println("8. Urutkan Ide (Upvotes - Selection Sort)")
// 		fmt.Println("9. Urutkan Ide (Judul - Insertion Sort)")
// 		fmt.Println("10. Tampilkan Ide Populer (Tanggal)")
// 		fmt.Println("0. Keluar")
// 		fmt.Print("Pilihan Anda: ")
// 		fmt.Scanln(&pilihan)

// 		if pilihan == 0 {
// 			fmt.Println("Keluar dari program.")
// 			return
// 		}

// 		switch pilihan {
// 		case 1:
// 			var judul, kategori string
// 			fmt.Print("Masukkan Judul Ide: ")
// 			fmt.Scan(&judul)
// 			fmt.Print("Masukkan Kategori: ")
// 			fmt.Scan(&kategori)
// 			tambahIde(judul, kategori)
// 		case 2:
// 			tampilkanSemuaIde()
// 		case 3:
// 			var keyword string
// 			fmt.Print("Masukkan kata kunci: ")
// 			fmt.Scan(&keyword)
// 			cariIdeSequential(keyword)
// 		case 4:
// 			var keyword string
// 			fmt.Print("Masukkan judul lengkap: ")
// 			fmt.Scan(&keyword)
// 			insertionSortByJudul()
// 			idx := binarySearch(keyword)
// 			if idx != -1 {
// 				fmt.Println("Ide ditemukan:")
// 				tampilkanIde(daftarIde[idx])
// 			} else {
// 				fmt.Println("Ide tidak ditemukan.")
// 			}
// 		case 5:
// 			var judul, baru, kategori string
// 			fmt.Print("Judul yang ingin diedit: ")
// 			fmt.Scan(&judul)
// 			fmt.Print("Judul baru: ")
// 			fmt.Scan(&baru)
// 			fmt.Print("Kategori baru: ")
// 			fmt.Scan(&kategori)
// 			editIde(judul, baru, kategori)
// 		case 6:
// 			var judul string
// 			fmt.Print("Masukkan judul yang ingin dihapus: ")
// 			fmt.Scan(&judul)
// 			hapusIde(judul)
// 		case 7:
// 			var judul string
// 			fmt.Print("Masukkan judul untuk di-upvote: ")
// 			fmt.Scan(&judul)
// 			upvoteIde(judul)
// 		case 8:
// 			var asc int
// 			fmt.Print("Urutan (1: ascending, 0: descending): ")
// 			fmt.Scan(&asc)
// 			selectionSortByUpvotes(asc == 1)
// 			tampilkanSemuaIde()
// 		case 9:
// 			insertionSortByJudul()
// 			tampilkanSemuaIde()
// 		case 10:
// 			var start, end string
// 			fmt.Print("Tanggal mulai (yyyy-mm-dd): ")
// 			fmt.Scan(&start)
// 			fmt.Print("Tanggal akhir (yyyy-mm-dd): ")
// 			fmt.Scan(&end)
// 			tampilkanIdePopuler(start, end)
// 		default:
// 			fmt.Println("Pilihan tidak valid.")
// 		}
// 	}
// }

func menu() {
    scanner := bufio.NewScanner(os.Stdin)
    
    for {
        var pilihan int
        fmt.Println("\n===== MENU =====")
        fmt.Println("1. Tambah Ide")
        fmt.Println("2. Tampilkan Semua Ide")
        fmt.Println("3. Cari Ide (Sequential Search)")
        fmt.Println("4. Cari Ide (Binary Search)")
        fmt.Println("5. Edit Ide")
        fmt.Println("6. Hapus Ide")
        fmt.Println("7. Upvote Ide")
        fmt.Println("8. Urutkan Ide (Upvotes - Selection Sort)")
        fmt.Println("9. Urutkan Ide (Judul - Insertion Sort)")
        fmt.Println("10. Tampilkan Ide Populer (Tanggal)")
        fmt.Println("0. Keluar")
        fmt.Print("Pilihan Anda: ")

        _, err := fmt.Scanln(&pilihan)
        if err != nil {
            // Clear input buffer jika ada error
            scanner.Scan()
            fmt.Println("Input tidak valid. Harap masukkan angka.")
            continue
        }

        if pilihan == 0 {
            fmt.Println("Keluar dari program.")
            return
        }

        switch pilihan {
        case 1:
            var judul, kategori string
            fmt.Print("Masukkan Judul Ide: ")
            scanner.Scan()
            judul = scanner.Text()
            fmt.Print("Masukkan Kategori: ")
            scanner.Scan()
            kategori = scanner.Text()
            tambahIde(judul, kategori)
        case 2:
            tampilkanSemuaIde()
        case 3:
            var keyword string
            fmt.Print("Masukkan kata kunci: ")
            scanner.Scan()
            keyword = scanner.Text()
            cariIdeSequential(keyword)
        case 4:
            var keyword string
            fmt.Print("Masukkan judul lengkap: ")
            scanner.Scan()
            keyword = scanner.Text()
            insertionSortByJudul()
            idx := binarySearch(keyword)
            if idx != -1 {
                fmt.Println("Ide ditemukan:")
                tampilkanIde(daftarIde[idx])
            } else {
                fmt.Println("Ide tidak ditemukan.")
            }
        case 5:
            var judul, baru, kategori string
            fmt.Print("Judul yang ingin diedit: ")
            scanner.Scan()
            judul = scanner.Text()
            fmt.Print("Judul baru: ")
            scanner.Scan()
            baru = scanner.Text()
            fmt.Print("Kategori baru: ")
            scanner.Scan()
            kategori = scanner.Text()
            editIde(judul, baru, kategori)
        case 6:
            var judul string
            fmt.Print("Masukkan judul yang ingin dihapus: ")
            scanner.Scan()
            judul = scanner.Text()
            hapusIde(judul)
        case 7:
            var judul string
            fmt.Print("Masukkan judul untuk di-upvote: ")
            scanner.Scan()
            judul = scanner.Text()
            upvoteIde(judul)
        case 8:
            var asc int
            fmt.Print("Urutan (1: ascending, 0: descending): ")
            fmt.Scanln(&asc)
            selectionSortByUpvotes(asc == 1)
            tampilkanSemuaIde()
        case 9:
            insertionSortByJudul()
            tampilkanSemuaIde()
        case 10:
            var start, end string
            fmt.Print("Tanggal mulai (yyyy-mm-dd): ")
            scanner.Scan()
            start = scanner.Text()
            fmt.Print("Tanggal akhir (yyyy-mm-dd): ")
            scanner.Scan()
            end = scanner.Text()
            tampilkanIdePopuler(start, end)
        default:
            fmt.Println("Pilihan tidak valid.")
        }
    }
}

func main() {
	menu()
}