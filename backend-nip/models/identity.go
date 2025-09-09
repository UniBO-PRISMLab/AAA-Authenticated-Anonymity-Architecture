package models

type PIDRequestModel struct {
	PublicKey string `json:"public_key" example:"abc123def456ghci78" binding:"required"`
}

type PIDResponseModel struct {
	PID     string `json:"pid" example:"a1b2c3d4-e5f6-7890-1234-567890abcdef"`
	Message string `json:"message" example:"PID successfully issued and saved"`
}

type User struct {
	PID       string
	PublicKey string
	Nonce     string
}
