package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"restapitry/config"
	"restapitry/models"
)

func GetTotalPrice(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	serviceName := r.URL.Query().Get("service_name")
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")

	if userID == "" || start == "" || end == "" {
		http.Error(w, "user_id, start, end", http.StatusBadRequest)
		return
	}

	var subs []models.Subscription
	query := config.DB.Where("user_id = ?", userID)

	if serviceName != "" {
		query = query.Where("service_name = ?", serviceName)
	}

	query = query.Where("start_date >= ? AND start_date <= ?", start, end)

	if err := query.Find(&subs).Error; err != nil {
		http.Error(w, "Ошибка при получении данных", http.StatusInternalServerError)
		return
	}

	total := 0
	for _, sub := range subs {
		total += sub.Price
	}

	log.Printf("Суммарная стоимость: %d", total)

	resp := map[string]int{"total_price": total}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
