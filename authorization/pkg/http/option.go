package http

import (
	"net"
	"time"
)

type Option func(*Server)

func FullAddress(addr string, port string) Option {
	return func(s *Server) {
		s.httpServer.Addr = net.JoinHostPort(addr, port)
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.httpServer.ReadTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.httpServer.WriteTimeout = timeout
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
