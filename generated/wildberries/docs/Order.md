# Order

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор сборочного задания в Маркетплейсе | [optional] [default to null]
**Rid** | **string** | Идентификатор сборочного задания в системе Wildberries | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Дата создания сборочного задания (RFC3339) | [optional] [default to null]
**WarehouseId** | **int32** | Идентификатор склада продавца, на который поступило сборочное задание | [optional] [default to null]
**SupplyId** | **string** | Идентификатор поставки. Возвращается, если заказ закреплён за поставкой | [optional] [default to null]
**PrioritySc** | **[]string** | Массив приоритетных СЦ для доставки сборочного задания. Если поле не заполнено или массив пустой, приоритетного СЦ для данного сборочного задания нет | [optional] [default to null]
**Offices** | **[]string** | Список офисов, куда следует привезти товар | [optional] [default to null]
**Address** | [***OrderAddress**](Order_address.md) |  | [optional] [default to null]
**User** | [***OrderUser**](Order_user.md) |  | [optional] [default to null]
**Skus** | **[]string** | Массив баркодов товара | [optional] [default to null]
**Price** | **int32** | Цена в валюте продажи с учетом всех скидок, умноженная на 100. Код валюты продажи в поле currencyCode. | [optional] [default to null]
**ConvertedPrice** | **int32** | Цена в валюте продажи с учетом всех скидок, сконвертированная по курсу на момент продажи в российские копейки. Предоставляется в информационных целях. | [optional] [default to null]
**CurrencyCode** | **int32** | Код валюты продажи (ISO 4217) | [optional] [default to null]
**ConvertedCurrencyCode** | **int32** | Код валюты страны поставщика (ISO 4217) | [optional] [default to null]
**OrderUid** | **string** | Идентификатор транзакции для группировки сборочных заданий. Сборочные задания в одной корзине покупателя будут иметь одинаковый orderUID | [optional] [default to null]
**DeliveryType** | **string** | Тип доставки: fbs - доставка на склад Wildberries, dbs - доставка силами продавца  | [optional] [default to null]
**NmId** | **int32** | Артикул WB | [optional] [default to null]
**ChrtId** | **int32** | Идентификатор размера товара в системе Wildberries | [optional] [default to null]
**Article** | **string** | Артикул продавца | [optional] [default to null]
**IsLargeCargo** | **bool** | сКГТ-признак товара, на который был сделан заказ | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

