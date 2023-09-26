package converter

import (
	"converter_currency/pkg/models"
	"fmt"
)

func Convert_value(response models.Response, CurrencyNew string, amount int) int {

	// Получение значения по ключу "CurrencyNew"
	usdRate := response.Rates[CurrencyNew]
	res := int(usdRate * float64(amount))
	fmt.Println("Значение CurrencyNew:", res)
	return res
}
