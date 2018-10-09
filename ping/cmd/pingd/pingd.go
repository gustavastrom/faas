package main

import (
	"fmt"
	"github.com/sparrc/go-ping"
	"net/http"
)

func response(w http.ResponseWriter, r *http.Request){
	value, ok := r.URL.Query()["p"]
	if !ok || len(value[0]) < 1 {
		fmt.Fprintf(w, "Error: param p is missing")
		return
	}

	p := value[0]
	pinger, err := ping.NewPinger(p)
	if err != nil {
		panic(err)
	}

	pinger.Count = 1
	pinger.Run()
	stats := pinger.Statistics()
	fmt.Fprint(w, stats)
}

func main() {
	http.HandleFunc("/", response)
	http.ListenAndServe(":80", nil)
}