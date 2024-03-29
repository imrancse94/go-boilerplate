package localize

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go-boilerplate/pkg/cache"
	"golang.org/x/text/language"
	"strings"
)

func SetLocale(lang string) {
	cache.Set("locale", lang, 0)
}

func GetLocale() string {
	return cache.Get("locale")
}

func initLocalizer(langs ...string) *i18n.Localizer {
	// Create a new i18n bundle with English as default language.
	bundle := i18n.NewBundle(language.English)

	// Register a json unmarshal function for i18n bundle.
	// This is to enable usage of json format
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// Load source language
	bundle.LoadMessageFile("./locales/en.json")
	bundle.LoadMessageFile("./locales/es.json")
	bundle.LoadMessageFile("./locales/bn.json")

	// Initialize localizer which will look for phrase keys in passed languages
	// in a strict order (first language is searched first)
	// When no key in any of the languages is found, it fallbacks to default - English language
	if langs == nil {
		cache.Set("locale", "en", 0)
	}
	localizer := i18n.NewLocalizer(bundle, langs...)

	return localizer
}

func Trans(key string, vars string) string {
	localize := initLocalizer(GetLocale())
	var x map[string]interface{}

	json.Unmarshal([]byte(vars), &x)
	for k, v := range x {
		var str = strings.Split(v.(string), "")
		var finalString = ""
		for _, value := range str {
			currentString, _ := localize.Localize(&i18n.LocalizeConfig{
				MessageID: value,
			})

			if currentString == "" {
				currentString = value
			}
			finalString += currentString
		}

		x[k] = finalString
	}
	simpleMessage, _ := localize.Localize(&i18n.LocalizeConfig{
		MessageID:    key, // source key identifier
		TemplateData: x,
	})

	if simpleMessage == "" {
		simpleMessage = key
		if x != nil {
			for k, v := range x {
				simpleMessage = strings.Replace(simpleMessage, fmt.Sprintf("{{.%s}}", k), v.(string), -1)
			}
		}

		fmt.Println("x", x)
	}
	return simpleMessage
}

func SetPlural(key string, pluralCount int) string {
	localize := initLocalizer(GetLocale())

	pluralMessage, _ := localize.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{ID: key}, // another source key identifier
		PluralCount:    pluralCount,            // would use "one" variant if the count was 1
	})

	fmt.Println(pluralMessage)
	return pluralMessage
}
