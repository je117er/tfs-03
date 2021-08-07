package main

import (
	"net/http"

	"./server"
)

func main() {

	http.HandleFunc("/calc", server.Calc)

	http.ListenAndServe(":8090", nil)
}
