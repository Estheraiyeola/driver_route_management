package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	connect("driver_route_management")
}

func ConnectTestDB() {
	connect("driver_route_management_test")
}

func connect(dbName string) {
	username := "root"
	password := "A#1234Esther"
	host := "localhost"
	port := "3306"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		username, password, host, port, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to %s database: %v", dbName, err)
	}

	log.Printf("✅ Connected to %s successfully!\n", dbName)
}
