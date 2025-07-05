package giorno_i18n

import (
	"encoding/json"
	"fmt"
	fsp "io/fs"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gopkg.in/yaml.v3"

	"git.jojoxd.nl/projects/go-giorno/localizer/locale"
)

type Loader struct {
	bundle Bundle
}

func NewLoader() *Loader {
	return &Loader{
		bundle: make(Bundle),
	}
}

func (l *Loader) Load(fs fsp.FS, locales []locale.Locale) error {
	for _, localeKey := range locales {
		i18nBundle, err := l.load(fs, localeKey)
		if err != nil {
			return err
		}

		l.bundle[localeKey] = i18nBundle
	}

	return nil
}

func (l *Loader) load(fs fsp.FS, locale locale.Locale) (*i18n.Bundle, error) {
	bundle := i18n.NewBundle(locale.Tag())
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	glob := fmt.Sprintf("**/*.%s.yaml", strings.ToLower(locale.String()))
	matches, err := fsp.Glob(fs, glob)
	if err != nil {
		return nil, err
	}

	for _, match := range matches {
		fmt.Printf("loading localization %s\n", match)
		_, err := bundle.LoadMessageFileFS(fs, match)
		if err != nil {
			return nil, err
		}
	}

	return bundle, nil
}

func (l *Loader) Bundle() Bundle {
	return l.bundle
}
