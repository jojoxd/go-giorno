package glayout

import "gioui.org/layout"

type Helper struct {
	aspectRatio *AspectRatio
}

func NewHelper() *Helper {
	return &Helper{
		aspectRatio: &AspectRatio{},
	}
}

func (h Helper) Update(gtx layout.Context) {
	h.aspectRatio.Update(gtx)
}

func (h Helper) AspectRatio() *AspectRatio {
	return h.aspectRatio
}
