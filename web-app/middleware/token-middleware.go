package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"web/lib"

	"github.com/dgrijalva/jwt-go"
)

func Token() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tokenString := lib.ExtractToken(r)

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(os.Getenv("API_SECRET")), nil
			})

			if err != nil {
				fmt.Fprintln(w, "Unauthorized")
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if ok && token.Valid {
				uId, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["userId"]), 10, 32)

				if err != nil {
					fmt.Fprintln(w, "Unauthorized")
					return
				}

				ctx := context.WithValue(r.Context(), lib.UserId{}, uId) // TODO: Add more user info.
				r = r.WithContext(ctx)

				f(w, r)
			} else {
				fmt.Fprintln(w, "Unauthorized")
			}
		}
	}
}
