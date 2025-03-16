package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

const jwtSecretKeyEnvVariable = "JWT_SECRET_KEY"

const accessTokenDuration = time.Minute * 15
const refreshTokenDuration = time.Hour * 24 * 21

type tokenClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func generateJwtToken(username string, tokenDuration time.Duration) (string, error) {
	currentTime := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(currentTime),
			NotBefore: jwt.NewNumericDate(currentTime),
		},
	})

	secretKey := os.Getenv(jwtSecretKeyEnvVariable)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.Wrap(err, "can't sign jwt token")
	}

	return tokenString, nil
}
