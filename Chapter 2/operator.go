package main

import "fmt"

func main() {
	var a = 10
	var b = 5

	fmt.Println("Penjumlahan:", a+b)
	fmt.Println("Pengurangan:", a-b)
	fmt.Println("Perkalian:", a*b)
	fmt.Println("Pembagian:", a/b)
	fmt.Println("Modulus:", a%b)
	fmt.Println("Pembanding:", a == b)
	fmt.Println("Logika:", a > 5 && b < 10)
	fmt.Println("Bitwise:", a&b)
}
