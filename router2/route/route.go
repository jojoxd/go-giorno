package route

import (
	intent2 "git.jojoxd.nl/projects/go-giorno/router2/intent"
	"git.jojoxd.nl/projects/go-giorno/router2/internal"
	"git.jojoxd.nl/projects/go-giorno/router2/view"
)

type Factory func() view.View

type Route struct {
	target internal.Target
}

func New(id string) *Route {
	return &Route{
		target: internal.Target(id),
	}
}

func (r Route) Target() internal.Target {
	return r.target
}

func (r Route) Intent() intent2.Base {
	return intent2.New(r.target)
}

func (r Route) ApplyIntent(it intent2.Base, vw view.View) {
	vw.OnIntent(it)
}
