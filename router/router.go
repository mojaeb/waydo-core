package router

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


func helloHandler(w http.ResponseWriter,_r *http.Request) {
	fmt.Fprintf(w, "hello i am jafar ebrahimi\n")
}



func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloHandler).Methods("GET", "OPTIONS")
	return router
}