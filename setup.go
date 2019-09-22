package supervisor

import (
	"github.com/FlorianPerrot/caddy-supervisor/httpplugin"
	"github.com/FlorianPerrot/caddy-supervisor/servertype"
)

func init() {

	httpplugin.Register()
	servertype.Register()

}