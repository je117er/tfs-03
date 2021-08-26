package router

import (
	"exercises/config"
	"exercises/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitServer() {
	config.ConnectDB()
	r := mux.NewRouter()

	userHandler := handler.UserHandler{}
	productHandler := handler.ProductHandler{}
	cartHandler := handler.CartHandler{}
	orderHandler := handler.OrderHandler{}
	paymentHandler := handler.PaymentHandler{}

	// handles users
	r.HandleFunc("/user/{id:\\d+}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/user/{id:\\d+}", userHandler.PatchUser).Methods("PATCH")
	r.HandleFunc("/user/", userHandler.PostUser).Methods("POST")
	r.HandleFunc("/user/{id:\\d+}", userHandler.DeleteUser).Methods("DELETE")

	// handles products
	r.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	r.HandleFunc("/product/{id:\\d+}", productHandler.GetProduct).Methods("GET")
	r.HandleFunc("/product/{id:\\d+}", productHandler.PatchProduct).Methods("PATCH")
	r.HandleFunc("/product/", productHandler.PostProduct).Methods("POST")
	r.HandleFunc("/product/{id:\\d+}", productHandler.DeleteProduct).Methods("DELETE")

	// handles carts
	r.HandleFunc("/cart", cartHandler.PostCart).Methods("POST")
	r.HandleFunc("/cart/{id:\\d+}", cartHandler.PatchCart).Methods("PATCH")
	r.HandleFunc("/cart/{id:\\d+}", cartHandler.GetCart).Methods("GET")
	r.HandleFunc("/cart{id:\\d+}/checkout", cartHandler.PostCartCheckout).Methods("POST")

	// handles orders and payments
	r.HandleFunc("/orders", orderHandler.GetOrder).Methods("GET")
	r.HandleFunc("/order/{id:\\d+}", orderHandler.GetOrders).Methods("GET")
	r.HandleFunc("/order/{id:\\d+}/pay", paymentHandler.CreatePayment).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
