package config

import (
	"fmt"
	"goauth/model"

	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgresDB() *gorm.DB {
	godotenv.Load()

	PORT := os.Getenv("DBPORT")

	port, _ := strconv.ParseUint(PORT, 10, 64)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DBHOST"), os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME"),port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	db.AutoMigrate(&model.Users{})

	return db
}

func GoogleConfig() (oauth2.Config) {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Some error occured. Err: %s", err)
    }

    GoogleLoginConfig := oauth2.Config{
        RedirectURL:  "http://localhost:8080/callback",
        ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
        ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
        Scopes: []string{"email", "profile"},
        Endpoint: google.Endpoint,
    }

	log.Println(os.Getenv("GOOGLE_CLIENT_ID"))
	log.Println(os.Getenv("GOOGLE_CLIENT_SECRET"))

    return GoogleLoginConfig
}