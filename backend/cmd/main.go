package main

import (
	"log"
	"net/http"
	"backend/internal/handlers"
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/repositories"
	"backend/internal/middleware"
	"backend/internal/utils"
)

func main() {
	utils.LoadEnv()

	config.LoadConfig()
	db.InitDB()
	defer db.CloseDB()

	preferenceRepo := repositories.NewPreferenceRepository(db.DB)
	preferenceHandler := handlers.NewPreferenceHandler(preferenceRepo)
	authRepo := repositories.NewAuthRepository(db.DB)
	authHandler := handlers.NewAuthHandler(authRepo)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/signup", authHandler.SignupHandler)
	mux.HandleFunc("/api/login", authHandler.LoginHandler)
	mux.HandleFunc("/api/check-session", middleware.AuthMiddleware(http.HandlerFunc(authHandler.CheckSessionHandler)))
	mux.HandleFunc("/api/logout", middleware.AuthMiddleware(authHandler.LogoutHandler))
	mux.HandleFunc("/api/fetch-track-data", middleware.AuthMiddleware(handlers.FetchTrackData))
	mux.HandleFunc("/api/preferences", middleware.AuthMiddleware(preferenceHandler.UpdatePreferences))
	mux.HandleFunc("/api/preferences/fetch", middleware.AuthMiddleware(preferenceHandler.FetchPreferences))


	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", middleware.CORSMiddleware(mux)))
}

// package main

// import (
// 	"log"
// 	"net/http"
// 	"backend/internal/handlers"
// 	"backend/internal/config"
// 	"backend/internal/db"
// 	"backend/internal/repositories"
// 	"backend/internal/utils"
// 	"backend/internal/middleware"
// )

// func main() {
// 	utils.LoadEnv()

// 	config.LoadConfig()
// 	db.InitDB()
// 	defer db.CloseDB()

// 	preferenceRepo := repositories.NewPreferenceRepository(db.DB)
// 	preferenceHandler := handlers.NewPreferenceHandler(preferenceRepo)

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/api/fetch-track-data", handlers.FetchTrackData)
// 	mux.HandleFunc("/api/preferences", preferenceHandler.UpdatePreferences)
// 	mux.HandleFunc("/api/preferences/fetch", preferenceHandler.FetchPreferences)

// 	log.Println("Server is running on port 8080...")
// 	log.Fatal(http.ListenAndServe(":8080", middleware.CORSMiddleware(mux)))
// }