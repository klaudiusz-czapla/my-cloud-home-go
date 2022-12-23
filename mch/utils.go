package mch

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/klaudiusz-czapla/my-cloud-home-go/mch/serde"
	"github.com/klaudiusz-czapla/my-cloud-home-go/utils"
	"github.com/mitchellh/mapstructure"
)

func DecodeIdToken(tokenString string) (*jwt.MapClaims, *serde.IdTokenPayload, error) {
	claims := jwt.MapClaims{}
	_, _, err := new(jwt.Parser).ParseUnverified(tokenString, &claims)

	if err != nil {
		return nil, nil, err
	}

	// if !token.Valid {
	// 	return nil, fmt.Errorf("passed token is not valid")
	// }

	if err := claims.Valid(); err != nil {
		return nil, nil, fmt.Errorf("claims inside the token are not valid")
	}

	var idTokenPayload = serde.IdTokenPayload{}
	mapstructure.Decode(claims, &idTokenPayload)

	json, _ := utils.ToJson(&idTokenPayload)
	fmt.Print(json)

	return &claims, &idTokenPayload, nil
}

func DecodeAccessToken(tokenString string) (*jwt.MapClaims, *serde.AccessTokenPayload, error) {
	claims := jwt.MapClaims{}
	_, _, err := new(jwt.Parser).ParseUnverified(tokenString, &claims)

	if err != nil {
		return nil, nil, err
	}

	// if !token.Valid {
	// 	return nil, fmt.Errorf("passed token is not valid")
	// }

	if err := claims.Valid(); err != nil {
		return nil, nil, fmt.Errorf("claims inside the token are not valid")
	}

	var accTokenPayload = serde.AccessTokenPayload{}
	mapstructure.Decode(claims, &accTokenPayload)

	json, _ := utils.ToJson(&accTokenPayload)
	fmt.Print(json)

	return &claims, &accTokenPayload, nil
}
