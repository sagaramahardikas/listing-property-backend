package config

import (
	"net/http"

	"example.com/dana/module/product/internal/handler"
	"example.com/dana/module/product/internal/repository"
	"example.com/dana/module/product/internal/usecase"
)

func RegisterProductGatewayHandler(mux *http.ServeMux) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}

	db, err := initializeDatabase(cfg)
	if err != nil {
		return err
	}

	propertyRepo := repository.NewPropertyRepository(db)
	propertyUsecase := usecase.NewPropertyUsecase(propertyRepo)
	propertyHandler := handler.NewPropertyHandler(propertyUsecase)
	mux.HandleFunc("/products/properties", propertyHandler.ListProperties())
	mux.HandleFunc("/products/properties/{id}", propertyHandler.GetProperty())

	return nil
}
