package base

import (
	"log/slog"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
)

type FrameFn func(gtx layout.Context)

func Run(window *app.Window, frame FrameFn) {
	var ops op.Ops

	go func() {
		for {
			switch ev := window.Event().(type) {
			case app.DestroyEvent:
				return

			case app.FrameEvent:
				gtx := app.NewContext(&ops, ev)

				frame(gtx)
				ev.Frame(gtx.Ops)
			}
		}
	}()

	app.Main()
}

func init() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
}
