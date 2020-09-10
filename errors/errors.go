package errors

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// ExitCoder defines the interface for exit code handling.
type ExitCoder interface {
	error
	Code() int
	Fields() logrus.Fields
}

// ExitError simply implements the defined interface.
type ExitError struct {
	message interface{}
	code    int
	fields  logrus.Fields
}

// Error implements the ExitCoder interface.
func (e ExitError) Error() string {
	return fmt.Sprintf("%v", e.message)
}

// Code implements the ExitCoder interface.
func (e ExitError) Code() int {
	return e.code
}

// Fields implements the ExitCoder interface.
func (e ExitError) Fields() logrus.Fields {
	return e.fields
}

// ExitMessage initializes a new ExitCoder implementation.
func ExitMessage(message interface{}) ExitError {
	return ExitError{
		message: message,
		code:    1,
	}
}

// ExitMessagef initializes a new ExitCoder implementation.
func ExitMessagef(format string, a ...interface{}) ExitError {
	return ExitError{
		message: fmt.Errorf(format, a...),
		code:    1,
	}
}

// WithFields initializes a new ExitCoder implementation.
func WithFields(message interface{}, fields logrus.Fields) ExitError {
	return ExitError{
		message: message,
		code:    1,
		fields:  fields,
	}
}

// HandleExit ist used within the main handler to exit properly.
func HandleExit(err error) {
	if err == nil {
		return
	}

	if e, ok := err.(ExitCoder); ok {
		if e.Error() != "" {
			logrus.WithFields(
				e.Fields(),
			).Error(
				e.Error(),
			)
		}

		os.Exit(e.Code())
	}
}
