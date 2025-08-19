package main

import (
	"github.com/MohdMusaiyab/cardyfy/api"
	"github.com/MohdMusaiyab/cardyfy/db"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	api.RegisterRoutes(mux)
	db.Connect()
	defer db.Disconnect()

	// Your server logic here
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
