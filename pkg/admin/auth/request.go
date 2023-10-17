package auth

import "net/url"

type LoginRequest struct {
	User string
	Pass string
}

func NewLoginRequest(form url.Values) LoginRequest {
	return LoginRequest{
		User: form.Get("user"),
		Pass: form.Get("pass"),
	}
}
