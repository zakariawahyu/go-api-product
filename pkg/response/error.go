package response

import (
	"github.com/pkg/errors"
	"net/http"
)

var (
	ErrNotFound = errors.New("Your requested item is not found")
	ErrConflict = errors.New("Your item already exist")
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
