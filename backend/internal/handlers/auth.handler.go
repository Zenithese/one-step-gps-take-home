package handlers

import (
	"encoding/json"
	"net/http"
	"backend/internal/repositories"
	"backend/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"backend/internal/models"
)

type AuthHandler struct {
	AuthRepo *repositories.AuthRepository
}

func NewAuthHandler(repo *repositories.AuthRepository) *AuthHandler {
	return &AuthHandler{AuthRepo: repo}
}

func (h *AuthHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	err = h.AuthRepo.CreateUser(creds.Username, string(hashedPassword))
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	userID, _, err := h.AuthRepo.GetUserByUsername(creds.Username)
	if err != nil || userID == "" {
		http.Error(w, "Error retrieving user ID", http.StatusInternalServerError)
		return
	}

	token, err := utils.GenerateJWT(userID, creds.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created and logged in successfully"})
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID, passwordHash, err := h.AuthRepo.GetUserByUsername(creds.Username)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if userID == "" || bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(creds.Password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(userID, creds.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}

func (a *AuthHandler) CheckSessionHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	claims, err := utils.ValidateJWT(cookie.Value)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Session valid",
		"user_id":  claims.UserID,
		"username": claims.Username,
	})
}

func (a *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1, // Expire immediately
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
}