package event

import (
	"git.jojoxd.nl/projects/go-giorno/localizer/locale"
)

type nilManager struct{}

func NewNilManager() Manager {
	return &nilManager{}
}

func (n nilManager) LocaleChangedEvent(_, _ locale.Locale) {}

func (n nilManager) LocalizationNotFoundEvent(_ string, _ locale.Locale) {}

func (n nilManager) Channel() <-chan Event {
	return nil
}
