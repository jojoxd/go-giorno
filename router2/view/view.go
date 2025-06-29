package view

import (
	"gioui.org/layout"

	"git.jojoxd.nl/projects/go-giorno/router2/intent"
)

type View interface {
	Layout(gtx layout.Context) layout.Dimensions
	OnIntent(intent intent.Base)
}
