package main

import (
	// Caddy
	"github.com/caddyserver/caddy/caddy/caddymain"

	// Plugins
	_ "github.com/hacdias/caddy-service"
	_ "github.com/FlorianPerrot/caddy-supervisor"
)

func main() {
	caddymain.Run()
}
