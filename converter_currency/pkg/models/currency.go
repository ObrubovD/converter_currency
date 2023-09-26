package models

// Структура для декодирования JSON
type Response struct {
	Rates map[string]float64 `json:"rates"`
}
