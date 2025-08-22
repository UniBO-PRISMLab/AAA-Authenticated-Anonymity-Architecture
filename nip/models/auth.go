package models

type PACRequestModel struct {
	SignedPID string `json:"signed_pid" example:"signed_pid_payload_base64"`
}

type PACResponseModel struct {
	PAC string `json:"pac" example:"signed_sid_payload_base64"`
}

type SACRequestModel struct {
	SignedSID string `json:"signed_sid" example:"signed_sid_payload_base64"`
}

type SACResponseModel struct {
	SAC string `json:"sac" example:"abc123def456ghci78"`
}
