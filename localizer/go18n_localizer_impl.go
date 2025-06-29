package localizer

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type goI18nLocalizerImpl struct {
	locale   Locale
	inner    *i18n.Localizer
	fallback *i18n.Localizer
	eventCh  chan<- LocalizerManagerEvent
}

func NewGoI18nLocalizer(
	locale Locale,
	inner *i18n.Localizer,
	fallback *i18n.Localizer,
	eventCh chan<- LocalizerManagerEvent,
) Localizer {
	return &goI18nLocalizerImpl{
		locale:   locale,
		inner:    inner,
		fallback: fallback,
		eventCh:  eventCh,
	}
}

func (l goI18nLocalizerImpl) T(key string) string {
	return l.localizeWithFallback(key, &i18n.LocalizeConfig{
		MessageID: key,
	})
}

func (l goI18nLocalizerImpl) Tf(key string, templateData any) string {
	return l.localizeWithFallback(key, &i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: templateData,
	})
}

func (l goI18nLocalizerImpl) Tc(key string, pluralCount any) string {
	return l.localizeWithFallback(key, &i18n.LocalizeConfig{
		MessageID:   key,
		PluralCount: pluralCount,
	})
}

func (l goI18nLocalizerImpl) Tfc(key string, pluralCount any, templateData any) string {
	return l.localizeWithFallback(key, &i18n.LocalizeConfig{
		MessageID:    key,
		PluralCount:  pluralCount,
		TemplateData: templateData,
	})
}

func (l goI18nLocalizerImpl) Tl(localizable Localizable) string {
	if text, ok := localizable.Localize(l.locale); ok {
		return text
	}

	if stringer, ok := localizable.(fmt.Stringer); ok {
		return stringer.String()
	}

	// TODO: Better handling
	panic(fmt.Errorf("localizer: could not localize %#v", localizable))
}

func (l goI18nLocalizerImpl) Inner() *i18n.Localizer {
	return l.inner
}

func (l goI18nLocalizerImpl) Locale() Locale {
	return l.locale
}

func (l goI18nLocalizerImpl) localizeWithFallback(key string, lc *i18n.LocalizeConfig) string {
	text, err := l.inner.Localize(lc)
	if err == nil {
		return text
	}

	l.eventCh <- LocalizationNotFoundEvent{
		Key:    key,
		Locale: l.locale,
	}

	if l.fallback != nil {
		text, err := l.fallback.Localize(lc)
		if err == nil {
			return text
		}
	}

	return key
}
