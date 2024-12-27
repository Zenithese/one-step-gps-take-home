package handlers

import (
    "encoding/json"
    "net/http"
    "backend/internal/models"
    "backend/internal/repositories"
)

type PreferencesHandler struct {
    Repo *repositories.PreferenceRepository
}

func NewPreferenceHandler(repo *repositories.PreferenceRepository) *PreferencesHandler {
    return &PreferencesHandler{Repo: repo}
}

func (h *PreferencesHandler) UpdatePreferences(w http.ResponseWriter, r *http.Request) {
    var update models.Preference
    if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    if err := h.Repo.SavePreferences(&update); err != nil {
        http.Error(w, "Failed to update preferences", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Preferences updated successfully"})
}

func (h *PreferencesHandler) FetchPreferences(w http.ResponseWriter, r *http.Request) {
    preferences, err := h.Repo.GetPreferences()
    if err != nil {
        http.Error(w, "Failed to fetch preferences", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(preferences)
}