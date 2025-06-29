package event

import (
	"git.jojoxd.nl/projects/go-giorno/router2/intent"
	"git.jojoxd.nl/projects/go-giorno/router2/view"
)

type Manager interface {
	NavigationEvent(NavigationType, view.View, intent.Base)
	Channel() <-chan Event
}
