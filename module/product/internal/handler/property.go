package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"example.com/dana/module/product/entity"
	"example.com/dana/module/product/internal/usecase"
)

type PropertyHandler struct {
	usecase usecase.PropertyUsecase
}

func (h *PropertyHandler) ListProperties() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")
		properties, err := h.usecase.List(context.Background(), entity.ListPayload{Search: search})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := entity.ListResponse{Properties: properties}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *PropertyHandler) GetProperty() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		property, err := h.usecase.GetByID(context.Background(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := entity.GetResponse{Property: property}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func NewPropertyHandler(usecase usecase.PropertyUsecase) *PropertyHandler {
	return &PropertyHandler{
		usecase: usecase,
	}
}
