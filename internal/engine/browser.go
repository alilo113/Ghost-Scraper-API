package engine

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	// "github.com/go-rod/rod/lib/proto" // Add this for the UserAgent type
	"github.com/go-rod/stealth"
)

func BrowserBotEvation() (*rod.Page, func()) {
    chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"

    l := launcher.New().
        Bin(chromePath).
        // ADD THIS LINE BELOW:
        Leakless(false). 
        Headless(false).
        Set("no-sandbox")

    // Use Launch() instead of MustLaunch() to see the error clearly if it fails
    u, err := l.Launch()
    if err != nil {
        panic(err)
    }

    browser := rod.New().ControlURL(u).MustConnect()
    page := stealth.MustPage(browser)

    cleanup := func() {
        browser.MustClose()
        l.Cleanup()
    }

    return page, cleanup
}

// we gonna take the url and the data from the req
type ScrapeRequest struct {
    URL      string `json:"url"`      // The website to visit
    Selector string `json:"selector"` // The CSS selector (e.g., ".price" or "h1")
}

func ScrapeData(targetUrl string, selector string) (string, error) {
    page, cleanup := BrowserBotEvation()
    defer cleanup()

    // 1. Set a longer deadline for Sétif internet
    page = page.Timeout(60 * time.Second) 

    fmt.Println("🌐 Navigating to:", targetUrl)
    err := page.Navigate(targetUrl)
    if err != nil {
        return "", err
    }
    
    fmt.Println("⏳ Waiting for selector:", selector)
    
    // FIX: Assign both the element (el) and the error (err)
    // We use '=' here because 'err' was already declared above
    el, err := page.Element(selector) 
    if err != nil {
        return "", fmt.Errorf("selector not found: %s", selector)
    }

    // Now that we know 'el' is safe, get the text
    result := el.MustText()
    fmt.Println("✅ Data extracted:", result)

    return result, nil
}