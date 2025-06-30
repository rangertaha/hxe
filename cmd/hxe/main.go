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
	"context"
	"fmt"
	"os"

	"github.com/rangertaha/hxe/internal"
	"github.com/rangertaha/hxe/internal/config"
	"github.com/rangertaha/hxe/internal/engine"
	"github.com/urfave/cli/v3"
)

var (
	CfgOption *config.Config
	HxeConfig *config.Config
	newAgent  *engine.Agent
	serverURL string
	username  string
	password  string
	err       error
)

func main() {
	app := &cli.Command{
		Name:                  "hxe",
		Usage:                 "Hxe task execution tool",
		Description:           `Hxe task execution tool`,
		Version:               internal.VERSION,
		Authors:               []any{"Rangertaha"},
		Copyright:             "Rangertaha",
		EnableShellCompletion: true,
		Suggest:               true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Hidden:  false,
				Usage:   "Configuration from `FILE`",
			},
			&cli.BoolFlag{
				Name:   "debug",
				Value:  false,
				Hidden: false,
				Usage:  "Log debug messages for development",
			},
			&cli.StringFlag{
				Name:        "username",
				Aliases:     []string{"u"},
				Value:       "admin",
				Hidden:      false,
				Usage:       "Username for authentication",
				Destination: &username,
			},
			&cli.StringFlag{
				Name:        "password",
				Aliases:     []string{"p"},
				Value:       "password",
				Hidden:      false,
				Usage:       "Password for authentication",
				Destination: &password,
			},
			&cli.StringFlag{
				Name:        "url",
				Value:       "http://0.0.0.0:9090",
				Hidden:      false,
				Usage:       "Hxe API server URL",
				Destination: &serverURL,
			},
		},
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			// Load configuration file
			cliOptions := config.CliOptions(ctx, cmd)
			
			// Create new config
			if HxeConfig, err = config.New(cliOptions); err != nil {
				return ctx, err
			}

			return ctx, nil
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			cli.ShowAppHelpAndExit(cmd, 1)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:        "server",
				Usage:       "Start the HXE server",
				Description: `Start the HXE server.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					// Create new agent
					if newAgent, err = engine.New(HxeConfig); err != nil {
						return err
					}

					// Start the agent
					return newAgent.Start()
				},
			},
			serviceCmd,
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
