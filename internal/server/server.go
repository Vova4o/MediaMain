package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Vova4o/Go-REST-API/internal/config"
	"github.com/Vova4o/Go-REST-API/internal/handlers"
	"github.com/Vova4o/Go-REST-API/internal/services"
	"github.com/gin-gonic/gin"
)

// Server структура, которая содержит адрес сервера и его обработчик
type Server struct {
	Address string
	Handler *gin.Engine
}

// New функция, которая создает новый экземпляр структуры Server
func New() *Server {
	addr := config.Address()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(handlers.LoggerMiddleware())

	serv := services.New()
	handl := handlers.New(serv) // pass the router to the handlers

	handlers.SetupRoutes(router, handl) // setup routes

	return &Server{
		Address: addr,
		Handler: router, // use the router as the server handler
	}
}

// NewServer функция, которая создает новый экземпляр http.Server
func (s *Server) NewServer() *http.Server {
	return &http.Server{
		Addr:    s.Address,
		Handler: s.Handler,
	}
}

// Run метод, который запускает сервер
func (s *Server) Run() error {
	srv := s.NewServer()
	log.Printf("Server is starting on %s\n", s.Address)
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s", err)
	}

	log.Println("Server exiting")

	return nil
}
