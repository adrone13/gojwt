package jwt

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
