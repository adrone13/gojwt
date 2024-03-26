package jwt

import (
	"crypto/hmac"
	"fmt"
	"strings"
)

type Token[C any] struct {
	Raw       string
	Header    Header
	Claims    C
	Signature []byte
}

// IsValid validates Token.Signature against computed one
func (t *Token[C]) IsValid(secret string) bool {
	var parts []string = strings.Split(t.Raw, ".")

	expected := computeSignature(fmt.Sprintf("%s.%s", parts[0], parts[1]), secret)

	return hmac.Equal(t.Signature, expected)
}
