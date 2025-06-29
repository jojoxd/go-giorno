package localizer

import "github.com/nicksnyder/go-i18n/v2/i18n"

type Localizable interface {
	Localize(locale Locale) (string, bool)
}

type Localizer interface {
	// T translates a key
	T(key string) string

	// Tf translates a key with formatting (template data)
	Tf(key string, templateData any) string

	// Tc translates a key with a pluralCount
	Tc(key string, pluralCount any) string

	// Tfc translates a key with a pluralCount and formatting (template data)
	Tfc(key string, pluralCount any, templateData any) string

	// Tl localizes a Localizable
	Tl(localizable Localizable) string

	// Inner gives access to the inner i18n.Localizer
	Inner() *i18n.Localizer

	// Locale gives the locale that this Localizer is for
	Locale() Locale
}
