package main

import (
	"fmt"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
	"strings"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

const dayDuration = time.Hour * 24

func main() {
	bundle := i18n.NewBundle(language.Russian)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	_, err := bundle.LoadMessageFile("translations/ru.yaml")
	if err != nil {
		panic(err)
	}

	localizer := NewLocalizer(i18n.NewLocalizer(bundle, "ru"))

	fmt.Println(
		FormatDuration(time.Hour*25+time.Minute*10, localizer),
	)
}

func FormatDuration(d time.Duration, localizer *Localizer) string {
	d = d.Round(time.Minute)
	days := d / dayDuration

	d -= days * dayDuration
	hours := d / time.Hour

	d -= hours * time.Hour
	minutes := d / time.Minute

	count := 0
	if days > 0 {
		count++
	}
	if hours > 0 {
		count++
	}
	if minutes > 0 {
		count++
	}
	parts := make([]string, count)

	if days > 0 {
		parts = append(parts, localizer.LocalizePlural("X days", int(days)))
	}

	if hours > 0 {
		parts = append(parts, localizer.LocalizePlural("X hours", int(hours)))
	}

	if minutes > 0 {
		parts = append(parts, localizer.LocalizePlural("X minutes", int(minutes)))
	}

	return strings.Join(parts, " ")
}

type (
	Localizer struct {
		internal *i18n.Localizer
	}
)

func NewLocalizer(localizer *i18n.Localizer) *Localizer {
	return &Localizer{
		internal: localizer,
	}
}

func (l *Localizer) Localize(messageId string) string {
	return l.internal.MustLocalize(&i18n.LocalizeConfig{MessageID: messageId})
}

func (l *Localizer) LocalizePlural(messageId string, pluralCount int) string {
	return l.internal.MustLocalize(&i18n.LocalizeConfig{MessageID: messageId, PluralCount: pluralCount})
}
