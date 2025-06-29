package event

import (
	"git.jojoxd.nl/projects/go-giorno/router/intent"
	"git.jojoxd.nl/projects/go-giorno/router/view"
)

type Manager interface {
	NavigationEvent(NavigationType, view.View, intent.Base)
	Channel() <-chan Event
}
