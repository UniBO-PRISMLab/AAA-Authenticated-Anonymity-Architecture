package models

type PIDRequestModel struct {
	PublicKey string `json:"public_key" example:"-----BEGIN RSA PUBLIC KEY-----MIIBCgKCAQEA2aOVUGMBXS07BIKnFq0w/enJLXtQ1K/Yq0ep+ov68R+TPO7EafnZ/fzZSFygAVjYTkgHWr6fXwp4xtPdUUVYMYJLQG5gqK6oOlONZO1NKwjb24Ww6ViRnxYUJOld/6AlZ6kKHDneVI1aINkjJqx6YMK82u5m/7G1/xeHQ9geeQL/CKCAIf+rKGTHYVYQRPIvKhoEMWmYlrAsHpzW462UZDu/zRMqoQXn+KvXI/WrIvisOwXZLoBve8gA6aLbtxOCCmZ/ARK/SeEJrp1mcnRdHVkUHvveQARkoT3dLEeRmLQf9N1HYcLC8GwcWaY0v7cCv4nCVgvJaJfQzA8SngSQqQIDAQAB-----END RSA PUBLIC KEY-----" binding:"required"`
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
