package config

import (
	"net/http"

	"example.com/dana/module/product/internal/handler"
	"example.com/dana/module/product/internal/repository"
	"example.com/dana/module/product/internal/usecase"
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

	listingRepo := repository.NewListingRepository(db)
	listingUsecase := usecase.NewListingUsecase(listingRepo)
	listingHandler := handler.NewListingHandler(listingUsecase)
	mux.HandleFunc("/listings", listingHandler.ListListings())
	mux.HandleFunc("/listings/{id}", listingHandler.GetListing())

	return nil
}
