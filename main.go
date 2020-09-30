package main

import (
	"net/http"
	"log"
	"fmt"
	"waydo/router"
)


func main() {
	r := router.Router()
	fmt.Println("Waydo is running on localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}