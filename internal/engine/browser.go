package engine

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

// BrowserBotEvation now returns the Page and a function to close everything
func BrowserBotEvation() (*rod.Page, func()) {
	// 1. Get our "Mask" and "Path" from Phase 1
	persona := personaCreation()
	proxy := GetRandomProxy()

	// 2. Setup the Launcher with the Proxy
	l := launcher.New().
		Headless(true).
		Proxy(proxy.Address). // Applying the "Path"
		Set("disable-blink-features", "AutomationControlled").
		Set("no-sandbox")

	// 3. Launch and Connect
	browser := rod.New().ControlURL(l.MustLaunch()).MustConnect()
	
	// 4. Perform the Proxy Handshake
	browser.MustSetAuth(proxy.Username, proxy.Password)

	// 5. Create the Stealth Page (The Mask)
	page := stealth.MustPage(browser)

	// 6. Apply the Persona details from our Struct
	page.MustSetUserAgent(persona.UserAgent)
	page.MustSetViewport(persona.Width, persona.Height, 1, false)

	// 7. Define the Cleanup function
	cleanup := func() {
		browser.MustClose()
		l.Cleanup()
	}

	return page, cleanup
