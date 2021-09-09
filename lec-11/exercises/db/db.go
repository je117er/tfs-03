package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/je117er/tfs-03/lec-11/exercises/config"
)

func ConnectDB() (*sql.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USERNAME"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_NAME"))

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTable(DB *sql.DB) error {
	_, err := DB.Exec(`
    CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(60) NOT NULL,
   	sent_confirm_email BOOLEAN DEFAULT 0
)`)
	if err != nil {
		return err
	}
	return nil
}
