package localizer

import (
	"fmt"
	"slices"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type GoI18nBundle map[Locale]*i18n.Bundle

type goI18nLocalizerManager struct {
	localizer         Localizer
	fallbackLocalizer Localizer
	bundle            GoI18nBundle
	eventCh           chan LocalizerManagerEvent
}

func NewGoI18nManager(bundle GoI18nBundle, defaultLocale Locale) (Manager, error) {
	manager := &goI18nLocalizerManager{
		bundle:  bundle,
		eventCh: make(chan LocalizerManagerEvent),
	}

	defaultLocalizer, err := manager.LocalizerFor(defaultLocale)
	if err != nil {
		return nil, err
	}
	manager.fallbackLocalizer = defaultLocalizer
	manager.localizer = defaultLocalizer

	return manager, nil
}

func (g *goI18nLocalizerManager) SetLocale(locale Locale) error {
	localizer, err := g.LocalizerFor(locale)
	if err != nil {
		return fmt.Errorf("failed to fetch localizer for locale %s: %w", locale, err)
	}

	oldLocale := g.localizer.Locale()
	g.localizer = localizer

	g.eventCh <- LocaleChangedEvent{
		OldLocale: oldLocale,
		NewLocale: locale,
	}

	return nil
}

func (g *goI18nLocalizerManager) LocalizerFor(locale Locale) (Localizer, error) {
	bundle, ok := g.bundle[locale]
	if !ok {
		return nil, fmt.Errorf("locale not found: %w", ErrLocaleNotFound)
	}

	inner := i18n.NewLocalizer(bundle)

	var fallback *i18n.Localizer
	if g.fallbackLocalizer != nil {
		fallback = g.fallbackLocalizer.Inner()
	}

	localizer := NewGoI18nLocalizer(locale, inner, fallback, g.eventCh)

	return localizer, nil
}

func (g *goI18nLocalizerManager) Localizer() Localizer {
	return g.localizer
}

func (g *goI18nLocalizerManager) Locales() []Locale {
	locales := make([]Locale, 0, len(g.bundle))
	for locale := range g.bundle {
		locales = append(locales, locale)
	}

	// We want to ensure these are in a stable order
	slices.SortFunc(locales, func(a, b Locale) int {
		return strings.Compare(a.String(), b.String())
	})

	return locales
}

func (g *goI18nLocalizerManager) Events() <-chan LocalizerManagerEvent {
	return g.eventCh
}
