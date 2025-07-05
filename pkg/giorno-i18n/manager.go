package giorno_i18n

import (
	"fmt"
	"slices"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"git.jojoxd.nl/projects/go-giorno/localizer"
	"git.jojoxd.nl/projects/go-giorno/localizer/event"
	"git.jojoxd.nl/projects/go-giorno/localizer/locale"
)

type Bundle map[locale.Locale]*i18n.Bundle

type localizerManager struct {
	*config
	localizer         localizer.Localizer
	fallbackLocalizer localizer.Localizer
}

func NewManager(bundle BundleOption, opts ...Option) (localizer.Manager, error) {
	config := newDefaultConfig()
	config.load(bundle, opts...)

	manager := &localizerManager{
		config: config,
	}

	if config.fallbackLocale != nil {
		defaultLocalizer, err := manager.LocalizerFor(*config.fallbackLocale)
		if err != nil {
			return nil, err
		}
		manager.fallbackLocalizer = defaultLocalizer
		manager.localizer = defaultLocalizer
	}

	return manager, nil
}

func (mgr *localizerManager) SetLocale(locale locale.Locale) error {
	newLocalizer, err := mgr.LocalizerFor(locale)
	if err != nil {
		return fmt.Errorf("failed to fetch localizer for locale %s: %w", locale, err)
	}

	oldLocale := mgr.localizer.Locale()
	mgr.localizer = newLocalizer

	mgr.eventMgr.LocaleChangedEvent(oldLocale, locale)

	return nil
}

func (mgr *localizerManager) LocalizerFor(locale locale.Locale) (localizer.Localizer, error) {
	bundle, ok := mgr.bundle[locale]
	if !ok {
		return nil, fmt.Errorf("locale not found: %w", localizer.ErrLocaleNotFound)
	}

	inner := i18n.NewLocalizer(bundle)

	var fallback *i18n.Localizer
	if mgr.fallbackLocalizer != nil {
		fallback = mgr.fallbackLocalizer.Inner()
	}

	localizerForLocale := newLocalizer(locale, inner, fallback, mgr.config)

	return localizerForLocale, nil
}

func (mgr *localizerManager) Localizer() localizer.Localizer {
	return mgr.localizer
}

func (mgr *localizerManager) Locales() []locale.Locale {
	locales := make([]locale.Locale, 0, len(mgr.bundle))
	for localeKey := range mgr.bundle {
		locales = append(locales, localeKey)
	}

	// We want to ensure these are in a stable order
	slices.SortFunc(locales, func(a, b locale.Locale) int {
		return strings.Compare(a.String(), b.String())
	})

	return locales
}

func (mgr *localizerManager) Events() <-chan event.Event {
	return mgr.eventMgr.Channel()
}
