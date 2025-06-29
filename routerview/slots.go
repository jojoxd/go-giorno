package routerview

import (
	"gioui.org/layout"

	"git.jojoxd.nl/projects/go-giorno/router"
)

type Slots struct {
	Empty func(gtx layout.Context) layout.Dimensions
	View  func(gtx layout.Context, view router.RouteView) layout.Dimensions
}

func (s Slots) Layout(gtx layout.Context, currentView router.RouteView) layout.Dimensions {
	if currentView == nil {
		return s.Empty(gtx)
	}

	return s.View(gtx, currentView)
}
