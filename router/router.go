package router

import (
	"git.jojoxd.nl/projects/go-giorno/router/event"
	"git.jojoxd.nl/projects/go-giorno/router/intent"
	"git.jojoxd.nl/projects/go-giorno/router/route"
	"git.jojoxd.nl/projects/go-giorno/router/view"
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
