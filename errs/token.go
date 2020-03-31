package errs

import "github.com/pkg/errors"

var (
	TokenIsNotValid = errors.New("The token is not valid")
)
