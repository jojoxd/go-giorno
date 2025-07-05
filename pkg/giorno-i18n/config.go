package giorno_i18n

import (
	"git.jojoxd.nl/projects/go-giorno/contract"
	"git.jojoxd.nl/projects/go-giorno/internal"
	"git.jojoxd.nl/projects/go-giorno/localizer/event"
	"git.jojoxd.nl/projects/go-giorno/localizer/locale"
)

type Option func(config *config)
type BundleOption func(config *config)

type config struct {
	bundle         Bundle
	logger         contract.Logger
	eventMgr       event.Manager
	fallbackLocale *locale.Locale
}

func (c *config) load(bundle BundleOption, opts ...Option) {
	bundle(c)

	for _, opt := range opts {
		opt(c)
	}
}

func newDefaultConfig() *config {
	return &config{
		logger:   internal.NewNilLogger(),
		eventMgr: event.NewNilManager(),
	}
}

// WithBundle sets the translation Bundle for a localizer.Manager
func WithBundle(bundle Bundle) BundleOption {
	return func(config *config) {
		config.bundle = bundle
	}
}

// WithEventing enables the events channel
func WithEventing(buffer int) Option {
	return func(config *config) {
		config.eventMgr = event.NewManager(buffer)
	}
}

// WithFallbackLocale sets the fallback locale. It should exist in the Bundle.
// It will also set the starting localizer.Localizer, so no call to SetLocale() is required.
func WithFallbackLocale(fallbackLocale locale.Locale) Option {
	return func(config *config) {
		config.fallbackLocale = &fallbackLocale
	}
}

// WithLogger sets the contract.Logger to be used in the localizer
func WithLogger(logger contract.Logger) Option {
	return func(config *config) {
		config.logger = internal.NewPrefixLogger("go-giorno:localizer-i18n", logger)
	}
}
