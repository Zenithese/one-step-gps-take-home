package handlers

import (
    "encoding/json"
    "net/http"
    "backend/internal/models"
    "backend/internal/repositories"
    "backend/internal/utils"
)

type PreferencesHandler struct {
    Repo *repositories.PreferenceRepository
}

func NewPreferenceHandler(repo *repositories.PreferenceRepository) *PreferencesHandler {
    return &PreferencesHandler{Repo: repo}
}

func (h *PreferencesHandler) UpdatePreferences(w http.ResponseWriter, r *http.Request) {
    userID, err := utils.GetUserIDFromRequest(r)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    var update models.Preference
    if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    if err := h.Repo.SavePreferences(userID, &update); err != nil {
        http.Error(w, "Failed to update preferences", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Preferences updated successfully"})
}

func (h *PreferencesHandler) FetchPreferences(w http.ResponseWriter, r *http.Request) {
    userID, err := utils.GetUserIDFromRequest(r)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    preferences, err := h.Repo.GetPreferences(userID)
    if err != nil {
        http.Error(w, "Failed to fetch preferences", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(preferences)
}