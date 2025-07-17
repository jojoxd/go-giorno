// Package glayout extends Gio's layout package
package glayout

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
)

type SlotWidget func(gtx layout.Context, ar *AspectRatio) layout.Dimensions

type ConstrainedWidget func(layout.Context, layout.Constraints) layout.Dimensions

// AspectRatio Defines the aspect ratio of a layout.Context
type AspectRatio struct {
	Width, Height int
	Ratio         float32
}

func (ar *AspectRatio) Portrait() bool {
	return ar.Ratio < 1
}

func (ar *AspectRatio) Landscape() bool {
	return ar.Ratio >= 1
}

func (ar *AspectRatio) Update(gtx layout.Context) {
	ar.Width = gtx.Constraints.Max.X
	ar.Height = gtx.Constraints.Max.Y
	ar.Ratio = float32(ar.Width) / float32(ar.Height)
}

func (ar *AspectRatio) Layout(gtx layout.Context, portrait, landscape layout.Widget) layout.Dimensions {
	if ar.Landscape() {
		return landscape(gtx)
	}

	return portrait(gtx)
}

func (ar *AspectRatio) LayoutSlot(gtx layout.Context, slots SlotWidget) layout.Dimensions {
	return slots(gtx, ar)
}

func (ar *AspectRatio) LayoutBounded(gtx layout.Context, ratio float32, widget ConstrainedWidget) layout.Dimensions {
	width, height := ar.Width, ar.Height

	if ar.Portrait() {
		width = int(float32(width) * ratio)
	} else {
		height = int(float32(height) * ratio)
	}

	constraints := layout.Constraints{
		Min: gtx.Constraints.Min,
		Max: image.Point{
			X: width,
			Y: height,
		},
	}

	dimensions := widget(gtx, constraints)

	if dimensions.Size.X > constraints.Max.X || dimensions.Size.Y > constraints.Max.Y {
		clip.Rect(constraints).Push(gtx.Ops).Pop()

		return layout.Dimensions{
			Size:     constraints.Max,
			Baseline: dimensions.Baseline,
		}
	}

	return dimensions
}
