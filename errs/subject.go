package errs

import "errors"

var (
	SubjectDoesntExist = errors.New("Such subject does not exist!")
)
