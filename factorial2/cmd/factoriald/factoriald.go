package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func factorial(n uint64)(result uint64) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func response(w http.ResponseWriter, r *http.Request){
	value, ok := r.URL.Query()["n"]
	if !ok || len(value[0]) < 1 {
		fmt.Fprintf(w, "Error: param n is missing")
		return
	}

	n, err := strconv.Atoi(value[0])
	if err != nil {
		fmt.Fprintf(w, "Error: param n is not an integer")
		return
	}

	if n > 64 {
		fmt.Fprintf(w, "Error: max value for param n is 64")
		return
	}

	fmt.Fprint(w, factorial(uint64(n)))
}

func main() {
	http.HandleFunc("/", response)
	http.ListenAndServe(":80", nil)
}