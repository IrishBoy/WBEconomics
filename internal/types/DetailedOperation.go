package types

import (
	"time"
)

type DetailedOperation struct {
	No                         int
	DeliveryNumber             int
	Item                       string
	NomenclatureCode           int
	Brand                      string
	ArticleOfSupplier          string
	Name                       string
	Size                       string
	Barkod                     int
	TypeOfDocument             string
	JustificationForPayment    string
	OrderDateByBuyer           time.Time
	DateOfSale                 time.Time
	Amount                     int
	RetailPrice                float64
	WildberriesSoldGoods       float64
	CoordinatedGroceryDiscount float64
	PromoCode                  float64
	FinalAgreedDiscount        float64
	RetailPriceWithDiscount    float64
	BuyerDiscount              float64
	KVVSize                    float64
	KVVSizeWithoutVAT          float64
	FinalKVVWithoutVAT         float64
	RemunerationFromSales      float64
	Compensation               float64
	Acquiring                  float64
	RemunerationWildberry      float64
	VATWithRewards             float64
	TransferToSeller           float64
	DeliveryNumberT            int
	ReturnNumber               int
	ServicesDeliveryToBuyer    float64
	TotalFines                 float64
	Surcharges                 float64
	LogisticsTypes             string
	StickerMP                  string
	AcquiringBankName          string
	OfficeNumber               string
	DeliveryOfficeName         string
	PartnerTIN                 string
	Partner                    string
	Stock                      string
	Country                    string
	BoxType                    string
	NumberOfCustoms            string
	MarkingCode                string
	ShK                        int
	Rid                        string
	Srid                       string
}
