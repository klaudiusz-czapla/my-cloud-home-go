package serde

import (
	"time"
)

type AccessTokenPayload struct {
	Exp int64
}

// maybe better will be to check whether we are about to expire soon (like few minutes before expiration time at least)
func (t *AccessTokenPayload) IsExpired() bool {
	return time.Now().UTC().Unix() > t.Exp
}
