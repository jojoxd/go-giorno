package router

import (
	"git.jojoxd.nl/projects/go-giorno/contract"
	"git.jojoxd.nl/projects/go-giorno/internal"
)

type config struct {
	logger             contract.Logger
	manageWindowTitle  bool
	defaultWindowTitle string
}

func newDefaultConfig() *config {
	return &config{
		logger:             internal.NewNilLogger(),
		manageWindowTitle:  false,
		defaultWindowTitle: "Gio",
	}
}

func (c *config) Load(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
}

type Option func(*config)

func Logger(logger contract.Logger) Option {
	return func(c *config) {
		c.logger = internal.NewPrefixLogger("gkrouter", logger)
	}
}

func ManageWindowTitle(defaultWindowTitle string) Option {
	return func(c *config) {
		c.manageWindowTitle = true
		c.defaultWindowTitle = defaultWindowTitle
	}
}
