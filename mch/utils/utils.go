package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/klaudiusz-czapla/my-cloud-home-go/mch/models"
	"github.com/mitchellh/mapstructure"
)

func DecodeToken(tokenString string) (*jwt.MapClaims, *models.IdTokenPayload, error) {
	claims := jwt.MapClaims{}
	token, parts, err := new(jwt.Parser).ParseUnverified(tokenString, &claims)

	if err != nil {
		return nil, nil, err
	}

	fmt.Print(token)
	fmt.Print(parts)

	// if !token.Valid {
	// 	return nil, fmt.Errorf("passed token is not valid")
	// }

	if err := claims.Valid(); err != nil {
		return nil, nil, fmt.Errorf("claims inside the token are not valid")
	}

	var idTokenPayload = models.IdTokenPayload{}
	mapstructure.Decode(claims, &idTokenPayload)

	return &claims, &idTokenPayload, nil
}
