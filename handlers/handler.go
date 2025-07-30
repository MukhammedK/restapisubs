package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"restapitry/config"
	"restapitry/models"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "subs.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Ошибка в директории", http.StatusInternalServerError)
		return
	}
	var subs []models.Subscription
	if err := config.DB.Find(&subs).Error; err != nil {
		http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, subs)
	if err != nil {
		http.Error(w, "Ошибка передачи данных", http.StatusInternalServerError)
	}
}
