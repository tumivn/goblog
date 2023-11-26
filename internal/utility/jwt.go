package utility

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/legangs/cms/internal/domain/auth/dtos"
	"time"
)

func GenerateJwt(issuer string, jwtSecret string, expirationTime time.Time) (string, error) {
	// Create the JWT claims, which includes the username and expiry time
	claims := &dtos.Claims{
		Issuer: issuer,
		Claims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", errors.New("unable to create token")
	}
	return tokenString, nil
}
