package middleware

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"web/lib"
	"web/model"

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

				http.SetCookie(w, &http.Cookie{Name: "errorMsg", Value: "Unauthorized", Expires: time.Now().Add(time.Second)})
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

				userCookie, err := r.Cookie("user")
				if err != nil {
					http.SetCookie(w, &http.Cookie{Name: "errorMsg", Value: "Unauthorized", Expires: time.Now().Add(time.Second)})
					http.Redirect(w, r, "/error", http.StatusSeeOther)
					return
				}
				var user model.User

				decodedUser, err := base64.StdEncoding.DecodeString(userCookie.Value)

				if err != nil {
					http.Redirect(w, r, "/error", http.StatusSeeOther)
					return
				}

				err = json.Unmarshal(decodedUser, &user)

				if err != nil {
					http.Redirect(w, r, "/error", http.StatusSeeOther)
					return
				}

				ctx = context.WithValue(r.Context(), lib.User{}, &user)
				r = r.WithContext(ctx)

				f(w, r)
			} else {
				if allow {
					f(w, r)
					return
				}

				http.SetCookie(w, &http.Cookie{Name: "errorMsg", Value: "Unauthorized", Expires: time.Now().Add(time.Second)})
				http.Redirect(w, r, "/error", http.StatusSeeOther)
				return
			}
		}
	}
}
