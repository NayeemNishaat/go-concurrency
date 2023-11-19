package lib

import (
	"net/http"
	"os"
	"strconv"
	"time"
)

func GenerateCookie(name string, value string, expires time.Time) (*http.Cookie, error) {
	if expires.IsZero() {
		expires = time.Now().AddDate(0, 0, 1)
	}

	secure, err := strconv.ParseBool(os.Getenv("SECURE_COOKIE"))

	if err != nil {
		return &http.Cookie{}, err
	}

	maxAge, err := strconv.ParseUint(os.Getenv("COOKIE_MAX_AGE"), 10, 32)

	if err != nil {
		return &http.Cookie{}, err
	}

	return &http.Cookie{
		Name:     name,
		Domain:   os.Getenv("COOKIE_DOMAIN"),
		Path:     "/",
		Secure:   secure,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Value:    value,
		MaxAge:   int(maxAge),
		Expires:  expires,
	}, nil
}
