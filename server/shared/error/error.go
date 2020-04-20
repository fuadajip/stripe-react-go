package error

import (
	"github.com/fuadajip/stripe-react-go/server/models"
	Errors "github.com/pkg/errors"
)

type (
	expError models.ExperimentError
)

// Error returns error type as a string
func (q *expError) Error() string {
	return q.Message
}

// New returnns new error message in standard pkg errors new
func New(msg string) error {
	return Errors.New(msg)
}

// Wrap returns a new error that adds context to the original error
func Wrap(code int, errorCode string, err error, msg string, status string) error {
	return Errors.Wrap(&expError{
		Code:      code,
		ErrorCode: errorCode,
		Message:   msg,
		Status:    status,
	}, err.Error())
}
