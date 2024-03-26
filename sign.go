package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func Sign[C any](c C, secret string) string {
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	payload := encodeClaims(c)

	message := fmt.Sprintf("%s.%s", header, payload)

	signature := computeSignature(message, secret)

	return fmt.Sprintf("%s.%s", message, base64.RawURLEncoding.EncodeToString(signature))
}

func encodeClaims[C any](c C) []byte {
	JSON, err := json.Marshal(c)
	if err != nil {
		fmt.Println("failed to marshal claims")

		panic(err)
	}

	b := make([]byte, base64.RawURLEncoding.EncodedLen(len(JSON)))
	base64.RawURLEncoding.Encode(b, JSON)

	return b
}
