package api

import (
	"net/http"

	"github.com/MohdMusaiyab/cardyfy/handlers"
)

// RegisterRoutes maps all endpoints to their handlers
func RegisterRoutes(mux *http.ServeMux) {
	// Health check (optional - root already in main.go)
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Main image generation endpoint
	mux.HandleFunc("/api/generate", handlers.GenerateCard)
	mux.HandleFunc("/preview", handlers.PreviewCard)

}
