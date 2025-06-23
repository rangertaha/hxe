/*
Copyright © 2025 Rangertaha <rangertaha@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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
