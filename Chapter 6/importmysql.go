package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Konfigurasi koneksi ke database MySQL
	dsn := "user:password@tcp(localhost)/database_name?charset=utf8mb4&parseTime=True&loc=Local"

	// Membuka koneksi ke database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Lakukan operasi CRUD atau query lainnya dengan GORM
	// ...
}
