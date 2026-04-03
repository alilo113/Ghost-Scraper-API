package internal

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

func BrowserBotEvation(){
	path, _ := launcher.LookPath()
	l := launcher.New().
		Bin(path).
		Headless(true). 
		Set("disable-blink-features", "AutomationControlled").
		Set("no-sandbox")

	
	browser := rod.New().ControlURL(l.MustLaunch()).MustConnect()
	defer brower.MustClose()

	page := stealth.MustPage(browser)

	cleanup := func() {
		browser.MustClose()
		l.Cleanup()
	}

	return page, cleanup
}