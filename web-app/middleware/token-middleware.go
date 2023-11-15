package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
	"web/lib"

	"github.com/dgrijalva/jwt-go"
)

func Token(allow bool) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tokenString := lib.ExtractToken(r)

			token, err := jwt.ParseWithClaims(tokenString, &lib.CustomClaims{}, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if err != nil {
				if allow {
					f(w, r)
					return
				}

				http.SetCookie(w, &http.Cookie{Name: "msg", Value: "Unauthorized", Expires: time.Now().Add(time.Second)})
				http.Redirect(w, r, "/error", http.StatusSeeOther)
				return
			}

			claims, ok := token.Claims.(*lib.CustomClaims)

			if ok && token.Valid {
				uId := claims.UserId
				activationToken := claims.ActivationToken

				ctx := context.WithValue(r.Context(), lib.UserId{}, uId)
				ctx = context.WithValue(ctx, lib.ActivationToken{}, activationToken)
				r = r.WithContext(ctx)

				f(w, r)
			} else {
				if allow {
					f(w, r)
					return
				}

				http.SetCookie(w, &http.Cookie{Name: "msg", Value: "Unauthorized", Expires: time.Now().Add(time.Second)})
				http.Redirect(w, r, "/error", http.StatusSeeOther)
				return
			}
		}
	}
}
