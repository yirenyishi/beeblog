package models

type Error struct {
	Status int
	Msg    string
}

func ReurnError(status int, msg string) *Error {
	if msg == "" {
		msg = "error"
	}
	return &Error{Status: status, Msg: msg}
}

func ReurnSuccess(msg string) *Error {
	if msg == "" {
		msg = "success"
	}
	return &Error{Status: 0, Msg: msg}
}

type Success struct {
	Status int
	Msg    string
	Data   interface{}
}

func ReurnData(msg string, data interface{}) *Success {
	if msg == "" {
		msg = "success"
	}
	return &Success{Status: 0, Msg: msg, Data: data}
}
