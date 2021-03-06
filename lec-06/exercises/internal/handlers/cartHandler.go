package handlers

import (
	"encoding/json"
	"exercises/internal/models/api"
	database2 "exercises/internal/models/database"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"strconv"
)

type CartHandler struct {
	storage database2.CartStorage
}

func (h CartHandler) PostCart(w http.ResponseWriter, r *http.Request) {
	var createCartRequest api.CreateCartRequest
	if err := json.NewDecoder(r.Body).Decode(&createCartRequest); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cart := api.CartDBModelFromCreateRequest(createCartRequest)
	if err := h.storage.Add(cart); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h CartHandler) PatchCart(w http.ResponseWriter, r *http.Request) {
	var updateCartRequest api.UpdateCartRequest
	err := json.NewDecoder(r.Body).Decode(&updateCartRequest)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// checks if request is empty
	emptyRequest := api.UpdateCartRequest{}
	if emptyRequest == updateCartRequest {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// checks if cart exists
	vars := mux.Vars(r)
	cartID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cart, err := h.storage.ByID(cartID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// now updating
	if updateCartRequest.Status != nil {
		cart.Status = *updateCartRequest.Status
	}

	if updateCartRequest.Products != nil {
		var cartItems []database2.CartItemDBModel
		var productHandler ProductHandler
		var totalAmount float64

		for _, e := range *updateCartRequest.Products {
			// checks if product exists
			product, err := productHandler.storage.ByID(e.ID)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// updates product inventory
			cartItemQuantity := int8(math.Min(float64(product.Quantity), float64(e.Quantity)))
			product.Quantity -= cartItemQuantity

			// update total amount in cart
			totalAmount += float64(cartItemQuantity) * e.Price
			cartItems = append(cartItems, api.ProductToCartItemDBModelRequest(e))
		}
		cart.Total = totalAmount
		cart.CartItems = cartItems
	}
	// now saves the cart to database
	if err := h.storage.Update(cart); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	// gets the cart's id first
	vars := mux.Vars(r)
	cartID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// now gets the cart
	cart, err := h.storage.ByID(cartID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cartResponse := api.CartResponseFromDBModel(cart)
	if err := json.NewEncoder(w).Encode(cartResponse); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h CartHandler) PostCartCheckout(w http.ResponseWriter, r *http.Request) {
	// unpacks json
	var createCheckoutRequest api.CreateCheckoutRequest
	err := json.NewDecoder(r.Body).Decode(createCheckoutRequest)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// converts cart items to order items
	order := api.OrderDBModelFromCreateCheckOutRequest(createCheckoutRequest)

	// saves order
	var orderHandler OrderHandler
	if err := orderHandler.storage.Add(order); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// deletes ordered cart
	if err := h.storage.Delete(createCheckoutRequest.CartID); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
