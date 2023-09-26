package main

import (
	"converter_currency/internal/config"
	"converter_currency/pkg/api"
	"converter_currency/pkg/converter"
	"converter_currency/pkg/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server is listening...")

	configValue, err := config.ReadConf()
	if err != nil {
		fmt.Println("Ошибка чтения Config", err)
		return
	}

	log.Fatal(http.ListenAndServe(configValue.Hostname, nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Получаем значения числа и состояния чекбокса из параметров запроса
	curOld := r.URL.Query().Get("currencyOld")
	curNew := r.URL.Query().Get("currencyNew")
	amount := r.URL.Query().Get("amount")

	if amount == "" {
		fmt.Fprintf(w, "0")
		return
	}
	formDataReq := models.FormDataRequest{
		CurrencyOld: curOld,
		CurrencyNew: curNew,
		Amount:      amount,
	}
	// Обрабатываем число
	result := processNumber(formDataReq)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprintf(w, result)
}

func processNumber(formDataReq models.FormDataRequest) string {

	res, err := strconv.Atoi(formDataReq.Amount)

	if formDataReq.CurrencyNew == formDataReq.CurrencyOld {
		return formDataReq.Amount
	}

	if err != nil {
		fmt.Println("Ошибка преобразования строки в число:", err)
		return ""
	}

	body := api.Get_currency(formDataReq, res)

	result := converter.Convert_value(body, formDataReq.CurrencyNew, res)

	str := strconv.Itoa(result)

	return "" + str
}
