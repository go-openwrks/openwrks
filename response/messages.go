package response

import (
	"fmt"
	"strings"
)

type (
	// Error structure.
	Error struct {
		Code    int64    `json:"errorCode"`
		Status  int64    `json:"status"`
		Message string   `json:"message"`
		Links   []string `json:"_links"`
	}
)

// Transform a help to ensure an error type is an Openwrks response error.
func Transform(err error) (*Error, bool) {
	e, ok := err.(*Error)
	return e, ok
}

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) String() string {
	return fmt.Sprintf(
		"Openwrks error [code:%d] [status:%d] %s %s",
		e.Code,
		e.Status,
		e.Message,
		strings.Join(e.Links, ", "),
	)
}
