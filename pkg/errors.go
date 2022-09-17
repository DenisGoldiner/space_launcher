package pkg

import "fmt"

const (
	errWrapFILOFormat = "%s; %w"
)

func WrapErr(cause string, err error) error {
	return fmt.Errorf(errWrapFILOFormat, cause, err)
}
