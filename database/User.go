package database

import (
	"fmt"
	"log"
	"os"
	"rsv-system-go/config"
	"rsv-system-go/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// const DB_USERNAME =
// const DB_PASSWORD = "123qwe"
// const DB_NAME = "rsv"
// const DB_HOST = "127.0.0.1"
// const DB_PORT = 3306

func InitDb() *gorm.DB {
	db := ConnectDB()
	db.AutoMigrate(&models.User{})
	return db
}

func ConnectDB() *gorm.DB {
	e := godotenv.Load("lokal.env")
	if e != nil {
		log.Fatalf("Erorr env. Err: %s", e)
	}

	conf := config.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	db, err := gorm.Open(mysql.Open(conf.MysqlUrl()), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database: error=%v\n", err)
	}

	return db
}
