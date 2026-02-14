package config

import (
	"net/http"

	"example.com/dana/module/product/internal/handler"
	"example.com/dana/module/product/internal/repository"
)

func RegisterTransactionServiceTextHandler(mux *http.ServeMux) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}

	db, err := initializeDatabase(cfg)
	if err != nil {
		return err
	}

	listingHandler := handler.NewListingHandler()
	_ = repository.NewListingRepository(db)

	mux.HandleFunc("/listings", listingHandler.ListListings())
	mux.HandleFunc("/listings/{id}", listingHandler.GetListing())

	return nil
}
