package localizer

type LocalizerManagerEvent interface{}

// LocaleChangedEvent signals that the internal locale was changed by a call to Manager.SetLocale
type LocaleChangedEvent struct {
	LocalizerManagerEvent
	OldLocale Locale
	NewLocale Locale
}

// LocalizationNotFoundEvent signals that a localization was not found
type LocalizationNotFoundEvent struct {
	LocalizerManagerEvent
	Key    string
	Locale Locale
}
