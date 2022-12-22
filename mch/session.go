package mch

import "github.com/klaudiusz-czapla/my-cloud-home-go/mch/models"

type MchSession struct {
	Config *models.MchConfig
	Token  *models.MchToken
	UserId string
}
