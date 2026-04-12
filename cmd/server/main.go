package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
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

	// C. Call engine function with the data from the request
	result, err := engine.ScrapeData(req.URL, req.Selector)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// D. Send back the success result
	json.NewEncoder(w).Encode(map[string]string{"data": result})
}


type GhostWorker struct {
	ID int
	ActiveWorker int
}

type LoadBalancer struct {
	mu      sync.Mutex
	Workers []*GhostWorker
}

func (lb *LoadBalancer) GetWorker() *GhostWorker {
	lb.mu.lock()
	defer lb.mu.unlock()

	var bestWorker *GhostWorker
	minTasks = 1000 

	for _, w in range lb.workers{
		if w.ActiveTasks < minTasks {
			minTasks = w.ActiveTasks
			bestWorker = w 
		}
	}

	GhostWorker.ActiveWorker++
	return bestWorker
}

func (lb *LoadBalancer) ReleaseWorker(w *GhostWorker) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	w.ActiveTasks--
}

func main() {
	lb := &LoadBalancer{
		Workers: []*GhostWorker{
			{ID: 1, ActiveWorker: 0}
			{ID: 2, ActiveWorker: 0}
			{ID: 3, ActiveWorker: 0}
		}
	}

	worker := lb.GetWorker()
	fmt.Printf("👻 Request assigned to Ghost #%d (Active: %d)\n", worker.ID, worker.ActiveTasks)

	lb.ReleaseWorker(worker)
	fmt.Printf("✅ Ghost #%d is free again (Active: %d)\n", worker.ID, worker.ActiveTasks)

	http.HandleFunc("/Scrape-Data", scrapingHandler)

	fmt.Println("🚀 Ghost-Scraper is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}