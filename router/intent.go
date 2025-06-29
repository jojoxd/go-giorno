package router

import "fmt"

type Intent struct {
	Target  Route
	Params  RouteParams
	Referer RouteLocation
	// indicates the provider to create a new view instance and show up in a new tab
	RequireNewStack bool
}

func (i Intent) Location() RouteLocation {
	return buildURL(i.Target, i.Params)
}

type RouteLocation string

func buildURL(target Route, params RouteParams) RouteLocation {
	return RouteLocation(fmt.Sprintf("%s:%s:%+v", target.path, target.name, params))
}

func Params[T any](intent Intent) (T, bool) {
	params, ok := intent.Params.(T)
	return params, ok
}
