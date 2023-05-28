package types

import (
	"time"
)

type SoldUnit struct {
	unit_name string
	unis_sold int
	total_sum float64
}

type OutputReport struct {
	Start_period                       time.Time
	End_period                         time.Time
	Week_number                        int
	Comission                          float64
	Fullfilment_amount                 float64
	Packaging_amount                   float64
	Taxes                              float64
	Prime_cost                         float64
	Logistc_amount                     float64
	Sold_units                         []SoldUnit
	Shop_sold_amount                   float64
	Return_number                      int
	Return_amount                      float64
	Self_purchase_number               int
	Self_purchase_amount               float64
	Self_purchase_services_amount      float64
	Self_purchase_amount_percent       float64
	Sales_amount_without_self_purchase float64
	Ads_amount                         float64
	Other_expenses_amount              float64
	Salaries                           float64
	Storage                            float64
	Net_profit                         float64
}

func (outputReport *OutputReport) CalculateSalestWithoutSelfPurchase() {
	outputReport.Sales_amount_without_self_purchase = outputReport.Shop_sold_amount - outputReport.Self_purchase_amount
}

func (outputReport *OutputReport) CalculateNetProfit() {
	outputReport.Net_profit = outputReport.Sales_amount_without_self_purchase - outputReport.Comission -
		outputReport.Fullfilment_amount - outputReport.Packaging_amount -
		outputReport.Taxes - outputReport.Prime_cost -
		outputReport.Logistc_amount - outputReport.Ads_amount -
		outputReport.Other_expenses_amount - outputReport.Salaries -
		outputReport.Salaries - outputReport.Storage
}
