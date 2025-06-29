package loader

import (
	"context"
	"sync"

	"git.jojoxd.nl/projects/go-giorno/async"
)

type loaderController[TArg comparable, TData any] struct {
	config    controllerConfig[TData]
	scheduler async.Scheduler
	loader    Loader[TArg, TData]
	state     State[TData]
	stateMu   sync.RWMutex

	// TODO: Needs to be changed to caching
	currentArg           TArg
	cancelCurrentContext context.CancelFunc
}

func NewLoaderController[TArg comparable, TData any](
	scheduler async.Scheduler,
	loader Loader[TArg, TData],
	opts ...ControllerOption[TData],
) Controller[TArg, TData] {
	config := newControllerConfig[TData]()
	config.load(opts...)

	return &loaderController[TArg, TData]{
		config:    config,
		scheduler: scheduler,
		loader:    loader,
		state:     StateInitial[TData]{},
	}
}

func (c *loaderController[TArg, TData]) Load(arg TArg) {
	// @TODO Replace with caching
	if c.currentArg == arg {
		return
	}
	c.currentArg = arg

	if c.cancelCurrentContext != nil {
		c.config.logger.Debug("canceled previous context")
		c.cancelCurrentContext()
	}

	ctx := context.TODO()
	workContext, cancelWorkContext := context.WithCancel(ctx)
	c.cancelCurrentContext = cancelWorkContext

	c.setState(StateQueued[TData]{})

	c.scheduler.Schedule(workContext,
		async.ScheduleFn(func(ctx context.Context) {
			c.setState(StateLoading[TData]{})

			data, err := c.loader.Load(ctx, arg)
			if err != nil {
				c.setState(StateError[TData]{Error: err})
				return
			}

			c.setState(StateLoaded[TData]{Data: data})
		}),
	)
}

func (c *loaderController[TArg, TData]) setState(newState State[TData]) {
	c.stateMu.Lock()
	defer c.stateMu.Unlock()

	c.state = newState
}

func (c *loaderController[TArg, TData]) State() State[TData] {
	c.stateMu.RLock()
	defer c.stateMu.RUnlock()

	return c.state
}

func (c *loaderController[TArg, TData]) Data() (TData, bool) {
	c.stateMu.RLock()
	defer c.stateMu.RUnlock()

	if state, ok := c.state.(StateLoaded[TData]); ok {
		return state.Data, true
	}

	return *new(TData), false
}
