package router

// RouteProvider is used to construct a new RouteView instance. Each view should have its own provider.
// Usually this is the constructor of the view, but this can be extended to give more context.
// Example: func() { return MyView{} }
// Example: func() { return ContextualView{app} }
type RouteProvider func() RouteView
