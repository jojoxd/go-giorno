package loader

import "gioui.org/layout"

type Style[TArg comparable, TData any] struct {
	controller Controller[TArg, TData]
}

func New[TArg comparable, TData any](controller Controller[TArg, TData]) *Style[TArg, TData] {
	return &Style[TArg, TData]{
		controller: controller,
	}
}

func (s *Style[TArg, TData]) Load(arg TArg) {
	s.controller.Load(arg)
}

type Widget[TData any] func(gtx layout.Context, state State[TData]) layout.Dimensions

func (s *Style[TArg, TData]) Layout(gtx layout.Context, widget Widget[TData]) layout.Dimensions {
	return widget(gtx, s.controller.State())
}
