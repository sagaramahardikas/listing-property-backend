package config

import (
	"net/http"

	"example.com/dana/module/product/internal/handler"
)

func RegisterTransactionServiceTextHandler(mux *http.ServeMux) error {
	listingHandler := handler.NewListingHandler()
	mux.HandleFunc("/listings", listingHandler.ListListings())
	mux.HandleFunc("/listings/{id}", listingHandler.GetListing())

	return nil
}
