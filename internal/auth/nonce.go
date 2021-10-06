package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func generateNonce(secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(10 * time.Minute).Unix(),
	})
	return token.SignedString(secret)
}

func verifyNonce(tokenString string, secret []byte) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate the alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected nonce signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	// expiration is already checked by JWT parser
	if token.Valid {
		return nil
	} else {
		return fmt.Errorf("nonce validation failed; %s", err)
	}
}
