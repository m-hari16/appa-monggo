package helper

type Response struct {
	Code    string      `json:"code"`
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func HasOk(data interface{}) Response {
	return Response{
		Code:    "200",
		Success: true,
		Message: "Ok",
		Data:    data,
	}
}

func HasStore(data interface{}) Response {
	messages := []string{"Data has been store"}
	return Response{
		Code:    "201",
		Success: true,
		Message: messages,
		Data:    data,
	}
}

func UnprocessableEntity() Response {
	messages := []string{"Unprocessable Entity"}
	return Response{
		Code:    "422",
		Success: false,
		Message: messages,
		Data:    nil,
	}
}

func BadRequest(message string) Response {
	var messages []string
	messages = append(messages, message)
	return Response{
		Code:    "400",
		Success: false,
		Message: messages,
		Data:    nil,
	}
}

func OnlyMessage(message string) Response {
	var messages []string
	messages = append(messages, message)
	return Response{
		Code:    "200",
		Success: true,
		Message: messages,
		Data:    nil,
	}
}

func Unauthorized(message string) Response {
	var messages []string
	messages = append(messages, message)
	return Response{
		Code:    "401",
		Success: false,
		Message: messages,
		Data:    nil,
	}
}
