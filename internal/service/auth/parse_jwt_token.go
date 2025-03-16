package auth

import (
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

var errInvalidToken = errors.New("token is invalid")
var errTokenExpired = errors.New("token has expired")

// returns username
func parseJwtToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(jwtSecretKeyEnvVariable)), nil
	})
	if err != nil {
		if strings.Contains(err.Error(), jwt.ErrTokenExpired.Error()) {
			return "", errTokenExpired
		}
		return "", errors.Wrap(err, "can't parse jwt token")
	}

	if claims, ok := token.Claims.(*tokenClaims); !ok || !token.Valid {
		return "", errInvalidToken
	} else {
		return claims.Username, nil
	}
}
