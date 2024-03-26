package jwt

// Claims for JWT token (RFC 7519)
// https://www.iana.org/assignments/jwt/jwt.xhtml
type Claims struct {
	Issuer     string   `json:"iss"`
	Expiration int64    `json:"exp"`
	Audience   string   `json:"aud"`
	Subject    string   `json:"sub"`
	Name       string   `json:"name"`
	Roles      []string `json:"roles"`
}
