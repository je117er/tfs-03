package database

import (
	"exercises/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
)

// ConnectDB connects to db
func ConnectDB() {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Fatal("Error parsing database port")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USERNAME"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database %s", err)
	}

	err = db.AutoMigrate(&UserDBModel{})
	if err != nil {
		log.Fatal("failed to apply migrations")
	}

	err = db.AutoMigrate(&ProductDBModel{})
	if err != nil {
		log.Fatal("failed to apply migrations")
	}

	err = db.AutoMigrate(&CartDBModel{})
	if err != nil {
		log.Fatal("failed to apply migrations")
	}
	err = db.AutoMigrate(&CartItemDBModel{})
	if err != nil {
		log.Fatal("failed to apply migrations")
	}
	err = db.AutoMigrate(&OrderDBModel{})
	if err != nil {
		log.Fatal("failed to apply migrations")
	}
	err = db.AutoMigrate(&OrderItemDBModel{})
	if err != nil {
		log.Fatal("failed to apply migrations")
	}

	log.Println("successfully migrated")
	/*
		user1 := UserDBModel{
			Username: "jodi",
			Email: "jodi@gmail.com",
			PasswordHash: "hash",
			Carts: []CartDBModel{
				{
					CartItems: []CartItemDBModel{
						{
							ProductID: 1,
							Quantity: 2,
							Price: 15.0,
						},
					},
					Status: "new",
					Total: 30.0,
				},
			},
			Orders: []OrderDBModel{
				{
					OrderItems: []OrderItemDBModel{
						{
							ProductID: 1,
							Quantity: 2,
							Price: 15.0,
						},
					},
					Status: "Paid",
					Total: 30.0,
				},
			},

		}
	*/
	//user2 := models.User{Username: "jo", Email: "jodi@gmail.com", PasswordHash: "hash1"}
	//product1 := ProductDBModel{
	//	Title: "men's shoes",
	//	Description: "",
	//	Quantity: 4,
	//	Price: 15.0,
	//}
	//product2 := models.ProductDBModel{Title: "women's shoes", Description: "", Quantity: 4, Price: 25.0}

	//result := db.Create(&user1)
	//result := db.Delete(&models.User{}, 1)
	//result := db.Exec("delete from users where id=1")
	//log.Println(result.Error)
	//log.Println(result.RowsAffected)
	//log.Println("successfully inserted")
}
