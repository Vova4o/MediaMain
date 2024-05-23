package services

import (
	"errors"
	"log"

	"github.com/Vova4o/MediaMain/internal/models"
)

// Services структура, которая содержит методы для работы с бизнес-логикой
type Services struct{}

// New функция, которая создает новый экземпляр структуры Services
func New() *Services {
	return &Services{}
}

// splitMoney функция помошник, которая находит все возможные комбинации купюр, которые можно использовать для размена суммы
func (s *Services) splitMoney(amount int, banknotes []int, index int, combination []int, result *[][]int) {
	if amount == 0 {
		// If the amount is 0, we have found a valid combination
		combinationCopy := make([]int, len(combination))
		copy(combinationCopy, combination)
		*result = append(*result, combinationCopy)
		return
	}

	if index == len(banknotes) || amount < 0 {
		return
	}

	s.splitMoney(amount, banknotes, index+1, combination, result)

	s.splitMoney(amount-banknotes[index], banknotes, index, append(combination, banknotes[index]), result)
}

// SplitMoney функция, которая принимает сумму и номиналы купюр и возвращает все возможные комбинации купюр, которые можно использовать для размена суммы
func (s *Services) SplitMoney(amount models.Banknotes) ([][]int, error) {
	log.Println("SplitMoney service called")
	sum := amount.Amount
	banknotes := amount.Banknotes
	result := make([][]int, 0)

	if sum == 0 || banknotes == nil || len(banknotes) == 0 {
		log.Printf("Invalid input: sum=%d, banknotes=%v\n", sum, banknotes)
		return result, errors.New("invalid input")
	}

	s.splitMoney(sum, banknotes, 0, []int{}, &result)

	return result, nil
}
