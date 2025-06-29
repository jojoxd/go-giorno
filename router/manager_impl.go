package router

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"sync"

	"gioui.org/app"
)

var _ Manager = (*managerImpl)(nil)

type managerImpl struct {
	window         *app.Window
	stacks         []*ViewStack
	currentTabIdx  int
	routeProviders map[Route]RouteProvider
	// title of the window
	currentTitle  string
	dispatchMutex sync.Mutex
	config        *config
}

func NewManager(window *app.Window, options ...Option) Manager {
	conf := newDefaultConfig()
	conf.Load(options...)

	return &managerImpl{
		window: window,
		config: conf,
	}
}

func (vm *managerImpl) CurrentView() RouteView {
	if len(vm.stacks) <= 0 {
		return nil
	}

	stack := vm.stacks[vm.currentTabIdx]
	return stack.Peek()
}

func (vm *managerImpl) Update(ctx context.Context) {
	if vm.config.manageWindowTitle {
		vm.updateTitle()
	}

	updateRouteView(vm.CurrentView(), ctx)
}

func (vm *managerImpl) updateTitle() {
	newTitle, ok := titleRouteView(vm.CurrentView())
	if !ok {
		newTitle = vm.config.defaultWindowTitle
	}

	if vm.currentTitle != newTitle {
		vm.currentTitle = newTitle
		vm.window.Option(app.Title(vm.currentTitle))
	}
}

func (vm *managerImpl) CurrentViewIndex() int {
	return vm.currentTabIdx
}

func (vm *managerImpl) Register(Id Route, provider RouteProvider) error {
	vm.dispatchMutex.Lock()
	defer vm.dispatchMutex.Unlock()

	if Id == NilRoute {
		return errors.New("cannot register empty view Id")
	}

	if provider == nil {
		return errors.New("view provider is nil")
	}

	if vm.routeProviders == nil {
		vm.routeProviders = make(map[Route]RouteProvider)
	}

	vm.routeProviders[Id] = provider
	vm.config.logger.Info(fmt.Sprintf("registered view %s", Id))

	return nil
}

func (vm *managerImpl) NavBack() RouteView {
	if len(vm.stacks) <= 0 {
		return nil
	}

	stack := vm.stacks[vm.currentTabIdx]
	if stack.Depth() <= 1 {
		// keep the last view
		return stack.Peek()
	}

	vw := stack.Pop()
	finishRouteView(vw, vm.config.logger)

	return stack.Peek()
}

func (vm *managerImpl) HasPrev() bool {
	if len(vm.stacks) <= 0 {
		return false
	}

	stack := vm.stacks[vm.currentTabIdx]
	return stack.Depth() > 1
}

func (vm *managerImpl) RequestSwitch(intent Intent) error {
	// use mutex to guard the dispatching
	vm.dispatchMutex.Lock()
	defer vm.dispatchMutex.Unlock()

	// Even if using an empty intent, vm refreshes the window.
	defer vm.window.Invalidate()

	if intent.Target == NilRoute {
		return nil
	}

	provider, ok := vm.routeProviders[intent.Target]
	if !ok {
		return fmt.Errorf("no target view found: %v", intent.Target)
	}

	var targetView RouteView
	stack := vm.route(&intent)

	// get target view
	if topVw := stack.Peek(); topVw != nil && topVw.Location() == intent.Location() {
		targetView = topVw
	} else {
		targetView = provider()
		err := stack.Push(targetView)
		if err != nil {
			return fmt.Errorf("push to viewstack error: %w", err)
		}
	}

	err := targetView.OnIntent(intent)
	if err != nil {
		stack.Pop()
		return fmt.Errorf("error handling intent: %w", err)
	}

	location := intent.Location()
	vm.config.logger.Info(fmt.Sprintf("switching to %s", location))
	return nil
}

// routeView routes the intent to the proper viewstack/tab by intent.URL()
func (vm *managerImpl) routeView(intent *Intent) *ViewStack {
	if len(vm.stacks) <= vm.currentTabIdx {
		// try to fix the illegal state
		stack := NewViewStack(vm.config.logger)
		vm.stacks = append(vm.stacks, stack)
		vm.currentTabIdx = len(vm.stacks) - 1
		return stack
	}

	// Iterate through all the viewstacks to find the top view with the same location.
	// switch to and replace the existing view.
	for idx, s := range vm.stacks {
		if vm.compareLocations(s.Peek().Location(), intent.Location()) {
			// switch to the tab
			vm.currentTabIdx = idx
			return s
		}
	}

	if intent.RequireNewStack {
		stack := NewViewStack(vm.config.logger)
		vm.stacks = append(vm.stacks, stack)
		vm.currentTabIdx = len(vm.stacks) - 1
		return stack
	}

	// @todo should be actual nil check
	// Respect referer by checking its parent view.
	if intent.Referer != "" && vm.compareLocations(intent.Referer, vm.CurrentView().Location()) {
		// push to current view stack
		return vm.stacks[vm.currentTabIdx]
	}

	// then try to match the viewID:
	for idx, s := range vm.stacks {
		if intent.Target == s.Peek().Id() {
			vm.currentTabIdx = idx
			return s
		}
	}

	// create new stack
	stack := NewViewStack(vm.config.logger)
	vm.stacks = append(vm.stacks, stack)
	vm.currentTabIdx = len(vm.stacks) - 1

	return stack
}

// route the intent to the proper viewstack/tab
func (vm *managerImpl) route(intent *Intent) *ViewStack {
	return vm.routeView(intent)
}

func (vm *managerImpl) OpenedViews() []RouteView {
	views := make([]RouteView, len(vm.stacks))
	for idx, stack := range vm.stacks {
		views[idx] = stack.Peek()
	}

	return views
}

func (vm *managerImpl) CloseTab(idx int) {
	if idx < 0 || idx >= len(vm.stacks) {
		return
	}

	stack := vm.stacks[idx]
	stack.Clear()
	vm.stacks = slices.Delete[[]*ViewStack, *ViewStack](vm.stacks, idx, idx+1)
	if vm.currentTabIdx >= idx && vm.currentTabIdx > 0 {
		vm.currentTabIdx -= 1
	}
}

func (vm *managerImpl) SwitchTab(idx int) {
	if idx >= len(vm.stacks) || idx < 0 {
		return
	}

	vm.currentTabIdx = idx
}

func (vm *managerImpl) Invalidate() {
	vm.window.Invalidate()
}

func (vm *managerImpl) Reset() {
	for _, stack := range vm.stacks {
		stack.Clear()
	}

	vm.currentTabIdx = 0
	vm.stacks = vm.stacks[:0]
	vm.Invalidate()
}

func (vm *managerImpl) compareLocations(a, b RouteLocation) bool {
	vm.config.logger.Debug(fmt.Sprintf("compareLocations %#v, %#v", a, b))

	return a == b
}
