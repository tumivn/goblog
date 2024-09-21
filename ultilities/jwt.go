package ultilities

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateJwt(issuer string, jwtSecret string, expirationTime time.Time) (string, error) {
	claims := jwt.MapClaims{
		"issuer": issuer,
		"exp":    expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", errors.New("unable to create token")
	}
	return tokenString, nil
}

func GetIssuer(token string, secret string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !t.Valid {
		return "", err
	}
	issuer := t.Claims.(jwt.MapClaims)["issuer"].(string)

	return issuer, nil
}
