package server

import (
	"WBEconomics/internal/types"
	"time"
)

func getDetailedReport(user int, start_period time.Time, end_period time.Time) ([]types.DetailedOperation, error) {

	if validateTime := periodValidate(start_period, end_period); validateTime != nil {
		return nil, validateTime
	}

	return nil, nil

}

func getFinancialReport(user int, start_period time.Time, end_period time.Time) ([]types.DetailedOperation, error) {

	if validateTime := periodValidate(start_period, end_period); validateTime != nil {
		return nil, validateTime
	}

	return nil, nil

}
