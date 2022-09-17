package platform

import "fmt"

const (
	errWrapFILOFormat = "%s; %w"
)

var (
	BodyDecodeErr = "failed to decode the body"
)

func WrapErr(cause string, err error) error {
	return fmt.Errorf(errWrapFILOFormat, cause, err)
}
