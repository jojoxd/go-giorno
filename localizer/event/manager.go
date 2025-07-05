package event

import (
	"git.jojoxd.nl/projects/go-giorno/localizer/locale"
)

type Manager interface {
	LocaleChangedEvent(oldLocale, newLocale locale.Locale)
	LocalizationNotFoundEvent(key string, locale locale.Locale)

	Channel() <-chan Event
}
