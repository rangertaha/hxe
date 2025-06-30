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

package tui

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nats-io/nats.go"
	"github.com/rangertaha/hxe/internal/api/client"
	"github.com/rangertaha/hxe/internal/tui/app"
)

// Run starts the TUI application
func Run() error {
	return RunWithMockMode(false)
}

// RunWithMockMode starts the TUI application with optional mock mode
func RunWithMockMode(mockMode bool) error {
	var nc *nats.Conn
	var err error

	if !mockMode {
		// Connect to NATS only if not in mock mode
		nc, err = nats.Connect(nats.DefaultURL)
		if err != nil {
			return err
		}
		defer nc.Close()
	}

	// Create client (can be nil in mock mode)
	var c *client.Client
	if nc != nil {
		c = client.NewClient(nc)
	}

	// Create and run the TUI application
	p := tea.NewProgram(
		app.NewModel(c, mockMode),
		tea.WithAltScreen(),       // Use alternate screen buffer
		tea.WithMouseCellMotion(), // Turn on mouse support
		tea.WithMouseAllMotion(),  // Turn on mouse support for all motion
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running TUI: %v\n", err)
		os.Exit(1)
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal("Failed to run TUI:", err)
	}
}
