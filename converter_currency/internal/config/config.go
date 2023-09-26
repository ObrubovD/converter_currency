package config

import (
	"converter_currency/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadConf() (models.Config, error) {

	var config models.Config

	// Чтение файла config.json
	file, err := ioutil.ReadFile("/converter_currency/config.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return config, err
	}

	// Декодирование JSON-структуры в структуру данных Go
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("Ошибка декодирования JSON:", err)
		return config, err
	}

	return config, nil
}
