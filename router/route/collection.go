package route

import (
	"fmt"

	"git.jojoxd.nl/projects/go-giorno/router/internal"
)

var (
	ErrAlreadyRegistered = fmt.Errorf("route already registered")
	ErrNotRegistered     = fmt.Errorf("route not registered")
)

type Collection struct {
	registry map[internal.Target]Base
}

func NewRouteCollection() *Collection {
	return &Collection{
		registry: make(map[internal.Target]Base),
	}
}

func (c *Collection) Register(route Base) error {
	if _, ok := c.registry[route.Target()]; ok {
		return fmt.Errorf("failed to add route '%s': %w", route.Target(), ErrAlreadyRegistered)
	}

	c.registry[route.Target()] = route
	return nil
}

func (c *Collection) Get(target internal.Target) (Base, error) {
	if route, ok := c.registry[target]; ok {
		return route, nil
	}

	return nil, fmt.Errorf("failed to get route '%s': %w", target, ErrNotRegistered)
}
