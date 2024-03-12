package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MyselfRoshan/goHTMX/internals/router"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file.")
	}
	router := router.Routes()
	PORT := os.Getenv("PORT")
	http.ListenAndServe(":"+PORT, router)
	log.Printf("Starting server at http://localhost:%v", PORT)
}
