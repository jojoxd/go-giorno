package event

import (
	"git.jojoxd.nl/projects/go-giorno/localizer/locale"
)

// LocaleChangedEvent signals that the internal locale was changed by a call to Manager.SetLocale
type LocaleChangedEvent struct {
	OldLocale, NewLocale locale.Locale
}

func (e LocaleChangedEvent) implementsEvent() {}
