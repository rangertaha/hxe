package schedules

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rangertaha/hxe/internal/log"
	"github.com/rs/zerolog"
)

type Server struct {
	log zerolog.Logger
	nc  *nats.Conn
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
	s.log.Info().Msg("initializing scheduling service")
	return nil
}

func (s *Server) Load() error {
	s.log.Info().Msg("configuring scheduling service")
	return nil
}

func (s *Server) Start() error {
	s.log.Info().Msg("starting scheduling service")
	for {
		select {
		case <-s.done:
			s.log.Info().Msg("stopping scheduling service")
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
	s.log.Info().Msg("restarting scheduling service")
	return nil
}

func (s *Server) Status() error {
	s.log.Info().Msg("scheduling service status")
	return nil
}
