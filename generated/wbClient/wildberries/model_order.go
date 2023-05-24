/*
 * API продавца
 *
 * # Общее описание <style> .langChange {   text-align: left;   padding: 2px 222px 2px 18px; } </style> <style> .new {   display: inline-block;   padding: 2px 8px;   margin: 0;   background-color: #ffa500;   color: #ffffff;   font-size: 13px;   vertical-align: middle;   line-height: 1.6;   border-radius: 4px;   font-weight: 600;   font-size: 12px; } </style> <style> .newM {   padding: 5px 10px;   display: inline-block;   background-color: #ffa500;   color: #ffffff;   cursor: pointer;   text-align: center;   outline: none;   margin: 0 5px 5px 5px;   #border: 1px solid #07090b;   border-radius: 5px;   min-width: 60px;   font-size: 0.9em;   font-weight: bold; } </style>  Wildberries API предоставляет продавцам возможность управления магазином и получения оперативной и статистической информации по протоколу HTTP RestAPI. <br> Описание API предоставляется в формате [Swagger](https://swagger.io/) (Open API) и может быть использовано для импорта в другие инструменты (такие как PostMan) или генерации клиентского кода на различных языках программирования с помощью [Swagger CodeGen](https://swagger.io/tools/swagger-codegen/)  <ul> <li> Описание в оригинальном swagger-формате <a href=\"/swagger\">swagger</a> <li> OpenAPI-файл <a href=\"/swagger.yaml\">swagger.yaml</a> </ul>  <br> Для ручной проверки API вы можете использовать: <ul> <li> Под ОС Windows - [PostMan](https://www.postman.com/) <li> Под ОС Linux - [curl](https://curl.se/)  </ul> <br>  ## Поддержка Техническая поддержка осуществляется через диалоги в личном кабинете продавца. При создании нового обращения в техподдержку используйте категорию API. <br> Новости и изменения, касающиеся API, публикуются в [новостной ленте Wildberries](https://seller.wildberries.ru/news). <br> <br> <br>  ## Авторизация Авторизация осуществляется по токенам API, которые  владелец личного кабинета (главный пользователь) самостоятельно  генерирует в разделе   [Профиль --> Настройки --> Доступ к API](https://seller.wildberries.ru/supplier-settings/access-to-api).   <br>  <br>  Существует три типа токенов:   <ol>     <li>Стандартный. Используется для работы с методами из разделов: Цены, Промокоды и скидки, Контент, Marketplace, Рекомендации, Отзывы, Вопросы.</li>     <li>Статистика. Используется для работы с методами из раздела Статистика.</li>     <li>Реклама. Используется для работы с методами из раздела Реклама.</li>   </ol>    Обратите внимание, что токен отображается ТОЛЬКО в момент создания. Его надо сохранить, потому что больше его отобразить будет нельзя. <br>Созданный токен следует добавлять в каждый запрос, прибавляя к запросу заголовок (http-header) формата `Authorization: .........`. <br> <br> <br> ## Форматы ### Дата и время Во всех методах API статистики дата и время передаются в формате [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).  <br> В большинстве случаев вы можете передать дату или дату со временем. Если время не указано, оно принимается равным 00:00:00. Время можно указывать с точностью до секунд или миллисекунд.  Литера `Z` в конце строки означает часовой пояс UTC. При ее отсутствии время считается в часовом поясе МСК (UTC+3). <br> Примеры: <ul> <li> `2019-06-20` <li> `2019-06-20T00:00:00Z` <li> `2019-06-20T23:59:59` <li> `2019-06-20T00:00:00.12345Z` <li> `2019-06-20T00:00:00.12345` <li> `2017-03-25T00:00:00` </ul> <br> ## Release Notes <p>Для просмотра обновлений перейдите по <span><a href=\"./releasenotes/index.html\" target=\"_blank\">ссылке</a></span>.</p> 
 *
 * API version: 2.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package wildberries
import (
	"time"
)

type Order struct {
	// Идентификатор сборочного задания в Маркетплейсе
	Id int64 `json:"id,omitempty"`
	// Идентификатор сборочного задания в системе Wildberries
	Rid string `json:"rid,omitempty"`
	// Дата создания сборочного задания (RFC3339)
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Идентификатор склада продавца, на который поступило сборочное задание
	WarehouseId int32 `json:"warehouseId,omitempty"`
	// Идентификатор поставки. Возвращается, если заказ закреплён за поставкой
	SupplyId string `json:"supplyId,omitempty"`
	// Массив приоритетных СЦ для доставки сборочного задания. Если поле не заполнено или массив пустой, приоритетного СЦ для данного сборочного задания нет
	PrioritySc []string `json:"prioritySc,omitempty"`
	// Список офисов, куда следует привезти товар
	Offices []string `json:"offices,omitempty"`
	Address *OrderAddress `json:"address,omitempty"`
	User *OrderUser `json:"user,omitempty"`
	// Массив баркодов товара
	Skus []string `json:"skus,omitempty"`
	// Цена в валюте продажи с учетом всех скидок, умноженная на 100. Код валюты продажи в поле currencyCode.
	Price int32 `json:"price,omitempty"`
	// Цена в валюте продажи с учетом всех скидок, сконвертированная по курсу на момент продажи в российские копейки. Предоставляется в информационных целях.
	ConvertedPrice int32 `json:"convertedPrice,omitempty"`
	// Код валюты продажи (ISO 4217)
	CurrencyCode int32 `json:"currencyCode,omitempty"`
	// Код валюты страны поставщика (ISO 4217)
	ConvertedCurrencyCode int32 `json:"convertedCurrencyCode,omitempty"`
	// Идентификатор транзакции для группировки сборочных заданий. Сборочные задания в одной корзине покупателя будут иметь одинаковый orderUID
	OrderUid string `json:"orderUid,omitempty"`
	// Тип доставки: fbs - доставка на склад Wildberries, dbs - доставка силами продавца 
	DeliveryType string `json:"deliveryType,omitempty"`
	// Артикул WB
	NmId int32 `json:"nmId,omitempty"`
	// Идентификатор размера товара в системе Wildberries
	ChrtId int32 `json:"chrtId,omitempty"`
	// Артикул продавца
	Article string `json:"article,omitempty"`
	// сКГТ-признак товара, на который был сделан заказ
	IsLargeCargo bool `json:"isLargeCargo,omitempty"`
}