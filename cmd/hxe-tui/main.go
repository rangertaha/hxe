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

package main

import (
	"flag"
	"log"

	"github.com/rangertaha/hxe/internal/tui"
)

func main() {
	var mockMode bool
	flag.BoolVar(&mockMode, "mock", false, "Run in mock mode with dummy data (no NATS required)")
	flag.Parse()

	var err error
	if mockMode {
		log.Println("Starting HXE TUI in mock mode with dummy data...")
		err = tui.RunWithMockMode(true)
	} else {
		log.Println("Starting HXE TUI...")
		err = tui.Run()
	}

	if err != nil {
		log.Fatal("Failed to run TUI:", err)
	}
}
