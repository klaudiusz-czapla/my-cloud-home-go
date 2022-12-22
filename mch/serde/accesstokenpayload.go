package serde

import (
	"time"
)

type AccessTokenPayload struct {
	Exp int64
}

func (t *AccessTokenPayload) IsExpired() bool {
	return time.Now().UTC().Unix() > t.Exp
}
