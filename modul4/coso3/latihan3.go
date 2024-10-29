package main

import "fmt"

func main() {
	var beratBadan, tinggiBadan, bmi float64
	fmt.Println("masukan berat badan (kg): ")
	fmt.Scan(&beratBadan)
	fmt.Println("masukan tinggi badan (m)")
	fmt.Scan(&tinggiBadan)
	bmi = beratBadan / (tinggiBadan * tinggiBadan)
	fmt.Printf("BMI anda: %.2f", bmi)
}
