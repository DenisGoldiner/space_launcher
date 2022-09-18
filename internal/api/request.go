package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/DenisGoldiner/space_launcher/pkg"
)

func parseCreateBookingRequest(r *http.Request, dest any) error {
	if err := DecodeJSON(r, dest); err != nil {
		if errors.Is(err, io.EOF) {
			return EmptyBodyErr
		}

		return pkg.WrapErr(err.Error(), BodyDecodeErr)
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
