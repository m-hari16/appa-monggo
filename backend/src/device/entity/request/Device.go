package request

type DeviceId string

type Device struct {
	Id          string `json:"_id"`
	MacAddress  string `json:"mac_address" validate:"required"`
	UserId      string `json:"user_id" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	Model       string `json:"model" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}
