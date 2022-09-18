package pkg

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrapErr(t *testing.T) {
	err := errors.New("error message")
	wrappers := []string{"wrapper_1", "wrapper_2", "wrapper_3"}
	expectedErrMsg := "wrapper_3; wrapper_2; wrapper_1; error message"

	for _, wrp := range wrappers {
		err = WrapErr(wrp, err)
	}

	require.Equal(t, expectedErrMsg, err.Error())
}
