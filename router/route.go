package router

import (
	"fmt"
	"net/url"
	"runtime"
	"strings"
)

type Route struct {
	name string
	path string
}

type RouteParams interface{}

var NilRoute = Route{}

func (id Route) Name() string {
	return id.name
}

func (id Route) Path() url.URL {
	u, err := url.Parse(fmt.Sprintf("gio-router://%s/%s", id.path, id.name))
	if err != nil {
		panic(fmt.Sprintf("Invalid view id: %+v", err))
	}

	return url.URL{
		Scheme: "gio-router",
		Host:   u.Host,
		Path:   u.Path,
	}
}

func (id Route) String() string {
	return fmt.Sprintf("%s/%s", id.path, id.name)
}

func NewRoute(name string) Route {
	// todo: remove runtime stuff and just use name, force user to be unique
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		funcName := runtime.FuncForPC(pc).Name()
		lastSlash := strings.LastIndexByte(funcName, '/')
		if lastSlash < 0 {
			lastSlash = 0
		}
		lastDot := strings.LastIndexByte(funcName[lastSlash:], '.') + lastSlash
		return Route{name: name, path: funcName[:lastDot]}
	}

	return Route{name: name}
}
