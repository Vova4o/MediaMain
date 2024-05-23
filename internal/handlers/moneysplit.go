package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/Vova4o/MediaMain/internal/models"
	"github.com/gin-gonic/gin"
)

// Handler структура, которая содержит методы для обработки запросов
type Handler struct {
	Service Servicer
}

// Servicer интерфейс, который содержит методы для работы с бизнес-логикой
type Servicer interface {
	SplitMoney(models.Banknotes) ([][]int, error)
}

// New функция, которая создает новый экземпляр структуры Handler
func New(service Servicer) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes функция, которая настраивает маршруты
func SetupRoutes(mux *gin.Engine, h *Handler) {
	mux.POST("/split", h.Split)
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Log request details
		log.Printf("Request: %s %s %s %s %s %d %s",
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL,
			c.Request.Proto,
			c.Request.UserAgent(),
			c.Writer.Status(),
			time.Since(start),
		)
	}
}

// Split функция, которая обрабатывает запрос на размен суммы и возвращает ответ пользователю
func (h *Handler) Split(c *gin.Context) {
	var req models.Banknotes

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	banknotes, err := h.Service.SplitMoney(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"banknotes": banknotes})
}
