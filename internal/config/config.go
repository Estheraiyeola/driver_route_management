package config

import (
	"fmt"
	"github.com/Estheraiyeola/driver-route-management/internal/user/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	loadEnv()
	connect(os.Getenv("DB_NAME"))
}

func ConnectTestDB() {
	loadEnv()
	connect(os.Getenv("DB_NAME_TEST"))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}
}

func connect(dbName string) {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		username, password, host, port, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to %s database: %v", dbName, err)
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Driver{},
		&models.Customer{},
	)
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Printf("✅ Connected to %s successfully!\n", dbName)
}
