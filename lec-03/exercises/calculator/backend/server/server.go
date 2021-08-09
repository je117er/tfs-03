package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../utils"
)

type result struct {
	Result string `json:"result"`
}

func InitServer() {

	mux := http.NewServeMux()
	mux.HandleFunc("/eval", eval)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}

func eval(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//json.NewEncoder(w).Encode("OKOK")

	q := req.URL.Query()
	fmt.Printf("%v\n", q)
	exp := q["exp"][0]
	res, err := utils.Eval(exp)
	if err != nil {
		msg := fmt.Sprintf("Invalid expression: %s\n", exp)
		http.Error(w, msg, http.StatusInternalServerError)
	} else {
		res1D := &result{Result: res}
		json.NewEncoder(w).Encode(res1D)

		//fmt.Fprintf(w, "%s", res)

		fmt.Println(res)
	}
}
