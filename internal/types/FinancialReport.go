package types

import "time"

type FinancialReport struct {
	Commision        float64
	Shop_sold_amount float64
	Start_period     time.Time
	End_period       time.Time
}
