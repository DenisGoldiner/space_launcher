package pkg

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	// EmptyBodyErr empty body
	EmptyBodyErr = errors.New("there is no request body, but it is expected")
	// BodyDecodeErr invalid body
	BodyDecodeErr = errors.New("failed to decode the body")
)

// ParseRequestPayload decodes the request body into the dest.
// the dest must be a pointer type.
func ParseRequestPayload(r *http.Request, dest any) error {
	if err := DecodeJSON(r, dest); err != nil {
		if errors.Is(err, io.EOF) {
			return EmptyBodyErr
		}

		return WrapErr(err.Error(), BodyDecodeErr)
	}

	return nil
}

// DecodeJSON decode JSON from *http.Request.
func DecodeJSON(r *http.Request, dest any) error {
	if r.Body == nil {
		return io.EOF
	}

	err := json.NewDecoder(r.Body).Decode(dest)
	if err != nil {
		return err
	}
	defer func() { _ = r.Body.Close() }()

	return nil
}
