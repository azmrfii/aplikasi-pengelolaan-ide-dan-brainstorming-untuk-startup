package main

import "fmt"

const MAX_IDE = 100

type Ide struct {
	Judul    string
	Kategori string
	Tanggal  string
	Upvotes  int
}

var daftarIde [MAX_IDE]Ide
var jumlahIde int

func tambahIde(judul, kategori, tanggal string) {
	if jumlahIde < MAX_IDE {
		daftarIde[jumlahIde] = Ide{judul, kategori, tanggal, 0}
		jumlahIde++
		fmt.Println("Ide berhasil ditambahkan!")
	} else {
		fmt.Println("Kapasitas penyimpanan penuh!")
	}
}

func cariIdeSequential(keyword string) {
	fmt.Println("\nHasil Pencarian:")
	ketemu := false
	for i := 0; i < jumlahIde; i++ {
		if daftarIde[i].Judul == keyword {
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
			if (ascending && daftarIde[j].Upvotes < daftarIde[idx].Upvotes) ||
				(!ascending && daftarIde[j].Upvotes > daftarIde[idx].Upvotes) {
				idx = j
			}
		}
		daftarIde[i], daftarIde[idx] = daftarIde[idx], daftarIde[i]
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

func menu() {
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
		fmt.Println("8. Urutkan Ide (Upvotes)")
		fmt.Println("9. Urutkan Ide (Judul)")
		fmt.Println("10. Tampilkan Ide Populer (Tanggal)")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			var judul, kategori, tanggal string
			fmt.Print("Masukkan Judul: ")
			fmt.Scanln(&judul)
			fmt.Print("Masukkan Kategori: ")
			fmt.Scanln(&kategori)
			fmt.Print("Masukkan Tanggal (yyyy-mm-dd): ")
			fmt.Scanln(&tanggal)
			tambahIde(judul, kategori, tanggal)
		case 2:
			tampilkanSemuaIde()
		case 3:
			var keyword string
			fmt.Print("Masukkan Judul yang dicari: ")
			fmt.Scanln(&keyword)
			cariIdeSequential(keyword)
		case 4:
			var keyword string
			fmt.Print("Masukkan Judul lengkap: ")
			fmt.Scanln(&keyword)
			insertionSortByJudul()
			idx := binarySearch(keyword)
			if idx != -1 {
				tampilkanIde(daftarIde[idx])
			} else {
				fmt.Println("Ide tidak ditemukan.")
			}
		case 5:
			var judul, baru, kategori string
			fmt.Print("Judul yang ingin diedit: ")
			fmt.Scanln(&judul)
			fmt.Print("Judul baru: ")
			fmt.Scanln(&baru)
			fmt.Print("Kategori baru: ")
			fmt.Scanln(&kategori)
			editIde(judul, baru, kategori)
		case 6:
			var judul string
			fmt.Print("Judul yang ingin dihapus: ")
			fmt.Scanln(&judul)
			hapusIde(judul)
		case 7:
			var judul string
			fmt.Print("Judul untuk upvote: ")
			fmt.Scanln(&judul)
			upvoteIde(judul)
		case 8:
			var asc int
			fmt.Print("Urutkan by upvotes (1: Asc, 0: Desc): ")
			fmt.Scanln(&asc)
			selectionSortByUpvotes(asc == 1)
			tampilkanSemuaIde()
		case 9:
			insertionSortByJudul()
			tampilkanSemuaIde()
		case 10:
			var start, end string
			fmt.Print("Tanggal mulai (yyyy-mm-dd): ")
			fmt.Scanln(&start)
			fmt.Print("Tanggal akhir (yyyy-mm-dd): ")
			fmt.Scanln(&end)
			tampilkanIdePopuler(start, end)
		case 0:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func main() {
	menu()
}
