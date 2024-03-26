package jwt

import (
	"encoding/base64"
	"errors"
	"strings"
)

// Parse parses 3 parts of JWT (headers, claims and signature) and returns them
// as part of Token struct
func Parse(jwt string) (*Token[Claims], error) {
	return ParseCustomClaims(jwt, Claims{})
}

func ParseCustomClaims[C any](jwt string, c C) (*Token[C], error) {
	if jwt == "" {
		return nil, errors.New("invalid token")
	}
	var parts = strings.Split(jwt, ".")

	if len(parts) != 3 {
		return nil, errors.New("invalid token")
	}

	header := new(Header)
	err := decodeBase64JSON(parts[0], header)
	if err != nil {
		return nil, errors.New("invalid headers")
	}

	if header.Algorithm != "HS256" {
		return nil, errors.New("invalid alg")
	}
	if header.Type != "JWT" {
		return nil, errors.New("invalid typ")
	}

	claims := &c
	err = decodeBase64JSON(parts[1], claims)
	if err != nil {
		return nil, errors.New("invalid claims")
	}

	signature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, errors.New("invalid signature")
	}

	t := new(Token[C])
	t.Raw = jwt
	t.Header = *header
	t.Claims = *claims
	t.Signature = signature

	return t, nil
}
