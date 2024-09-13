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

//func GetTendersHadler(w http.ResponseWriter, r *http.Request) {
//	limitParam := r.URL.Query().Get("limit")
//	offsetParam := r.URL.Query().Get("offset")
//	serviceTypeParam := r.URL.Query()["service_type"] // Получаем все значения type как массив строк
//
//	// Значения по умолчанию для limit и offset
//	var limit int32 = 5
//	var offset int32 = 0
//
//	// Преобразуем limit и offset в целые числа, если они есть в запросе
//	if limitParam != "" {
//		l, err := strconv.Atoi(limitParam)
//		if err != nil {
//			http.Error(w, "Invalid 'limit' parameter. ", http.StatusBadRequest)
//			return
//		}
//		limit = int32(l)
//
//	}
//
//	if offsetParam != "" {
//		o, err := strconv.Atoi(offsetParam)
//		if err != nil {
//			http.Error(w, "Invalid 'offset' parameter. ", http.StatusBadRequest)
//			return
//		}
//		offset = int32(o)
//
//	}
//
//	if serviceTypeParam != nil {
//
//	}
//	validType := make(map[string]string)
//	validTypes := []string{"Construction", "Delivery", "Manufacture"}
//	for _, st := range serviceTypeParam {
//		if !contains(validTypes, st) {
//			http.Error(w, "Invalid 'service_type' parameter. Must be 'work' or 'personal'.", http.StatusBadRequest)
//			return
//		}
//	}
//}
//

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
