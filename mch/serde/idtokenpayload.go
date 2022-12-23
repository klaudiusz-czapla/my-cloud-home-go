package serde

type IdTokenPayload struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Iss      string `json:"iss"`
	Sub      string `json:"sub"`
	Aud      string `json:"aud"`
	// Issued At
	Iat int `json:"iat"`
	// Expiration Time
	Exp int `json:"exp"`
}
