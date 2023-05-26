# InlineResponse2004DataCards

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sizes** | [**[]InlineResponse2004DataSizes**](inline_response_200_4_data_sizes.md) | Массив размеров для номенклатуры (для безразмерного товара все равно нужно передавать данный массив с одним элементом и нулевым размером, но с ценой и баркодом)  | [optional] [default to null]
**MediaFiles** | **[]string** | Медиафайлы номенклатуры. &lt;br&gt;Фото, URL которого заканчивается на &lt;b&gt;1.jpg&lt;/b&gt; является главным в карточке.  | [optional] [default to null]
**Colors** | **[]string** | Цвета номенклатуры | [optional] [default to null]
**UpdateAt** | **string** | Дата обновления | [optional] [default to null]
**VendorCode** | **string** | Артикул продавца | [optional] [default to null]
**Brand** | **string** | Брэнд | [optional] [default to null]
**Object** | **string** | Категория для который создавалось КТ с данной НМ | [optional] [default to null]
**NmID** | **int32** | Артикул WB | [optional] [default to null]
**ImtID** | **int32** | Идентификатор карточки товара (нужен для группирования НМ в одну КТ) | [optional] [default to null]
**ChrtID** | **int32** | Числовой идентификатор размера для данной номенклатуры Wildberries | [optional] [default to null]
**IsProhibited** | **bool** | &#x60;true&#x60; - категория карточки запрещена к реализации&lt;br&gt; &#x60;false&#x60; - категория карточки разрешена к реализации  | [optional] [default to null]
**Tags** | **[]string** | Массив тегов | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

