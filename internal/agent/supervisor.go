/*
 * HXE - Host-based Process Execution Engine
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
	"fmt"

	"github.com/rangertaha/hxe/internal/log"
	"github.com/rs/zerolog"
)

// Supervisor manages the process lifecycle
type Supervisor struct {
	log zerolog.Logger
}

// NewSupervisor creates a new Supervisor instance
func NewSupervisor(agent *Agent) (*Supervisor, error) {
	pm := &Supervisor{
		log: log.With().Logger(),
	}

	if err := pm.Init(); err != nil {
		return nil, err
	}

	return pm, nil
}

// Init initializes the process manager
func (pm *Supervisor) Init() error {
	pm.log.Info().Msg("initializing process manager")

	return nil
}

// StartSupervisor starts the supervisor lifecycle
func (pm *Supervisor) StartSupervisor() error {
	pm.log.Info().Msg("starting supervisor")
	return nil
}

// StopSupervisor stops the supervisor lifecycle
func (pm *Supervisor) StopSupervisor() {
	pm.log.Info().Msg("stopping supervisor")
}

// // Start starts the process manager
// func (pm *ProcessManager) Start() error {
// 	pm.log.Info().Msg("starting process manager")
// 	return nil
// }

// func (pm *ProcessManager) Stop() {
// 	pm.log.Info().Msg("stopping process manager")
// }

// ---------------------------------------------------

// List lists all configured services
func (pm *Supervisor) List(name ...string) error {
	pm.log.Info().Msg("listing services")
	// TODO: Implement service listing
	fmt.Println("Services:")
	fmt.Println("  No services configured")
	return nil
}

// Start starts a service by name
func (pm *Supervisor) Start(name string) error {
	pm.log.Info().Str("service", name).Msg("starting service")
	// TODO: Implement service starting
	fmt.Printf("Starting service: %s\n", name)
	return nil
}

// Stop stops a service by name
func (pm *Supervisor) Stop(name string) error {
	pm.log.Info().Str("service", name).Msg("stopping service")
	// TODO: Implement service stopping
	fmt.Printf("Stopping service: %s\n", name)
	return nil
}

// Restart restarts a service by name
func (pm *Supervisor) Restart(name string) error {
	pm.log.Info().Str("service", name).Msg("restarting service")
	// TODO: Implement service restarting
	fmt.Printf("Restarting service: %s\n", name)
	return nil
}

// Status shows status of a specific service
func (pm *Supervisor) Status(name string) error {
	pm.log.Info().Str("service", name).Msg("showing service status")
	// TODO: Implement status for specific service
	fmt.Printf("Status of service '%s': Not found\n", name)
	return nil
}

// TailLogs follows the logs of a service
func (pm *Supervisor) Tail(name string, lines int, follow bool) error {
	pm.log.Info().Str("service", name).Int("lines", lines).Bool("follow", follow).Msg("tailing service logs")
	// TODO: Implement log tailing
	fmt.Printf("Tailing logs for service '%s' (lines: %d, follow: %t)\n", name, lines, follow)
	return nil
}

// Reload reloads the configuration
func (pm *Supervisor) Reload(name string) error {
	pm.log.Info().Msg("reloading configuration")
	// TODO: Implement configuration reloading
	fmt.Println("Reloading configuration...")
	return nil
}

// Enable enables a service to start automatically
func (pm *Supervisor) Enable(name string) error {
	pm.log.Info().Str("service", name).Msg("enabling service")
	// TODO: Implement service enabling
	fmt.Printf("Enabling service: %s\n", name)
	return nil
}

// Disable disables a service from starting automatically
func (pm *Supervisor) Disable(name string) error {
	pm.log.Info().Str("service", name).Msg("disabling service")
	// TODO: Implement service disabling
	fmt.Printf("Disabling service: %s\n", name)
	return nil
}

// Shell opens an interactive shell for a service
func (pm *Supervisor) Shell(name string) error {
	pm.log.Info().Str("service", name).Msg("opening shell for service")
	// TODO: Implement shell opening
	fmt.Printf("Opening shell for service: %s\n", name)
	return nil
}
