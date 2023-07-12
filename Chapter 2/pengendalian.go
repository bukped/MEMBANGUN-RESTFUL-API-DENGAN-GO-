package main

import "fmt"

func main() {
	var age = 18

	if age >= 18 {
		fmt.Println("Anda sudah dewasa")
	} else {
		fmt.Println("Anda masih anak-anak")
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	switch age {
	case 18:
		fmt.Println("Umur Anda 18 tahun")
	case 21:
		fmt.Println("Umur Anda 21 tahun")
	default:
		fmt.Println("Umur Anda tidak diketahui")
	}
}
