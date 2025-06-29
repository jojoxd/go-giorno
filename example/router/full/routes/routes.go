package routes

import (
	route2 "git.jojoxd.nl/projects/go-giorno/router/route"
)

var HomePage = route2.New("home")
var SubPage = route2.New("sub")
var TypedPage = route2.NewTyped[string]("typed")
