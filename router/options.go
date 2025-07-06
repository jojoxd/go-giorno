package router

import (
	"git.jojoxd.nl/projects/go-giorno/contract"
	"git.jojoxd.nl/projects/go-giorno/internal"
	"git.jojoxd.nl/projects/go-giorno/router/event"
	"git.jojoxd.nl/projects/go-giorno/router/history"
)

type Option func(config *config)

type config struct {
	logger   contract.Logger
	hist     history.History
	eventMgr event.Manager
}

func (c *config) applyOptions(options ...Option) {
	for _, option := range options {
		option(c)
	}
}

func newDefaultConfig() *config {
	return &config{
		logger:   internal.NewNilLogger(),
		hist:     history.NewSimple(),
		eventMgr: event.NewNilManager(),
	}
}

// WithEventing enables the Router.Events() channel
func WithEventing() Option {
	return func(config *config) {
		config.eventMgr = event.NewManager(4)
	}
}

// WithLogger enables the internal (debug) logger
func WithLogger(logger contract.Logger) Option {
	return func(config *config) {
		config.logger = internal.NewModulePrefixLogger("router", "router", logger)
	}
}

// WithHistory changes the history management to a specific history.History
func WithHistory(hist history.History) Option {
	return func(config *config) {
		config.hist = hist
	}
}
