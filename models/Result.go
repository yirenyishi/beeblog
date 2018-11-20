package models

type Error struct {
	Status byte
	Msg    string
}

func ReurnError(status int,msg string) *Error {
	if msg == "" {
		msg = "error"
	}
	return &Error{Status: byte(status), Msg: msg}
}

type Success struct {
	Status byte
	Msg    string
}

func ReurnSuccess(msg string) *Error {
	if msg == "" {
		msg = "success"
	}
	return &Error{Status: byte(0), Msg: msg}
}
