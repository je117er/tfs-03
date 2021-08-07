package server

import (
	"fmt"
	"net/http"
)

func Calc(w http.ResponseWriter, req *http.Request) {

	query := string(req.URL.RawQuery)
	if len(query) == 0 {
		fmt.Fprintln(w, "No parameters were supplied")
	} else {
		result, err := eval(string(req.URL.RawQuery))
		if err != nil {
			fmt.Fprintf(w, "%v\n", err)
		}
		fmt.Fprintf(w, "%s\n", result)
	}
}
