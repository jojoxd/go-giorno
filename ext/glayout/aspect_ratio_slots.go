package glayout

import "gioui.org/layout"

type AspectRatioSlots struct {
	Portrait, Landscape layout.Widget
}

func (slots AspectRatioSlots) Layout(gtx layout.Context, ar *AspectRatio) layout.Dimensions {
	if ar.Portrait() && slots.Portrait != nil {
		return slots.Portrait(gtx)
	}

	if ar.Landscape() && slots.Landscape != nil {
		return slots.Landscape(gtx)
	}

	panic("no slots defined")
}
