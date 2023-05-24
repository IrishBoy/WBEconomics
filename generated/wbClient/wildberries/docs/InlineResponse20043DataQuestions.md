# InlineResponse20043DataQuestions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id вопроса | [optional] [default to null]
**Text** | **string** | Текст вопроса | [optional] [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) | Дата и время создания вопроса | [optional] [default to null]
**State** | **string** | &lt;dt&gt;Статус вопроса:&lt;/dt&gt; &lt;dd&gt;&#x60;none&#x60; - вопрос отклонён продавцом (такой вопрос не отображается на портале покупателей)&lt;/dd&gt; &lt;dd&gt;&#x60;wbRu&#x60; - ответ предоставлен, вопрос отображается на сайте покупателей&lt;/dd&gt; &lt;dd&gt;&#x60;suppliersPortalSynch&#x60; - новый вопрос&lt;/dd&gt;  | [optional] [default to null]
**Answer** | [***InlineResponse20043DataAnswer**](inline_response_200_43_data_answer.md) |  | [optional] [default to null]
**ProductDetails** | [***InlineResponse20043DataProductDetails**](inline_response_200_43_data_productDetails.md) |  | [optional] [default to null]
**WasViewed** | **bool** | Просмотрен ли вопрос | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

