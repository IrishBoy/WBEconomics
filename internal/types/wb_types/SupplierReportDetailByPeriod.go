package wb_types

import (
	"time"

	"github.com/antihax/optional"
)

type SupplierReportDetailByPeriodOpts struct {
	Limit optional.Int32
	Rrdid optional.Int32
}

type DetailReportItem struct {
	// Номер отчёта
	RealizationreportId int32 `json:"realizationreport_id,omitempty"`
	// Дата начала отчётного периода
	DateFrom time.Time `json:"date_from,omitempty"`
	// Дата конца отчётного периода
	DateTo time.Time `json:"date_to,omitempty"`
	// Дата формирования отчёта
	CreateDt time.Time `json:"create_dt,omitempty"`
	// Договор
	SuppliercontractCode *interface{} `json:"suppliercontract_code,omitempty"`
	// Номер строки
	RrdId int32 `json:"rrd_id,omitempty"`
	// Номер поставки
	GiId int32 `json:"gi_id,omitempty"`
	// Предмет
	SubjectName string `json:"subject_name,omitempty"`
	// Артикул WB
	NmId int32 `json:"nm_id,omitempty"`
	// Бренд
	BrandName string `json:"brand_name,omitempty"`
	// Артикул продавца
	SaName string `json:"sa_name,omitempty"`
	// Размер
	TsName string `json:"ts_name,omitempty"`
	// Баркод
	Barcode string `json:"barcode,omitempty"`
	// Тип документа
	DocTypeName string `json:"doc_type_name,omitempty"`
	// Количество
	Quantity int32 `json:"quantity,omitempty"`
	// Цена розничная
	RetailPrice float64 `json:"retail_price,omitempty"`
	// Сумма продаж (возвратов)
	RetailAmount float64 `json:"retail_amount,omitempty"`
	// Согласованная скидка
	SalePercent int32 `json:"sale_percent,omitempty"`
	// Процент комиссии
	CommissionPercent float64 `json:"commission_percent,omitempty"`
	// Склад
	OfficeName string `json:"office_name,omitempty"`
	// Обоснование для оплаты
	SupplierOperName string `json:"supplier_oper_name,omitempty"`
	// Дата заказа. <br> Присылается с явным указанием часового пояса
	OrderDt time.Time `json:"order_dt,omitempty"`
	// Дата продажи. <br> Присылается с явным указанием часового пояса
	SaleDt time.Time `json:"sale_dt,omitempty"`
	// Дата операции. <br> Присылается с явным указанием часового пояса
	RrDt time.Time `json:"rr_dt,omitempty"`
	// Штрих-код
	ShkId int32 `json:"shk_id,omitempty"`
	// Цена розничная с учетом согласованной скидки
	RetailPriceWithdiscRub float64 `json:"retail_price_withdisc_rub,omitempty"`
	// Количество доставок
	DeliveryAmount int32 `json:"delivery_amount,omitempty"`
	// Количество возвратов
	ReturnAmount int32 `json:"return_amount,omitempty"`
	// Стоимость логистики
	DeliveryRub float64 `json:"delivery_rub,omitempty"`
	// Тип коробов
	GiBoxTypeName string `json:"gi_box_type_name,omitempty"`
	// Согласованный продуктовый дисконт
	ProductDiscountForReport float64 `json:"product_discount_for_report,omitempty"`
	// Промокод
	SupplierPromo float64 `json:"supplier_promo,omitempty"`
	// Уникальный идентификатор заказа
	Rid int32 `json:"rid,omitempty"`
	// Скидка постоянного покупателя
	PpvzSppPrc float64 `json:"ppvz_spp_prc,omitempty"`
	// Размер кВВ без НДС, % базовый
	PpvzKvwPrcBase float64 `json:"ppvz_kvw_prc_base,omitempty"`
	// Итоговый кВВ без НДС, %
	PpvzKvwPrc float64 `json:"ppvz_kvw_prc,omitempty"`
	// Вознаграждение с продаж до вычета услуг поверенного, без НДС
	PpvzSalesCommission float64 `json:"ppvz_sales_commission,omitempty"`
	// К перечислению продавцу за реализованный товар
	PpvzForPay float64 `json:"ppvz_for_pay,omitempty"`
	// Возмещение за выдачу и возврат товаров на ПВЗ
	PpvzReward float64 `json:"ppvz_reward,omitempty"`
	// Возмещение расходов по эквайрингу. <br>Расходы WB на услуги эквайринга: вычитаются из вознаграждения WB и не влияют на доход продавца
	AcquiringFee float64 `json:"acquiring_fee,omitempty"`
	// Наименование банка, предоставляющего услуги эквайринга
	AcquiringBank string `json:"acquiring_bank,omitempty"`
	// Вознаграждение WB без НДС
	PpvzVw float64 `json:"ppvz_vw,omitempty"`
	// НДС с вознаграждения WB
	PpvzVwNds float64 `json:"ppvz_vw_nds,omitempty"`
	// Номер офиса
	PpvzOfficeId int32 `json:"ppvz_office_id,omitempty"`
	// Наименование офиса доставки
	PpvzOfficeName string `json:"ppvz_office_name,omitempty"`
	// Номер партнера
	PpvzSupplierId int32 `json:"ppvz_supplier_id,omitempty"`
	// Партнер
	PpvzSupplierName string `json:"ppvz_supplier_name,omitempty"`
	// ИНН партнера
	PpvzInn string `json:"ppvz_inn,omitempty"`
	// Номер таможенной декларации
	DeclarationNumber string `json:"declaration_number,omitempty"`
	// Обоснование штрафов и доплат. <br> Поле будет в ответе при наличии значения
	BonusTypeName string `json:"bonus_type_name,omitempty"`
	// Цифровое значение стикера, который клеится на товар в процессе сборки заказа по схеме \"Маркетплейс\"
	StickerId string `json:"sticker_id,omitempty"`
	// Страна продажи
	SiteCountry string `json:"site_country,omitempty"`
	// Штрафы
	Penalty float64 `json:"penalty,omitempty"`
	// Доплаты
	AdditionalPayment float64 `json:"additional_payment,omitempty"`
	// Код маркировки. <br> Поле будет в ответе при наличии значения
	Kiz string `json:"kiz,omitempty"`
	// Уникальный идентификатор заказа. Примечание для использующих API Marketplace: `srid` равен `rid` в ответах методов сборочных заданий.
	Srid string `json:"srid,omitempty"`
}
