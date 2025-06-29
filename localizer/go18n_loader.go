package localizer

import (
	"encoding/json"
	"fmt"
	fsp "io/fs"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gopkg.in/yaml.v3"
)

type GoI18nBundleLoader struct {
	bundle GoI18nBundle
}

func NewGoI18nBundleLoader() *GoI18nBundleLoader {
	return &GoI18nBundleLoader{
		bundle: make(GoI18nBundle),
	}
}

func (l *GoI18nBundleLoader) Load(fs fsp.FS, locales []Locale) error {
	for _, locale := range locales {
		i18nBundle, err := l.load(fs, locale)
		if err != nil {
			return err
		}

		l.bundle[locale] = i18nBundle
	}

	return nil
}

func (l *GoI18nBundleLoader) load(fs fsp.FS, locale Locale) (*i18n.Bundle, error) {
	bundle := i18n.NewBundle(locale.Tag())
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	glob := fmt.Sprintf("**/*.%s.yaml", strings.ToLower(locale.String()))
	fmt.Printf(glob)
	matches, err := fsp.Glob(fs, glob)
	if err != nil {
		return nil, err
	}

	for _, match := range matches {
		fmt.Printf("loading localization %s", match)
		_, err := bundle.LoadMessageFileFS(fs, match)
		if err != nil {
			return nil, err
		}
	}

	return bundle, nil
}

func (l *GoI18nBundleLoader) Bundle() GoI18nBundle {
	return l.bundle
}
