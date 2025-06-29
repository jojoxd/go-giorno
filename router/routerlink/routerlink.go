package routerlink

import (
	"gioui.org/layout"
	"gioui.org/widget"

	"git.jojoxd.nl/projects/go-giorno/router"
	"git.jojoxd.nl/projects/go-giorno/router/intent"
)

type Style struct {
	state  *widget.Clickable
	router router.Router
}

func New(router router.Router) *Style {
	return &Style{
		state:  &widget.Clickable{},
		router: router,
	}
}

type Widget func(gtx layout.Context, button *widget.Clickable) layout.Dimensions

func (s Style) Layout(gtx layout.Context, it intent.Base, widget Widget) layout.Dimensions {
	if s.state.Clicked(gtx) {
		s.router.Push(it)
	}

	return widget(gtx, s.state)
}
