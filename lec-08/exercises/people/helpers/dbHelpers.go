package helpers

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"people/config"
	"people/models"
	"strconv"
)

func CreateTable(db *gorm.DB) error {
	if err := db.AutoMigrate(models.PersonDBModel{}); err != nil {
		return err
	}
	return nil
}

func OpenDB(dsn string) (*gorm.DB, error) {
	// Returns a pool of connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetDSN() string {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Fatal("error parsing database port")
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USERNAME"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))
}
