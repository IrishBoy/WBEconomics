# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2WarehousesGet**](Marketplace_Api.md#ApiV2WarehousesGet) | **Get** /api/v2/warehouses | Cписок складов
[**ApiV3OrdersGet**](Marketplace_Api.md#ApiV3OrdersGet) | **Get** /api/v3/orders | Получить информацию по сборочным заданиям
[**ApiV3OrdersNewGet**](Marketplace_Api.md#ApiV3OrdersNewGet) | **Get** /api/v3/orders/new | Получить список новых сборочных заданий
[**ApiV3OrdersOrderCancelPatch**](Marketplace_Api.md#ApiV3OrdersOrderCancelPatch) | **Patch** /api/v3/orders/{order}/cancel | Отменить сборочное задание
[**ApiV3OrdersOrderMetaSgtinPost**](Marketplace_Api.md#ApiV3OrdersOrderMetaSgtinPost) | **Post** /api/v3/orders/{order}/meta/sgtin | Закрепить за сборочным заданием КиЗ (маркировку Честного знака)
[**ApiV3OrdersStatusPost**](Marketplace_Api.md#ApiV3OrdersStatusPost) | **Post** /api/v3/orders/status | Получить статусы сборочных заданий
[**ApiV3OrdersStickersPost**](Marketplace_Api.md#ApiV3OrdersStickersPost) | **Post** /api/v3/orders/stickers | Получить этикетки для сборочных заданий
[**ApiV3StocksWarehouseDelete**](Marketplace_Api.md#ApiV3StocksWarehouseDelete) | **Delete** /api/v3/stocks/{warehouse} | Удалить остатки товаров
[**ApiV3StocksWarehousePost**](Marketplace_Api.md#ApiV3StocksWarehousePost) | **Post** /api/v3/stocks/{warehouse} | Получить остатки товаров
[**ApiV3StocksWarehousePut**](Marketplace_Api.md#ApiV3StocksWarehousePut) | **Put** /api/v3/stocks/{warehouse} | Обновить остатки товаров
[**ApiV3SuppliesGet**](Marketplace_Api.md#ApiV3SuppliesGet) | **Get** /api/v3/supplies | Получить список поставок
[**ApiV3SuppliesPost**](Marketplace_Api.md#ApiV3SuppliesPost) | **Post** /api/v3/supplies | Создать новую поставку
[**ApiV3SuppliesSupplyBarcodeGet**](Marketplace_Api.md#ApiV3SuppliesSupplyBarcodeGet) | **Get** /api/v3/supplies/{supply}/barcode | Получить QR поставки
[**ApiV3SuppliesSupplyDelete**](Marketplace_Api.md#ApiV3SuppliesSupplyDelete) | **Delete** /api/v3/supplies/{supply} | Удалить поставку
[**ApiV3SuppliesSupplyDeliverPatch**](Marketplace_Api.md#ApiV3SuppliesSupplyDeliverPatch) | **Patch** /api/v3/supplies/{supply}/deliver | Передать поставку в доставку
[**ApiV3SuppliesSupplyGet**](Marketplace_Api.md#ApiV3SuppliesSupplyGet) | **Get** /api/v3/supplies/{supply} | Получить информацию о поставке
[**ApiV3SuppliesSupplyOrdersGet**](Marketplace_Api.md#ApiV3SuppliesSupplyOrdersGet) | **Get** /api/v3/supplies/{supply}/orders | Получить сборочные задания в поставке
[**ApiV3SuppliesSupplyOrdersOrderPatch**](Marketplace_Api.md#ApiV3SuppliesSupplyOrdersOrderPatch) | **Patch** /api/v3/supplies/{supply}/orders/{order} | Добавить к поставке сборочное задание

# **ApiV2WarehousesGet**
> []InlineResponse20023 ApiV2WarehousesGet(ctx, )
Cписок складов

Cписок складов продавца.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse20023**](inline_response_200_23.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3OrdersGet**
> InlineResponse20029 ApiV3OrdersGet(ctx, limit, next, optional)
Получить информацию по сборочным заданиям

Возвращает информацию по сборочным заданиям без их актуального статуса.  Данные по сборочному заданию, возвращающиеся в данном методе, не меняются. Рекомендуется использовать для получения исторических данных. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **limit** | **int32**| Параметр пагинации. Устанавливает предельное количество возвращаемых данных. | 
  **next** | **int64**| Параметр пагинации. Устанавливает значение, с которого надо получить следующий пакет данных. Для получения полного списка данных должен быть равен 0 в первом запросе. Для следующих запросов необходимо брать значения из одноимённого поля в ответе. | 
 **optional** | ***Marketplace_ApiApiV3OrdersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Marketplace_ApiApiV3OrdersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dateFrom** | **optional.Int32**| Дата начала периода в формате Unix timestamp. Необязательный параметр. | 
 **dateTo** | **optional.Int32**| Дата конца периода в формате Unix timestamp. Необязательный параметр. | 

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3OrdersNewGet**
> InlineResponse20030 ApiV3OrdersNewGet(ctx, )
Получить список новых сборочных заданий

Возвращает список всех новых сборочных заданий у продавца на данный момент. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3OrdersOrderCancelPatch**
> ApiV3OrdersOrderCancelPatch(ctx, order)
Отменить сборочное задание

Переводит сборочное задание в статус **cancel** (\"Отменено продавцом\").

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **order** | **int64**| Идентификатор сборочного задания | 

### Return type

 (empty response body)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3OrdersOrderMetaSgtinPost**
> ApiV3OrdersOrderMetaSgtinPost(ctx, order, optional)
Закрепить за сборочным заданием КиЗ (маркировку Честного знака)

Метод позволяет закрепить за сборочным заданием КиЗ (маркировку Честного знака). У одного сборочного заказа не может быть больше 3 маркировок. <br> <br> Параметры `sid`, `numerator`, `denominator` опциональны. Заполняются в зависимости от специфики товара. <br> <br> `Важно!` Получить загруженные КиЗ можно только в личном кабинете. Для этого необходимо: <ol>  <li>Зайти в раздел Маркетплейс - Сборочные задания</li>  <li>Пройти в любую из перечисленных вкладок (<code>На сборке</code>, <code>В доставке</code>, <code>Архив</code>)</li>  <li>Зайти в детализацию поставки</li>  <li>Нажать кнопку \"Выгрузить в Excel\"</li>  <li>В полученном файле открыть лист КИЗы</li> </ol> Получить загруженные КиЗ можно по всем заказам, кроме: <code>Новые</code>, <code>Отменено продавцом</code>.<br> <br> С правилами работы с КиЗ можно ознакомиться тут: https://честныйзнак.рф <br> <br> О реализации API-функционала для получения загруженных КиЗ будет сообщено в разделе Новости, на портале продавцов. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **order** | **int64**| Идентификатор сборочного задания | 
 **optional** | ***Marketplace_ApiApiV3OrdersOrderMetaSgtinPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Marketplace_ApiApiV3OrdersOrderMetaSgtinPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MetaSgtinBody**](MetaSgtinBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3OrdersStatusPost**
> InlineResponse20028 ApiV3OrdersStatusPost(ctx, optional)
Получить статусы сборочных заданий

Возвращает статусы сборочных заданий по переданному списку идентификаторов сборочных заданий.  **supplierStatus** - статус сборочного задания, триггером изменения которого является сам продавец.<br> Возможны следующие значения данного поля: | Статус   | Описание            | Как перевести сборочное задание в данный статус | | -------  | ---------           | --------------------------------------| | new      | Новое сборочное задание |           | confirm  | На сборке            | При добавлении сборочного задания к поставке **PATCH** *_/api/v3/supplies/{supply}/orders/{order}* | complete | В доставке           | При переводе в доставку соответствующей поставки **PATCH** *_/api/v3/supplies/{supply}/deliver* | cancel   | Отменено продавцом   | **PATCH** *_/api/v3/orders/{order}/cancel*   **wbStatus** - статус сборочного задания в системе Wildberries.<br> Возможны следующие значения данного поля: - **waiting** - сборочное задание в работе - **sorted** - сборочное задание отсортировано - **sold** - сборочное задание получено покупателем - **canceled** - отмена сборочного задания - **canceled_by_client** - отмена сборочного задания покупателем - **defect** - отмена сборочного задания по причине брака - **ready_for_pickup** - сборочное задание прибыло на ПВЗ <span class=\"new\">new</span> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Marketplace_ApiApiV3OrdersStatusPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Marketplace_ApiApiV3OrdersStatusPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OrdersStatusBody**](OrdersStatusBody.md)|  | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3OrdersStickersPost**
> InlineResponse20031 ApiV3OrdersStickersPost(ctx, type_, width, height, optional)
Получить этикетки для сборочных заданий

Возвращает список этикеток по переданному массиву сборочных заданий. Можно запросить этикетку в формате svg, zplv (вертикальный), zplh (горизонтальный), png.  **Ограничения при работе с методом**: - Нельзя запросить больше 100 этикеток за раз (не более 100 идентификаторов сборочных заданий в запросе). - Метод возвращает этикетки только для сборочных заданий, находящихся на сборке (в статусе **confirm**). - Доступные размеры: <dd>580x400 пикселей, при параметрах width = 58, height = 40</dd> <dd>400x300 пикселей, при параметрах width = 40, height = 30</dd> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| Тип этикетки | 
  **width** | **int32**| Ширина этикетки | 
  **height** | **int32**| Высота этикетки | 
 **optional** | ***Marketplace_ApiApiV3OrdersStickersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Marketplace_ApiApiV3OrdersStickersPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of OrdersStickersBody**](OrdersStickersBody.md)|  | 

### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3StocksWarehouseDelete**
> ApiV3StocksWarehouseDelete(ctx, body, warehouse)
Удалить остатки товаров

Удаляет остатки товаров. **Внимание! Действие необратимо**. Удаленный остаток будет необходимо загрузить повторно для возобновления продаж.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StocksWarehouseBody2**](StocksWarehouseBody2.md)|  | 
  **warehouse** | **int32**| Идентификатор склада продавца | 

### Return type

 (empty response body)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3StocksWarehousePost**
> InlineResponse20024 ApiV3StocksWarehousePost(ctx, body, warehouse)
Получить остатки товаров

Возвращает остатки товаров.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StocksWarehouseBody1**](StocksWarehouseBody1.md)|  | 
  **warehouse** | **int32**| Идентификатор склада продавца | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3StocksWarehousePut**
> ApiV3StocksWarehousePut(ctx, warehouse, optional)
Обновить остатки товаров

Обновляет остатки товаров.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **warehouse** | **int32**| Идентификатор склада продавца | 
 **optional** | ***Marketplace_ApiApiV3StocksWarehousePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Marketplace_ApiApiV3StocksWarehousePutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of StocksWarehouseBody**](StocksWarehouseBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3SuppliesGet**
> InlineResponse20025 ApiV3SuppliesGet(ctx, limit, next)
Получить список поставок

Возвращает список поставок

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **limit** | **int32**| Параметр пагинации. Устанавливает предельное количество возвращаемых данных. | 
  **next** | **int64**| Параметр пагинации. Устанавливает значение, с которого надо получить следующий пакет данных. Для получения полного списка данных должен быть равен 0 в первом запросе. Для следующих запросов необходимо брать значения из одноимённого поля в ответе. | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3SuppliesPost**
> InlineResponse201 ApiV3SuppliesPost(ctx, body)
Создать новую поставку

**Ограничения работы с поставками**:  - Только для сборочных заданий по схеме \"Везу на склад WB\" - При добавлении в поставку все передаваемые сборочные задания в статусе **new** (\"Новое\") будут автоматически переведены в статус **confirm** (\"На сборке\"). - Обратите внимание, что если вы переведёте сборочное задание в статус **cancel** (\"Отмена продавцом\"), то сборочное задание автоматически удалится из поставки, если было прикреплено к ней. - Поставку можно собрать только из одного типа сборочных заданий: сКГТ (isLargeCargo = true) или обычный (isLargeCargo = false). Новая поставка не обладает сКГТ-признаком. При добавлении первого заказа в поставку она приобретает сКГТ-признак добавляемого товара в заказе. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V3SuppliesBody**](V3SuppliesBody.md)|  | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3SuppliesSupplyBarcodeGet**
> InlineResponse20027 ApiV3SuppliesSupplyBarcodeGet(ctx, supply, type_)
Получить QR поставки

Возвращает QR в svg, zplv (вертикальный), zplh (горизонтальный), png. <br> Можно получить, только если поставка передана в доставку. <dt>Доступные размеры:</dt> <dd>580x400 пикселей 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supply** | **string**| Идентификатор поставки | 
  **type_** | **string**| Тип этикетки | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3SuppliesSupplyDelete**
> ApiV3SuppliesSupplyDelete(ctx, supply)
Удалить поставку

Удаляет поставку, если она активна и за ней не закреплено ни одно сборочное задание.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supply** | **string**| Идентификатор поставки | 

### Return type

 (empty response body)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3SuppliesSupplyDeliverPatch**
> ApiV3SuppliesSupplyDeliverPatch(ctx, supply)
Передать поставку в доставку

Закрывает поставку и переводит все сборочные задания в ней в статус **complete** (\"В доставке\"). После закрытия поставки новые сборочные задания к ней добавить будет невозможно. Передать поставку в доставку можно только при наличии в ней хотя бы одного сборочного задания. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supply** | **string**| Идентификатор поставки | 

### Return type

 (empty response body)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3SuppliesSupplyGet**
> Supply ApiV3SuppliesSupplyGet(ctx, supply)
Получить информацию о поставке

Возвращает информацию о поставке

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supply** | **string**| Идентификатор поставки | 

### Return type

[**Supply**](Supply.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3SuppliesSupplyOrdersGet**
> InlineResponse20026 ApiV3SuppliesSupplyOrdersGet(ctx, supply)
Получить сборочные задания в поставке

Возвращает сборочные задания, закреплённые за поставкой.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supply** | **string**| Идентификатор поставки | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3SuppliesSupplyOrdersOrderPatch**
> ApiV3SuppliesSupplyOrdersOrderPatch(ctx, supply, order)
Добавить к поставке сборочное задание

Добавляет к поставке сборочное задание и переводит его в статус **confirm** (\"На сборке\").  Также может перемещать сборочное задание между активными поставками, либо из закрытой в активную при условии, что сборочное задание требует повторной отгрузки. Добавить в поставку возможно только задания с соответствующим сКГТ-признаком (isLargeCargo), либо если поставке ещё не присвоен сКГТ-признак (isLargeCargo = null). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supply** | **string**| Идентификатор поставки | 
  **order** | **int64**| Идентификатор сборочного задания | 

### Return type

 (empty response body)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

