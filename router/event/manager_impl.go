package event

import (
	"git.jojoxd.nl/projects/go-giorno/router/intent"
	"git.jojoxd.nl/projects/go-giorno/router/view"
)

type managerImpl struct {
	ch chan Event
}

func NewManager(buffer int) Manager {
	return &managerImpl{
		ch: make(chan Event, buffer),
	}
}

func (m *managerImpl) NavigationEvent(typ NavigationType, vw view.View, it intent.Base) {
	m.ch <- &NavigationEvent{
		Type:   typ,
		Intent: it,
		View:   vw,
	}
}

func (m *managerImpl) Channel() <-chan Event {
	return m.ch
}
