package localizer

import (
	"git.jojoxd.nl/projects/go-giorno/localizer/event"
	"git.jojoxd.nl/projects/go-giorno/localizer/locale"
)

type Manager interface {
	// SetLocale sets a new locale, it will error when the locale isn't available
	SetLocale(locale.Locale) error

	// LocalizerFor fetches a localizer for a specific locale
	// it will return an error when the locale isn't available
	LocalizerFor(locale.Locale) (Localizer, error)

	// Localizer returns the Localizer for the current Locale
	Localizer() Localizer

	// Locales returns all available Locale
	Locales() []locale.Locale

	// Events returns an event bus
	Events() <-chan event.Event
}
