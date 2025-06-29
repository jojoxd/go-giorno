package route

import (
	intent2 "git.jojoxd.nl/projects/go-giorno/router2/intent"
	"git.jojoxd.nl/projects/go-giorno/router2/internal"
	view2 "git.jojoxd.nl/projects/go-giorno/router2/view"
)

type TypedFactory[P any] func() view2.TypedView[P]

type TypedRoute[P any] struct {
	target internal.Target
}

func NewTyped[P any](id string) *TypedRoute[P] {
	return &TypedRoute[P]{
		target: internal.Target(id),
	}
}

func (r TypedRoute[P]) Target() internal.Target {
	return r.target
}

func (r TypedRoute[P]) Intent(param P) intent2.Base {
	return intent2.NewTyped[P](r.target, param)
}

func (r TypedRoute[P]) ApplyIntent(it intent2.Base, vw view2.View) {
	vw.OnIntent(it)

	// todo: maybe not panic?
	if it, ok := it.(intent2.TypedIntent[P]); ok {
		if vw, ok := vw.(view2.TypedView[P]); ok {
			vw.OnParameter(it.Param)
		} else {
			panic("failed to cast view")
		}
	} else {
		panic("failed to cast intent")
	}
}
