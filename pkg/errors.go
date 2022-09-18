package pkg

import "fmt"

const errWrapFILOFormat = "%s; %w"

// WrapErr formatted FILO wrap for errors.
func WrapErr(cause string, err error) error {
	return fmt.Errorf(errWrapFILOFormat, cause, err)
}
