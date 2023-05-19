package main

import (
	"WBEconomics/internal/types"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

func main() {

	// reports := loadDetailedReport()
	report := loadFinanceReport()
	fmt.Printf("%+v\n", report)
}

func loadFinanceReport() types.FinancialReport {
	f, err := excelize.OpenFile(FileNameFinancialReport)
	if err != nil {
		log.Fatal(err)
	}

	commision, err := f.GetCellValue(SheetNameFinancialReport, CellComissionsFinancialReport)
	shop_sold_amount, err := f.GetCellValue(SheetNameFinancialReport, CellShopSoldAmountFinancialReport)
	start_period, err := f.GetCellValue(SheetNameFinancialReport, CellStartPeriodFinancialReport)
	end_period, err := f.GetCellValue(SheetNameFinancialReport, CellEndPeriodFinancialReport)

	report := types.FinancialReport{
		Commision:        getFloatValue(commision),
		Shop_sold_amount: getFloatValue(shop_sold_amount),
		Start_period:     getDateValue(start_period),
		End_period:       getDateValue(end_period),
	}
	return report
}

func loadDetailedReport() []types.DetailedOperation {
	f, err := excelize.OpenFile(FileNameDetailedReport)
	if err != nil {
		log.Fatal(err)
	}

	// Read the column names from the first row
	rows, err := f.GetRows(SheetNameDetailedReport)
	if err != nil {
		log.Fatal(err)
	}

	// Map the column names to their respective indices
	columnIndices := make(map[string]int)
	for i, cell := range rows[0] {
		columnIndices[cell] = i
	}
	// Read data from each row and store in DetailedReport struct
	var reports []types.DetailedOperation
	if err != nil {
		log.Fatal(err)
	}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if len(row) < len(columnIndices) {
			for i := 1; i <= len(columnIndices)-len(row); i++ {
				row = append(row, " ")
			}
		}
		report := types.DetailedOperation{
			No:                         getIntValue(row, columnIndices["№"]),
			DeliveryNumber:             getIntValue(row, columnIndices["Номер поставки"]),
			Item:                       row[columnIndices["Предмет"]],
			NomenclatureCode:           getIntValue(row, columnIndices["Код номенклатуры"]),
			Brand:                      row[columnIndices["Бренд"]],
			ArticleOfSupplier:          row[columnIndices["Артикул поставщика"]],
			Name:                       row[columnIndices["Название"]],
			Size:                       row[columnIndices["Размер"]],
			Barkod:                     getIntValue(row, columnIndices["Баркод"]),
			TypeOfDocument:             row[columnIndices["Тип документа"]],
			JustificationForPayment:    row[columnIndices["Обоснование для оплаты"]],
			OrderDateByBuyer:           getDateValue(row[columnIndices["Дата заказа покупателем"]]),
			DateOfSale:                 getDateValue(row[columnIndices["Дата продажи"]]),
			Amount:                     getIntValue(row, columnIndices["Кол-во"]),
			RetailPrice:                getFloatValue(row[columnIndices["Цена розничная"]]),
			WildberriesSoldGoods:       getFloatValue(row[columnIndices["Вайлдберриз реализовал Товар (Пр)"]]),
			CoordinatedGroceryDiscount: getFloatValue(row[columnIndices["Согласованный продуктовый дисконт, %"]]),
			PromoCode:                  getFloatValue(row[columnIndices["Промокод %"]]),
			FinalAgreedDiscount:        getFloatValue(row[columnIndices["Итоговая согласованная скидка"]]),
			RetailPriceWithDiscount:    getFloatValue(row[columnIndices["Цена розничная с учетом согласованной скидки"]]),
			BuyerDiscount:              getFloatValue(row[columnIndices["Скидка постоянного Покупателя (СПП)"]]),
			KVVSize:                    getFloatValue(row[columnIndices["Размер кВВ, %"]]),
			KVVSizeWithoutVAT:          getFloatValue(row[columnIndices["Размер  кВВ без НДС, % Базовый"]]),
			FinalKVVWithoutVAT:         getFloatValue(row[columnIndices["Итоговый кВВ без НДС, %"]]),
			RemunerationFromSales:      getFloatValue(row[columnIndices["Вознаграждение с продаж до вычета услуг поверенного, без НДС"]]),
			Compensation:               getFloatValue(row[columnIndices["Возмещение за выдачу и возврат товаров на ПВЗ"]]),
			Acquiring:                  getFloatValue(row[columnIndices["Возмещение расходов по эквайрингу"]]),
			RemunerationWildberry:      getFloatValue(row[columnIndices["Вознаграждение Вайлдберриз (ВВ), без НДС"]]),
			VATWithRewards:             getFloatValue(row[columnIndices["НДС с Вознаграждения Вайлдберриз"]]),
			TransferToSeller:           getFloatValue(row[columnIndices["К перечислению Продавцу за реализованный Товар"]]),
			DeliveryNumberT:            getIntValue(row, columnIndices["Количество доставок"]),
			ReturnNumber:               getIntValue(row, columnIndices["Количество возврата"]),
			ServicesDeliveryToBuyer:    getFloatValue(row[columnIndices["Услуги по доставке товара покупателю"]]),
			TotalFines:                 getFloatValue(row[columnIndices["Общая сумма штрафов"]]),
			Surcharges:                 getFloatValue(row[columnIndices["Доплаты"]]),
			LogisticsTypes:             row[columnIndices["Виды логистики, штрафов и доплат"]],
			StickerMP:                  row[columnIndices["Стикер МП"]],
			AcquiringBankName:          row[columnIndices["Наименование банка эквайринга"]],
			OfficeNumber:               row[columnIndices["Номер офиса"]],
			DeliveryOfficeName:         row[columnIndices["Наименование офиса доставки"]],
			PartnerTIN:                 row[columnIndices["ИНН партнера"]],
			Partner:                    row[columnIndices["Партнер"]],
			Stock:                      row[columnIndices["Склад"]],
			Country:                    row[columnIndices["Страна"]],
			BoxType:                    row[columnIndices["Тип коробов"]],
			NumberOfCustoms:            row[columnIndices["Номер таможенной декларации"]],
			MarkingCode:                row[columnIndices["Код маркировки"]],
			ShK:                        getIntValue(row, columnIndices["ШК"]),
			Rid:                        row[columnIndices["Rid"]],
			Srid:                       row[columnIndices["Srid"]],
		}
		reports = append(reports, report)
	}
	return reports
}

func getIntValue(row []string, columnIndex int) int {
	if columnIndex < 1 || columnIndex > len(row) {
		return 0
	}

	value := row[columnIndex]
	if value == "" {
		return 0
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

// Helper function to convert string to float
func getFloatValue(value string) float64 {
	if value == "" {
		return 0.0
	}
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// Helper function to convert string to date
func getDateValue(value string) time.Time {
	if value == "" {
		return time.Time{}
	}
	result, err := time.Parse("2006-01-02", value)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
