package main

import (
	"fmt"
	"net/http"

	"example.com/dana/module/product/config"
)

func main() {
	mux := http.NewServeMux()
	err := config.RegisterTransactionServiceTextHandler(mux)
	if err != nil {
		fmt.Printf("failed to register handler: %v\n", err)
		return
	}

	http.ListenAndServe("0.0.0.0:8888", mux)
}
