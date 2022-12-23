package serde

import (
	"time"
)

type AccessTokenPayload struct {
	Iss string   `json:"iss"`
	Sub string   `json:"sub"`
	Aud []string `json:"aud"`
	// Issued At
	Iat int64 `json:"iat"`
	// Expiration Time
	Exp   int64  `json:"exp"`
	Azp   string `json:"azp"`
	Scope string `json:"scope"`
	Gty   string `json:"gty"`
}

// maybe better will be to check whether we are about to expire soon (like few minutes before expiration time at least)
func (t *AccessTokenPayload) IsExpired() bool {
	return time.Now().UTC().Unix() > t.Exp
}
