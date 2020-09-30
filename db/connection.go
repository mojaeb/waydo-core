package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
)


func New() *sql.DB{
	// get environment variable from .env file
	err := godotenv.Load(".env")
	// check env file received
	if err != nil {
		log.Fatalf("load .env file was denied")
	}
	db, err := sql.Open("postgres", os.Getenv("CONNECTION_STRING"))
	if err != nil {
		log.Fatal("error in connection to database")
	}
	//check databse is connected
	err = db.Ping()
	if err != nil {
		log.Fatal("db not response to server..")
	}
	fmt.Println("connection to database was successfuly :D")

	return db
}

func Migration() {}
func AutoMigration() {}
func SeedCustomData() {}