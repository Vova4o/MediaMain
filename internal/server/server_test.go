package server

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New()
	assert.NotNil(t, s)
	assert.NotNil(t, s.Handler)
	assert.NotEmpty(t, s.Address)
}

func TestNewServer(t *testing.T) {
	s := New()
	srv := s.NewServer()

	assert.NotNil(t, srv)
	assert.Equal(t, s.Address, srv.Addr)
	assert.Equal(t, s.Handler, srv.Handler)
}

func TestRun(t *testing.T) {
	s := New()
	go func() {
		err := s.Run()
		assert.Nil(t, err)
	}()

	// Let's wait a bit for the server to start
	time.Sleep(1 * time.Second)

	// Send a request to the server
	resp, err := http.Get("http://" + s.Address)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}
