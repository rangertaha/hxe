package services

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rangertaha/hxe/internal/log"
	"github.com/rs/zerolog"
)

type Server struct {
	service *Service
	log     zerolog.Logger
	done    chan bool
}

func New(nc *nats.Conn) (server *Server, err error) {
	svc, err := NewService(nc)
	if err != nil {
		return nil, err
	}
	server = &Server{
		log:     log.With().Logger(),
		done:    make(chan bool),
		service: svc,
	}

	if err := server.Init(); err != nil {
		return nil, err
	}

	return
}

func (s *Server) Init() error {
	s.log.Info().Msg("initializing services")
	return nil
}

func (s *Server) Load() error {
	s.log.Info().Msg("configuring services")
	return nil
}

func (s *Server) Start() error {
	s.log.Info().Msg("starting services")
	for {
		select {
		case <-s.done:
			s.log.Info().Msg("stopping services")
			return nil
		default:
			time.Sleep(1 * time.Second)
		}
	}
	return nil
}

func (s *Server) Stop() error {
	close(s.done)
	return nil
}

func (s *Server) Restart() error {
	s.log.Info().Msg("restarting services")
	return nil
}

func (s *Server) Status() error {
	s.log.Info().Msg("services status")
	return nil
}
