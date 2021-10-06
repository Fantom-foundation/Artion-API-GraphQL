package auth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang-jwt/jwt"
	"time"
)

func generateBearer(user *common.Address, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.String(),
		"exp": time.Now().Add(12 * time.Hour).Unix(),
	})
	return token.SignedString(secret)
}

func verifyBearer(tokenString string, secret []byte) (*common.Address, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate the alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	// expiration is already checked by JWT parser
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userHex, ok := claims["sub"].(string); !ok {
			return nil, fmt.Errorf("jwt claim does not contain subject")
		} else {
			user := common.HexToAddress(userHex)
			return &user, nil
		}
	} else {
		return nil, fmt.Errorf("jwt validation failed; %s", err)
	}
}
