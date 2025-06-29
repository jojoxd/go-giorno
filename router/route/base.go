package route

import (
	"git.jojoxd.nl/projects/go-giorno/router/intent"
	"git.jojoxd.nl/projects/go-giorno/router/internal"
	"git.jojoxd.nl/projects/go-giorno/router/view"
)

type Base interface {
	Target() internal.Target

	Create() view.View
	ApplyIntent(it intent.Base, view view.View)
}

func BindFactory(route *Route, factory Factory) Base {
	return &baseRoute{
		Route:   route,
		factory: factory,
	}
}

func BindTypedFactory[P any](route *TypedRoute[P], factory TypedFactory[P]) Base {
	return &baseTypedRoute[P]{
		TypedRoute: route,
		factory:    factory,
	}
}

type baseRoute struct {
	*Route
	factory Factory
}

func (b baseRoute) Create() view.View {
	return b.factory()
}

type baseTypedRoute[P any] struct {
	*TypedRoute[P]
	factory TypedFactory[P]
}

func (b baseTypedRoute[P]) Create() view.View {
	return b.factory()
}
