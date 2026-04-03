package main

import (
	"math/rand"
)

type Persona struct {
	UserAgent string
	Width     int
	Height    int
	Language  string
	Platform  string
}

func personaCreation() Persona {
	var personas = []Persona{
		{
			UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
			Width:     1920,
			Height:    1080,
			Language:  "en-US,en;q=0.9",
			Platform:  "Win32",
		},
		{
			UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36",
			Width:     2560,
			Height:    1440,
			Language:  "en-GB,en;q=0.8",
			Platform:  "MacIntel",
		},
		{
			UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Mobile/15E148 Safari/604.1",
			Width:     390,
			Height:    844,
			Language:  "fr-FR,fr;q=0.9",
			Platform:  "iPhone",
		},
	}

	randomIndex := rand.Intn(len(personas))
    
    // THIS IS THE FIX: Return the item at the index
	return personas[randomIndex] 
}

func main() {
	p := personaCreation()
	println("Success! Picked identity:", p.UserAgent)
}