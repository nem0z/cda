package scrapper

import (
	"fmt"
	"net/http"
)

type DoNotExistError struct {
}

type StatusError struct {
	res *http.Response
}

func (err DoNotExistError) Error() string {
	return fmt.Sprintf("Page don't exist")
}

func (err StatusError) Error() string {
	return fmt.Sprintf("Request status error : (%v) %v", err.res.StatusCode, err.res.Status)
}
