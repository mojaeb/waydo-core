package main

import (
	"fmt"
	"net/http"
	"log"
	"waydo/router"
)





func main() {
	r := router.Router()
	fmt.Println("this is sample app with golang...")
	log.Fatal(http.ListenAndServe(":8082", r))
}