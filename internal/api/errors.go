package api

import "errors"

var (
	BodyDecodeErr = errors.New("failed to decode the body")
	EmptyBodyErr  = errors.New("there is no request body, but it is expected")
)
