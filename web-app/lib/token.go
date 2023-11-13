package lib

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId int, activationToken bool) (string, error) {
	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := CustomClaims{
		userId,
		activationToken,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix(),
			Issuer:    "Laby",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ExtractToken(r *http.Request) string {
	query := r.URL.Query()
	queryToken := query["token"]

	if len(queryToken) != 0 && queryToken[0] != "" {
		return queryToken[0]
	}

	bearerToken := r.Header.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	cookieToken, err := r.Cookie("token")

	if err == nil {
		return cookieToken.Value
	}

	return ""
}
