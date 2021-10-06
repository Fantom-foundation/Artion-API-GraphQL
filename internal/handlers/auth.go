package handlers

import (
	"artion-api-graphql/internal/auth"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"net/http"
	"strings"
)

// AuthHandler creates a handler for chained HTTP requests resolving for validating
// the request authenticity; the request context is extended with the User information
// if the request is validated.
func AuthHandler(h http.Handler) http.Handler {
	// validate JWT token, if it's present
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := authGetIdentity(r)
		if err != nil {
			fmt.Errorf("unauthorized API request; %s", err)
		}
		if id != nil {
			// set the user identity (account address) into the request context
			h.ServeHTTP(w, r.WithContext(auth.SetIdentity(r.Context(), id)))
		} else {
			// pass the processing
			h.ServeHTTP(w, r)
		}
	})
}

// authGetIdentity extract authorized user identity from the request header.
func authGetIdentity(r *http.Request) (*common.Address, error) {
	tokenString, err := extractTokenString(r)
	if err != nil {
		return nil, err
	}
	return auth.GetAuthenticator().VerifyBearer(tokenString)
}

// extractTokenString extracts an auth token from the request, if available
func extractTokenString(r *http.Request) (string, error) {
	authr := r.Header.Get("Authorization")
	if authr == "" {
		return authr, fmt.Errorf("authorization header missing")
	}
	if len(authr) < 7 {
		return authr, fmt.Errorf("invalid authorization header length")
	}
	if !strings.EqualFold("Bearer", authr[0:6]) {
		return authr, fmt.Errorf("invalid format of authorization header")
	}
	return authr[7:], nil
}
