package main

import (
	"sync"

	"github.com/essial/AthenaEngine/aehttp"
	"github.com/essial/AthenaEngine/aenetwork"
)

// The AthenaEngine server.
type Server struct {
	isRunning bool
}

// Creates an instance of the AthenaEngine server.
func CreateServer() *Server {
	result := &Server{}
	return result
}

// Shuts down the server and closes it.
func (s *Server) Close() {
	s.isRunning = false
}

// Runs the server.
func (s *Server) Run() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		aehttp.Serve()
		wg.Done()
	}()
	wg.Wait()
}

// Called when a client connects.
func (s *Server) OnConnect(connection *aenetwork.ClientConnection) {

}

// Called when a client disconnects.
func (s *Server) OnDisconnect(connection *aenetwork.ClientConnection) {

}

func (s *Server) IsRunning() bool {
	return s.isRunning
}
