package localizer

import (
	"git.jojoxd.nl/projects/go-giorno/localizer/locale"
)

type Localizable interface {
	Localize(locale locale.Locale) (string, bool)
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

	// Locale gives the locale that this Localizer is for
	Locale() locale.Locale
}
