/*
 * HXE - Host-based Process Execution Agent
 * Copyright (C) 2025 Rangertaha <rangertaha@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package engine

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/rangertaha/hxe/internal/api"
	"github.com/rangertaha/hxe/internal/config"
	"github.com/rangertaha/hxe/internal/log"
	"github.com/rangertaha/hxe/internal/modules/alerts"
	"github.com/rangertaha/hxe/internal/modules/messaging"
	"github.com/rangertaha/hxe/internal/modules/metrics"
	"github.com/rangertaha/hxe/internal/modules/schedules"
	"github.com/rangertaha/hxe/internal/modules/services"
	"github.com/rs/zerolog"
)

// Agent is the main struct for the hxe backend
type Agent struct {
	Config    *config.Config    // Configuration
	Schedules *schedules.Server // Schedules service
	Messages  *messaging.Server // Messaging service
	Services  *services.Server  // Service service
	Alerts    *alerts.Server    // Alerts service
	Metrics   *metrics.Server   // Metrics service
	API       *api.Server       // API service

	sig  chan os.Signal
	conf chan config.Config
	done chan struct{}
	log  zerolog.Logger
	ns   *server.Server
	nc   *nats.Conn
}

// New creates a new Agent instance
func New(cfg *config.Config) (agent *Agent, err error) {
	agent = &Agent{
		Config: cfg,
		sig:    make(chan os.Signal, 1),
		done:   make(chan struct{}),
		conf:   make(chan config.Config),
		log:    log.With().Logger(),
	}

	// Create messaging with server options
	agent.ns, agent.nc, err = NewMessaging(&cfg.Mq)
	if err != nil {
		agent.log.Error().Err(err).Msg("failed to create messaging service")
		return nil, err
	}

	if err = agent.Init(); err != nil {
		return nil, err
	}

	return
}

// Init initializes the agent
func (a *Agent) Init() (err error) {
	a.log.Info().Msg("initializing agent")

	// Create messaging with server options
	a.Schedules, err = schedules.New(a.nc)
	if err != nil {
		return err
	}

	// Create process manager
	a.Services, err = services.New(a.nc)
	if err != nil {
		return err
	}

	// Create alerts
	a.Alerts, err = alerts.New(a.nc)
	if err != nil {
		return err
	}

	// Create metrics
	a.Metrics, err = metrics.New(a.nc)
	if err != nil {
		return err
	}

	// Create api server
	a.Config.API.NC = a.nc
	a.API, err = api.New(&a.Config.API)
	if err != nil {
		return err
	}

	return nil
}

// Stop the agent
func (a *Agent) Stop() {
	log.Info().Msg("stopping agent")

	a.Schedules.Stop()
	a.Services.Stop()
	a.Alerts.Stop()
	a.Metrics.Stop()
	a.API.Stop()
}

// Start the agent
func (a *Agent) Start() error {
	a.log.Info().Msg("starting agent")

	// Start all services concurrently using goroutines
	var wg sync.WaitGroup

	// Start scheduler service
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := a.Schedules.Start(); err != nil {
			a.log.Error().Err(err).Msg("failed to start scheduler service")
		}
		log.Debug().Msg("scheduler service stopped")
	}()

	// Start services manager
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := a.Services.Start(); err != nil {
			a.log.Error().Err(err).Msg("failed to start services manager")
		}
		log.Debug().Msg("services manager stopped")
	}()

	// Start alerts service
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := a.Alerts.Start(); err != nil {
			a.log.Error().Err(err).Msg("failed to start alerts service")
		}
		log.Debug().Msg("alerts service stopped")
	}()

	// Start metrics service
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := a.Metrics.Start(); err != nil {
			a.log.Error().Err(err).Msg("failed to start metrics service")
		}
		log.Debug().Msg("metrics service stopped")
	}()

	// Start API server
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := a.API.Start(); err != nil {
			a.log.Error().Err(err).Msg("failed to start API server")
		}
		log.Debug().Msg("API server stopped")
	}()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// Setup signal handling
	signal.Notify(a.sig, syscall.SIGINT, syscall.SIGTERM)

	// Wait for shutdown signal
	select {
	case sig := <-a.sig:
		log.Debug().Msgf("received system signal %v, initiating shutdown", sig)
		a.Stop()
	case <-a.done:
		log.Debug().Msg("received done signal, initiating shutdown")
		a.Stop()
	}
	// }()

	// // Wait for all services to complete
	// wg.Wait()
	log.Debug().Msg("agent shutdown complete")

	// // Close messaging service
	a.nc.Close()
	a.ns.Shutdown()

	return nil
}
