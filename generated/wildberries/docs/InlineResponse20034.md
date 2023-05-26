# InlineResponse20034

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertId** | **int32** | Идентификатор РК | [optional] [default to null]
**Type_** | **int32** | &lt;dl&gt; &lt;dt&gt;Тип РК:&lt;/dt&gt; &lt;dd&gt;&lt;code&gt;4&lt;/code&gt; - реклама в каталоге&lt;/dd&gt; &lt;dd&gt;&lt;code&gt;5&lt;/code&gt; - реклама в карточке товара&lt;/dd&gt; &lt;dd&gt;&lt;code&gt;6&lt;/code&gt; - реклама в поиске&lt;/dd&gt; &lt;dd&gt;&lt;code&gt;7&lt;/code&gt; - реклама в рекомендациях на главной странице&lt;/dd&gt; &lt;/dl&gt;  | [optional] [default to null]
**Status** | **int32** | &lt;dl&gt; &lt;dt&gt;Статус РК:&lt;/dt&gt; &lt;dd&gt;&lt;code&gt;7&lt;/code&gt; - РК завершена&lt;/dd&gt; &lt;dd&gt;&lt;code&gt;9&lt;/code&gt; - идут показы&lt;/dd&gt; &lt;dd&gt;&lt;code&gt;11&lt;/code&gt; - РК на паузе&lt;/dd&gt; &lt;/dl&gt;  | [optional] [default to null]
**DailyBudget** | **int32** | Дневной бюджет, если не установлен, то 0 | [optional] [default to null]
**CreateTime** | **string** | Время создания РК | [optional] [default to null]
**ChangeTime** | **string** | Время последнего изменения РК | [optional] [default to null]
**StartTime** | **string** | Дата запуска РК | [optional] [default to null]
**EndTime** | **string** | Дата завершения РК | [optional] [default to null]
**Name** | **string** | Название РК | [optional] [default to null]
**Params** | [**[]InlineResponse20034Params**](inline_response_200_34_params.md) | Параметры РК | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

