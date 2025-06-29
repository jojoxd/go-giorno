package intent

import (
	"git.jojoxd.nl/projects/go-giorno/router/internal"
)

type TypedIntent[P any] struct {
	target internal.Target
	Param  P
}

func NewTyped[P any](target internal.Target, param P) TypedIntent[P] {
	return TypedIntent[P]{
		target: target,
		Param:  param,
	}
}

func (i TypedIntent[P]) Target() internal.Target {
	return i.target
}

func (i TypedIntent[P]) implementsIntent() {}
