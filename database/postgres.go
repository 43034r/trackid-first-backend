package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// Environment variables
/*
var (
	host  = os.Getenv("TI_DB_HOST")
	port = os.Getenv("TI_DB_PORT")
	dbName    = os.Getenv("TI_DB_NAME")
	dbUser    = os.Getenv("TI_DB_USER")
	password    = os.Getenv("TI_DB_PASSWORD")
)
*/

type Trackid struct {
	gorm.Model
	Trackid string `json:"trackid"`
	Status  string `json:"status"`
}

func DatabaseConnection() {
	host := "192.168.1.150"
	port := "5432"
	dbName := "mydb"
	dbUser := "myuser"
	password := "mypass"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(Trackid{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
}
