package main

// columnNamesDetailedReport := map[string]int{
// 	"No.":                                1,
// 	"Delivery number":                    2,
// 	"Item":                               3,
// 	"Nomenclature code":                  4,
// 	"Brand":                              5,
// 	"The article of the supplier":        6,
// 	"Name":                               7,
// 	"Size":                               8,
// 	"Barkod":                             9,
// 	"Type of document":                   10,
// 	"Justification for payment":          11,
// 	"The date of the order by the buyer": 12,
// 	"Date of sale":                       13,
// 	"Amount":                             14,
// 	"Retail price":                       15,
// 	"Wildberris sold goods (pr)":         16,
// 	"Coordinated grocery discount, %":    17,
// 	"Promo code %":                       18,
// 	"The final agreed discount":          19,
// 	"Retail price, taking into account the agreed discount": 20,
// 	"Buyer discount (SPP)":          21,
// 	"KVV size, %":                   22,
// 	"KVV size without VAT, % basic": 23,
// 	"Final KVV without VAT, %":      24,
// 	"Remuneration from sales before deducting the services of an attorney, without VAT": 25,
// 	"Compensation for the issuance and return of goods to PVZ":                          26,
// 	"Acquiring": 27,
// 	"Remuneration Wildberry (BB), without VAT":         28,
// 	"VAT with rewards Wildberry":                       29,
// 	"To the transfer of the seller for the goods sold": 30,
// 	"The number of deliveryT":                          31,
// 	"he number of return":                              32,
// 	"Services for the delivery of goods to the buyer":  33,
// 	"The total amount of fines":                        34,
// 	"Surcharges":                                       35,
// 	"Types of logistics, fines and surcharges":         36,
// 	"Sticker MP":                      37,
// 	"The name of the acquiring bank":  38,
// 	"Office number":                   39,
// 	"The name of the delivery office": 40,
// 	"TIN of the partner":              41,
// 	"Partner":                         42,
// 	"Stock":                           43,
// 	"A country":                       44,
// 	"Type of boxes":                   45,
// 	"Number of customs declaration":   46,
// 	"Marking code":                    47,
// 	"ShK":                             48,
// 	"Rid":                             49,
// 	"Srid":                            50,
// }

// columnNamesDetailedReport := map[string]int{
// 	"№":                       1,
// 	"Номер поставки":          2,
// 	"Предмет":                 3,
// 	"Код номенклатуры":        4,
// 	"Бренд":                   5,
// 	"Артикул поставщика":      6,
// 	"Название":                7,
// 	"Размер":                  8,
// 	"Баркод":                  9,
// 	"Тип документа":           10,
// 	"Обоснование для оплаты":  11,
// 	"Дата заказа покупателем": 12,
// 	"Дата продажи":            13,
// 	"Кол-во":                  14,
// 	"Цена розничная":          15,
// 	"Вайлдберриз реализовал Товар (Пр)":    16,
// 	"Согласованный продуктовый дисконт, %": 17,
// 	"Промокод %": 18,
// 	"Итоговая согласованная скидка":                19,
// 	"Цена розничная с учетом согласованной скидки": 20,
// 	"Скидка постоянного Покупателя (СПП)":          21,
// 	"Размер кВВ, %":                  22,
// 	"Размер  кВВ без НДС, % Базовый": 23,
// 	"Итоговый кВВ без НДС, %":        24,
// 	"Вознаграждение с продаж до вычета услуг поверенного, без НДС": 25,
// 	"Возмещение за выдачу и возврат товаров на ПВЗ":                26,
// 	"Возмещение расходов по эквайрингу":                            27,
// 	"Вознаграждение Вайлдберриз (ВВ), без НДС":                     28,
// 	"НДС с Вознаграждения Вайлдберриз":                             29,
// 	"К перечислению Продавцу за реализованный Товар":               30,
// 	"Количество доставок":                                          31,
// 	"Количество возврата":                                          32,
// 	"Услуги по доставке товара покупателю":                         33,
// 	"Общая сумма штрафов":                                          34,
// 	"Доплаты":                                                      35,
// 	"Виды логистики, штрафов и доплат":                             36,
// 	"Стикер МП": 37,
// 	"Наименование банка эквайринга": 38,
// 	"Номер офиса":                 39,
// 	"Наименование офиса доставки": 40,
// 	"ИНН партнера":                41,
// 	"Партнер":                     42,
// 	"Склад":                       43,
// 	"Страна":                      44,
// 	"Тип коробов":                 45,
// 	"Номер таможенной декларации": 46,
// 	"Код маркировки":              47,
// 	"ШК":                          48,
// 	"Rid":                         49,
// 	"Srid":                        50,
// }
