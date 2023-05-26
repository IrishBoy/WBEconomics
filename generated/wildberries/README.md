# Go API client for wildberries

# Общее описание <style> .langChange {   text-align: left;   padding: 2px 222px 2px 18px; } </style> <style> .new {   display: inline-block;   padding: 2px 8px;   margin: 0;   background-color: #ffa500;   color: #ffffff;   font-size: 13px;   vertical-align: middle;   line-height: 1.6;   border-radius: 4px;   font-weight: 600;   font-size: 12px; } </style> <style> .newM {   padding: 5px 10px;   display: inline-block;   background-color: #ffa500;   color: #ffffff;   cursor: pointer;   text-align: center;   outline: none;   margin: 0 5px 5px 5px;   #border: 1px solid #07090b;   border-radius: 5px;   min-width: 60px;   font-size: 0.9em;   font-weight: bold; } </style>  Wildberries API предоставляет продавцам возможность управления магазином и получения оперативной и статистической информации по протоколу HTTP RestAPI. <br> Описание API предоставляется в формате [Swagger](https://swagger.io/) (Open API) и может быть использовано для импорта в другие инструменты (такие как PostMan) или генерации клиентского кода на различных языках программирования с помощью [Swagger CodeGen](https://swagger.io/tools/swagger-codegen/)  <ul> <li> Описание в оригинальном swagger-формате <a href=\"/swagger\">swagger</a> <li> OpenAPI-файл <a href=\"/swagger.yaml\">swagger.yaml</a> </ul>  <br> Для ручной проверки API вы можете использовать: <ul> <li> Под ОС Windows - [PostMan](https://www.postman.com/) <li> Под ОС Linux - [curl](https://curl.se/)  </ul> <br>  ## Поддержка Техническая поддержка осуществляется через диалоги в личном кабинете продавца. При создании нового обращения в техподдержку используйте категорию API. <br> Новости и изменения, касающиеся API, публикуются в [новостной ленте Wildberries](https://seller.wildberries.ru/news). <br> <br> <br>  ## Авторизация Авторизация осуществляется по токенам API, которые  владелец личного кабинета (главный пользователь) самостоятельно  генерирует в разделе   [Профиль --> Настройки --> Доступ к API](https://seller.wildberries.ru/supplier-settings/access-to-api).   <br>  <br>  Существует три типа токенов:   <ol>     <li>Стандартный. Используется для работы с методами из разделов: Цены, Промокоды и скидки, Контент, Marketplace, Рекомендации, Отзывы, Вопросы.</li>     <li>Статистика. Используется для работы с методами из раздела Статистика.</li>     <li>Реклама. Используется для работы с методами из раздела Реклама.</li>   </ol>    Обратите внимание, что токен отображается ТОЛЬКО в момент создания. Его надо сохранить, потому что больше его отобразить будет нельзя. <br>Созданный токен следует добавлять в каждый запрос, прибавляя к запросу заголовок (http-header) формата `Authorization: .........`. <br> <br> <br> ## Форматы ### Дата и время Во всех методах API статистики дата и время передаются в формате [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).  <br> В большинстве случаев вы можете передать дату или дату со временем. Если время не указано, оно принимается равным 00:00:00. Время можно указывать с точностью до секунд или миллисекунд.  Литера `Z` в конце строки означает часовой пояс UTC. При ее отсутствии время считается в часовом поясе МСК (UTC+3). <br> Примеры: <ul> <li> `2019-06-20` <li> `2019-06-20T00:00:00Z` <li> `2019-06-20T23:59:59` <li> `2019-06-20T00:00:00.12345Z` <li> `2019-06-20T00:00:00.12345` <li> `2017-03-25T00:00:00` </ul> <br> ## Release Notes <p>Для просмотра обновлений перейдите по <span><a href=\"./releasenotes/index.html\" target=\"_blank\">ссылке</a></span>.</p> 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.4
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./wildberries"
```

## Documentation for API Endpoints

All URIs are relative to */*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**AdvV0ActiveGet**](docs/DefaultApi.md#advv0activeget) | **Get** /adv/v0/active | Изменение активности предметной группы для РК в поиске
*DefaultApi* | [**AdvV0AdvertGet**](docs/DefaultApi.md#advv0advertget) | **Get** /adv/v0/advert | Информация о РК
*DefaultApi* | [**AdvV0AdvertsGet**](docs/DefaultApi.md#advv0advertsget) | **Get** /adv/v0/adverts | Список РК
*DefaultApi* | [**AdvV0AllcpmPost**](docs/DefaultApi.md#advv0allcpmpost) | **Post** /adv/v0/allcpm | Список ставок по типу размещения РК
*DefaultApi* | [**AdvV0CountGet**](docs/DefaultApi.md#advv0countget) | **Get** /adv/v0/count | Получение РК
*DefaultApi* | [**AdvV0CpmGet**](docs/DefaultApi.md#advv0cpmget) | **Get** /adv/v0/cpm | Список ставок
*DefaultApi* | [**AdvV0CpmPost**](docs/DefaultApi.md#advv0cpmpost) | **Post** /adv/v0/cpm | Изменение ставки у РК
*DefaultApi* | [**AdvV0DailybudgetPost**](docs/DefaultApi.md#advv0dailybudgetpost) | **Post** /adv/v0/dailybudget | Изменение дневного бюджета РК
*DefaultApi* | [**AdvV0IntervalsPost**](docs/DefaultApi.md#advv0intervalspost) | **Post** /adv/v0/intervals | Изменение интервалов показа РК
*DefaultApi* | [**AdvV0NmactivePost**](docs/DefaultApi.md#advv0nmactivepost) | **Post** /adv/v0/nmactive | Изменение активности номенклатур РК
*DefaultApi* | [**AdvV0ParamsMenuGet**](docs/DefaultApi.md#advv0paramsmenuget) | **Get** /adv/v0/params/menu | Словарь значений параметра menuId
*DefaultApi* | [**AdvV0ParamsSetGet**](docs/DefaultApi.md#advv0paramssetget) | **Get** /adv/v0/params/set | Словарь значений параметра setId
*DefaultApi* | [**AdvV0ParamsSubjectGet**](docs/DefaultApi.md#advv0paramssubjectget) | **Get** /adv/v0/params/subject | Словарь значений параметра subjectId
*DefaultApi* | [**AdvV0PauseGet**](docs/DefaultApi.md#advv0pauseget) | **Get** /adv/v0/pause | Пауза РК
*DefaultApi* | [**AdvV0RenamePost**](docs/DefaultApi.md#advv0renamepost) | **Post** /adv/v0/rename | Переименование РК
*DefaultApi* | [**AdvV0StartGet**](docs/DefaultApi.md#advv0startget) | **Get** /adv/v0/start | Запуск РК
*DefaultApi* | [**AdvV0StopGet**](docs/DefaultApi.md#advv0stopget) | **Get** /adv/v0/stop | Завершение РК
*DefaultApi* | [**ApiV1DelPost**](docs/DefaultApi.md#apiv1delpost) | **Post** /api/v1/del | Удаление рекомендаций
*DefaultApi* | [**ApiV1FeedbacksArchiveGet**](docs/DefaultApi.md#apiv1feedbacksarchiveget) | **Get** /api/v1/feedbacks/archive | Список архивных отзывов
*DefaultApi* | [**ApiV1FeedbacksCountUnansweredGet**](docs/DefaultApi.md#apiv1feedbackscountunansweredget) | **Get** /api/v1/feedbacks/count-unanswered | Необработанные отзывы
*DefaultApi* | [**ApiV1FeedbacksGet**](docs/DefaultApi.md#apiv1feedbacksget) | **Get** /api/v1/feedbacks | Список отзывов
*DefaultApi* | [**ApiV1FeedbacksPatch**](docs/DefaultApi.md#apiv1feedbackspatch) | **Patch** /api/v1/feedbacks | Работа с отзывом
*DefaultApi* | [**ApiV1FeedbacksProductsRatingGet**](docs/DefaultApi.md#apiv1feedbacksproductsratingget) | **Get** /api/v1/feedbacks/products/rating | Средняя оценка товаров по родительской категории
*DefaultApi* | [**ApiV1FeedbacksProductsRatingNmidGet**](docs/DefaultApi.md#apiv1feedbacksproductsratingnmidget) | **Get** /api/v1/feedbacks/products/rating/nmid | Средняя оценка товара по артикулу WB
*DefaultApi* | [**ApiV1FeedbacksProductsRatingTopGet**](docs/DefaultApi.md#apiv1feedbacksproductsratingtopget) | **Get** /api/v1/feedbacks/products/rating/top | Товары с наибольшей и наименьшей средней оценкой по родительской категории
*DefaultApi* | [**ApiV1FeedbacksReportGet**](docs/DefaultApi.md#apiv1feedbacksreportget) | **Get** /api/v1/feedbacks/report | Получение отзывов в формате XLSX
*DefaultApi* | [**ApiV1InsPost**](docs/DefaultApi.md#apiv1inspost) | **Post** /api/v1/ins | Добавление рекомендаций
*DefaultApi* | [**ApiV1ListGet**](docs/DefaultApi.md#apiv1listget) | **Get** /api/v1/list | Список рекомендаций
*DefaultApi* | [**ApiV1NewFeedbacksQuestionsGet**](docs/DefaultApi.md#apiv1newfeedbacksquestionsget) | **Get** /api/v1/new-feedbacks-questions | Непросмотренные отзывы и вопросы
*DefaultApi* | [**ApiV1ParentSubjectsGet**](docs/DefaultApi.md#apiv1parentsubjectsget) | **Get** /api/v1/parent-subjects | Родительские категории товаров
*DefaultApi* | [**ApiV1QuestionsCountUnansweredGet**](docs/DefaultApi.md#apiv1questionscountunansweredget) | **Get** /api/v1/questions/count-unanswered | Неотвеченные вопросы
*DefaultApi* | [**ApiV1QuestionsGet**](docs/DefaultApi.md#apiv1questionsget) | **Get** /api/v1/questions | Список вопросов
*DefaultApi* | [**ApiV1QuestionsPatch**](docs/DefaultApi.md#apiv1questionspatch) | **Patch** /api/v1/questions | Работа с вопросами
*DefaultApi* | [**ApiV1QuestionsProductsRatingGet**](docs/DefaultApi.md#apiv1questionsproductsratingget) | **Get** /api/v1/questions/products/rating | Часто спрашиваемые товары
*DefaultApi* | [**ApiV1QuestionsReportGet**](docs/DefaultApi.md#apiv1questionsreportget) | **Get** /api/v1/questions/report | Получение вопросов в формате XLSX
*DefaultApi* | [**ApiV1SetPost**](docs/DefaultApi.md#apiv1setpost) | **Post** /api/v1/set | Управление рекомендациями
*DefaultApi* | [**ApiV1SupplierIncomesGet**](docs/DefaultApi.md#apiv1supplierincomesget) | **Get** /api/v1/supplier/incomes | Поставки
*DefaultApi* | [**ApiV1SupplierOrdersGet**](docs/DefaultApi.md#apiv1supplierordersget) | **Get** /api/v1/supplier/orders | Заказы
*DefaultApi* | [**ApiV1SupplierReportDetailByPeriodGet**](docs/DefaultApi.md#apiv1supplierreportdetailbyperiodget) | **Get** /api/v1/supplier/reportDetailByPeriod | Отчет о продажах по реализации
*DefaultApi* | [**ApiV1SupplierSalesGet**](docs/DefaultApi.md#apiv1suppliersalesget) | **Get** /api/v1/supplier/sales | Продажи
*DefaultApi* | [**ApiV1SupplierStocksGet**](docs/DefaultApi.md#apiv1supplierstocksget) | **Get** /api/v1/supplier/stocks | Склад
*DefaultApi* | [**PublicApiV1InfoGet**](docs/DefaultApi.md#publicapiv1infoget) | **Get** /public/api/v1/info | Получение информации о ценах.
*DefaultApi* | [**PublicApiV1PricesPost**](docs/DefaultApi.md#publicapiv1pricespost) | **Post** /public/api/v1/prices | Загрузка цен
*Marketplace_Api* | [**ApiV2WarehousesGet**](docs/Marketplace_Api.md#apiv2warehousesget) | **Get** /api/v2/warehouses | Cписок складов
*Marketplace_Api* | [**ApiV3OrdersGet**](docs/Marketplace_Api.md#apiv3ordersget) | **Get** /api/v3/orders | Получить информацию по сборочным заданиям
*Marketplace_Api* | [**ApiV3OrdersNewGet**](docs/Marketplace_Api.md#apiv3ordersnewget) | **Get** /api/v3/orders/new | Получить список новых сборочных заданий
*Marketplace_Api* | [**ApiV3OrdersOrderCancelPatch**](docs/Marketplace_Api.md#apiv3ordersordercancelpatch) | **Patch** /api/v3/orders/{order}/cancel | Отменить сборочное задание
*Marketplace_Api* | [**ApiV3OrdersOrderMetaSgtinPost**](docs/Marketplace_Api.md#apiv3ordersordermetasgtinpost) | **Post** /api/v3/orders/{order}/meta/sgtin | Закрепить за сборочным заданием КиЗ (маркировку Честного знака)
*Marketplace_Api* | [**ApiV3OrdersStatusPost**](docs/Marketplace_Api.md#apiv3ordersstatuspost) | **Post** /api/v3/orders/status | Получить статусы сборочных заданий
*Marketplace_Api* | [**ApiV3OrdersStickersPost**](docs/Marketplace_Api.md#apiv3ordersstickerspost) | **Post** /api/v3/orders/stickers | Получить этикетки для сборочных заданий
*Marketplace_Api* | [**ApiV3StocksWarehouseDelete**](docs/Marketplace_Api.md#apiv3stockswarehousedelete) | **Delete** /api/v3/stocks/{warehouse} | Удалить остатки товаров
*Marketplace_Api* | [**ApiV3StocksWarehousePost**](docs/Marketplace_Api.md#apiv3stockswarehousepost) | **Post** /api/v3/stocks/{warehouse} | Получить остатки товаров
*Marketplace_Api* | [**ApiV3StocksWarehousePut**](docs/Marketplace_Api.md#apiv3stockswarehouseput) | **Put** /api/v3/stocks/{warehouse} | Обновить остатки товаров
*Marketplace_Api* | [**ApiV3SuppliesGet**](docs/Marketplace_Api.md#apiv3suppliesget) | **Get** /api/v3/supplies | Получить список поставок
*Marketplace_Api* | [**ApiV3SuppliesPost**](docs/Marketplace_Api.md#apiv3suppliespost) | **Post** /api/v3/supplies | Создать новую поставку
*Marketplace_Api* | [**ApiV3SuppliesSupplyBarcodeGet**](docs/Marketplace_Api.md#apiv3suppliessupplybarcodeget) | **Get** /api/v3/supplies/{supply}/barcode | Получить QR поставки
*Marketplace_Api* | [**ApiV3SuppliesSupplyDelete**](docs/Marketplace_Api.md#apiv3suppliessupplydelete) | **Delete** /api/v3/supplies/{supply} | Удалить поставку
*Marketplace_Api* | [**ApiV3SuppliesSupplyDeliverPatch**](docs/Marketplace_Api.md#apiv3suppliessupplydeliverpatch) | **Patch** /api/v3/supplies/{supply}/deliver | Передать поставку в доставку
*Marketplace_Api* | [**ApiV3SuppliesSupplyGet**](docs/Marketplace_Api.md#apiv3suppliessupplyget) | **Get** /api/v3/supplies/{supply} | Получить информацию о поставке
*Marketplace_Api* | [**ApiV3SuppliesSupplyOrdersGet**](docs/Marketplace_Api.md#apiv3suppliessupplyordersget) | **Get** /api/v3/supplies/{supply}/orders | Получить сборочные задания в поставке
*Marketplace_Api* | [**ApiV3SuppliesSupplyOrdersOrderPatch**](docs/Marketplace_Api.md#apiv3suppliessupplyordersorderpatch) | **Patch** /api/v3/supplies/{supply}/orders/{order} | Добавить к поставке сборочное задание
*_Api* | [**ContentV1BarcodesPost**](docs/_Api.md#contentv1barcodespost) | **Post** /content/v1/barcodes | Генерация баркодов
*_Api* | [**ContentV1CardsCursorListPost**](docs/_Api.md#contentv1cardscursorlistpost) | **Post** /content/v1/cards/cursor/list | Список НМ
*_Api* | [**ContentV1CardsErrorListGet**](docs/_Api.md#contentv1cardserrorlistget) | **Get** /content/v1/cards/error/list | Список несозданных НМ с ошибками
*_Api* | [**ContentV1CardsFilterPost**](docs/_Api.md#contentv1cardsfilterpost) | **Post** /content/v1/cards/filter | Получение КТ по артикулам продавца
*_Api* | [**ContentV1CardsLimitsGet**](docs/_Api.md#contentv1cardslimitsget) | **Get** /content/v1/cards/limits | Лимиты по КТ
*_Api* | [**ContentV1CardsTrashListPost**](docs/_Api.md#contentv1cardstrashlistpost) | **Post** /content/v1/cards/trash/list | Список НМ, находящихся в корзине
*_Api* | [**ContentV1CardsUpdatePost**](docs/_Api.md#contentv1cardsupdatepost) | **Post** /content/v1/cards/update | Редактирование КТ
*_Api* | [**ContentV1CardsUploadAddPost**](docs/_Api.md#contentv1cardsuploadaddpost) | **Post** /content/v1/cards/upload/add | Добавление НМ к КТ
*_Api* | [**ContentV1CardsUploadPost**](docs/_Api.md#contentv1cardsuploadpost) | **Post** /content/v1/cards/upload | Создание  КТ
*_Api* | [**ContentV1DirectoryBrandsGet**](docs/_Api.md#contentv1directorybrandsget) | **Get** /content/v1/directory/brands | Бренд
*_Api* | [**ContentV1DirectoryColorsGet**](docs/_Api.md#contentv1directorycolorsget) | **Get** /content/v1/directory/colors | Цвет
*_Api* | [**ContentV1DirectoryCountriesGet**](docs/_Api.md#contentv1directorycountriesget) | **Get** /content/v1/directory/countries | Страна Производства
*_Api* | [**ContentV1DirectoryKindsGet**](docs/_Api.md#contentv1directorykindsget) | **Get** /content/v1/directory/kinds | Пол
*_Api* | [**ContentV1DirectorySeasonsGet**](docs/_Api.md#contentv1directoryseasonsget) | **Get** /content/v1/directory/seasons | Сезон
*_Api* | [**ContentV1DirectoryTnvedGet**](docs/_Api.md#contentv1directorytnvedget) | **Get** /content/v1/directory/tnved | ТНВЭД код
*_Api* | [**ContentV1MediaFilePost**](docs/_Api.md#contentv1mediafilepost) | **Post** /content/v1/media/file | Добавление медиа контента в КТ
*_Api* | [**ContentV1MediaSavePost**](docs/_Api.md#contentv1mediasavepost) | **Post** /content/v1/media/save | Изменение медиа контента КТ
*_Api* | [**ContentV1ObjectAllGet**](docs/_Api.md#contentv1objectallget) | **Get** /content/v1/object/all | Категория товаров
*_Api* | [**ContentV1ObjectCharacteristicsListFilterGet**](docs/_Api.md#contentv1objectcharacteristicslistfilterget) | **Get** /content/v1/object/characteristics/list/filter | Характеристики для создания КТ по всем подкатегориям
*_Api* | [**ContentV1ObjectCharacteristicsObjectNameGet**](docs/_Api.md#contentv1objectcharacteristicsobjectnameget) | **Get** /content/v1/object/characteristics/{objectName} | Характеристики для создания КТ для категории товара
*_Api* | [**ContentV1ObjectParentAllGet**](docs/_Api.md#contentv1objectparentallget) | **Get** /content/v1/object/parent/all | Родительские категории товаров
*_Api* | [**ContentV1TagIdDelete**](docs/_Api.md#contentv1tagiddelete) | **Delete** /content/v1/tag/{id} | Удаление тега
*_Api* | [**ContentV1TagIdPatch**](docs/_Api.md#contentv1tagidpatch) | **Patch** /content/v1/tag/{id} | Изменение тега
*_Api* | [**ContentV1TagNomenclatureLinkPost**](docs/_Api.md#contentv1tagnomenclaturelinkpost) | **Post** /content/v1/tag/nomenclature/link | Управление тегами в КТ
*_Api* | [**ContentV1TagPost**](docs/_Api.md#contentv1tagpost) | **Post** /content/v1/tag | Создание тега
*_Api* | [**ContentV1TagsGet**](docs/_Api.md#contentv1tagsget) | **Get** /content/v1/tags | Список тегов
*_Api* | [**PublicApiV1RevokeDiscountsPost**](docs/_Api.md#publicapiv1revokediscountspost) | **Post** /public/api/v1/revokeDiscounts | Сброс скидок для номенклатур
*_Api* | [**PublicApiV1RevokePromocodesPost**](docs/_Api.md#publicapiv1revokepromocodespost) | **Post** /public/api/v1/revokePromocodes | Сброс промокодов для номенклатур
*_Api* | [**PublicApiV1UpdateDiscountsPost**](docs/_Api.md#publicapiv1updatediscountspost) | **Post** /public/api/v1/updateDiscounts | Установка скидок
*_Api* | [**PublicApiV1UpdatePromocodesPost**](docs/_Api.md#publicapiv1updatepromocodespost) | **Post** /public/api/v1/updatePromocodes | Установка промокодов для номенклатур

## Documentation For Models

 - [Advv0allcpmCpm](docs/Advv0allcpmCpm.md)
 - [Advv0intervalsIntervals](docs/Advv0intervalsIntervals.md)
 - [Advv0nmactiveActive](docs/Advv0nmactiveActive.md)
 - [Apiv1questionsAnswer](docs/Apiv1questionsAnswer.md)
 - [Apiv3ordersordermetasgtinSgtin](docs/Apiv3ordersordermetasgtinSgtin.md)
 - [Apiv3stockswarehouseStocks](docs/Apiv3stockswarehouseStocks.md)
 - [CardsFilterBody](docs/CardsFilterBody.md)
 - [CardsUpdateBody](docs/CardsUpdateBody.md)
 - [CardsUploadBody](docs/CardsUploadBody.md)
 - [Contentv1cardscursorlistSort](docs/Contentv1cardscursorlistSort.md)
 - [Contentv1cardscursorlistSortCursor](docs/Contentv1cardscursorlistSortCursor.md)
 - [Contentv1cardscursorlistSortFilter](docs/Contentv1cardscursorlistSortFilter.md)
 - [Contentv1cardscursorlistSortSort](docs/Contentv1cardscursorlistSortSort.md)
 - [Contentv1cardstrashlistSort](docs/Contentv1cardstrashlistSort.md)
 - [Contentv1cardsupdateSizes](docs/Contentv1cardsupdateSizes.md)
 - [Contentv1cardsuploadaddCards](docs/Contentv1cardsuploadaddCards.md)
 - [Contentv1cardsuploadaddSizes](docs/Contentv1cardsuploadaddSizes.md)
 - [CursorListBody](docs/CursorListBody.md)
 - [DetailReportItem](docs/DetailReportItem.md)
 - [FeedbackErr400](docs/FeedbackErr400.md)
 - [FeedbackErr403](docs/FeedbackErr403.md)
 - [FeedbackErr404](docs/FeedbackErr404.md)
 - [IncomesItem](docs/IncomesItem.md)
 - [IncorrectRequest](docs/IncorrectRequest.md)
 - [IncorrectRequestAdditionalErrors](docs/IncorrectRequestAdditionalErrors.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse20010](docs/InlineResponse20010.md)
 - [InlineResponse20010Data](docs/InlineResponse20010Data.md)
 - [InlineResponse20011](docs/InlineResponse20011.md)
 - [InlineResponse20011Data](docs/InlineResponse20011Data.md)
 - [InlineResponse20012](docs/InlineResponse20012.md)
 - [InlineResponse20012Data](docs/InlineResponse20012Data.md)
 - [InlineResponse20013](docs/InlineResponse20013.md)
 - [InlineResponse20014](docs/InlineResponse20014.md)
 - [InlineResponse20014Data](docs/InlineResponse20014Data.md)
 - [InlineResponse20015](docs/InlineResponse20015.md)
 - [InlineResponse20016](docs/InlineResponse20016.md)
 - [InlineResponse20017](docs/InlineResponse20017.md)
 - [InlineResponse20017Data](docs/InlineResponse20017Data.md)
 - [InlineResponse20018](docs/InlineResponse20018.md)
 - [InlineResponse20019](docs/InlineResponse20019.md)
 - [InlineResponse20019Data](docs/InlineResponse20019Data.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse20020](docs/InlineResponse20020.md)
 - [InlineResponse20021](docs/InlineResponse20021.md)
 - [InlineResponse20022](docs/InlineResponse20022.md)
 - [InlineResponse20022Data](docs/InlineResponse20022Data.md)
 - [InlineResponse20022DataCards](docs/InlineResponse20022DataCards.md)
 - [InlineResponse20022DataCursor](docs/InlineResponse20022DataCursor.md)
 - [InlineResponse20022DataSizes](docs/InlineResponse20022DataSizes.md)
 - [InlineResponse20023](docs/InlineResponse20023.md)
 - [InlineResponse20024](docs/InlineResponse20024.md)
 - [InlineResponse20024Stocks](docs/InlineResponse20024Stocks.md)
 - [InlineResponse20025](docs/InlineResponse20025.md)
 - [InlineResponse20026](docs/InlineResponse20026.md)
 - [InlineResponse20027](docs/InlineResponse20027.md)
 - [InlineResponse20028](docs/InlineResponse20028.md)
 - [InlineResponse20028Orders](docs/InlineResponse20028Orders.md)
 - [InlineResponse20029](docs/InlineResponse20029.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse20030](docs/InlineResponse20030.md)
 - [InlineResponse20031](docs/InlineResponse20031.md)
 - [InlineResponse20031Stickers](docs/InlineResponse20031Stickers.md)
 - [InlineResponse20032](docs/InlineResponse20032.md)
 - [InlineResponse20032Adverts](docs/InlineResponse20032Adverts.md)
 - [InlineResponse20033](docs/InlineResponse20033.md)
 - [InlineResponse20034](docs/InlineResponse20034.md)
 - [InlineResponse20034Intervals](docs/InlineResponse20034Intervals.md)
 - [InlineResponse20034Nms](docs/InlineResponse20034Nms.md)
 - [InlineResponse20034Params](docs/InlineResponse20034Params.md)
 - [InlineResponse20035](docs/InlineResponse20035.md)
 - [InlineResponse20036](docs/InlineResponse20036.md)
 - [InlineResponse20037](docs/InlineResponse20037.md)
 - [InlineResponse20038](docs/InlineResponse20038.md)
 - [InlineResponse20039](docs/InlineResponse20039.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse20040](docs/InlineResponse20040.md)
 - [InlineResponse20040Data](docs/InlineResponse20040Data.md)
 - [InlineResponse20041](docs/InlineResponse20041.md)
 - [InlineResponse20041Data](docs/InlineResponse20041Data.md)
 - [InlineResponse20042](docs/InlineResponse20042.md)
 - [InlineResponse20042Data](docs/InlineResponse20042Data.md)
 - [InlineResponse20042DataProducts](docs/InlineResponse20042DataProducts.md)
 - [InlineResponse20043](docs/InlineResponse20043.md)
 - [InlineResponse20043Data](docs/InlineResponse20043Data.md)
 - [InlineResponse20043DataAnswer](docs/InlineResponse20043DataAnswer.md)
 - [InlineResponse20043DataProductDetails](docs/InlineResponse20043DataProductDetails.md)
 - [InlineResponse20043DataQuestions](docs/InlineResponse20043DataQuestions.md)
 - [InlineResponse20044](docs/InlineResponse20044.md)
 - [InlineResponse20045](docs/InlineResponse20045.md)
 - [InlineResponse20045Data](docs/InlineResponse20045Data.md)
 - [InlineResponse20046](docs/InlineResponse20046.md)
 - [InlineResponse20046Data](docs/InlineResponse20046Data.md)
 - [InlineResponse20047](docs/InlineResponse20047.md)
 - [InlineResponse20047Data](docs/InlineResponse20047Data.md)
 - [InlineResponse20048](docs/InlineResponse20048.md)
 - [InlineResponse20048Data](docs/InlineResponse20048Data.md)
 - [InlineResponse20049](docs/InlineResponse20049.md)
 - [InlineResponse20049Data](docs/InlineResponse20049Data.md)
 - [InlineResponse2004Data](docs/InlineResponse2004Data.md)
 - [InlineResponse2004DataCards](docs/InlineResponse2004DataCards.md)
 - [InlineResponse2004DataCursor](docs/InlineResponse2004DataCursor.md)
 - [InlineResponse2004DataSizes](docs/InlineResponse2004DataSizes.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse20050](docs/InlineResponse20050.md)
 - [InlineResponse20050Data](docs/InlineResponse20050Data.md)
 - [InlineResponse20051](docs/InlineResponse20051.md)
 - [InlineResponse20051Data](docs/InlineResponse20051Data.md)
 - [InlineResponse20052](docs/InlineResponse20052.md)
 - [InlineResponse20052Data](docs/InlineResponse20052Data.md)
 - [InlineResponse20053](docs/InlineResponse20053.md)
 - [InlineResponse20053Data](docs/InlineResponse20053Data.md)
 - [InlineResponse20054](docs/InlineResponse20054.md)
 - [InlineResponse20054Data](docs/InlineResponse20054Data.md)
 - [InlineResponse2005Data](docs/InlineResponse2005Data.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse2006Data](docs/InlineResponse2006Data.md)
 - [InlineResponse2006Sizes](docs/InlineResponse2006Sizes.md)
 - [InlineResponse2007](docs/InlineResponse2007.md)
 - [InlineResponse2007Data](docs/InlineResponse2007Data.md)
 - [InlineResponse2008](docs/InlineResponse2008.md)
 - [InlineResponse2008Data](docs/InlineResponse2008Data.md)
 - [InlineResponse2009](docs/InlineResponse2009.md)
 - [InlineResponse2009Data](docs/InlineResponse2009Data.md)
 - [InlineResponse201](docs/InlineResponse201.md)
 - [InlineResponse400](docs/InlineResponse400.md)
 - [InlineResponse4001](docs/InlineResponse4001.md)
 - [InlineResponse40010](docs/InlineResponse40010.md)
 - [InlineResponse40011](docs/InlineResponse40011.md)
 - [InlineResponse40012](docs/InlineResponse40012.md)
 - [InlineResponse4002](docs/InlineResponse4002.md)
 - [InlineResponse4003](docs/InlineResponse4003.md)
 - [InlineResponse4004](docs/InlineResponse4004.md)
 - [InlineResponse4005](docs/InlineResponse4005.md)
 - [InlineResponse4006](docs/InlineResponse4006.md)
 - [InlineResponse4007](docs/InlineResponse4007.md)
 - [InlineResponse4008](docs/InlineResponse4008.md)
 - [InlineResponse4009](docs/InlineResponse4009.md)
 - [InlineResponse401](docs/InlineResponse401.md)
 - [InlineResponse4011](docs/InlineResponse4011.md)
 - [InlineResponse4012](docs/InlineResponse4012.md)
 - [InlineResponse422](docs/InlineResponse422.md)
 - [InlineResponse4221](docs/InlineResponse4221.md)
 - [InlineResponse4222](docs/InlineResponse4222.md)
 - [InlineResponse4223](docs/InlineResponse4223.md)
 - [InlineResponse4224](docs/InlineResponse4224.md)
 - [InlineResponse4225](docs/InlineResponse4225.md)
 - [MediaFileBody](docs/MediaFileBody.md)
 - [MediaSaveBody](docs/MediaSaveBody.md)
 - [MetaSgtinBody](docs/MetaSgtinBody.md)
 - [ModelError](docs/ModelError.md)
 - [NomenclatureLinkBody](docs/NomenclatureLinkBody.md)
 - [Order](docs/Order.md)
 - [OrderAddress](docs/OrderAddress.md)
 - [OrderUser](docs/OrderUser.md)
 - [OrdersItem](docs/OrdersItem.md)
 - [OrdersStatusBody](docs/OrdersStatusBody.md)
 - [OrdersStickersBody](docs/OrdersStickersBody.md)
 - [ProductRating](docs/ProductRating.md)
 - [ResponseBody400LenName](docs/ResponseBody400LenName.md)
 - [ResponseBodyError400](docs/ResponseBodyError400.md)
 - [ResponseBodyError403](docs/ResponseBodyError403.md)
 - [ResponseBodyError500](docs/ResponseBodyError500.md)
 - [ResponseBodyLimitTag](docs/ResponseBodyLimitTag.md)
 - [ResponseBodyNameNotExist](docs/ResponseBodyNameNotExist.md)
 - [ResponseBodyTagNotExist](docs/ResponseBodyTagNotExist.md)
 - [ResponseBodyTagNotExistAdditionalErrors](docs/ResponseBodyTagNotExistAdditionalErrors.md)
 - [ResponseFeadbackInner](docs/ResponseFeadbackInner.md)
 - [ResponseNotFound200](docs/ResponseNotFound200.md)
 - [ResponseNotFound200AdditionalErrors](docs/ResponseNotFound200AdditionalErrors.md)
 - [ResponseNotFound200Del](docs/ResponseNotFound200Del.md)
 - [ResponseOk200](docs/ResponseOk200.md)
 - [SalesItem](docs/SalesItem.md)
 - [StocksItem](docs/StocksItem.md)
 - [StocksWarehouseBody](docs/StocksWarehouseBody.md)
 - [StocksWarehouseBody1](docs/StocksWarehouseBody1.md)
 - [StocksWarehouseBody2](docs/StocksWarehouseBody2.md)
 - [Supply](docs/Supply.md)
 - [SupplyOrder](docs/SupplyOrder.md)
 - [TagIdBody](docs/TagIdBody.md)
 - [TrashListBody](docs/TrashListBody.md)
 - [UploadAddBody](docs/UploadAddBody.md)
 - [V0AllcpmBody](docs/V0AllcpmBody.md)
 - [V0CpmBody](docs/V0CpmBody.md)
 - [V0DailybudgetBody](docs/V0DailybudgetBody.md)
 - [V0IntervalsBody](docs/V0IntervalsBody.md)
 - [V0NmactiveBody](docs/V0NmactiveBody.md)
 - [V0RenameBody](docs/V0RenameBody.md)
 - [V1BarcodesBody](docs/V1BarcodesBody.md)
 - [V1DelBody](docs/V1DelBody.md)
 - [V1FeedbacksBody](docs/V1FeedbacksBody.md)
 - [V1InsBody](docs/V1InsBody.md)
 - [V1PricesBody](docs/V1PricesBody.md)
 - [V1QuestionsBody](docs/V1QuestionsBody.md)
 - [V1SetBody](docs/V1SetBody.md)
 - [V1TagBody](docs/V1TagBody.md)
 - [V3SuppliesBody](docs/V3SuppliesBody.md)

## Documentation For Authorization

## HeaderApiKey
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

