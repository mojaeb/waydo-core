package main

import (
	"fmt"
	"net/http"
	"log"
	"waydo/router"
	"waydo/db"
)


func main() {
	r := router.Router()

	db.New()
	fmt.Println("this is sample app with golang...")
	log.Fatal(http.ListenAndServe(":8082", r))
}