package pages

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"git.jojoxd.nl/projects/go-giorno/example/router/full/routes"
	"git.jojoxd.nl/projects/go-giorno/router2"
	"git.jojoxd.nl/projects/go-giorno/router2/intent"
	"git.jojoxd.nl/projects/go-giorno/router2/routerlink"
)

type TypedPage struct {
	th             *material.Theme
	param          string
	homeRouterLink *routerlink.Style
	subRouterLink  *routerlink.Style
}

func NewTypedPage(th *material.Theme, router router2.Router) *TypedPage {
	return &TypedPage{
		th:             th,
		homeRouterLink: routerlink.New(router),
		subRouterLink:  routerlink.New(router),
	}
}

func (t *TypedPage) Layout(gtx layout.Context) layout.Dimensions {
	fmt.Printf("TypedPage$Layout\n")
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Body1(t.th, t.param).Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: unit.Dp(16)}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return t.homeRouterLink.Layout(gtx, routes.HomePage.Intent(),
				func(gtx layout.Context, button *widget.Clickable) layout.Dimensions {
					return material.Button(t.th, button, "Home").Layout(gtx)
				})
		}),
		layout.Rigid(layout.Spacer{Height: unit.Dp(16)}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return t.subRouterLink.Layout(gtx, routes.SubPage.Intent(),
				func(gtx layout.Context, button *widget.Clickable) layout.Dimensions {
					return material.Button(t.th, button, "SubPage").Layout(gtx)
				})
		}),
	)
}

func (t *TypedPage) OnIntent(intent intent.Base) {
	fmt.Printf("TypedPage$OnIntent %#v\n", intent)
}

func (t *TypedPage) OnParameter(p string) {
	fmt.Printf("TypedPage$OnParameter %#v\n", p)
	t.param = p
}
