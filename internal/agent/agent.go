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

package agent

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/rangertaha/hxe/internal/config"
	"github.com/rangertaha/hxe/internal/interfaces"
	"github.com/rangertaha/hxe/internal/log"
	"github.com/rs/zerolog"

	"github.com/rangertaha/hxe/internal/services"
	_ "github.com/rangertaha/hxe/internal/services/all"
)

// Agent is the main struct for the hxe backend
type Agent struct {
	Services []interfaces.Service

	conf *config.AgentConfig
	sig  chan os.Signal
	done chan struct{}
	log  zerolog.Logger
	ns   *server.Server
	nc   *nats.Conn
}

// New creates a new Agent instance
func New(cfg *config.AgentConfig) (agent *Agent, err error) {
	agent = &Agent{
		conf: cfg,
		sig:  make(chan os.Signal, 1),
		done: make(chan struct{}),
		log:  log.With().Logger(),
		// Services: make([]interfaces.Service, len(cfg.Services)),
	}

	// Create messaging with server options
	agent.ns, agent.nc, err = NewMessaging(&cfg.Server)
	if err != nil {
		agent.log.Error().Err(err).Msg("failed to create messaging service")
		return nil, err
	}

	if err = agent.Load(); err != nil {
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

	for _, service := range a.Services {
		service.Init()
	}

	return nil
}

func (a *Agent) Load() (err error) {
	a.log.Info().Msg("loading services")

	for _, svc := range a.conf.Services {
		creator, err := services.Get(svc.ID)
		if err != nil {
			return err
		}

		srv := creator(a.nc)

		diags := gohcl.DecodeBody(svc.Config, config.CtxFunctions, srv)
		for _, diag := range diags {
			return errors.New(diag.Error())
		}

		a.Services = append(a.Services, srv)
	}

	return nil
}

// Stop the agent
func (a *Agent) Stop() {
	log.Info().Msg("stopping agent")

	for _, service := range a.Services {
		service.Stop()
	}

}

// Start the agent
func (a *Agent) Start() error {
	a.log.Info().Msg("starting agent")

	for _, service := range a.Services {
		service.Start()
	}

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

func NewMessaging(cfg *config.Server) (ns *server.Server, nc *nats.Conn, err error) {
	log.Info().Msg("initializing messaging service")
	opts := &server.Options{Port: cfg.Port, Host: cfg.Host}
	ns, err = server.NewServer(opts)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create messaging server: %w", err)
	}
	// Start server in background
	go ns.Start()

	// Wait for server to be ready
	if !ns.ReadyForConnections(5 * time.Second) {
		return nil, nil, fmt.Errorf("server not ready for connections")
	}

	clientOpts := []nats.Option{
		nats.Name("hxe-server"),
		nats.InProcessServer(ns),
		nats.FlusherTimeout(5 * time.Second),
	}
	nc, err = nats.Connect(ns.ClientURL(), clientOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to messaging server: %w", err)
	}

	log.Info().Msgf("connected to messaging server: %s", ns.ClientURL())

	return ns, nc, nil

}
