package router

import (
	"github.com/gorilla/mux"
)




func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/plan", GetAllPlansHandler).Methods("GET", "OPTIONS")
	return router
}