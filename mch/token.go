package mch

type MchToken struct {
	IdToken      string `json:"id_token"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int32  `json:"expires_in"`
}
