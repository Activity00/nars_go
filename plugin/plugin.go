package plugin

import "nars_go/nars"

type Plugin interface {
	SetEnabled(n *nars.Nars, enabled bool)
	Name()
}
