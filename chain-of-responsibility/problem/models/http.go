package models

type Request struct {
	User *User
}

type Response struct {
	Code    int
	Message string
}

func NewRequest(user *User) *Request {
	return &Request{User: user}
}

func NewResponse(code int, message string) *Response {
	return &Response{Code: code, Message: message}
}
