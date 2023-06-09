# DetailReportItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RealizationreportId** | **int32** | Номер отчёта | [optional] [default to null]
**DateFrom** | [**time.Time**](time.Time.md) | Дата начала отчётного периода | [optional] [default to null]
**DateTo** | [**time.Time**](time.Time.md) | Дата конца отчётного периода | [optional] [default to null]
**CreateDt** | [**time.Time**](time.Time.md) | Дата формирования отчёта | [optional] [default to null]
**SuppliercontractCode** | [***interface{}**](interface{}.md) | Договор | [optional] [default to null]
**RrdId** | **int32** | Номер строки | [optional] [default to null]
**GiId** | **int32** | Номер поставки | [optional] [default to null]
**SubjectName** | **string** | Предмет | [optional] [default to null]
**NmId** | **int32** | Артикул WB | [optional] [default to null]
**BrandName** | **string** | Бренд | [optional] [default to null]
**SaName** | **string** | Артикул продавца | [optional] [default to null]
**TsName** | **string** | Размер | [optional] [default to null]
**Barcode** | **string** | Баркод | [optional] [default to null]
**DocTypeName** | **string** | Тип документа | [optional] [default to null]
**Quantity** | **int32** | Количество | [optional] [default to null]
**RetailPrice** | **float64** | Цена розничная | [optional] [default to null]
**RetailAmount** | **float64** | Сумма продаж (возвратов) | [optional] [default to null]
**SalePercent** | **int32** | Согласованная скидка | [optional] [default to null]
**CommissionPercent** | **float64** | Процент комиссии | [optional] [default to null]
**OfficeName** | **string** | Склад | [optional] [default to null]
**SupplierOperName** | **string** | Обоснование для оплаты | [optional] [default to null]
**OrderDt** | [**time.Time**](time.Time.md) | Дата заказа. &lt;br&gt; Присылается с явным указанием часового пояса | [optional] [default to null]
**SaleDt** | [**time.Time**](time.Time.md) | Дата продажи. &lt;br&gt; Присылается с явным указанием часового пояса | [optional] [default to null]
**RrDt** | [**time.Time**](time.Time.md) | Дата операции. &lt;br&gt; Присылается с явным указанием часового пояса | [optional] [default to null]
**ShkId** | **int32** | Штрих-код | [optional] [default to null]
**RetailPriceWithdiscRub** | **float64** | Цена розничная с учетом согласованной скидки | [optional] [default to null]
**DeliveryAmount** | **int32** | Количество доставок | [optional] [default to null]
**ReturnAmount** | **int32** | Количество возвратов | [optional] [default to null]
**DeliveryRub** | **float64** | Стоимость логистики | [optional] [default to null]
**GiBoxTypeName** | **string** | Тип коробов | [optional] [default to null]
**ProductDiscountForReport** | **float64** | Согласованный продуктовый дисконт | [optional] [default to null]
**SupplierPromo** | **float64** | Промокод | [optional] [default to null]
**Rid** | **int32** | Уникальный идентификатор заказа | [optional] [default to null]
**PpvzSppPrc** | **float64** | Скидка постоянного покупателя | [optional] [default to null]
**PpvzKvwPrcBase** | **float64** | Размер кВВ без НДС, % базовый | [optional] [default to null]
**PpvzKvwPrc** | **float64** | Итоговый кВВ без НДС, % | [optional] [default to null]
**PpvzSalesCommission** | **float64** | Вознаграждение с продаж до вычета услуг поверенного, без НДС | [optional] [default to null]
**PpvzForPay** | **float64** | К перечислению продавцу за реализованный товар | [optional] [default to null]
**PpvzReward** | **float64** | Возмещение за выдачу и возврат товаров на ПВЗ | [optional] [default to null]
**AcquiringFee** | **float64** | Возмещение расходов по эквайрингу. &lt;br&gt;Расходы WB на услуги эквайринга: вычитаются из вознаграждения WB и не влияют на доход продавца  | [optional] [default to null]
**AcquiringBank** | **string** | Наименование банка, предоставляющего услуги эквайринга | [optional] [default to null]
**PpvzVw** | **float64** | Вознаграждение WB без НДС | [optional] [default to null]
**PpvzVwNds** | **float64** | НДС с вознаграждения WB | [optional] [default to null]
**PpvzOfficeId** | **int32** | Номер офиса | [optional] [default to null]
**PpvzOfficeName** | **string** | Наименование офиса доставки | [optional] [default to null]
**PpvzSupplierId** | **int32** | Номер партнера | [optional] [default to null]
**PpvzSupplierName** | **string** | Партнер | [optional] [default to null]
**PpvzInn** | **string** | ИНН партнера | [optional] [default to null]
**DeclarationNumber** | **string** | Номер таможенной декларации | [optional] [default to null]
**BonusTypeName** | **string** | Обоснование штрафов и доплат. &lt;br&gt; Поле будет в ответе при наличии значения  | [optional] [default to null]
**StickerId** | **string** | Цифровое значение стикера, который клеится на товар в процессе сборки заказа по схеме \&quot;Маркетплейс\&quot; | [optional] [default to null]
**SiteCountry** | **string** | Страна продажи | [optional] [default to null]
**Penalty** | **float64** | Штрафы | [optional] [default to null]
**AdditionalPayment** | **float64** | Доплаты | [optional] [default to null]
**Kiz** | **string** | Код маркировки. &lt;br&gt; Поле будет в ответе при наличии значения  | [optional] [default to null]
**Srid** | **string** | Уникальный идентификатор заказа. Примечание для использующих API Marketplace: &#x60;srid&#x60; равен &#x60;rid&#x60; в ответах методов сборочных заданий.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

