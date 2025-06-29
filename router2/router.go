package router2

import (
	"git.jojoxd.nl/projects/go-giorno/router2/event"
	"git.jojoxd.nl/projects/go-giorno/router2/intent"
	"git.jojoxd.nl/projects/go-giorno/router2/route"
	"git.jojoxd.nl/projects/go-giorno/router2/view"
)

type Router interface {
	// Register registers a route.Route
	Register(route.Base) error

	// Push pushes an Intent onto the current stack
	Push(it intent.Base) error

	Back() error
	Current() view.View

	// Replace replaces the current stack with an Intent
	Replace(it intent.Base) error

	Events() <-chan event.Event
}
