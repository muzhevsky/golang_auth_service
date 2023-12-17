package http

import (
	"context"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 15 * time.Second
	_defaultWriteTimeout    = 15 * time.Second
	_defaultAddr            = "127.0.0.1:8000"
	_defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	httpServer      *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func New(handler http.Handler, opts ...Option) *Server {
	httpSrv := &http.Server{
		Handler:      handler,
		Addr:         _defaultAddr,
		WriteTimeout: _defaultWriteTimeout,
		ReadTimeout:  _defaultReadTimeout,
	}
	s := &Server{
		httpServer:      httpSrv,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}
func (s *Server) start() {
	go func() {
		s.notify <- s.httpServer.ListenAndServe()
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}
