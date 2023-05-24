# IncomesItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomeId** | **int32** | Номер поставки | [optional] [default to null]
**Number** | **string** | Номер УПД | [optional] [default to null]
**Date** | **string** | Дата поступления. Если часовой пояс не указан, то берется Московское время UTC+3. | [optional] [default to null]
**LastChangeDate** | [**time.Time**](time.Time.md) | Дата и время обновления информации в сервисе. Это поле соответствует параметру &#x60;dateFrom&#x60; в запросе. Если часовой пояс не указан, то берется Московское время UTC+3. | [optional] [default to null]
**SupplierArticle** | **string** | Артикул продавца | [optional] [default to null]
**TechSize** | **string** | Размер товара (пример S, M, L, XL, 42, 42-43) | [optional] [default to null]
**Barcode** | **string** | Бар-код | [optional] [default to null]
**Quantity** | **int32** | Количество | [optional] [default to null]
**TotalPrice** | **float64** | Цена из УПД | [optional] [default to null]
**DateClose** | **string** | Дата принятия (закрытия) в WB. Если часовой пояс не указан, то берется Московское время UTC+3. | [optional] [default to null]
**WarehouseName** | **string** | Название склада | [optional] [default to null]
**NmId** | **int32** | Артикул WB | [optional] [default to null]
**Status** | **string** | Текущий статус поставки | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

