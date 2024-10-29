package main

import "fmt"

func main() {
	var totalBelanja, diskon int

	fmt.Print("Masukkan harga belanja: ")
	fmt.Scan(&totalBelanja)

	fmt.Print("Masukkan besar diskon(dalam persen): ")
	fmt.Scan(&diskon)

	potonganDiskon := (totalBelanja * diskon) / 100
	totalAkhir := totalBelanja - potonganDiskon

	fmt.Println("Total belanja setelah diskon:", totalAkhir)
}

