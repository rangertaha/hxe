package agent

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/rangertaha/hxe/internal/api"
	"github.com/rangertaha/hxe/internal/config"
	"github.com/rangertaha/hxe/internal/log"
	"github.com/rs/zerolog"
)

type APIServer struct {
	config config.API
	broker *Broker
	log    zerolog.Logger
	router *echo.Echo
}

// NewAPIServer creates a new API server
func NewAPIServer(a *Agent) (*APIServer, error) {
	svc := &APIServer{
		config: a.Config.API,
		log:    log.With().Logger(),
		router: api.New(a.Broker),
	}

	if err := svc.Init(); err != nil {
		return nil, err
	}

	return svc, nil
}

func (a *APIServer) Init() error {
	a.log.Info().Msg("initializing api")

	return nil
}

// Start starts the task
func (a *APIServer) Start() error {
	a.log.Info().Msg("starting api")

	// Start server on port
	go func() {
		a.log.Info().Str("host", a.config.Host).Int("port", a.config.Port).Msg("starting api")
		if err := a.router.Start(fmt.Sprintf("%s:%d", a.config.Host, a.config.Port)); err != nil {
			a.log.Fatal().Err(err).Msg("failed to start api")
		}
	}()

	return nil
}

func (a *APIServer) Stop() {
	a.log.Info().Msg("stopping api")

	a.router.Close()
}
