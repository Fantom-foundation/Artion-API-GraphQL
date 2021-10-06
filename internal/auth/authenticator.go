package auth

import (
	"artion-api-graphql/internal/config"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strings"
)

// challengeTemplate is template of message to be signed using Metamask
const challengeTemplate = "Click \"Sign\" to sign in into Artion.\n\nNonce:%s"

type Authenticator struct {
	BearerSecret []byte
	NonceSecret  []byte
}

func NewAuthenticator(cfg config.Auth) Authenticator {
	return Authenticator{
		BearerSecret: hexutil.MustDecode(cfg.BearerSecret),
		NonceSecret: hexutil.MustDecode(cfg.NonceSecret),
	}
}

// GenerateChallenge provides message to be signed using Metamask
func (a Authenticator) GenerateChallenge() (string, error) {
	nonce, err := generateNonce(a.NonceSecret)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(challengeTemplate, nonce), nil
}

// GenerateBearer verifies signed challenge and issues bearer token against it
func (a Authenticator) GenerateBearer(challenge string, address common.Address, signatureHex string) (string, error) {
	err := a.verifySignedChallenge(challenge, address, signatureHex)
	if err != nil {
		return "", err
	}

	return generateBearer(&address, a.BearerSecret)
}

// VerifyBearer verifies issued bearer token and returns proved user address
func (a Authenticator) VerifyBearer(bearer string) (*common.Address, error) {
	return verifyBearer(bearer, a.BearerSecret)
}

func (a Authenticator) verifySignedChallenge(challenge string, address common.Address, signatureHex string) error {
	err := a.verifyNonceInChallenge(challenge)
	if err != nil {
		return fmt.Errorf("nonce verification failed; %s", err)
	}
	signature, err := hexutil.Decode(signatureHex)
	if err != nil {
		return fmt.Errorf("signature hex decoding failed; %s", err)
	}
	ok, err := verifySignature(challenge, address, signature)
	if err != nil || !ok {
		return fmt.Errorf("signature verification failed; %s", err)
	}
	return nil
}

func (a Authenticator) verifyNonceInChallenge(challenge string) error {
	// needs to be in sync with challengeTemplate
	index := strings.LastIndex(challenge, "Nonce:") + 6
	if index == -1 || len(challenge) <= index {
		return fmt.Errorf("no nonce found in challenge [%s]", challenge)
	}
	nonce := challenge[index:]
	return verifyNonce(nonce, a.NonceSecret)
}
