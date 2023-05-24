# InlineResponse20022DataCards

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NmID** | **int32** | Артикул WB | [optional] [default to null]
**VendorCode** | **string** | Артикул продавца | [optional] [default to null]
**Object** | **string** | Категория для который создавалось КТ с данной НМ | [optional] [default to null]
**Brand** | **string** | Брэнд | [optional] [default to null]
**UpdateAt** | **string** | Дата обновления | [optional] [default to null]
**Colors** | **[]string** | Цвета номенклатуры | [optional] [default to null]
**MediaFiles** | **[]string** | Медиафайлы номенклатуры. &lt;br&gt;Фото, URL которого заканчивается на &lt;b&gt;1.jpg&lt;/b&gt; является главным в карточке.  | [optional] [default to null]
**Sizes** | [**[]InlineResponse20022DataSizes**](inline_response_200_22_data_sizes.md) | Массив размеров для номенклатуры (для безразмерного товара все равно нужно передавать данный массив с одним элементом и нулевым размером, но с ценой и баркодом)  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

