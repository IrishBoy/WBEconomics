# StocksItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastChangeDate** | [**time.Time**](time.Time.md) | Дата и время обновления информации в сервисе. Это поле соответствует параметру &#x60;dateFrom&#x60; в запросе. Если часовой пояс не указан, то берется Московское время UTC+3. | [optional] [default to null]
**SupplierArticle** | **string** | Артикул продавца | [optional] [default to null]
**TechSize** | **string** | Размер товара (пример S, M, L, XL, 42, 42-43) | [optional] [default to null]
**Barcode** | **string** | Бар-код | [optional] [default to null]
**Quantity** | **int32** | Количество, доступное для продажи (сколько можно добавить в корзину) | [optional] [default to null]
**IsSupply** | **bool** | Договор поставки | [optional] [default to null]
**IsRealization** | **bool** | Договор реализации | [optional] [default to null]
**QuantityFull** | **int32** | Полное (непроданное) количество, которое числится за складом (&#x3D; &#x60;quantity&#x60; + в пути) | [optional] [default to null]
**WarehouseName** | **string** | Название склада | [optional] [default to null]
**NmId** | **int32** | Артикул WB | [optional] [default to null]
**Subject** | **string** | Предмет | [optional] [default to null]
**Category** | **string** | Категория | [optional] [default to null]
**DaysOnSite** | **int32** | Количество дней на сайте | [optional] [default to null]
**Brand** | **string** | Бренд | [optional] [default to null]
**SCCode** | **string** | Код контракта | [optional] [default to null]
**Price** | **float64** | Цена | [optional] [default to null]
**Discount** | **float64** | Скидка | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

