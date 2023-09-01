package sdk

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"strings"
)

var (
	// go-snowflake errors.
	ErrObjectNotExistOrAuthorized = errors.New("object does not exist or not authorized")
	ErrAccountIsEmpty             = errors.New("account is empty")

	// snowflake-sdk errors.
	ErrInvalidObjectIdentifier = errors.New("invalid object identifier")
)

func errOneOf(fieldNames ...string) error {
	return fmt.Errorf("fields %v are incompatible and cannot be set at once", fieldNames)
}

func errExactlyOneOf(fieldNames ...string) error {
	return fmt.Errorf("exactly one of %v must be set", fieldNames)
}

func errAtLeastOneOf(fieldNames ...string) error {
	return fmt.Errorf("at least one of %v must be set", fieldNames)
}

func decodeDriverError(err error) error {
	if err == nil {
		return nil
	}
	log.Printf("[DEBUG] err: %v\n", err)
	m := map[string]error{
		"does not exist or not authorized": ErrObjectNotExistOrAuthorized,
		"account is empty":                 ErrAccountIsEmpty,
	}
	for k, v := range m {
		if strings.Contains(err.Error(), k) {
			return v
		}
	}

	return err
}

type errorSDK struct {
	file         string
	line         int
	message      string
	nestedErrors []error
}

func NewError(message string) error {
	_, file, line, _ := runtime.Caller(1)
	return &errorSDK{
		file:         file,
		line:         line,
		message:      message,
		nestedErrors: make([]error, 0),
	}
}

func (e *errorSDK) Error() string {
	errorMessage := fmt.Sprintf("[%s: line %d] %s", e.file, e.line, e.message)
	if len(e.nestedErrors) > 0 {
		var b []byte
		for i, err := range e.nestedErrors {
			if i > 0 {
				b = append(b, '\n')
			}
			b = append(b, err.Error()...)
		}
		return fmt.Sprintf("%s\n%s", errorMessage, string(b))
	}
	return errorMessage
}
