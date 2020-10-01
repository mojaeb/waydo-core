package main

import (
	"fmt"
	"log"
	"net/http"
	"waydo/router"
)

func main() {
	r := router.Router()
	fmt.Println("Waydo is running on localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
