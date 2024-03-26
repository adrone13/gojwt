# Go JWT package

## Usage
```go
package main

import (
	"fmt"
	"time"

	"github.com/adrone13/gojwt"
)

var secret = "your_jwt_secret"

func main() {
	exp, _ := time.Parse(time.RFC3339, "2024-01-01T12:00:00.000Z")

	// Registered claims
	c := jwt.Claims{
		Issuer:     "auth",
		Expiration: exp.Unix(), // 120 sec ttl
		Audience:   "app",
		Subject:    "user_id",
		Name:       "User Name",
		Roles:      []string{"user"},
	}

	t1 := jwt.Sign(c, secret)
	fmt.Println(t1) // eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhdXRoIiwiZXhwIjoxNzA0MTEwNDAwLCJhdWQiOiJhcHAiLCJzdWIiOiJ1c2VyX2lkIiwibmFtZSI6IlVzZXIgTmFtZSIsInJvbGVzIjpbInVzZXIiXX0.wdpRLHtuIT6-Rm3p84UBMdD2j1DIzIvxtYFz_Ud9zGU

	parsedToken1, err := jwt.Parse(t1)
	if err != nil {
		log.Fatalln("Failed to parse JWT. Error:", err)
	}

	fmt.Println(parsedToken1.IsValid(secret)) // true

	// Custom claims struct
	cc := CustomClaims{
		Expiration: exp.Unix(),
		Subject:    "user_id",
		LastName:   "User Last Name",
	}
	fmt.Printf("Custom claims: %+v\n", cc)

	t2 := jwt.Sign(cc, secret)

	fmt.Println("JWT with custom claims:", t2) // eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOjE3MDQxMTA0MDAsImF1ZCI6InVzZXJfaWQiLCJsYXN0X25hbWUiOiJVc2VyIExhc3QgTmFtZSJ9.UqfgRY36pxe42WIQW8JCqa4eUrAu9iAn5ckP8pps3Bo

	parsedToken2, err := jwt.ParseCustomClaims(t2, CustomClaims{})
	if err != nil {
		log.Fatalln("Failed to parse JWT with custom claims. Error:", err)
	}

	fmt.Println(parsedToken2.IsValid(secret)) // true
}

```

