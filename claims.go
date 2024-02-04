package gojwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// Claims for JWT token (RFC 7519)
// https://www.iana.org/assignments/jwt/jwt.xhtml
type Claims struct {
	Issuer     string   `json:"iss"` // Issuer (e.g. Auth service)
	Expiration int64    `json:"exp"` // Expiration timestamp
	Audience   string   `json:"aud"` // Audience - service for which JWT is intended
	Subject    string   `json:"sub"` // Subject - user identity
	Name       string   `json:"name"`
	Roles      []string `json:"roles"`
}

func (c *Claims) encode() []byte {
	JSON, err := json.Marshal(c)
	if err != nil {
		fmt.Println("failed to marshal claims")

		panic(err)
	}

	b := make([]byte, base64.RawURLEncoding.EncodedLen(len(JSON)))
	base64.RawURLEncoding.Encode(b, JSON)

	return b
}
