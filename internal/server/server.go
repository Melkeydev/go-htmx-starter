package server

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var port = 8080

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increase() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) Decrease() {
	c.mu.Lock()
	c.value--
	c.mu.Unlock()
}

func (c *Counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

type Server struct {
	port    int
	Counter *Counter
}

func NewServer() *http.Server {

	NewServer := &Server{
		port:    port,
		Counter: &Counter{},
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
