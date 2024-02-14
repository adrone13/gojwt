package main

import (
	"fmt"
	"time"

	"github.com/adrone13/gojwt"
)

var secret = "your_jwt_secret"

func main() {
	e := time.Now().Add(time.Second * time.Duration(120)).Unix() // 120 sec ttl
	c := jwt.Claims{
		Issuer:     "auth",
		Expiration: e,
		Audience:   "app",
		Subject:    "user_id",
		Name:       "User Name",
		Roles:      []string{"user"},
	}

	token := jwt.Sign(c, secret)
	fmt.Println(token) // eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhdXRoIiwiZXhwIjoxNzA3OTIxNjUyLCJhdWQiOiJhcHAiLCJzdWIiOiJ1c2VyX2lkIiwibmFtZSI6IlVzZXIgTmFtZSIsInJvbGVzIjpbInVzZXIiXX0.Y1psCCMOA1AZLopIkJ8aW2cfr1YDPF27iYw7poeVfR4

	parsed, err := jwt.Parse(token)
	if err != nil {
		panic("Invalid JWT")
	}

	fmt.Println(parsed.IsValid(secret)) // true
}
