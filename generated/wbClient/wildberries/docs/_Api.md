# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentV1BarcodesPost**](_Api.md#ContentV1BarcodesPost) | **Post** /content/v1/barcodes | Генерация баркодов
[**ContentV1CardsCursorListPost**](_Api.md#ContentV1CardsCursorListPost) | **Post** /content/v1/cards/cursor/list | Список НМ
[**ContentV1CardsErrorListGet**](_Api.md#ContentV1CardsErrorListGet) | **Get** /content/v1/cards/error/list | Список несозданных НМ с ошибками
[**ContentV1CardsFilterPost**](_Api.md#ContentV1CardsFilterPost) | **Post** /content/v1/cards/filter | Получение КТ по артикулам продавца
[**ContentV1CardsLimitsGet**](_Api.md#ContentV1CardsLimitsGet) | **Get** /content/v1/cards/limits | Лимиты по КТ
[**ContentV1CardsTrashListPost**](_Api.md#ContentV1CardsTrashListPost) | **Post** /content/v1/cards/trash/list | Список НМ, находящихся в корзине
[**ContentV1CardsUpdatePost**](_Api.md#ContentV1CardsUpdatePost) | **Post** /content/v1/cards/update | Редактирование КТ
[**ContentV1CardsUploadAddPost**](_Api.md#ContentV1CardsUploadAddPost) | **Post** /content/v1/cards/upload/add | Добавление НМ к КТ
[**ContentV1CardsUploadPost**](_Api.md#ContentV1CardsUploadPost) | **Post** /content/v1/cards/upload | Создание  КТ
[**ContentV1DirectoryBrandsGet**](_Api.md#ContentV1DirectoryBrandsGet) | **Get** /content/v1/directory/brands | Бренд
[**ContentV1DirectoryColorsGet**](_Api.md#ContentV1DirectoryColorsGet) | **Get** /content/v1/directory/colors | Цвет
[**ContentV1DirectoryCountriesGet**](_Api.md#ContentV1DirectoryCountriesGet) | **Get** /content/v1/directory/countries | Страна Производства
[**ContentV1DirectoryKindsGet**](_Api.md#ContentV1DirectoryKindsGet) | **Get** /content/v1/directory/kinds | Пол
[**ContentV1DirectorySeasonsGet**](_Api.md#ContentV1DirectorySeasonsGet) | **Get** /content/v1/directory/seasons | Сезон
[**ContentV1DirectoryTnvedGet**](_Api.md#ContentV1DirectoryTnvedGet) | **Get** /content/v1/directory/tnved | ТНВЭД код
[**ContentV1MediaFilePost**](_Api.md#ContentV1MediaFilePost) | **Post** /content/v1/media/file | Добавление медиа контента в КТ
[**ContentV1MediaSavePost**](_Api.md#ContentV1MediaSavePost) | **Post** /content/v1/media/save | Изменение медиа контента КТ
[**ContentV1ObjectAllGet**](_Api.md#ContentV1ObjectAllGet) | **Get** /content/v1/object/all | Категория товаров
[**ContentV1ObjectCharacteristicsListFilterGet**](_Api.md#ContentV1ObjectCharacteristicsListFilterGet) | **Get** /content/v1/object/characteristics/list/filter | Характеристики для создания КТ по всем подкатегориям
[**ContentV1ObjectCharacteristicsObjectNameGet**](_Api.md#ContentV1ObjectCharacteristicsObjectNameGet) | **Get** /content/v1/object/characteristics/{objectName} | Характеристики для создания КТ для категории товара
[**ContentV1ObjectParentAllGet**](_Api.md#ContentV1ObjectParentAllGet) | **Get** /content/v1/object/parent/all | Родительские категории товаров
[**ContentV1TagIdDelete**](_Api.md#ContentV1TagIdDelete) | **Delete** /content/v1/tag/{id} | Удаление тега
[**ContentV1TagIdPatch**](_Api.md#ContentV1TagIdPatch) | **Patch** /content/v1/tag/{id} | Изменение тега
[**ContentV1TagNomenclatureLinkPost**](_Api.md#ContentV1TagNomenclatureLinkPost) | **Post** /content/v1/tag/nomenclature/link | Управление тегами в КТ
[**ContentV1TagPost**](_Api.md#ContentV1TagPost) | **Post** /content/v1/tag | Создание тега
[**ContentV1TagsGet**](_Api.md#ContentV1TagsGet) | **Get** /content/v1/tags | Список тегов
[**PublicApiV1RevokeDiscountsPost**](_Api.md#PublicApiV1RevokeDiscountsPost) | **Post** /public/api/v1/revokeDiscounts | Сброс скидок для номенклатур
[**PublicApiV1RevokePromocodesPost**](_Api.md#PublicApiV1RevokePromocodesPost) | **Post** /public/api/v1/revokePromocodes | Сброс промокодов для номенклатур
[**PublicApiV1UpdateDiscountsPost**](_Api.md#PublicApiV1UpdateDiscountsPost) | **Post** /public/api/v1/updateDiscounts | Установка скидок
[**PublicApiV1UpdatePromocodesPost**](_Api.md#PublicApiV1UpdatePromocodesPost) | **Post** /public/api/v1/updatePromocodes | Установка промокодов для номенклатур

# **ContentV1BarcodesPost**
> InlineResponse2003 ContentV1BarcodesPost(ctx, body)
Генерация баркодов

Метод позволяет сгенерировать массив уникальных баркодов для создания размера НМ в КТ.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1BarcodesBody**](V1BarcodesBody.md)|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1CardsCursorListPost**
> InlineResponse2004 ContentV1CardsCursorListPost(ctx, body)
Список НМ

Метод позволяет получить список созданых НМ по фильтру (баркод, артикул продавца, артикул WB (nmId), тег) с пагинацией и сортировкой. <br> <br>Порядок работы с `cursor/list`: <br> Чтобы получить полный список номенклатур, <b>если их > 1000</b>, необходимо воспользоваться пагинацией.   <ol>     <li>Cделать первый запрос (все указанные ниже параметры обязательны): <br>       <pre style=\"background-color: rgb(38 50 56 / 5%); color: #e53935\">         {           \"sort\": {               \"cursor\": {                   \"limit\": 1000               },               \"filter\": {                   \"withPhoto\": -1               }           }         }</pre>       По желанию можно добавить поиск по <code>\"textSearch\"</code> и сортировку.       <pre style=\"background-color: rgb(38 50 56 / 5%); color: #e53935\">         \"sort\": {           \"sortColumn\": \"\",           \"ascending\": false         }</pre>     </li>     <li>Пройти в конец полученного списка номенклатур, скопировать из <code>cursor</code> две строки:       <ul>         <li><code>\"updatedAt\": \"***\"</code>,</li>         <li><code>\"nmID\": ***</code>,</li>       </ul>     <li>Вставить скопированные строки в <code>cursor</code> запроса, повторить вызов метода. </li>     <li>Повторять пункты <b>2</b> и <b>3</b>, пока <code>total</code> в ответе не станет меньше чем <code>limit</code> в запросе.       <br>Это будет означать, что Вы получили все карточки.   </ol> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CursorListBody**](CursorListBody.md)|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1CardsErrorListGet**
> InlineResponse2005 ContentV1CardsErrorListGet(ctx, )
Список несозданных НМ с ошибками

Метод позволяет получить список НМ и список ошибок которые произошли во время создания КТ. <br>`ВАЖНО`: Для того чтобы убрать НМ из ошибочных, надо повторно сделать запрос с исправленными ошибками на создание КТ. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1CardsFilterPost**
> InlineResponse2006 ContentV1CardsFilterPost(ctx, body)
Получение КТ по артикулам продавца

Метод позволяет получить полную информацию по КТ с помощью артикула(-ов) продавца. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CardsFilterBody**](CardsFilterBody.md)|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1CardsLimitsGet**
> InlineResponse2007 ContentV1CardsLimitsGet(ctx, )
Лимиты по КТ

Метод позволяет получить отдельно бесплатные и платные лимиты продавца на создание карточек товаров.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1CardsTrashListPost**
> InlineResponse20022 ContentV1CardsTrashListPost(ctx, body)
Список НМ, находящихся в корзине

Метод позволяет получить список НМ, находящихся в корзине.<br> Метод позволяет получить список НМ, которые находятся в корзине по фильтру (баркод (<code>skus</code>), артикул продавца(<code>vendorCode</code>), артикул WB(<code>nmID</code>)) с пагинацией и сортировкой. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TrashListBody**](TrashListBody.md)|  | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1CardsUpdatePost**
> InlineResponse2002 ContentV1CardsUpdatePost(ctx, body)
Редактирование КТ

Метод позволяет отредактировать несколько карточек за раз. <br> Редактирование КТ происходит асинхронно, после отправки запрос становится в очередь на обработку. <br> `Важно`: Баркоды (skus) не подлежат удалению или замене. Попытка заменить существующий баркод приведет к добавлению нового баркода к существующему. <br> Если запрос прошел успешно, а информация в карточке не обновилась, значит были допущены ошибки и карточка попала в \"Черновики\" (метод `cards/error/list`) с описанием ошибок. <br>Необходимо исправить ошибки в запросе и отправить его повторно. <br> <br> Для успешного обновления карточки рекомендуем Вам придерживаться следующего порядка действий: <br> 1. Сначала существующую карточку необходимо запросить методом cards/filter. <br> 2. Забираем из ответа массив data. <br> 3. В этом массиве вносим необходимые изменения и отправляем его в cards/update <br> <br>За раз можно отредактировать 1000 КТ по 5 НМ в каждой. <br> <br> `Важно`: Изменение цен данным методом невозможно, используйте метод Загрузка цен, раздел документации Цены. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]CardsUpdateBody**](cards_update_body.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1CardsUploadAddPost**
> InlineResponse2002 ContentV1CardsUploadAddPost(ctx, body)
Добавление НМ к КТ

Метод позволяет добавить к карточке товара новую номенклатуру. <br>Добавление НМ к КТ происходит асинхронно, после отправки запрос становится в очередь на обработку. <br>`Важно`: Если после успешной отправки запроса номенклатура не создалась, то необходимо проверить раздел \"Список несозданных НМ с ошибками\". Для того чтобы убрать НМ из ошибочных, необходимо повторно сделать запрос с исправленными ошибками. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UploadAddBody**](UploadAddBody.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1CardsUploadPost**
> InlineResponse2002 ContentV1CardsUploadPost(ctx, body)
Создание  КТ

Создание карточки товара происходит асинхронно, при отправке запроса на создание КТ ваш запрос становится в очередь на создание КТ. <br>`ПРИМЕЧАНИЕ`: Карточка товара считается `созданной`, если успешно создалась хотя бы одна `НМ`. <br>`ВАЖНО`: Если во время обработки запроса в очереди выявляются ошибки, то НМ считается ошибочной. <br>Если запрос на создание прошел успешно, а карточка не создалась, то необходимо в первую очередь проверить наличие карточки в методе `cards/error/list`. Если карточка попала в ответ к этому методу, то необходимо исправить описанные ошибки в запросе на создание карточки и отправить его повторно. <br>За раз можно создать 1000 КТ по 5 НМ в каждой.  <br>Если Вам требуется больше 5 НМ в КТ, то после создания карточки Вы можете добавить их методом \"Добавление НМ к КТ\". 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CardsUploadBody**](CardsUploadBody.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1DirectoryBrandsGet**
> InlineResponse20016 ContentV1DirectoryBrandsGet(ctx, optional)
Бренд

Получение значения характеристики Бренд.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***_ApiContentV1DirectoryBrandsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a _ApiContentV1DirectoryBrandsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **optional.Int32**| Количество запрашиваемых значений (максимум 5000) | 
 **pattern** | **optional.String**| Поиск по наименованию значения характеристики | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1DirectoryColorsGet**
> InlineResponse20012 ContentV1DirectoryColorsGet(ctx, )
Цвет

Получение значения характеристики цвет.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1DirectoryCountriesGet**
> InlineResponse20014 ContentV1DirectoryCountriesGet(ctx, )
Страна Производства

Получение значения характеристики Страна Производства.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1DirectoryKindsGet**
> InlineResponse20013 ContentV1DirectoryKindsGet(ctx, )
Пол

Получение значения характеристики пол.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1DirectorySeasonsGet**
> InlineResponse20015 ContentV1DirectorySeasonsGet(ctx, )
Сезон

Получение значения характеристики Сезон.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1DirectoryTnvedGet**
> InlineResponse20017 ContentV1DirectoryTnvedGet(ctx, optional)
ТНВЭД код

С помощью данного метода можно получить список ТНВЭД кодов по имени категории и фильтру по тнвэд коду.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***_ApiContentV1DirectoryTnvedGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a _ApiContentV1DirectoryTnvedGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectName** | **optional.String**| Поиск по наименованию категории | 
 **tnvedsLike** | **optional.String**| Поиск по коду ТНВЭД | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1MediaFilePost**
> InlineResponse20018 ContentV1MediaFilePost(ctx, uploadfile, xVendorCode, xPhotoNumber)
Добавление медиа контента в КТ

Метод позволяет загрузить и добавить один медиафайл за запрос, к НМ в КТ. <br>Требования к медиафайлам: <br>`Фото`: минимальное разрешение – 450х450. <br>Максимально допустимое количество фото в КТ 30. <br>Допустимые форматы изображений - jpg и png. <br>`Видео`: максимальный размер 50 мб. Форматы MOV, MP4. <br>Максимально допустимое количество видео в КТ 1. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uploadfile** | ***os.File*****os.File**|  | 
  **xVendorCode** | **string**| Артикул продавца | 
  **xPhotoNumber** | **int32**| Номер медиафайла на загрузку. &lt;b&gt;Начинать с 1&lt;/b&gt;. &lt;br&gt;Чтобы добавить фото к уже загруженным в НМ, номер медиафайла должен быть больше кол-ва загруженных в НМ медиафайлов.  | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1MediaSavePost**
> InlineResponse20018 ContentV1MediaSavePost(ctx, body)
Изменение медиа контента КТ

Метод позволяет изменить порядок изображений или удалить медиафайлы с НМ в КТ, а также загрузить изображения в НМ со сторонних ресурсов по URL. <br>Текущие изображения заменяются на переданные в массиве data. <br> <br>Требования к медиафайлам: <br>`Фото`: минимальное разрешение – 450х450. <br>Максимально допустимое количество фото в КТ 30.  <br>Допустимые форматы изображений - jpg и png. <br> <br>Если хотя бы одно изображение в запросе не соответствует требованиям к медиафайлам, то даже при коде ответа 200 ни одно изображение не загрузится в КТ. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MediaSaveBody**](MediaSaveBody.md)|  | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1ObjectAllGet**
> InlineResponse2008 ContentV1ObjectAllGet(ctx, optional)
Категория товаров

С помощью данного метода можно получить список категорий товаров по текстовому фильтру (названию категории). <br> <br> Чтобы получить список всех категорий необходимо указать `top=8000`, к примеру. <br> <br> По состоянию на `27.03.2023` в списке `7440` категорий. Количество доступных категорий может меняться.       

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***_ApiContentV1ObjectAllGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a _ApiContentV1ObjectAllGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Поиск по названию категории | 
 **top** | **optional.Int32**| Количество запрашиваемых значений | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1ObjectCharacteristicsListFilterGet**
> InlineResponse20010 ContentV1ObjectCharacteristicsListFilterGet(ctx, optional)
Характеристики для создания КТ по всем подкатегориям

С помощью данного метода можно получить список характеристик, которые можно или нужно заполнить при создании КТ в подкатегории определенной родительской категории.   <br>   <br>Характеристики с  `charcType=4`, имеющие единственное значение, <b>указывать строго без массива</b>, в противном случае карточка не будет создана:   <br><b>Правильно:</b>       <br>`{\"Высота упаковки\": 4}`   <br><b>Не правильно:</b>       <br>`{\"Высота упаковки\": [4]}` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***_ApiContentV1ObjectCharacteristicsListFilterGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a _ApiContentV1ObjectCharacteristicsListFilterGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Поиск по родительской категории. | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1ObjectCharacteristicsObjectNameGet**
> InlineResponse20011 ContentV1ObjectCharacteristicsObjectNameGet(ctx, objectName)
Характеристики для создания КТ для категории товара

С помощью данного метода можно получить список характеристик, которые можно или нужно заполнить при создании КТ для определенной категории товаров. <br> <br> `Важно`: обязательная к заполнению характеристика при создании карточки любого товара - `Предмет`. <br>Значение характеристики `Предмет` соответствует значению параметра `objectName` в запросе. <br> <br>Характеристики с  `charcType=4`, имеющие единственное значение, <b>указывать строго без массива</b>, в противном случае карточка не будет создана: <br><b>Правильно:</b>     <br>`{\"Высота упаковки\": 4}` <br><b>Не правильно:</b>     <br>`{\"Высота упаковки\": [4]}` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectName** | **string**| Поиск по наименованию категории | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1ObjectParentAllGet**
> InlineResponse2009 ContentV1ObjectParentAllGet(ctx, )
Родительские категории товаров

С помощью данного метода можно получить список всех родительских категорий товаров.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1TagIdDelete**
> InlineResponse20020 ContentV1TagIdDelete(ctx, id)
Удаление тега

Метод позволяет удалить тег.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Числовой идентификатор тега | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1TagIdPatch**
> InlineResponse20021 ContentV1TagIdPatch(ctx, body, id)
Изменение тега

Метод позволяет изменять информацию о теге (имя и цвет).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagIdBody**](TagIdBody.md)|  | 
  **id** | **int32**| Числовой идентификатор тега | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1TagNomenclatureLinkPost**
> ResponseOk200 ContentV1TagNomenclatureLinkPost(ctx, body)
Управление тегами в КТ

Метод позволяет добавить теги к КТ и снять их с КТ.<br> При снятии тега с КТ сам тег не удаляется.<br> К карточке можно добавить 8 тегов. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NomenclatureLinkBody**](NomenclatureLinkBody.md)|  | 

### Return type

[**ResponseOk200**](responseOK200.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1TagPost**
> ResponseOk200 ContentV1TagPost(ctx, body)
Создание тега

Метод позволяет создать тег.<br> Завести можно 8 тегов.<br> Максимальная длина тега 15 символов. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1TagBody**](V1TagBody.md)|  | 

### Return type

[**ResponseOk200**](responseOK200.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentV1TagsGet**
> InlineResponse20019 ContentV1TagsGet(ctx, )
Список тегов

Метод позволяет получить список существующих тегов продавца.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicApiV1RevokeDiscountsPost**
> string PublicApiV1RevokeDiscountsPost(ctx, body)
Сброс скидок для номенклатур

Сброс скидок для номенклатур. При первом запросе ответ будет пустым. При повторной попытке сбросить скидку вернет JSON с id первичного запроса.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int32**](int32.md)| Перечень номенклатур к отмене скидок | 

### Return type

**string**

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicApiV1RevokePromocodesPost**
> string PublicApiV1RevokePromocodesPost(ctx, body)
Сброс промокодов для номенклатур

Сброс промокодов для номенклатур

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int32**](int32.md)| Перечень номенклатур к отмене промокодов | 

### Return type

**string**

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicApiV1UpdateDiscountsPost**
> string PublicApiV1UpdateDiscountsPost(ctx, body, optional)
Установка скидок

Установка скидок для номенклатур. Максимальное количество номенклатур на запрос - 1000

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]interface{}**](interface{}.md)| Перечень номенклатур | 
 **optional** | ***_ApiPublicApiV1UpdateDiscountsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a _ApiPublicApiV1UpdateDiscountsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activateFrom** | **optional.**| Дата активации скидки в формате &#x60;YYYY-MM-DD&#x60; или &#x60;YYYY-MM-DD HH:MM:SS&#x60;. Если не указывать, скидка начнет действовать сразу. | 

### Return type

**string**

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicApiV1UpdatePromocodesPost**
> string PublicApiV1UpdatePromocodesPost(ctx, body, optional)
Установка промокодов для номенклатур

Установка промокодов для номенклатур. Максимальное количество номенклатур на запрос - 1000

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]interface{}**](interface{}.md)| Перечень номенклатур | 
 **optional** | ***_ApiPublicApiV1UpdatePromocodesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a _ApiPublicApiV1UpdatePromocodesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activateFrom** | **optional.**| Дата активации промокада в формате &#x60;YYYY-MM-DD&#x60; или &#x60;YYYY-MM-DD HH:MM:SS&#x60;. Если не указывать, промокод начнет действовать сразу | 

### Return type

**string**

### Authorization

[HeaderApiKey](../README.md#HeaderApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

