package database

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	if godotenv.Load() != nil {
		panic("Gagal loading file env")
	}

	DbUser := os.Getenv("DB_USER")
	DbPass := os.Getenv("DB_PASS")
	DbHost := os.Getenv("DB_HOST")
	DbName := os.Getenv("DB_NAME")

	conn := fmt.Sprint("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&local=Local", DbUser, DbPass, DbHost, DbName)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		panic("Gagal koneksi ke database")
	}

	db.AutoMigrate(&models.Photo{}, &models.User{})
	return db
}

func CloseConnection(db *gorm.DB) {
	sql, err := db.DB()
	if err != nil {
		panic("Gagal menutup koneksi dari database")
	}
	sql.Close()
}