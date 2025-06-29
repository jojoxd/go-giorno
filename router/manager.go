package router

import "context"

// Manager is the core of gio_router. It can be used to:
//  1. manages routeProviders;
//  2. dispatch view requests via the Intent object.
//  3. routeProviders navigation.
//
// Views can have a bounded history stack depending on its intent request.
// The history stack is bounded to a tab widget which is part of a tab bar.
// TODO: Remove the tab stuff and be a pure router
// Most of the API of the view manager handles navigating between routeProviders.
type Manager interface {
	// Register is used to register routeProviders before the view rendering happens.
	// Use provider to enable us to use dynamically constructed routeProviders.
	Register(ID Route, provider RouteProvider) error

	// RequestSwitch tries to switch the current view to the requested view.
	// If referer of the intent equals to the current viewID of the current tab,
	//   the requested view should be routed and pushed to the existing viewstack(current tab).
	//   Otherwise a new viewstack for the intent is created(a new tab),
	//   if there's no duplicate active view (first routeProviders of the stacks).
	RequestSwitch(intent Intent) error

	// OpenedViews return the routeProviders on top of the stack of each tab.
	OpenedViews() []RouteView

	// CloseTab Closes the current tab and move backwards to the previous one if there's any.
	CloseTab(idx int)

	// SwitchTab switches the current tab to the requested tab
	SwitchTab(idx int)

	// CurrentView returns the top most view of the current tab.
	CurrentView() RouteView

	// CurrentViewIndex gives the current tab index
	CurrentViewIndex() int

	Update(ctx context.Context)

	// NavBack navigates back to the last view if there's any and pop out the current view.
	// It returns the view that is to be rendered (which might not be the current view)
	NavBack() RouteView

	// HasPrev checks if there is any naviBack-able routeProvider in the current stack or not.
	// This should not count for the current view.
	HasPrev() bool

	// Invalidate refreshes the window
	Invalidate()

	// Reset resets internal states of the Manager
	Reset()
}
