package event

import (
	"git.jojoxd.nl/projects/go-giorno/localizer/locale"
)

type managerImpl struct {
	ch chan Event
}

func NewManager(buffer int) Manager {
	return &managerImpl{
		ch: make(chan Event, buffer),
	}
}

func (m *managerImpl) LocaleChangedEvent(oldLocale, newLocale locale.Locale) {
	m.ch <- LocaleChangedEvent{
		OldLocale: oldLocale,
		NewLocale: newLocale,
	}
}

func (m *managerImpl) LocalizationNotFoundEvent(key string, locale locale.Locale) {
	m.ch <- LocalizationNotFoundEvent{
		Key:    key,
		Locale: locale,
	}
}

func (m *managerImpl) Channel() <-chan Event {
	return m.ch
}
