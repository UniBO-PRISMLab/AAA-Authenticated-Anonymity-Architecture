package models

import "time"

type PACRequestModel struct {
	PID       string `json:"pid" example:"DmIq8x2JNs+E2qGBr16LxP6lUK+i/5nJRLuzvYEymtY=" binding:"required"`
	SignedPID string `json:"signed_pid" example:"Tn7dpsKMqbhC3j8/o7OH5juVrkejhhZYkDHNqVyp8Am8uRktEDKO5i09iH9GZO0NOQHRRKD6lpwD0wC5GEPiuHN6aAof3WtPO9bqEB8ZLup4FjnyDh3mMX1UlvlYjKi6eTLMD2J8dwObz0nkmZOVsjk51o8jMapJtfnzMkYhADh4vZVDLpWYWbtBsmmXhNHS4SuWc+K3ZKMLtCYE/MFK3JT+Zcyabmrd9jmqvHRgi9XT+kABZ6XnyUSp4VWo1M7pl767/hZM5CxqgXORUMk8z9M9lHCDOsLfjBOwi81ObMtd+4oVgXgAbzBBHXnicB6X5PEBb5Qyh0RNyWFFNOEvtA==" binding:"required"`
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
	SID       string `json:"sid" example:"jEA4/2q3Z4DcSllWmXtCbOaRrzyYAn3VXDHW5AN5U/8=" binding:"required"`
	SignedSID string `json:"signed_sid" example:"lOY7KkjuT3HV7R0/l66LR/0vR5URqy6i0cMH4nEHebks4ZSIHPO63FY+DCZX9QRu3VPt/eO21Fwm69pXk1lWbETll5vXbpChkKjaOAsuJfIRqWBX+bXHfiBCPvdjGr0f5YfWDx8xL1ndbXntNYF1WBt02Z/2JG06XsPsYwrT1FFDw6GiJiliJrBBJ7CYj4GhgdrgD70ZPdadKl7ChAgp9erc9s1nZ3Kdhnp9xmUR+aHcxi0RfikCvhgZvVvbHneSviI8tjp8zxT4VJhX2gbIJ/0fp4CVIaiHxHsMzCPB3rFdNNo5Pqr4TyDG/IP5k3QLD1yQXY+a9wlw6MGUBEiXrw==" binding:"required"`
}

type SACResponseModel struct {
	SAC        string    `json:"sac" example:"7MWorstwntw="`
	Expiration time.Time `json:"expiration" example:"2024-12-31T23:59:59Z"`
}

type SACVerificationRequestModel struct {
	SAC       string `json:"sac" example:"7MWorstwntw=" binding:"required"`
	PublicKey string `json:"public_key" example:"LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF1aVRTYzZMY3NmaFRZbEY0WDk2OApEcStXVzNrRmx3U21UR2NJQXJ5eFZ2OVhBQ2VFSXJuL0JYZE9McTFNSU52enc5WTZvNmhTd3hqUkp2SzlvdFpMCnlXYzFvc2RUNGdlTDJaT2xEMHUwakFvbEpqVU9hVWVyMmJKQTUzQzhXMjd1MHJyQzVvc3Z3THRNVWYxWnpTR0MKL2hrY0VOcjNjRGl6YWJibkw2dzVFUEdCVjlNU2RnaUdlelJXaXVEeGIzcldEOFNwSlpNa2c5b3ZSK3dOMlI2YwpOK2RwTkNjdDdWYjQxaWhXU0VXT1o2djhFdTREV3FjVW0vQzBpLzZ3eVJkWU93Y0E3NktkMmkxcnNZdHFlQS84CnZuMklTTEp2eGdkM3ZCdXY1ZnFkS0kzRm0rd3NIVDhHdHlkQUw5VjF4V3g4bnFlcFlwMlIvaXJMNHFUeGFnVzkKMXdJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg==" binding:"required"`
	SignedSAC string `json:"signed_sac" example:"Tn7dpsKMqbhC3j8/o7OH5juVrkejhhZYkDHNqVyp8Am8uRktEDKO5i09iH9GZO0NOQHRRKD6lpwD0wC5GEPiuHN6aAof3WtPO9bqEB8ZLup4FjnyDh3mMX1UlvlYjKi6eTLMD2J8dwObz0nkmZOVsjk51o8jMapJtfnzMkYhADh4vZVDLpWYWbtBsmmXhNHS4SuWc+K3ZKMLtCYE/MFK3JT+Zcyabmrd9jmqvHRgi9XT+kABZ6XnyUSp4VWo1M7pl767/hZM5CxqgXORUMk8z9M9lHCDOsLfjBOwi81ObMtd+4oVgXgAbzBBHXnicB6X5PEBb5Qyh0RNyWFFNOEvtA==" binding:"required"`
}

type SACVerificationResponseModel struct {
	Valid bool `json:"valid" example:"true"`
}
