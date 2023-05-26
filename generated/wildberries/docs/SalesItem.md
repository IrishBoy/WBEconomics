# SalesItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GNumber** | **string** | Номер заказа. Объединяет все позиции одного заказа. | [optional] [default to null]
**Date** | [**time.Time**](time.Time.md) | Дата и время продажи. Это поле соответствует параметру &#x60;dateFrom&#x60; в запросе, если параметр &#x60;flag&#x3D;1&#x60;. Если часовой пояс не указан, то берется Московское время UTC+3. | [optional] [default to null]
**LastChangeDate** | [**time.Time**](time.Time.md) | Дата и время обновления информации в сервисе. Это поле соответствует параметру &#x60;dateFrom&#x60; в запросе, если параметр &#x60;flag&#x3D;0&#x60; или не указан. Если часовой пояс не указан, то берется Московское время UTC+3. | [optional] [default to null]
**SupplierArticle** | **string** | Артикул продавца | [optional] [default to null]
**TechSize** | **string** | Размер товара (пример S, M, L, XL, 42, 42-43) | [optional] [default to null]
**Barcode** | **string** | Бар-код | [optional] [default to null]
**TotalPrice** | **float64** | Цена до согласованной скидки/промо/спп. Для получения цены со скидкой можно воспользоваться формулой &#x60;priceWithDiscount &#x3D; totalPrice * (1 - discountPercent/100)&#x60; | [optional] [default to null]
**DiscountPercent** | **int32** | Согласованный итоговый дисконт | [optional] [default to null]
**IsSupply** | **bool** | Договор поставки | [optional] [default to null]
**IsRealization** | **bool** | Договор реализации | [optional] [default to null]
**PromoCodeDiscount** | **float64** | Скидка по промокоду | [optional] [default to null]
**WarehouseName** | **string** | Название склада отгрузки | [optional] [default to null]
**CountryName** | **string** | Страна | [optional] [default to null]
**OblastOkrugName** | **string** | Округ | [optional] [default to null]
**RegionName** | **string** | Регион | [optional] [default to null]
**IncomeID** | **int32** | Номер поставки (от продавца на склад) | [optional] [default to null]
**SaleID** | **string** | Уникальный идентификатор продажи/возврата. &lt;ul&gt;  &lt;li&gt; &#x60;SXXXXXXXXXX&#x60; — продажа  &lt;li&gt; &#x60;RXXXXXXXXXX&#x60; — возврат  &lt;li&gt; &#x60;DXXXXXXXXXXX&#x60; — доплата &lt;li&gt; &#x60;AXXXXXXXXX&#x60; – сторно продаж (все значения полей как у продажи, но поля с суммами и кол-вом с минусом как в возврате) &lt;li&gt; &#x60;BXXXXXXXXX&#x60; - сторно возврата (все значения полей как у возврата, но поля с суммами и кол-вом с плюсом, в противоположность возврату) &lt;/ul&gt;  | [optional] [default to null]
**Odid** | **int32** | Уникальный идентификатор позиции заказа. Может использоваться для поиска соответствия между заказами и продажами. | [optional] [default to null]
**Spp** | **float64** | Согласованная скидка постоянного покупателя | [optional] [default to null]
**ForPay** | **float64** | К перечислению продавцку | [optional] [default to null]
**FinishedPrice** | **float64** | Фактическая цена заказа с учетом всех скидок | [optional] [default to null]
**PriceWithDisc** | **float64** | Цена, от которой считается вознаграждение продавца &#x60;forpay&#x60; (с учетом всех согласованных скидок) | [optional] [default to null]
**NmId** | **int32** | Артикул WB | [optional] [default to null]
**Subject** | **string** | Предмет | [optional] [default to null]
**Category** | **string** | Категория | [optional] [default to null]
**Brand** | **string** | Бренд | [optional] [default to null]
**IsStorno** | **int32** | Для сторно-операций &#x60;1&#x60;, для остальных &#x60;0&#x60; | [optional] [default to null]
**Sticker** | **string** | Цифровое значение стикера, который клеится на товар в процессе сборки заказа по системе Маркетплейс. | [optional] [default to null]
**Srid** | **string** | Уникальный идентификатор заказа, функционально аналогичный &#x60;odid&#x60;/&#x60;rid&#x60;.  Данный параметр введен в июле&#x27;22 и в течение переходного периода может быть заполнен не во всех ответах. Примечание для работающих по системе Маркетплейс: &#x60;srid&#x60; равен &#x60;rid&#x60; в ответе на метод &#x60;GET /api/v2/orders&#x60;.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

