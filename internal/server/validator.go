package server

import (
	"errors"
	"time"
)

func periodValidate(start_period time.Time, end_period time.Time) error {
	if end_period.Before(start_period) {
		return errors.New("Start period must be before end period")
	}
	if time.Now().Before(end_period) {
		return errors.New("End period must not be in the future")
	}
	if start_period.IsZero() {
		return errors.New("Start period is null")
	}

	if end_period.IsZero() {
		return errors.New("End period is null")
	}

	return nil
}
