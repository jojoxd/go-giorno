package router

import (
	"fmt"

	"git.jojoxd.nl/projects/go-giorno/router/event"
	"git.jojoxd.nl/projects/go-giorno/router/history"
	"git.jojoxd.nl/projects/go-giorno/router/intent"
	"git.jojoxd.nl/projects/go-giorno/router/route"
	"git.jojoxd.nl/projects/go-giorno/router/view"
)

type router struct {
	*config
	routes *route.Collection
}

func NewRouter(options ...Option) Router {
	config := newDefaultConfig()
	config.applyOptions(options...)

	return &router{
		config: config,
		routes: route.NewRouteCollection(),
	}
}

func (router router) Register(r route.Base) error {
	router.logger.Debug(fmt.Sprintf("registering route '%s'", r.Target()))

	return router.routes.Register(r)
}

func (router router) Push(it intent.Base) error {
	target, err := router.routes.Get(it.Target())
	if err != nil {
		return err
	}

	vw := target.Create()
	target.ApplyIntent(it, vw)

	router.eventMgr.NavigationEvent(event.NavigationPush, vw, it)

	router.hist.Push(&history.Item{
		View:   vw,
		Intent: it,
	})

	router.logger.Debug(fmt.Sprintf("pushed %T with %#v\n", vw, it))

	return nil
}

func (router router) Replace(it intent.Base) error {
	target, err := router.routes.Get(it.Target())
	if err != nil {
		return err
	}

	vw := target.Create()
	target.ApplyIntent(it, vw)

	router.eventMgr.NavigationEvent(event.NavigationReplace, vw, it)

	// todo: finishers?
	router.hist.Clear()
	router.hist.Push(&history.Item{
		View:   vw,
		Intent: it,
	})

	router.logger.Debug(fmt.Sprintf("replaced %T with %#v\n", vw, it))

	return nil
}

func (router router) Current() view.View {
	item := router.hist.Peek()
	if item == nil {
		return nil
	}

	return item.View
}

func (router router) Back() error {
	current := router.hist.Peek()
	defer router.hist.Pop()

	if current == nil {
		return fmt.Errorf("no current view")
	}

	view.Finish(current.View)
	// todo: resolver stuff

	return nil
}

func (router router) Events() <-chan event.Event {
	return router.eventMgr.Channel()
}
