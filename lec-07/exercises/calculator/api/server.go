package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/je117er/tfs-03/lec-07/exercises/calculator/api/utils"
)

type result struct {
	Result string `json:"result"`
}

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/api/calc" || req.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
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
