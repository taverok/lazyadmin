package rest

import (
	"net/http"
)

var (
	ErrGeneric          = NewError(nil, "error occurred")
	ErrNotImplemented   = NewError(nil, "not implemented")
	ErrNotAuthorized    = NewError(nil, "not authorized").SetHttpCode(http.StatusUnauthorized)
	ErrForbidden        = NewError(nil, "forbidden").SetHttpCode(http.StatusForbidden)
	ErrNotFound         = NewError(nil, "not found").SetHttpCode(http.StatusNotFound)
	ErrResourceNotFound = NewError(nil, "resource not found").SetHttpCode(http.StatusNotFound)
	ErrInvalidParams    = NewError(nil, "invalid params")
	ErrInternal         = NewError(nil, "internal error").SetHttpCode(http.StatusInternalServerError)
	ErrNotOwnedResource = NewError(nil, "this resource is not owned by this author")
	ErrInvalidRequest   = NewError(nil, "request format is not valid")
	ErrAuth             = NewError(nil, "wrong password or email")
)

type PublicError struct {
	Cause     error
	PublicMsg string
	HttpCode  int
}

func NewError(err error, publicMsg string) *PublicError {
	return &PublicError{
		Cause:     err,
		PublicMsg: publicMsg,
		HttpCode:  http.StatusBadRequest,
	}
}

func (it *PublicError) SetHttpCode(code int) *PublicError {
	it.HttpCode = code
	return it
}

func (it *PublicError) Error() string {
	if it.Cause != nil {
		it.Cause.Error()
	}
	return it.PublicMsg
}

func (it *PublicError) Unwrap() error {
	return it.Cause
}

func (it *PublicError) PublicError() string {
	return it.PublicMsg
}
