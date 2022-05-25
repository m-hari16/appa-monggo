package request

type Device struct {
	Id          string `json:"_id"`
	Brand       string `json:"brand" validate:"required"`
	Model       string `json:"model" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type DeviceId string
