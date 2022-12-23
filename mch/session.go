package mch

import (
	"github.com/klaudiusz-czapla/my-cloud-home-go/mch/serde"
)

type MchSession struct {
	Config *serde.MchConfig
	Token  *serde.MchToken
	// some of the values coming from token will be denormalized and stored here to speed up the access..
	UserId string
}
