package event

import (
	"git.jojoxd.nl/projects/go-giorno/localizer/locale"
)

// LocalizationNotFoundEvent signals that a localization was not found
type LocalizationNotFoundEvent struct {
	Key    string
	Locale locale.Locale
}

func (e LocalizationNotFoundEvent) implementsEvent() {}
