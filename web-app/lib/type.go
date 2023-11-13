package lib

import jwt "github.com/dgrijalva/jwt-go"

type Error struct{}
type Warning struct{}
type Success struct{}
type UserId struct{}
type ActivationToken struct{}

type CustomClaims struct {
	UserId          int  `json:"userId"`
	ActivationToken bool `json:"activationToken"`
	jwt.StandardClaims
}

// type userIdT string
// const userId userIdT = "userId"
// const userId userIdT = userIdT("userId")
// const userId = userIdT("userId")
