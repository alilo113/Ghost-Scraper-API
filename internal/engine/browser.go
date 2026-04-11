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
        Leakless(false). 
        Headless(false).
        Set("no-sandbox")

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

type ScrapeRequest struct {
    URL      string `json:"url"`      
    Selector string `json:"selector"`
}

func ScrapeData(targetUrl string, selector string) (string, error) {
    page, cleanup := BrowserBotEvation()
    defer cleanup()

    page = page.Timeout(60 * time.Second) 

    fmt.Println("🌐 Navigating to:", targetUrl)
    err := page.Navigate(targetUrl)
    if err != nil {
        return "", err
    }
    
    fmt.Println("⏳ Waiting for selector:", selector)
    
    el, err := page.Element(selector) 
    if err != nil {
        return "", fmt.Errorf("selector not found: %s", selector)
    }

    result := el.MustText()
    fmt.Println("✅ Data extracted:", result)

    return result, nil
}