package config

import (
	database2 "exercises/internal/models/database"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
)

// ConnectDB connects to database
func ConnectDB() {
	p := Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Fatal("Error parsing database port")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", Config("DB_USERNAME"), Config("DB_PASSWORD"), Config("DB_HOST"), port, Config("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database %s", err)
	}

	err = db.AutoMigrate(&database2.UserDBModel{}, &database2.ProductDBModel{}, &database2.PaymentDBModel{}, &database2.OrderDBModel{}, &database2.OrderItemDBModel{}, &database2.CartDBModel{}, &database2.CartItemDBModel{})
	if err != nil {
		log.Fatal("failed to apply migrations")
	}

	log.Println("successfully migrated")
}
