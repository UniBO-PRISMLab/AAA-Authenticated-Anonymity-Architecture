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
	SID       string `json:"sid" example:"MHhhOTcwZWRkMTEyNTRjNTQ3OTIyMWVjOTFjZGU0YjNiMGI2MjQ1NzI3NmVmMzA0MmMzN2M2MDJkMWRlMTg4Njc5" binding:"required"`
	SignedSID string `json:"signed_sid" example:"QEZZSiBAF1TPbSgHAlcwlrG4unRGZBGYOf8yjiSfP8wUcnkvdm+yQTOzGpqWG/m4/qCG7stlDYqZrQIyRWC7mb0sbTe2MhMYrZOzyz6uH8hCxIJh5JQ6ivsntufaeg9HOoZ9M6wdhp1/BsbwuGG27CQ7BcuuKZu2Mi1cNu32dnTGxv4PORoCy1JqgV519JsuwL3aekdZ2jVrYaepTTDT1vJR4vhLZ76gV1ywUAIEMm8q+R12ydntAroOd8EaiCXx0pyM3zVAE1G0Qv1g0qRGrnogmMhCIpljJ79KEFtOQDHLd0RnbPuoh99eszx99qc911VXTjiHMNdIxhanp31odA==" binding:"required"`
}

type SACResponseModel struct {
	SAC        int64     `json:"sac" example:"874532"`
	Expiration time.Time `json:"expiration" example:"2024-12-31T23:59:59Z"`
}
