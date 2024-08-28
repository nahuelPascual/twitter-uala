package rest

import "fmt"

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	ErrMsg     string `json:"err_msg"`
}

func (err ErrorResponse) Error() string {
	return fmt.Sprintf("[%d] %s", err.StatusCode, err.ErrMsg)
}
