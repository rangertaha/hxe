package agent

import (
	"github.com/rangertaha/hxe/internal/api"
	"github.com/rangertaha/hxe/internal/log"
)

// NewAPIServer creates a new API server
func NewAPIServer(a *Agent) (*APIServer, error) {
	svc := &APIServer{
		log:    log.With().Logger(),
		router: api.New(),
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
		if err := a.router.Start(":8080"); err != nil {
			a.log.Fatal().Err(err).Msg("failed to start api")
		}
	}()

	return nil
}

func (a *APIServer) Stop() {
	a.log.Info().Msg("stopping api")

	a.router.Close()
}
