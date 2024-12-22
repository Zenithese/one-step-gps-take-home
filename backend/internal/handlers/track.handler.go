package handlers

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
)

func FetchTrackData(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	url := fmt.Sprintf("https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=%s", os.Getenv("OSGPS_API_KEY"))
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch external data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}