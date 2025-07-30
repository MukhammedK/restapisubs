package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restapitry/config"
	"restapitry/models"
)

// Создание новой подписки
// Метод добавляет запись о подписке в базу данных
// Требует JSON с названием сервиса, ценой, ID пользователя и датой начала
// @Tags Подписки
// @Accept json
// @Produce json
// @Param subscription body models.Subscription true "Новая подписка"
// @Success 201 {object} models.Subscription
// @Failure 400 {string} string "Ошибка в теле запроса"
// @Router /subscriptions [post]
func CreateSubs(w http.ResponseWriter, r *http.Request) {
	var subs models.Subscription

	if err := json.NewDecoder(r.Body).Decode(&subs); err != nil {

		http.Error(w, "Неправильный ввод данных", http.StatusBadRequest)
	}

	if err := config.DB.Create(&subs).Error; err != nil {

		http.Error(w, "Не удалось добавить в таблицу", http.StatusBadRequest)
	}
	log.Printf("Подписка создана: service=%s, user_id=%s, price=%.2f\n", subs.ServiceName, subs.UserID, subs.Price)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(subs)
}

// Удалить подписку
// Удаляет подписку по переданному ID
// @Tags Подписки
// @Produce plain
// @Param id path string true "ID подписки"
// @Success 204 {string} string "Удалено"
// @Failure 500 {string} string "Ошибка при удалении"
// @Router /subscriptions/{id} [delete]
func DeleteSubs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := config.DB.Delete(&models.Subscription{}, "id = ?", id).Error; err != nil {
		http.Error(w, "Не удалось удалить", http.StatusInternalServerError)
		return
	}

	log.Println("Подписка удалена:", id)
	w.WriteHeader(http.StatusNoContent)

}

// Обновить данные подписки
// Принимает JSON и обновляет существующую подписку по ID
// @Tags Подписки
// @Accept json
// @Produce json
// @Param id path string true "ID подписки"
// @Param subscription body models.Subscription true "Обновлённые данные"
// @Success 200 {object} models.Subscription
// @Failure 400 {string} string "Ошибка в теле запроса"
// @Router /subscriptions/{id} [put]
func AlterSubs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var sub models.Subscription
	if err := config.DB.First(&sub, "id = ?", id).Error; err != nil {
		http.Error(w, "Подписка не найдена", http.StatusNotFound)
		return
	}

	var updated models.Subscription
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {

		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	sub.ServiceName = updated.ServiceName
	sub.Price = updated.Price
	sub.UserID = updated.UserID
	sub.StartDate = updated.StartDate
	sub.EndDate = updated.EndDate

	if err := config.DB.Save(&sub).Error; err != nil {
		http.Error(w, "Ошибка при обновлении", http.StatusInternalServerError)
		return
	}

	log.Println("Подписка обновлена:", sub.ID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sub)

}

// GetSubs godoc
// @Summary Получить список всех подписок
// @Tags subscriptions
// @Produce json
// @Success 200 {array} models.Subscription
// @Router /subscriptions [get]
func GetSubs(w http.ResponseWriter, r *http.Request) {
	var subs []models.Subscription
	if err := config.DB.Find(&subs).Error; err != nil {
		http.Error(w, "Не нашли данные", http.StatusInternalServerError)
	}
	log.Printf("Нашли %d записи\n", len(subs))

	for _, sub := range subs {
		log.Println(sub.ID)
		log.Println(sub.ServiceName)
		log.Println(sub.Price)
		log.Println(sub.UserID)
		log.Println(sub.StartDate)
		log.Println()
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subs)

}

// Получить подписку по ID
// Ищет подписку по UUID и возвращает её, если найдена
// @Tags Подписки
// @Produce json
// @Param id path string true "ID подписки (UUID)"
// @Success 200 {object} models.Subscription
// @Failure 404 {string} string "Не найдено"
// @Router /subscriptions/{id} [get]
func GetSubsByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var sub models.Subscription
	if err := config.DB.First(&sub, "id = ?", id).Error; err != nil {
		http.Error(w, "Подписка не найден", http.StatusNotFound)
		return
	}

	log.Println("Найдена подписка:")
	log.Println(sub.ServiceName)
	log.Println(sub.Price)
	log.Println(sub.UserID)
	log.Println(sub.StartDate)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sub)

}
