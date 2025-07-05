package giorno_i18n

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"git.jojoxd.nl/projects/go-giorno/localizer"
	"git.jojoxd.nl/projects/go-giorno/localizer/locale"
)

type localizerImpl struct {
	*config
	locale   locale.Locale
	inner    *i18n.Localizer
	fallback *i18n.Localizer
}

func newLocalizer(
	locale locale.Locale,
	inner *i18n.Localizer,
	fallback *i18n.Localizer,
	config *config,
) localizer.Localizer {
	return &localizerImpl{
		locale:   locale,
		inner:    inner,
		fallback: fallback,
		config:   config,
	}
}

func (l localizerImpl) T(key string) string {
	return l.localizeWithFallback(key, &i18n.LocalizeConfig{
		MessageID: key,
	})
}

func (l localizerImpl) Tf(key string, templateData any) string {
	return l.localizeWithFallback(key, &i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: templateData,
	})
}

func (l localizerImpl) Tc(key string, pluralCount any) string {
	return l.localizeWithFallback(key, &i18n.LocalizeConfig{
		MessageID:   key,
		PluralCount: pluralCount,
	})
}

func (l localizerImpl) Tfc(key string, pluralCount any, templateData any) string {
	return l.localizeWithFallback(key, &i18n.LocalizeConfig{
		MessageID:    key,
		PluralCount:  pluralCount,
		TemplateData: templateData,
	})
}

func (l localizerImpl) Tl(localizable localizer.Localizable) string {
	if text, ok := localizable.Localize(l.locale); ok {
		return text
	}

	if stringer, ok := localizable.(fmt.Stringer); ok {
		return stringer.String()
	}

	l.logger.Warn("failed to localize localizable", "localizable", localizable)
	return fmt.Sprintf("Unlocalizable:%#v", localizable)
}

func (l localizerImpl) Inner() *i18n.Localizer {
	return l.inner
}

func (l localizerImpl) Locale() locale.Locale {
	return l.locale
}

func (l localizerImpl) localizeWithFallback(key string, lc *i18n.LocalizeConfig) string {
	text, err := l.inner.Localize(lc)
	if err == nil {
		return text
	}

	l.eventMgr.LocalizationNotFoundEvent(key, l.locale)

	if l.fallback != nil {
		text, err := l.fallback.Localize(lc)
		if err == nil {
			return text
		}
	}

	return key
}
