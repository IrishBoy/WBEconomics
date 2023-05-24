# InlineResponse2006Data

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImtID** | **int32** | Идентификатор карточки товара (нужен для группирования НМ в одно КТ) | [optional] [default to null]
**NmID** | **int32** | Артикул WB | [optional] [default to null]
**VendorCode** | **string** | Артикул продавца | [optional] [default to null]
**IsProhibited** | **bool** | &#x60;true&#x60; - категория карточки запрещена к реализации&lt;br&gt; &#x60;false&#x60; категория карточки разрешена к реализации  | [optional] [default to null]
**MediaFiles** | **[]string** | Медиафайлы номенклатуры. &lt;br&gt;Фото, URL которого заканчивается на &lt;b&gt;1.jpg&lt;/b&gt; является главным в карточке.  | [optional] [default to null]
**Sizes** | [**[]InlineResponse2006Sizes**](inline_response_200_6_sizes.md) |  | [optional] [default to null]
**Characteristics** | [**[]interface{}**](interface{}.md) | Массив характеристик, индивидуальный для каждой категории | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

