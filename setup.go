package caddysupervisor

import (
	"caddysupervisor/httpplugin"
	"caddysupervisor/servertype"
)

func init() {

	httpplugin.Register()
	servertype.Register()

}