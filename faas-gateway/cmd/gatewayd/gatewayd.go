package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func info(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Available services: /factorial?n=<int>, /ping?u=<URL>")
}

func factorial(w http.ResponseWriter, r *http.Request){
	value, ok := r.URL.Query()["n"]
	if !ok || len(value[0]) < 1 {
		fmt.Fprintf(w, "Error: param n is missing")
		return
	}

	n := value[0]

	resp, err := http.Get("http://factorial?n=" + n)
	if err != nil {
		fmt.Fprintf(w, "Error: Can't establish connection to factorial service")
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(body))
}

func ping(w http.ResponseWriter, r *http.Request) {
	value, ok := r.URL.Query()["u"]
	if !ok || len(value[0]) < 1 {
		fmt.Fprintf(w, "Error: param u is missing")
		return
	}

	u := value[0]

	resp, err := http.Get("http://ping?u=" + u)
	if err != nil {
		fmt.Fprintf(w, "Error: Can't establish connection to ping service or " + u)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(body))
}

func main() {
	http.HandleFunc("/lambda", info)
	http.HandleFunc("/lambda/info", info)
	http.HandleFunc("/lambda/factorial", factorial)
	http.HandleFunc("/lambda/ping", ping)

	http.ListenAndServe(":80", nil)
}