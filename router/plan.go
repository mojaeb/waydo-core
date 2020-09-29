package router

import (
	"net/http"
	"fmt"
)

func GetAllPlansHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "plans of your mind")
}
