package routerview

import (
	"gioui.org/layout"

	"git.jojoxd.nl/projects/go-giorno/router"
)

type Widget func(gtx layout.Context, currentView router.RouteView) layout.Dimensions

type Style struct {
	mgr router.Manager
}

func New(mgr router.Manager) *Style {
	return &Style{
		mgr: mgr,
	}
}

func (w *Style) Layout(gtx layout.Context, widget Widget) layout.Dimensions {
	return widget(gtx, w.mgr.CurrentView())
}
