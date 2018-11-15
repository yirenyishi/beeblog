package models

type Error struct {
	Status byte
	Msg    string
}

func ReurnError(msg string) *Error {
	return &Error{Status: byte(1), Msg: msg}
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
