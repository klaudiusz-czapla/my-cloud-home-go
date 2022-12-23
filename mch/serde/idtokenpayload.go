package serde

type IdTokenPayload struct {
	Iss string `json:"iss"`
	Sub string `json:"sub"`
}
