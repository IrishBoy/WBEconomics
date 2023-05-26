# ResponseFeadbackInner

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id отзыва | [optional] [default to null]
**UserName** | **string** | Имя автора отзыва | [optional] [default to null]
**Text** | **string** | Текст отзыва | [optional] [default to null]
**ProductValuation** | **int32** | Оценка товара | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) | Дата и время создания отзыва | [optional] [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) | Дата и время обновления отзыва | [optional] [default to null]
**State** | **string** | &lt;dt&gt;Статус вопроса:&lt;/dt&gt; &lt;dd&gt;&#x60;none&#x60; - не обработан (новый)&lt;/dd&gt; &lt;dd&gt;&#x60;wbRu&#x60; - обработан&lt;/dd&gt; &lt;dd&gt;&#x60;deleted&#x60; - удален&lt;/dd&gt;  | [optional] [default to null]
**Answer** | [***interface{}**](interface{}.md) | Структура ответа | [optional] [default to null]
**ProductDetails** | [***interface{}**](interface{}.md) | Структура товара | [optional] [default to null]
**PhotoLinks** | [**[]interface{}**](interface{}.md) | Массив структур фотографий | [optional] [default to null]
**Video** | [***interface{}**](interface{}.md) | Структура видео | [optional] [default to null]
**WasViewed** | **bool** | Просмотрен ли отзыв | [optional] [default to null]
**IsCreationSupplierComplaint** | **bool** | Можно ли оставить жалобу (&#x60;true&#x60; - да, &#x60;false&#x60; - нет) &lt;br&gt; Жалобу можно оставить, если с момента создания отзыва прошло менее 30 дней.  | [optional] [default to null]
**SupplierComplaint** | [***interface{}**](interface{}.md) | Жалоба на отзыв | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

