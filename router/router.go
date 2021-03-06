package router

import (
	"github.com/gorilla/mux"
)


func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/plan", GetAllPlans).Methods("GET", "OPTIONS")
	router.HandleFunc("/plan", AddPlan).Methods("POST", "OPTIONS")
	router.HandleFunc("/plan-types",GetAllPlanTypes).Methods("GET", "OPTIONS")
	return router
}