package database

import (
	"exercises/config"
	models2 "exercises/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"strconv"
)

// ConnectDB connects to db
func ConnectDB() {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Println("Error parsing database port")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USERNAME"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))
	//fmt.Println(dsn)
	db, err := gorm.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		log.Fatalf("failed to connect to database %s", err)
	}
	log.Println("successfully connected to database")

	db.DropTableIfExists(&models2.User{}, &models2.Product{}, &models2.Cart{}, &models2.CartItem{}, &models2.Order{}, &models2.OrderItem{})
	db.AutoMigrate(&models2.User{}, &models2.Product{}, &models2.Cart{}, &models2.CartItem{}, &models2.Order{}, &models2.OrderItem{})
	log.Println("successfully migrated")

	//user1 := models.User{
	//	Username: "jodi",
	//	Email: "jodi@gmail.com",
	//	PasswordHash: "hash",
	//	Carts: []models.Cart{
	//		{
	//			CartItems: []models.CartItem{
	//				{
	//					ProductID: 1,
	//					Quantity: 2,
	//					Price: 15.0,
	//				},
	//			},
	//			Status: "new",
	//			Total: 30.0,
	//		},
	//	},
	//	Orders: []models.Order{
	//		{
	//			OrderItems: []models.OrderItem{
	//				{
	//					ProductID: 1,
	//					Quantity: 2,
	//					Price: 15.0,
	//				},
	//			},
	//			Status: "Paid",
	//			Total: 30.0,
	//		},
	//	},
	//
	//}
	//user2 := models.User{Username: "jo", Email: "jodi@gmail.com", PasswordHash: "hash1"}
	//product1 := models.Product{
	//	Title: "men's shoes",
	//	Description: "",
	//	Quantity: 4,
	//	Price: 15.0,
	//}
	//product2 := models.Product{Title: "women's shoes", Description: "", Quantity: 4, Price: 25.0}

	//result := db.Create(&user1)
	////result := db.Delete(&models.User{}, 1)
	//result := db.Exec("delete from users where id=1")
	//log.Println(result.Error)
	//log.Println(result.RowsAffected)
	//log.Println("successfully inserted")
}
