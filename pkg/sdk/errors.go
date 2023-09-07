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

type sdkError struct {
	file         string
	line         int
	message      string
	nestedErrors []error
}

func NewError(message string) error {
	_, file, line, _ := runtime.Caller(1)
	fileSplit := strings.Split(file, "/")
	var filename string
	if len(fileSplit) > 1 {
		filename = fileSplit[len(fileSplit)-1]
	} else {
		filename = fileSplit[0]
	}
	return &sdkError{
		file:         filename,
		line:         line,
		message:      message,
		nestedErrors: make([]error, 0),
	}
}

// TODO We can force to use sdkError
func WrapError(wrapper error, errs ...error) error {
	if sdkErr, ok := wrapper.(*sdkError); ok {
		sdkErr.nestedErrors = append(sdkErr.nestedErrors, errs...)
		return sdkErr
	} else {
		joinedErrs := []error{wrapper}
		joinedErrs = append(joinedErrs, errs...)
		return errors.Join(joinedErrs...)
	}
}

//func (e *sdkError) errorIndented(indent int) string {
//	errorMessage := fmt.Sprintf("%s[%s:%d] %s", strings.Repeat(" ", indent), e.file, e.line, e.message)
//	if len(e.nestedErrors) > 0 {
//		var b []byte
//		for i, err := range e.nestedErrors {
//			if i > 0 {
//				b = append(b, '\n')
//			}
//
//			b = append(b, "├──"...)
//
//			var sdkErr *sdkError
//			if errors.As(err, &sdkErr) {
//				b = append(b, sdkErr.errorIndented(indent+2)...)
//				continue
//			}
//
//			if joinedErr, ok := err.(interface{ Unwrap() []error }); ok {
//				errs := joinedErr.Unwrap()
//				for j, e := range errs {
//					if j > 0 {
//						b = append(b, '\n')
//					}
//					b = append(b, strings.Repeat(" ", indent+2)...)
//					b = append(b, e.Error()...)
//				}
//				continue
//			}
//
//			b = append(b, strings.Repeat(" ", indent+2)...)
//			b = append(b, err.Error()...)
//		}
//		return fmt.Sprintf("%s\n%s", errorMessage, string(b))
//	}
//	return errorMessage
//}

func (e *sdkError) writeTree(builder *strings.Builder, indent int) {
	builder.WriteString(strings.Repeat("› ", indent) + fmt.Sprintf("[%s:%d] %s\n", e.file, e.line, e.message))

	for _, err := range e.nestedErrors {
		var sdkErr *sdkError
		if errors.As(err, &sdkErr) {
			sdkErr.writeTree(builder, indent+2)
		}
	}
}

func (e *sdkError) Error() string {
	builder := new(strings.Builder)
	e.writeTree(builder, 0)
	return builder.String()
}
