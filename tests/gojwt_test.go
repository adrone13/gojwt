package tests

import (
	"fmt"
	"github.com/adrone13/gojwt"
	"testing"
)

var (
	validToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhdXRoIiwiZXhwIjoxNzAyMTYxMzg4LCJhdWQiOiJ0b2RvIiwic3ViIjoidXVpZCIsIm5hbWUiOiJBbGV4IFRoZSBNYWQiLCJyb2xlcyI6WyJUT0RPIl19.DSIhbioL9esS0gsiliNl9rUFYaLaZAciVvNG7e7OxyI"
	secret     = "AixSmyAlU0Gh-Tvpw_ytFPLtc2GyVCPG9uxlsBDsmy4"
)

func TestParseResult(t *testing.T) {
	token, err := jwt.Parse(validToken)
	if err != nil {
		t.Errorf("Received error: %s", err)

		return
	}

	assert(t, "HS256", token.Header.Algorithm)
	assert(t, "JWT", token.Header.Type)
	assert(t, "auth", token.Claims.Issuer)
	assert(t, 1702161388, token.Claims.Expiration)
	assert(t, "todo", token.Claims.Audience)
	assert(t, "uuid", token.Claims.Subject)
	assert(t, "Alex The Mad", token.Claims.Name)
	assertCol(t, []string{"TODO"}, token.Claims.Roles)
}

func TestParseErrors(t *testing.T) {
	_, err := jwt.Parse("abc.def")
	if err == nil {
		t.Errorf(`Expected to return error: "invalid token"`)
	} else {
		assert(t, "invalid token", err.Error())
	}

	_, err = jwt.Parse("abc.def.ghi")
	if err == nil {
		t.Errorf(`Expected to return error: "invalid headers"`)
	} else {
		assert(t, "invalid headers", err.Error())
	}

	_, err = jwt.Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.def.ghi")
	if err == nil {
		t.Errorf(`Expected to return error: "invalid claims"`)
	} else {
		assert(t, "invalid claims", err.Error())
	}

	_, err = jwt.Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhdXRoIiwiZXhwIjoxNzAyMTYxMzg4LCJhdWQiOiJ0b2RvIiwic3ViIjoidXVpZCIsIm5hbWUiOiJBbGV4IFRoZSBNYWQiLCJyb2xlcyI6WyJUT0RPIl19.abc:")
	if err == nil {
		t.Errorf(`Expected to return error: "invalid signature"`)
	} else {
		assert(t, "invalid signature", err.Error())
	}

	_, err = jwt.Parse("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhdXRoIiwiZXhwIjoxNzAyMTYxMzg4LCJhdWQiOiIiLCJzdWIiOiIiLCJuYW1lIjoiIiwicm9sZXMiOm51bGx9.CX3vnWXs-hb0AYQWP0RqVjcHbSzTKWwgIfyCq1fwVPo")
	if err == nil {
		t.Errorf(`Expected to return error: "invalid alg"`)
	} else {
		assert(t, "invalid alg", err.Error())
	}

	_, err = jwt.Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IklOVkFMSURfVFlQIn0.eyJpc3MiOiJhdXRoIiwiZXhwIjoxNzAyMTYxMzg4LCJhdWQiOiIiLCJzdWIiOiIiLCJuYW1lIjoiIiwicm9sZXMiOm51bGx9.6BPLmsnhQETeKeb7BzpBkEj71JPp-3wPlHOSZw2m1Rg")
	if err == nil {
		t.Errorf(`Expected to return error: "invalid typ"`)

		return
	} else {
		assert(t, "invalid typ", err.Error())
	}
}

func TestIsValid(t *testing.T) {
	token, err := jwt.Parse(validToken)
	if err != nil {
		t.Errorf("Received error: %s", err)

		return
	}

	assert(t, token.IsValid(secret), true)
	assert(t, token.IsValid("invalid-secret"), false)
}

func TestSign(t *testing.T) {
	c := jwt.Claims{
		Issuer:     "auth",
		Expiration: 1702161388,
	}

	token := jwt.Sign(c, secret)

	assert(
		t,
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhdXRoIiwiZXhwIjoxNzAyMTYxMzg4LCJhdWQiOiIiLCJzdWIiOiIiLCJuYW1lIjoiIiwicm9sZXMiOm51bGx9.q2deoYFUCyNJFjeibJf7XDneL1QUNmTrkZfDq01qojE",
		token,
	)
}

func assert[T comparable](t *testing.T, expected, received T) {
	if expected != received {
		t.Errorf(fmt.Sprintf("\nExpected: \u001B[1;31m%+v\u001B[0m\nReceived: \u001B[1;32m%+v\u001B[0m\n", expected, received))
	}
}

func assertCol[T comparable](t *testing.T, expected, received []T) {
	if len(expected) != len(received) {
		t.Errorf(fmt.Sprintf("\nExpected: \u001B[1;31m%+v\u001B[0m\nReceived: \u001B[1;32m%+v\u001B[0m\n", expected, received))

		return
	}

	for i, v := range expected {
		if v != received[i] {
			t.Errorf(fmt.Sprintf("\nExpected: \u001B[1;31m%+v\u001B[0m\nReceived: \u001B[1;32m%+v\u001B[0m\n", expected, received))

			return
		}
	}
}
