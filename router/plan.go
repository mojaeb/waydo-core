package router

import (
	"net/http"
	"fmt"
)

func GetAllPlans(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "plans of your mind")
}

func AddPlan(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a plan")
}