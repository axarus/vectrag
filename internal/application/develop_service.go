package application

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

type ListenerProvider interface {
	Listen(basePort int) (port int, ln net.Listener, err error)
}

type ServerStarter interface {
	Start(register func(mux *http.ServeMux), ln net.Listener) *http.Server
}

type AdminHandlerProvider interface {
	Handler() http.Handler
}

type DevelopService struct {
	listener ListenerProvider
	server   ServerStarter
	admin    AdminHandlerProvider
}

func NewDevelopService(listener ListenerProvider, server ServerStarter, admin AdminHandlerProvider) *DevelopService {
	return &DevelopService{listener: listener, server: server, admin: admin}
}

func (s *DevelopService) Start(basePort int, host string) (url string, shutdown func(ctx context.Context) error, err error) {
	port, ln, err := s.listener.Listen(basePort)
	if err != nil {
		return "", nil, err
	}

	srv := s.server.Start(func(mux *http.ServeMux) {
		mux.Handle("/", s.admin.Handler())
	}, ln)

	url = fmt.Sprintf("http://%s:%d", host, port)
	shutdown = srv.Shutdown
	return url, shutdown, nil
}
