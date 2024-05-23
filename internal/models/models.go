package models

type Banknotes struct {
	Amount    int   `json:"amount"`
	Banknotes []int `json:"banknotes"`
}
