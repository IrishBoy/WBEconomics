# CardsUpdateBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImtID** | **int32** | Идентификатор карточки товара | [optional] [default to null]
**NmID** | **int32** | Артикул WB | [default to null]
**VendorCode** | **string** | Артикул продавца | [default to null]
**Sizes** | [**[]Contentv1cardsupdateSizes**](contentv1cardsupdate_sizes.md) | Массив размеров для номенклатуры (для безразмерного товара все равно нужно передавать данный массив без параметров (&#x60;wbSize&#x60; и &#x60;techSize&#x60;), но с &#x60;chrtID&#x60; и баркодом) | [default to null]
**Characteristics** | [**[]interface{}**](interface{}.md) | Массив характеристик, индивидуальный для каждой категории | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

