package localizer

type Manager interface {
	// SetLocale sets a new locale, it will error when the locale isn't available
	SetLocale(Locale) error

	// LocalizerFor fetches a localizer for a specific locale
	// it will return an error when the locale isn't available
	LocalizerFor(Locale) (Localizer, error)

	// Localizer returns the Localizer for the current Locale
	Localizer() Localizer

	// Locales returns all available Locale
	Locales() []Locale

	// Events returns an event bus
	Events() <-chan LocalizerManagerEvent
}
