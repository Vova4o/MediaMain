package services

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/Vova4o/MediaMain/internal/models"
)

func TestServices_SplitMoney(t *testing.T) {
	s := New()
	tests := []struct {
		name    string
		amount  models.Banknotes
		want    [][]int
		wantErr bool
	}{
		{
			name: "Test 1",
			amount: models.Banknotes{
				Amount:    400,
				Banknotes: []int{5000, 2000, 1000, 500, 200, 100, 50},
			},
			want: [][]int{
				{200, 200},
				{200, 100, 100},
				{200, 100, 50, 50},
				{200, 50, 50, 50, 50},
				{100, 100, 100, 100},
				{100, 100, 100, 50, 50},
				{100, 100, 50, 50, 50, 50},
				{100, 50, 50, 50, 50, 50, 50},
				{50, 50, 50, 50, 50, 50, 50, 50},
			},
			wantErr: false,
		},
		{
			name: "Test 2",
			amount: models.Banknotes{
				Amount:    0,
				Banknotes: []int{5000, 2000, 1000, 500, 200, 100, 50},
			},
			want:    [][]int{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.SplitMoney(tt.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("Services.SplitMoney() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			sort.Slice(got, func(i, j int) bool {
				return fmt.Sprint(got[i]) < fmt.Sprint(got[j])
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return fmt.Sprint(tt.want[i]) < fmt.Sprint(tt.want[j])
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.SplitMoney() = %v, want %v", got, tt.want)
			}
		})
	}
}
