package models

import "time"

type PACRequestModel struct {
	PID       string `json:"pid" example:"DmIq8x2JNs+E2qGBr16LxP6lUK+i/5nJRLuzvYEymtY=" binding:"required"`
	SignedPID string `json:"signed_pid" example:"signed_pid_payload_base64" binding:"required"`
}

type PACVerificationRequestModel struct {
	PID string `json:"pid" example:"DmIq8x2JNs+E2qGBr16LxP6lUK+i/5nJRLuzvYEymtY=" binding:"required"`
	PAC int64  `json:"pac" example:"874532" binding:"required"`
}

type PACResponseModel struct {
	PAC        int64     `json:"pac" example:"874532"`
	Expiration time.Time `json:"expiration" example:"2024-12-31T23:59:59Z"`
}

type PACVerificationResponseModel struct {
	Valid      bool      `json:"valid" example:"true"`
	Expiration time.Time `json:"expiration" example:"2024-12-31T23:59:59Z"`
}

type SACRequestModel struct {
	SignedSID string `json:"signed_sid" example:"signed_sid_payload_base64"`
}

type SACResponseModel struct {
	SAC        int64     `json:"sac" example:"874532"`
	Expiration time.Time `json:"expiration" example:"2024-12-31T23:59:59Z"`
}
