package api

import "errors"

var (
	BodyDecodeErr       = errors.New("failed to decode the body")
	EmptyBodyErr        = errors.New("there is no request body, but it is expected")
	EmptyFieldErr       = errors.New("field should not be empty")
	ValidationFailedErr = errors.New("the request validation failed")
)
