package utils

import (
	"errors"
	"time"
)

var ErrCannotConnectionToService = errors.New("error, can't connection to a service")

// DoWithTries tries to connect a service (like PostgreSQL, MongoDB, etc...) equal attempts,
// with delay interval with tries
// If can't connect to the service, return ErrCannotConnectToService
func DoWithTries(fn func() error, attempts int, delay time.Duration) error {
	for attempts > 0 {
		if err := fn(); err != nil {
			time.Sleep(delay)
			attempts--
			continue
		}
		return nil
	}
	return ErrCannotConnectionToService
}
