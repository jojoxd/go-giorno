package event

import (
	"git.jojoxd.nl/projects/go-giorno/router/intent"
	"git.jojoxd.nl/projects/go-giorno/router/view"
)

type NavigationType int

const (
	NavigationReplace NavigationType = iota
	NavigationPush
)

type NavigationEvent struct {
	Type   NavigationType
	Intent intent.Base
	View   view.View
}

func (NavigationEvent) implementsEvent() {}
