package middleware

import (
	"MyMiniProject/constant"
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateToken(authID, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authID"] = authID
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.SECRET_JWT))
}
