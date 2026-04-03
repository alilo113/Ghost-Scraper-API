package engine

import (
	"fmt"
	"math/rand"
)

// Proxy represents our "Path" credentials
type Proxy struct {
	Address  string // format: "ip:port"
	Username string
	Password string
}

func GetRandomProxy() Proxy {
	// In a real app, you'd load this from a JSON file or Environment Variables
	var proxies = []Proxy{
		{
			Address:  "123.45.67.89:8080",
			Username: "user123",
			Password: "pass_secret_1",
		},
		{
			Address:  "98.76.54.32:9000",
			Username: "user456",
			Password: "pass_secret_2",
		},
	}

	return proxies[rand.Intn(len(proxies))]
}

// ApplyProxy prepares the "Tunnel" for the browser
func ApplyProxy(proxy Proxy) string {
	// We return a string that the Launcher understands
	// Format: --proxy-server=http://123.45.67.89:8080
	return fmt.Sprintf("--proxy-server=%s", proxy.Address)
}