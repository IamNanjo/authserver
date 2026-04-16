package middleware

import (
	"fmt"
	"net/http"
	"slices"
	"strings"
)

const defaultLocale = "en-FI"

var validLanguages = []string{"en", "fi"}
var validLocales = []string{"FI", "GB", "US"}
var defaultLanguageLocale = strings.SplitN(defaultLocale, "-", 2)

var defaultLocaleCookie = http.Cookie{
	Name:     "locale",
	Value:    defaultLocale,
	Path:     "/",
	HttpOnly: true,
	SameSite: http.SameSiteStrictMode,
}

// Updates locale cookie
func Locale(w http.ResponseWriter, r *http.Request) bool {
	localeCookie, err := r.Cookie("locale")
	if err != nil || !localeCookie.HttpOnly {
		updatedCookie := defaultLocaleCookie
		updatedCookie.Value = parseLanguageHeader(r)
		http.SetCookie(w, &updatedCookie)
		return true
	}

	localeCookie.Value = defaultLocale
	http.SetCookie(w, localeCookie)

	return true
}

// Selects first valid language and locale from Accept-Language header.
// Returns defaultLocale if no value is set.
func parseLanguageHeader(r *http.Request) string {
	languageHeader := r.Header.Get("Accept-Language")
	if languageHeader == "" {
		return defaultLocale
	}

	languages := strings.Split(languageHeader, ",")
	for i := range languages {
		languages[i] = strings.TrimSpace(languages[i])
	}

	language := ""
	locale := ""

	for i := range languages {
		languageLocaleWeight := strings.Split(languages[i], ";")
		if len(languageLocaleWeight) < 2 {
			continue
		}
		if language == "" && slices.Contains(validLanguages, languageLocaleWeight[0]) {
			language = languageLocaleWeight[0]
		}
		if locale == "" && slices.Contains(validLocales, languageLocaleWeight[1]) {
			locale = languageLocaleWeight[1]
		}
	}

	if language == "" {
		language = defaultLanguageLocale[0]
	}
	if locale == "" {
		locale = defaultLanguageLocale[1]
	}

	return fmt.Sprintf("%s-%s", language, locale)
}
