package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"zadanie-6105/internal/models"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))

}

func (i *Implementation) CreateTenderHandler(w http.ResponseWriter, r *http.Request) {
	tender := &models.Tender{}

	if err := json.NewDecoder(r.Body).Decode(&tender.Info); err != nil {
		http.Error(w, "Failed to decode tender data.", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	response, err := i.tenderService.Create(ctx, tender)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode tender data.", http.StatusInternalServerError)
		return
	}
}
