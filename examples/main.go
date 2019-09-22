package main

import (
	// Caddy
	"github.com/caddyserver/caddy/caddy/caddymain"

	// Plugins
	_ "github.com/hacdias/caddy-service"
	_ "github.com/FlorianPerrot/caddy-supervisor/servertype"
	_ "github.com/FlorianPerrot/caddy-supervisor/httpplugin"
)

func main() {
	caddymain.Run()
}
