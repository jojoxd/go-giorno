package loader

import (
	"git.jojoxd.nl/projects/go-giorno/contract"
	"git.jojoxd.nl/projects/go-giorno/internal"
)

type ControllerOption[TData any] func(c *controllerConfig[TData])

type controllerConfig[TData any] struct {
	logger                      contract.Logger
	keepStaleDataWhileReloading bool
}

func newControllerConfig[TData any]() controllerConfig[TData] {
	return controllerConfig[TData]{
		logger:                      internal.NewNilLogger(),
		keepStaleDataWhileReloading: false,
	}
}

func (c *controllerConfig[TData]) load(opts ...ControllerOption[TData]) {
	for _, opt := range opts {
		opt(c)
	}
}

func WithLogger[TData any](logger contract.Logger) ControllerOption[TData] {
	return func(c *controllerConfig[TData]) {
		c.logger = internal.NewModulePrefixLogger("loader", "controller", logger)
	}
}

func WithStaleDataWhileReloading[TData any]() ControllerOption[TData] {
	return func(c *controllerConfig[TData]) {
		c.keepStaleDataWhileReloading = true
	}
}
