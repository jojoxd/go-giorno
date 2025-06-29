package router

import (
	"context"
	"fmt"

	"gioui.org/layout"

	"git.jojoxd.nl/projects/go-giorno/contract"
)

type RouteView interface {
	Id() Route
	Layout(gtx layout.Context) layout.Dimensions
	OnIntent(intent Intent) error
	Location() RouteLocation
}

type RouteViewTitler interface {
	Title() string
}

func titleRouteView(view RouteView) (string, bool) {
	if titler, ok := view.(RouteViewTitler); ok {
		return titler.Title(), true
	}

	return "", false
}

type RouteViewFinisher interface {
	OnFinish()
	Finished() bool
}

func finishRouteView(rv RouteView, logger contract.Logger) {
	if finishable, ok := rv.(RouteViewFinisher); ok {
		logger.Debug(fmt.Sprintf("finishing view %T", rv))
		finishable.OnFinish()
		logger.Debug(fmt.Sprintf("finished view %T", rv))
	}
}

type RouteViewUpdater interface {
	Update(ctx context.Context)
}

func updateRouteView(rv RouteView, ctx context.Context) {
	if updater, ok := rv.(RouteViewUpdater); ok {
		updater.Update(ctx)
	}
}
