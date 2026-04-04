package main

import (
	"fmt"
	"Ghost-Scraper-API/internal/engine"
)

func main() {
	fmt.Println("👻 Starting the Ghost Engine...")

	page, cleanup := engine.BrowserBotEvation()
	
	defer cleanup()

	fmt.Println("🕵️ Testing stealth on 'sannysoft'...")
	page.MustNavigate("https://bot.sannysoft.com")

	fmt.Println("✅ Success! The Ghost is live.")
}