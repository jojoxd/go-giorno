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

type SubPage struct {
	th              *material.Theme
	homeRouterLink  *routerlink.Style
	typedRouterLink *routerlink.Style
	typedParam      *widget.Editor
}

func NewSubPage(th *material.Theme, router router2.Router) *SubPage {
	return &SubPage{
		th:              th,
		homeRouterLink:  routerlink.New(router),
		typedRouterLink: routerlink.New(router),
		typedParam:      &widget.Editor{},
	}
}

func (s SubPage) Layout(gtx layout.Context) layout.Dimensions {
	fmt.Printf("SubPage$Layout\n")

	s.typedParam.Update(gtx)

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return s.homeRouterLink.Layout(gtx, routes.HomePage.Intent(),
				func(gtx layout.Context, button *widget.Clickable) layout.Dimensions {
					return material.Button(s.th, button, "Home").Layout(gtx)
				})
		}),
		layout.Rigid(layout.Spacer{Height: unit.Dp(16)}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Editor(s.th, s.typedParam, "Param").Layout(gtx)
		}),
		layout.Rigid(layout.Spacer{Height: unit.Dp(16)}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return s.typedRouterLink.Layout(gtx, routes.TypedPage.Intent(s.typedParam.Text()),
				func(gtx layout.Context, button *widget.Clickable) layout.Dimensions {
					return material.Button(s.th, button, "Typed Page").Layout(gtx)
				})
		}),
	)
}

func (s SubPage) OnIntent(intent intent.Base) {
	fmt.Printf("SubPage$OnIntent: %#v\n", intent)
}
