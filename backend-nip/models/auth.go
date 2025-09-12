package models

import "time"

type PACRequestModel struct {
	PID       string `json:"pid" example:"DmIq8x2JNs+E2qGBr16LxP6lUK+i/5nJRLuzvYEymtY=" binding:"required"`
	SignedPID string `json:"signed_pid" example:"fzLD1/PAJoq7EtWqIXLOYMhxNJCK9L4w4yniXZSxu0r3+oZ84jreOAhAIlrodSdg0/4EaKYWLmyf0D0uD0eozzwhvVYPDZIrLSxllH+HeEWWuOm1hnPVOOUwo7G4ROemojqwNQ6uvSB+AxoAPiSfm01Rgi0BPoaH9X8H6r94O9Hm3ZrgA9zoxQhMTx/eDpsW2yNbfPFXd18IAuHqLJSCBOSl3U4PfHEad3xFw6H04s2hr3Rqlsb2WsJmS6rf0mPEPAKjq1fSvBIcxg6PCJdfw/GhZEByrP0qc5DllkYSYZaddM/pAh+8LA16YmN1eO9VthsQxGt/WHma/5xsChOaCMd+fgk+Yjio56LoHrC86SqKsS4gkfa2jZoyagwNCah/nKAQkeVzQjk5+WLdCfMh71v82lw9L3rwlng2KhLC4z19k7wheoKQipJaxFR72T7PDftt1TvttK5q6F8x+sV6cTC6roRkgZG03X4tgLJrkBLjF9w3KbcKhPzLyu9ehYtTqFe9hBuR11UXirTi96bvmT9wWDVoTxv+ZGvn59e0VLBcQp67W2m3CZgmIs5eksQPhsFufRwYq49kXxFqnX5OZdqmVVhbQM8X9b8dTpgxd+tIuPzrYcG3ts5ApPCdXMoM/s1Tp6x+QQtq6uOD44V5v5E9tHd9GAH5cYKDfPvdA/w=" binding:"required"`
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
