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
	start_period                       time.Time
	end_period                         time.Time
	week_number                        int
	comission                          float32
	fullfilment_amount                 float32
	packaging_amount                   float32
	taxes                              float32
	prime_cost                         float32
	logistc_amount                     float32
	sold_units                         []SoldUnit
	shop_sold_amount                   float32
	return_number                      int
	return_amount                      float32
	self_purchase_number               int
	self_purchase_amount               float32
	self_purchase_services_amount      float32
	self_purchase_amount_percent       float32
	sales_amount_without_self_purchase float32
	ads_amount                         float32
	other_expenses_amount              float32
	salaries                           float32
	storage                            float32
	net_profit                         float32
}

func (outputReport *OutputReport) calculateSalestWithoutSelfPurchase() {
	outputReport.sales_amount_without_self_purchase = outputReport.shop_sold_amount - outputReport.self_purchase_amount
}

func (outputReport *OutputReport) calculateNetProfit() {
	outputReport.net_profit = outputReport.sales_amount_without_self_purchase - outputReport.comission -
		outputReport.fullfilment_amount - outputReport.packaging_amount -
		outputReport.taxes - outputReport.prime_cost -
		outputReport.logistc_amount - outputReport.ads_amount -
		outputReport.other_expenses_amount - outputReport.salaries -
		outputReport.salaries - outputReport.storage
}
