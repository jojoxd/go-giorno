package gkinset

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

// Uniform applies a layout.UniformInset
func Uniform(v unit.Dp, widget layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(v).Layout(gtx, widget)
	}
}

// Horizontal applies a layout.Inset with equal Left and Right values
func Horizontal(v unit.Dp, widget layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{Left: v, Right: v}.Layout(gtx, widget)
	}
}

// Vertical applies a layout.Inset with equal Top and Bottom values
func Vertical(v unit.Dp, widget layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{Top: v, Bottom: v}.Layout(gtx, widget)
	}
}

func Top(v unit.Dp, widget layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{Top: v}.Layout(gtx, widget)
	}
}

func Right(v unit.Dp, widget layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{Right: v}.Layout(gtx, widget)
	}
}

func Bottom(v unit.Dp, widget layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{Bottom: v}.Layout(gtx, widget)
	}
}

func Left(v unit.Dp, widget layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{Left: v}.Layout(gtx, widget)
	}
}
