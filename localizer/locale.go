package localizer

import "golang.org/x/text/language"

type Locale language.Tag

func (l Locale) String() string {
	return language.Tag(l).String()
}

func (l Locale) Tag() language.Tag {
	return language.Tag(l)
}
