package loader

import (
	"errors"
	"fmt"

	"gioui.org/layout"
)

type Slots[TData any] struct {
	Initial func(gtx layout.Context) layout.Dimensions
	Error   func(gtx layout.Context, err error) layout.Dimensions
	Queued  func(gtx layout.Context) layout.Dimensions
	Loading func(gtx layout.Context) layout.Dimensions
	Loaded  func(gtx layout.Context, data TData) layout.Dimensions
}

var ErrSlotFailed = errors.New("slot failed")
var ErrSlotUndefined = errors.New("slot undefined")

func (s Slots[TData]) Layout(gtx layout.Context, state State[TData]) layout.Dimensions {
	switch state := state.(type) {
	case StateInitial[TData]:
		return s.layoutInitial(gtx)

	case StateError[TData]:
		return s.layoutError(gtx, state.Error)

	case StateQueued[TData]:
		return s.layoutQueued(gtx)

	case StateLoading[TData]:
		return s.layoutLoading(gtx)

	case StateLoaded[TData]:
		return s.layoutLoaded(gtx, state.Data)
	}

	return s.layoutError(gtx, ErrSlotFailed)
}

func (s Slots[TData]) layoutInitial(gtx layout.Context) layout.Dimensions {
	if s.Initial == nil {
		return layout.Dimensions{}
	}

	return s.Initial(gtx)
}

func (s Slots[TData]) layoutError(gtx layout.Context, err error) layout.Dimensions {
	if s.Error == nil {
		// TODO should not panic
		panic(err)
	}

	return s.Error(gtx, err)
}

func (s Slots[TData]) layoutQueued(gtx layout.Context) layout.Dimensions {
	if s.Queued == nil {
		return s.layoutLoading(gtx)
	}

	return s.Queued(gtx)
}

func (s Slots[TData]) layoutLoading(gtx layout.Context) layout.Dimensions {
	if s.Loading == nil {
		return s.layoutInitial(gtx)
	}

	return s.Loading(gtx)
}

func (s Slots[TData]) layoutLoaded(gtx layout.Context, data TData) layout.Dimensions {
	if s.Loaded == nil {
		return s.layoutError(gtx, fmt.Errorf("%w: no loaded slot defined", ErrSlotUndefined))
	}

	return s.Loaded(gtx, data)
}
