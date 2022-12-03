package mch

import "github.com/klaudiusz-czapla/my-cloud-home-go/mch"

const (
	clientID     = "9B0Gi617tROKHc2rS95sT1yJzR6MkQDm"
	clientSecret = "oSJOB1KOWnLVZm11DVknu2wZkTj5AGKxcINEDtEUPE30jHKvEqorM8ocWbyo17Hd"
)

type Proxy struct {
}

func Test(username string, password string) {
	config, err := mch.GetConfiguration()
	if err != nil {
		return nil, err
	}

	req := map[string]string{
		"grant_type":    "http://auth0.com/oauth/grant-type/password-realm",
		"realm":         "Username-Password-Authentication",
		"audience":      "mycloud.com",
		"username":      username,
		"password":      password,
		"scope":         "openid offline_access nas_read_write nas_read_only user_read device_read",
		"client_id":     clientID,
		"client_secret": clientSecret,
	}
}
