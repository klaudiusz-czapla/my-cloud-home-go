package serde

type IdTokenPayload struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Iss      string `json:"iss"`
	Sub      string `json:"sub"`
	Aud      string `json:"aud"`
	Iat      int    `json:"iat"`
	Exp      int    `json:"exp"`
}
