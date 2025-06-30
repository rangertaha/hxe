package metrics

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rangertaha/hxe/internal/log"
	"github.com/rs/zerolog"
)

type Server struct {
	nc  *nats.Conn
	log zerolog.Logger
	done chan bool
}

func New(nc *nats.Conn) (server *Server, err error) {
	server = &Server{
		nc:  nc,
		log: log.With().Logger(),
		done: make(chan bool),
	}

	if err := server.Init(); err != nil {
		return nil, err
	}

	return
}

func (s *Server) Init() error {
	s.log.Info().Msg("initializing metrics service")
	return nil
}

func (s *Server) Load() error {
	s.log.Info().Msg("configuring metrics service")
	return nil
}

func (s *Server) Start() error {
	s.log.Info().Msg("starting metrics service")
	for {
		select {
		case <-s.done:
			return nil
		default:
			time.Sleep(1 * time.Second)
		}
	}
	return nil
}

func (s *Server) Stop() error {
	s.log.Info().Msg("stopping metrics service")
	close(s.done)
	return nil
}

func (s *Server) Restart() error {
	s.log.Info().Msg("restarting metrics service")
	return nil
}

func (s *Server) Status() error {
	s.log.Info().Msg("metrics service status")
	return nil
}
