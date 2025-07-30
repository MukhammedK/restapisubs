package routers

import (
	"github.com/gorilla/mux"
	"restapitry/handlers"
)

func Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.MainHandler).Methods("GET")
	r.HandleFunc("/subscriptions", handlers.CreateSubs).Methods("POST")
	r.HandleFunc("/subscriptions/summary", handlers.GetTotalPrice).Methods("GET")
	r.HandleFunc("/subscriptions/{id}", handlers.GetSubsByID).Methods("GET")
	r.HandleFunc("/subscriptions/{id}", handlers.AlterSubs).Methods("PUT")
	r.HandleFunc("/subscriptions/{id}", handlers.DeleteSubs).Methods("DELETE")
	r.HandleFunc("/subscriptions", handlers.GetSubs).Methods("GET")

	return r
}
