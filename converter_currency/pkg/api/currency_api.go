package api

import (
	"converter_currency/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get_currency(formDataRequest models.FormDataRequest, amount int) models.Response {

	var response models.Response
	// Создаем HTTP-клиент
	client := &http.Client{}

	strAdr := "https://api.frankfurter.app/latest?from="
	strAdr += formDataRequest.CurrencyOld

	// Создаем GET-запрос
	req, err := http.NewRequest("GET", strAdr, nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return response
	}

	// Выполняем запрос
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return response
	}

	// Закрываем тело ответа после завершения
	defer resp.Body.Close()

	// Читаем содержимое ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return response
	}

	ex := json.Unmarshal(body, &response)
	if ex != nil {
		fmt.Println("Ошибка декодирования JSON:", err)
		return response
	}

	return response
}
