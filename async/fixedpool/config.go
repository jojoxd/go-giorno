package fixedpool

import (
	"runtime"

	"git.jojoxd.nl/projects/go-giorno/contract"
	"git.jojoxd.nl/projects/go-giorno/internal"
)

type Option func(c *config)

type config struct {
	workers int
	logger  contract.Logger
}

func newDefaultConfig() *config {
	conf := &config{
		workers: runtime.NumCPU(),
		logger:  internal.NewNilLogger(),
	}

	return conf
}

func (c *config) Load(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
}

func Logger(logger contract.Logger) Option {
	return func(c *config) {
		c.logger = internal.NewModulePrefixLogger("async", "fixedPoolScheduler", logger)
	}
}

func Workers(workers int) Option {
	return func(c *config) {
		c.workers = workers
	}
}
