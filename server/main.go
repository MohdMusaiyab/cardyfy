package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MohdMusaiyab/cardyfy/api"
)

func main() {
	// Load port from env (default :8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register routes
	api.RegisterRoutes(mux)

	// Serve static files (generated images, assets)
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Root health check
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "✅ Cardyfy server is running")
	})

	// Start server
	log.Printf("🚀 Server running on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
