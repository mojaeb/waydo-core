package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	_ "database/sql"
	_ "github.com/lib/pq"
)

var (
	host
)

func New() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("load .env file was denied")
	}
	const CONNECTION_STRING string = fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DB_NAME"))
	fmt.Printf("create connection to database %v \n", os.Getenv("USER"))
}

func Migration() {}
func AutoMigration() {}
func SeedCustomData() {}