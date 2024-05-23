package models

// Banknotes структура, которая содержит сумму и номиналы купюр
type Banknotes struct {
	Amount    int   `json:"amount"`
	Banknotes []int `json:"banknotes"`
}
