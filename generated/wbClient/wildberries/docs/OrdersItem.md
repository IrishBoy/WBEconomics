# OrdersItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GNumber** | **string** | Номер заказа. Объединяет все позиции одного заказа. | [optional] [default to null]
**Date** | [**time.Time**](time.Time.md) | Дата и время заказа. Это поле соответствует параметру &#x60;dateFrom&#x60; в запросе, если параметр &#x60;flag&#x3D;1&#x60;. Если часовой пояс не указан, то берется Московское время UTC+3. | [optional] [default to null]
**LastChangeDate** | [**time.Time**](time.Time.md) | Дата и время обновления информации в сервисе. Это поле соответствует параметру &#x60;dateFrom&#x60; в запросе, если параметр &#x60;flag&#x3D;0&#x60; или не указан. Если часовой пояс не указан, то берется Московское время UTC+3. | [optional] [default to null]
**SupplierArticle** | **string** | Артикул продавца | [optional] [default to null]
**TechSize** | **string** | Размер товара (пример S, M, L, XL, 42, 42-43) | [optional] [default to null]
**Barcode** | **string** | Бар-код | [optional] [default to null]
**TotalPrice** | **float64** | Цена до согласованной итоговой скидки/промо/спп. Для получения цены со скидкой можно воспользоваться формулой &#x60;priceWithDiscount &#x3D; totalPrice * (1 - discountPercent/100)&#x60; | [optional] [default to null]
**DiscountPercent** | **int32** | Согласованный итоговый дисконт. Будучи примененным к &#x60;totalPrice&#x60;, даёт сумму к оплате. | [optional] [default to null]
**WarehouseName** | **string** | Название склада отгрузки | [optional] [default to null]
**Oblast** | **string** | Область | [optional] [default to null]
**IncomeID** | **int32** | Номер поставки (от продавца на склад) | [optional] [default to null]
**Odid** | **int32** | Уникальный идентификатор позиции заказа. Может использоваться для поиска соответствия между заказами и продажами. | [optional] [default to null]
**NmId** | **int32** | Артикул WB | [optional] [default to null]
**Subject** | **string** | Предмет | [optional] [default to null]
**Category** | **string** | Категория | [optional] [default to null]
**Brand** | **string** | Бренд | [optional] [default to null]
**IsCancel** | **bool** | Отмена заказа. true - заказ отменен до оплаты. | [optional] [default to null]
**CancelDt** | [**time.Time**](time.Time.md) | Дата и время отмены заказа. Если заказ не был отменен, то &#x60;\&quot;0001-01-01T00:00:00\&quot;&#x60;. Если часовой пояс не указан, то берется Московское время UTC+3. | [optional] [default to null]
**Sticker** | **string** | Цифровое значение стикера, который клеится на товар в процессе сборки заказа по системе Маркетплейс. | [optional] [default to null]
**Srid** | **string** | Уникальный идентификатор заказа, функционально аналогичный &#x60;odid&#x60;/&#x60;rid&#x60;.  Данный параметр введен в июле&#x27;22 и в течение переходного периода может быть заполнен не во всех ответах. Примечание для работающих по системе Маркетплейс: &#x60;srid&#x60; равен &#x60;rid&#x60; в ответе на метод &#x60;GET /api/v2/orders&#x60;.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

