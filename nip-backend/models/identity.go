package models

type PIDRequestModel struct {
	PublicKey string `json:"public_key" example:"LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF1aVRTYzZMY3NmaFRZbEY0WDk2OApEcStXVzNrRmx3U21UR2NJQXJ5eFZ2OVhBQ2VFSXJuL0JYZE9McTFNSU52enc5WTZvNmhTd3hqUkp2SzlvdFpMCnlXYzFvc2RUNGdlTDJaT2xEMHUwakFvbEpqVU9hVWVyMmJKQTUzQzhXMjd1MHJyQzVvc3Z3THRNVWYxWnpTR0MKL2hrY0VOcjNjRGl6YWJibkw2dzVFUEdCVjlNU2RnaUdlelJXaXVEeGIzcldEOFNwSlpNa2c5b3ZSK3dOMlI2YwpOK2RwTkNjdDdWYjQxaWhXU0VXT1o2djhFdTREV3FjVW0vQzBpLzZ3eVJkWU93Y0E3NktkMmkxcnNZdHFlQS84CnZuMklTTEp2eGdkM3ZCdXY1ZnFkS0kzRm0rd3NIVDhHdHlkQUw5VjF4V3g4bnFlcFlwMlIvaXJMNHFUeGFnVzkKMXdJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg==" binding:"required"`
}

type PIDResponseModel struct {
	PID     string `json:"pid" example:"Nifl3y+2jmuAxF26jqpjogu0ZYnA6IxSikjmTnnjm7k="`
	Message string `json:"message" example:"PID successfully issued and saved"`
}

type User struct {
	PID       string
	PublicKey string
	Nonce     string
}
