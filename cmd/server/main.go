package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"Ghost-Scraper-API/internal/engine"
)

// 1. Define what the user SHOULD send us
type ScrapeRequest struct {
	URL      string `json:"url"`
	Selector string `json:"selector"`
}

func scrapingHandler(w http.ResponseWriter, r *http.Request) {
	// A. Set the response header (Capital 'H' and 'S')
	w.Header().Set("Content-Type", "application/json")

	// B. Read the JSON sent by the user
	var req ScrapeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	// C. Call your engine function with the data from the request
	result, err := engine.ScrapeData(req.URL, req.Selector)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// D. Send back the success result
	json.NewEncoder(w).Encode(map[string]string{"data": result})
}

func main() {
	// Note: HandleFunc (Capital 'H')
	// If you are using Go 1.22+, "POST /Scrape-Data" works. 
	// If older, just use "/Scrape-Data" and check r.Method inside the handler.
	http.HandleFunc("/Scrape-Data", scrapingHandler)

	fmt.Println("🚀 Ghost-Scraper is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}