package main

import (
	"log"
	"log/slog"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"

	pages2 "git.jojoxd.nl/projects/go-giorno/example/router/full/pages"
	"git.jojoxd.nl/projects/go-giorno/example/router/full/routes"
	"git.jojoxd.nl/projects/go-giorno/router2"
	"git.jojoxd.nl/projects/go-giorno/router2/route"
	"git.jojoxd.nl/projects/go-giorno/router2/view"
)

type context struct {
	router router2.Router
}

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	ctx := context{
		router: router2.NewRouter(router2.WithLogger(slog.Default())),
	}

	th := material.NewTheme()

	ctx.router.Register(route.BindFactory(routes.HomePage, func() view.View {
		return pages2.NewHomePage(th, ctx.router)
	}))

	ctx.router.Register(route.BindFactory(routes.SubPage, func() view.View {
		return pages2.NewSubPage(th, ctx.router)
	}))

	ctx.router.Register(route.BindTypedFactory(routes.TypedPage, func() view.TypedView[string] {
		return pages2.NewTypedPage(th, ctx.router)
	}))

	go func() {
		window := new(app.Window)
		if err := run(window, ctx); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()

	ctx.router.Push(routes.HomePage.Intent())

	app.Main()
}

func run(window *app.Window, ctx context) error {
	var ops op.Ops

	for {
		switch ev := window.Event().(type) {
		case app.DestroyEvent:
			return ev.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, ev)

			frame(gtx, ctx)

			ev.Frame(gtx.Ops)
		}
	}
}

func frame(gtx layout.Context, ctx context) layout.Dimensions {
	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return ctx.router.Current().Layout(gtx)
	})
}
