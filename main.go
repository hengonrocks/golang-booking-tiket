package main

import "fmt"

type Booking struct {
	namaDepan, namaBelakang, email string
	jumlahTiket uint
}

func main() {
		a := Booking {
		namaDepan	: "",
		namaBelakang 	: "",
		email		: "",
		jumlahTiket	: 0,
	        }
	
		fmt.Println("Aplikasi Booking Tiket")
		fmt.Println("======================")

		fmt.Println("Masukan Nama Anda : ")
		fmt.Scan(&a.namaDepan)

		fmt.Println("Masukan Nama Belakang : ")
		fmt.Scan(&a.namaBelakang)

		fmt.Println("Masukan Email Anda : ")
		fmt.Scan(&a.email)

		fmt.Println("Masukan Jumlah Tiket : ")
		fmt.Scan(&a.jumlahTiket)
	
	fmt.Printf("Terima Kasih %v %v Sudah Memesan Tiket Sebanyak %v\n", a.namaDepan, a.namaBelakang, a.jumlahTiket)
	fmt.Printf("Konfirmasi Pembayaran Sudah Kami Kirimkan Ke Alamat Email %v\n", a.email)
	fmt.Printf("Tiket Tersedia Yang Dapat Dipesan adalah %v Tiket\n", tiketAkhir(a.jumlahTiket))

}

func tiketAkhir(x uint) uint {
	//b := Booking { jumlahTiket }
	x = 50 - x
	return x
}
