package mch

import (
	"github.com/klaudiusz-czapla/my-cloud-home-go/mch/serde"
)

type MchSession struct {
	Config *serde.MchConfig
	Token  *serde.MchToken
	UserId string
}
