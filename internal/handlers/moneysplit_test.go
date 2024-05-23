package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Vova4o/MediaMain/internal/models"
	"github.com/gin-gonic/gin"
)

type mockService struct{}

func (m *mockService) SplitMoney(models.Banknotes) ([][]int, error) {
	return [][]int{{200, 200}}, nil
}

func TestHandler_Split(t *testing.T) {
	service := &mockService{}
	handler := New(service)
	router := gin.Default()
	SetupRoutes(router, handler)

	tests := []struct {
		name       string
		body       models.Banknotes
		wantStatus int
		wantBody   string
	}{
		{
			name: "Test 1",
			body: models.Banknotes{
				Amount:    400,
				Banknotes: []int{5000, 2000, 1000, 500, 200, 100, 50},
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"banknotes":[[200,200]]}`,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest(http.MethodPost, "/split", bytes.NewBuffer(body))
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			if resp.Code != tt.wantStatus {
				t.Errorf("Handler.Split() status = %v, want %v", resp.Code, tt.wantStatus)
			}

			if !strings.Contains(resp.Body.String(), tt.wantBody) {
				t.Errorf("Handler.Split() body = %v, want %v", resp.Body.String(), tt.wantBody)
			}
		})
	}
}
