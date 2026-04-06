package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"Ghost-Scraper-API/internal/engine"
)


func scrapingHandler(w http.ResponseWriter, r *http.Request){
	w.header().set("Content-type", "application/json")
	json.NewEncoder(w).Encode(engine.scrapeData())
}

func main() {
	http.HadleFunc("POST /Srape-Data", scrapingHandler)

	fmt.Println("Server is runing...")
	http.ListenAndServe(":8080", nil)
}