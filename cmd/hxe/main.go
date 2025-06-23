/*
 * HXE - Host-based Process Execution Engine
 * Copyright (C) 2025 rangertaha@gmail.com
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
	"strings"

	"github.com/rangertaha/hxe/internal"
	"github.com/rangertaha/hxe/internal/agent"
	"github.com/rangertaha/hxe/internal/config"
	"github.com/rangertaha/hxe/pkg/client"
	"github.com/urfave/cli/v3"
)

var (
	cfgOption *config.Config
	newAgent  *agent.Agent
	username  string
	password  string
	clientURL string
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
				Value:       "",
				Hidden:      false,
				Usage:       "Hxe API server URL",
				Destination: &clientURL,
			},
		},
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			// Load configuration file
			cfgFile := cmd.String("config")
			cliOptions := config.CliOptions(ctx, cmd)
			fileOptions := config.FileOption(cfgFile)

			// Create new config
			if cfgOption, err = config.New(fileOptions, cliOptions); err != nil {
				return ctx, err
			}
			if clientURL == "" {
				clientURL = cfgOption.API.Client.URL
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
					if newAgent, err = agent.New(cfgOption); err != nil {
						return err
					}

					// Start the agent
					return newAgent.Start()
				},
			},
			{
				Name:  "run",
				Usage: "Run a program",
				Description: `Run a program. If no program name is provided, 
				executes the command directly.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					hxeClient := client.NewClient(clientURL, username, password)
					if hxeClient, err = hxeClient.Login(); err != nil {
						return err
					}

					command := strings.Join(cmd.Args().Slice(), " ")
					prog, err := hxeClient.Program.Run(command)
					if err != nil {
						return fmt.Errorf("failed to run command: %w", err)
					}

					hxeClient.Program.PrintDetail(prog)
					return nil
				},
			},
			{
				Name:        "list",
				Usage:       "List all programs",
				Description: `List all programs with their current status.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					hxeClient := client.NewClient(clientURL, username, password)
					if hxeClient, err = hxeClient.Login(); err != nil {
						return err
					}

					programs, err := hxeClient.Program.List()
					if err != nil {
						return fmt.Errorf("failed to run command: %w", err)
					}

					hxeClient.Program.Print(programs)
					return nil
				},
			},
			{
				Name:        "start",
				Usage:       "Start a program",
				Description: `Start a program by name.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					hxeClient := client.NewClient(clientURL, username, password)
					if hxeClient, err = hxeClient.Login(); err != nil {
						return err
					}
					progs, err := hxeClient.Program.MultiStart(cmd.Args().Slice()...)
					if err != nil {
						return fmt.Errorf("failed to run command: %w", err)
					}
					hxeClient.Program.Print(progs)

					return nil
				},
			},
			{
				Name:        "stop",
				Usage:       "Stop a program",
				Description: `Stop a program by name.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					hxeClient := client.NewClient(clientURL, username, password)
					if hxeClient, err = hxeClient.Login(); err != nil {
						return err
					}

					progs, err := hxeClient.Program.MultiStop(cmd.Args().Slice()...)
					if err != nil {
						return fmt.Errorf("failed to run command: %w", err)
					}
					hxeClient.Program.Print(progs)
					return nil
				},
			},
			{
				Name:        "restart",
				Usage:       "Restart a program",
				Description: `Restart a program by name.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					hxeClient := client.NewClient(clientURL, username, password)
					if hxeClient, err = hxeClient.Login(); err != nil {
						return err
					}

					progs, err := hxeClient.Program.MultiRestart(cmd.Args().Slice()...)
					if err != nil {
						return fmt.Errorf("failed to run command: %w", err)
					}
					hxeClient.Program.Print(progs)
					return nil
				},
			},
			{
				Name:        "status",
				Usage:       "Show program status",
				Description: `Show the status of a specific program or all programs.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					hxeClient := client.NewClient(clientURL, username, password)
					if hxeClient, err = hxeClient.Login(); err != nil {
						return err
					}

					progs, err := hxeClient.Program.MultiStatus(cmd.Args().Slice()...)
					if err != nil {
						return fmt.Errorf("failed to run command: %w", err)
					}
					hxeClient.Program.Print(progs)
					return nil
				},
			},
			{
				Name:        "reload",
				Usage:       "Reload configuration",
				Description: `Reload the configuration without restarting services.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					hxeClient := client.NewClient(clientURL, username, password)
					if hxeClient, err = hxeClient.Login(); err != nil {
						return err
					}

					prog, err := hxeClient.Program.Reload(cmd.Args().First())
					if err != nil {
						return fmt.Errorf("failed to run command: %w", err)
					}

					hxeClient.Program.PrintDetail(prog)
					return nil
				},
			},
			{
				Name:        "enable",
				Usage:       "Enable a program",
				Description: `Enable a program to start automatically.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					hxeClient := client.NewClient(clientURL, username, password)
					if hxeClient, err = hxeClient.Login(); err != nil {
						return err
					}

					prog, err := hxeClient.Program.Enable(cmd.Args().First())
					if err != nil {
						return fmt.Errorf("failed to run command: %w", err)
					}

					hxeClient.Program.PrintDetail(prog)

					return nil
				},
			},
			{
				Name:        "disable",
				Usage:       "Disable a program",
				Description: `Disable a program from starting automatically.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					hxeClient := client.NewClient(clientURL, username, password)
					if hxeClient, err = hxeClient.Login(); err != nil {
						return err
					}

					prog, err := hxeClient.Program.Disable(cmd.Args().First())
					if err != nil {
						return fmt.Errorf("failed to run command: %w", err)
					}

					hxeClient.Program.PrintDetail(prog)
					return nil
				},
			},
			{
				Name:        "delete",
				Usage:       "Delete registered program by ID",
				Description: `Delete a registered program by ID.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					hxeClient := client.NewClient(clientURL, username, password)
					if hxeClient, err = hxeClient.Login(); err != nil {
						return err
					}

					progs, err := hxeClient.Program.MultiDelete(cmd.Args().Slice()...)
					if err != nil {
						return fmt.Errorf("failed to run command: %w", err)
					}

					hxeClient.Program.Print(progs)
					return nil
				},
			},
			{
				Name:        "tail",
				Usage:       "Tail service logs",
				Description: `Follow the logs of a service in real-time.`,
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "lines",
						Aliases: []string{"n"},
						Usage:   "Number of lines to show",
						Value:   50,
					},
					&cli.BoolFlag{
						Name:    "follow",
						Aliases: []string{"f"},
						Usage:   "Follow log output",
						Value:   true,
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					hxeClient := client.NewClient(clientURL, username, password)
					if hxeClient, err = hxeClient.Login(); err != nil {
						return err
					}

					prog, err := hxeClient.Program.Tail(cmd.Args().First())
					if err != nil {
						return fmt.Errorf("failed to run command: %w", err)
					}

					hxeClient.Program.PrintDetail(prog)
					return nil
				},
			},
			{
				Name:        "shell",
				Usage:       "Open shell for a service",
				Description: `Open an interactive shell in the context of a service.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					hxeClient := client.NewClient(clientURL, username, password)
					if hxeClient, err = hxeClient.Login(); err != nil {
						return err
					}

					prog, err := hxeClient.Program.Shell(cmd.Args().First())
					if err != nil {
						return fmt.Errorf("failed to run command: %w", err)
					}

					hxeClient.Program.PrintDetail(prog)
					return nil
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
