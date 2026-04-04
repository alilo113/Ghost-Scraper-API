package engine

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto" // Add this for the UserAgent type
	"github.com/go-rod/stealth"
)

func BrowserBotEvation() (*rod.Page, func()) {
	persona := personaCreation()
	proxy := GetRandomProxy()

	l := launcher.New().
		Headless(true).
		Proxy(proxy.Address).
		Set("disable-blink-features", "AutomationControlled").
		Set("no-sandbox")

	browser := rod.New().ControlURL(l.MustLaunch()).MustConnect()
	
	// Use 'go' keyword for Auth to prevent hanging during the handshake
	go browser.MustHandleAuth(proxy.Username, proxy.Password)()

	page := stealth.MustPage(browser)

	// FIX: Wrap the string in the proto object
	page.MustSetUserAgent(&proto.NetworkSetUserAgentOverride{
		UserAgent: persona.UserAgent,
	})
	
	page.MustSetViewport(persona.Width, persona.Height, 1, false)

	cleanup := func() {
		browser.MustClose()
		l.Cleanup()
	}

	return page, cleanup
}