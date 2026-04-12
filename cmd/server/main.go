package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	// Ensure this path matches your go.mod folder structure
	"Ghost-Scraper-API/internal/engine" 
)

type ScrapeRequest struct {
	URL      string `json:"url"`
	Selector string `json:"selector"`
}

type GhostWorker struct {
	ID          int
	ActiveTasks int // Changed to ActiveTasks to match your logic below
}

type LoadBalancer struct {
	mu      sync.Mutex
	Workers []*GhostWorker
}

func (lb *LoadBalancer) GetWorker() *GhostWorker {
	// 1. Fixed capitalization: Lock() and Unlock() must be Capitalized
	lb.mu.Lock()
	defer lb.mu.Unlock()

	var bestWorker *GhostWorker
	minTasks := 1000 // 2. Added ':' for short variable declaration

	// 3. Fixed 'range' syntax: for index, value := range slice
	for _, w := range lb.Workers { 
		if w.ActiveTasks < minTasks {
			minTasks = w.ActiveTasks
			bestWorker = w
		}
	}

	// 4. Update the actual worker found, not the Type name
	if bestWorker != nil {
		bestWorker.ActiveTasks++
	}
	return bestWorker
}

func (lb *LoadBalancer) ReleaseWorker(w *GhostWorker) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	w.ActiveTasks--
}

func scrapingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req ScrapeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	result, err := engine.ScrapeData(req.URL, req.Selector)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"data": result})
}

func main() {
	// 5. Fixed missing commas in the slice literal
	lb := &LoadBalancer{
		Workers: []*GhostWorker{
			{ID: 1, ActiveTasks: 0},
			{ID: 2, ActiveTasks: 0},
			{ID: 3, ActiveTasks: 0},
		},
	}

	// Just a test print to show it works
	testWorker := lb.GetWorker()
	fmt.Printf("👻 Test: Assigned Ghost #%d (Active: %d)\n", testWorker.ID, testWorker.ActiveTasks)
	lb.ReleaseWorker(testWorker)

	http.HandleFunc("/Scrape-Data", scrapingHandler)

	fmt.Println("🚀 Ghost-Scraper is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}