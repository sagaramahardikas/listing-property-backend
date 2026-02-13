package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/listings", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "handling task status\n")
	})

	mux.HandleFunc("/listings/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
		fmt.Fprintf(w, "handling task status with id=%v\n", id)
	})

	http.ListenAndServe("localhost:8888", mux)
}
