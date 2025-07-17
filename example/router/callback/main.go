package main

import (
	"fmt"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"git.jojoxd.nl/projects/go-giorno/router"
	"git.jojoxd.nl/projects/go-giorno/router/intent"
	"git.jojoxd.nl/projects/go-giorno/router/route"
	"git.jojoxd.nl/projects/go-giorno/router/routerlink"
	"git.jojoxd.nl/projects/go-giorno/router/view"
)

var Main = route.New("main")

type MainStyle struct {
	theme  *material.Theme
	router router.Router

	callableRouterLink *routerlink.Style
	resolvedValue      string
}

func (m *MainStyle) Layout(gtx layout.Context) layout.Dimensions {
	it := Callable.Intent(CallableParams{
		Call: m.OnCallableReturn,
	})

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Body1(m.theme, fmt.Sprintf("Resolved Value: %s", m.resolvedValue)).Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return m.callableRouterLink.Layout(gtx, it, func(gtx layout.Context, button *widget.Clickable) layout.Dimensions {
				return material.Button(m.theme, button, "Go To Callable").Layout(gtx)
			})
		}),
	)

}

func (m *MainStyle) OnCallableReturn(data string) {
	fmt.Printf("Received back callable: %T:%#v", data, data)
	m.resolvedValue = data
}

func (m *MainStyle) OnIntent(intent intent.Base) {}

type CallableParams struct {
	Call func(data string)
}

var Callable = route.NewTyped[CallableParams]("callable")

type CallableStyle struct {
	theme  *material.Theme
	router router.Router

	param        CallableParams
	submitButton *widget.Clickable
}

func (c *CallableStyle) Layout(gtx layout.Context) layout.Dimensions {
	if c.submitButton.Clicked(gtx) {
		c.param.Call("Resolved")
		c.router.Back()
	}

	return material.Button(c.theme, c.submitButton, "Submit").Layout(gtx)
}

func (c *CallableStyle) OnIntent(intent intent.Base) {}

func (c *CallableStyle) OnParameter(param CallableParams) {
	c.param = param
}

func main() {
	theme := material.NewTheme()

	router := router.NewRouter()
	router.Register(route.BindFactory(Main, func() view.View {
		return &MainStyle{
			theme:              theme,
			router:             router,
			callableRouterLink: routerlink.New(router),
			resolvedValue:      "NONE",
		}
	}))
	router.Register(route.BindTypedFactory(Callable, func() view.TypedView[CallableParams] {
		return &CallableStyle{
			theme:        theme,
			router:       router,
			param:        CallableParams{},
			submitButton: &widget.Clickable{},
		}
	}))

	go func() {
		window := new(app.Window)
		if err := run(window, router); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()

	router.Push(Main.Intent())

	app.Main()
}

func run(window *app.Window, router router.Router) error {
	var ops op.Ops

	for {
		switch ev := window.Event().(type) {
		case app.DestroyEvent:
			return ev.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, ev)

			router.Current().Layout(gtx)

			ev.Frame(gtx.Ops)
		}
	}
}
