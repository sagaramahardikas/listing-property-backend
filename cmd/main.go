package main

import (
	"fmt"
	"net/http"
	"os"

	"example.com/dana/module/product/config"
)

func main() {
	mux := http.NewServeMux()
	err := config.RegisterProductGatewayHandler(mux)
	if err != nil {
		fmt.Printf("failed to register handler: %v\n", err)
		return
	}

	err = http.ListenAndServe("0.0.0.0:8888", mux)
	if err != nil {
		fmt.Printf("failed to start server: %v\n", err)
		os.Exit(1)
	}
}
