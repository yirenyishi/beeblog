package models

type Error struct {
	Status int
	Msg    string
}

func ReurnError(status int,msg string) *Error {
	if msg == "" {
		msg = "error"
	}
	return &Error{Status: status, Msg: msg}
}

type Success struct {
	Status int
	Msg    string
}

func ReurnSuccess(msg string) *Error {
	if msg == "" {
		msg = "success"
	}
	return &Error{Status: 0, Msg: msg}
}
