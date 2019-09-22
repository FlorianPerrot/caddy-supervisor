package main

import (
	// Caddy
	"github.com/caddyserver/caddy/caddy/caddymain"

	// Plugins
	_ "github.com/hacdias/caddy-service"
	_ "github.com/FlorianPerrot/caddy-supervisor/httpplugin"
	_ "github.com/FlorianPerrot/caddy-supervisor/servertype"
)

func main() {
	caddymain.Run()
}
