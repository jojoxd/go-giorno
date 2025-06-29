package event

import (
	"git.jojoxd.nl/projects/go-giorno/router2/intent"
	"git.jojoxd.nl/projects/go-giorno/router2/view"
)

type nilManager struct{}

func NewNilManager() Manager {
	return &nilManager{}
}

func (n nilManager) NavigationEvent(NavigationType, view.View, intent.Base) {}

func (n nilManager) Channel() <-chan Event {
	return nil
}
