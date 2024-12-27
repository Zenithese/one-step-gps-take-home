package main

import (
	"log"
	"net/http"
	"backend/internal/handlers"
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/repositories"
)

func main() {
	config.LoadConfig()
	db.InitDB()
	defer db.CloseDB()

	preferenceRepo := repositories.NewPreferenceRepository(db.DB)
	preferenceHandler := handlers.NewPreferenceHandler(preferenceRepo)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/fetch-track-data", handlers.FetchTrackData)
	mux.HandleFunc("/api/preferences", preferenceHandler.UpdatePreferences)
	mux.HandleFunc("/api/preferences/fetch", preferenceHandler.FetchPreferences)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(mux)))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}