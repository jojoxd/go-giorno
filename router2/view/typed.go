package view

type TypedView[P any] interface {
	View
	OnParameter(P)
}
