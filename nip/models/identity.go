package models

type PIDRequestModel struct {
	PublicKey string `json:"public_key" example:"-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQ...\n-----END PUBLIC KEY-----"`
}

type PIDResponseModel struct {
	PID     string `json:"pid" example:"a1b2c3d4-e5f6-7890-1234-567890abcdef"`
	Message string `json:"message" example:"PID successfully issued and saved"`
}
