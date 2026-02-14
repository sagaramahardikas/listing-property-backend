package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"example.com/dana/module/product/entity"
	"example.com/dana/module/product/internal/usecase"
)

type ListingHandler struct {
	usecase usecase.ListingUsecase
}

func (h *ListingHandler) ListListings() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		listings, err := h.usecase.List(context.Background())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := entity.ListResponse{Listings: listings}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *ListingHandler) GetListing() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		listing, err := h.usecase.GetByID(context.Background(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := entity.GetResponse{Listing: listing}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func NewListingHandler(usecase usecase.ListingUsecase) *ListingHandler {
	return &ListingHandler{
		usecase: usecase,
	}
}
