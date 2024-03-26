package main

import (
	"fmt"
	"github.com/adrone13/gojwt"
	"log"
	"time"
)

var secret = "your_jwt_secret"

type CustomClaims struct {
	Expiration int64  `json:"iss"`
	Audience   string `json:"aud"`
	LastName   string `json:"last_name"`
}

func main() {
	c := jwt.Claims{
		Issuer:     "auth",
		Expiration: time.Now().Add(time.Second * time.Duration(120)).Unix(), // 120 sec ttl
		Audience:   "app",
		Subject:    "user_id",
		Name:       "User Name",
		Roles:      []string{"user"},
	}

	t1 := jwt.Sign(c, secret)
	fmt.Println(t1) // eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhdXRoIiwiZXhwIjoxNzA3OTIxNjUyLCJhdWQiOiJhcHAiLCJzdWIiOiJ1c2VyX2lkIiwibmFtZSI6IlVzZXIgTmFtZSIsInJvbGVzIjpbInVzZXIiXX0.Y1psCCMOA1AZLopIkJ8aW2cfr1YDPF27iYw7poeVfR4

	parsedToken1, err := jwt.Parse(t1)
	if err != nil {
		log.Fatalln("Failed to parse JWT. Error:", err)
	}

	fmt.Println(parsedToken1.IsValid(secret)) // true

	cc := CustomClaims{
		Expiration: time.Now().Add(time.Second * time.Duration(120)).Unix(),
		Audience:   "user_id",
		LastName:   "User Last Name",
	}
	fmt.Printf("Custom claims: %+v\n", cc)

	t2 := jwt.Sign(cc, secret)

	fmt.Println("JWT with custom claims:", t2) // eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOjE3MTE0NzU1MzEsImF1ZCI6InVzZXJfdXVpZCIsInNlc3Npb25faWQiOiJzZXNzaW9uX3V1aWQifQ.bUCazfzGc3ASXNexE-hK1CbKc1bpXNtJZn8jvErLMe0

	parsedToken2, err := jwt.ParseCustomClaims(t2, CustomClaims{})
	if err != nil {
		log.Fatalln("Failed to parse JWT with custom claims. Error:", err)
	}

	fmt.Println(parsedToken2.IsValid(secret)) // true

	fmt.Printf("Token: %+v\n", parsedToken2)
}
