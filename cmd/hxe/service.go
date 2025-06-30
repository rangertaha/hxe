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
	"strings"

	"github.com/rangertaha/hxe/internal"
	"github.com/rangertaha/hxe/pkg/client"
	"github.com/urfave/cli/v3"
)

var serviceCmd *cli.Command = &cli.Command{
	Name:                  "service",
	Aliases:               []string{"s", "svc"},
	Usage:                 "Hxe service management",
	Description:           `Hxe service management`,
	Version:               internal.VERSION,
	Authors:               []any{"Rangertaha"},
	Copyright:             "Rangertaha",
	EnableShellCompletion: true,
	Suggest:               true,

	// Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
	// 	// Load configuration file
	// 	cfgFile := cmd.String("config")
	// 	cliOptions := config.CliOptions(ctx, cmd)
	// 	fileOptions := config.FileOption(cfgFile)

	// 	// Create new config
	// 	if CfgOption, err = config.New(fileOptions, cliOptions); err != nil {
	// 		return ctx, err
	// 	}

	// 	return ctx, nil
	// },
	Action: func(ctx context.Context, cmd *cli.Command) error {
		cli.ShowSubcommandHelpAndExit(cmd, 1)
		return nil
	},
	Commands: []*cli.Command{

		{
			Name:  "run",
			Usage: "Run a service",
			Description: `Run a service. If no service name is provided, 
				executes the command directly.`,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				hxeClient := client.NewClient(serverURL, username, password)
				if hxeClient, err = hxeClient.Login(); err != nil {
					return err
				}

				command := strings.Join(cmd.Args().Slice(), " ")
				prog, err := hxeClient.Service.Run(command)
				if err != nil {
					return fmt.Errorf("failed to run command: %w", err)
				}

				hxeClient.Service.PrintDetail(prog)
				return nil
			},
		},
		{
			Name:        "list",
			Usage:       "List all services",
			Description: `List all services with their current status.`,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				hxec := client.NewClient(serverURL, username, password)
				if _, err := hxec.Login(); err != nil {
					return fmt.Errorf("failed to login: %w", err)
				}

				services, err := hxec.Service.List()
				if err != nil {
					return fmt.Errorf("failed to list services: %w", err)
				}

				hxec.Service.Print(services)
				return nil
			},
		},
		{
			Name:        "start",
			Usage:       "Start a service",
			Description: `Start a service by name.`,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				hxeClient := client.NewClient(serverURL, username, password)
				if hxeClient, err = hxeClient.Login(); err != nil {
					return err
				}
				progs, err := hxeClient.Service.MultiStart(cmd.Args().Slice()...)
				if err != nil {
					return fmt.Errorf("failed to run command: %w", err)
				}
				hxeClient.Service.Print(progs)

				return nil
			},
		},
		{
			Name:        "stop",
			Usage:       "Stop a service",
			Description: `Stop a service by name.`,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				hxeClient := client.NewClient(serverURL, username, password)
				if hxeClient, err = hxeClient.Login(); err != nil {
					return err
				}

				progs, err := hxeClient.Service.MultiStop(cmd.Args().Slice()...)
				if err != nil {
					return fmt.Errorf("failed to run command: %w", err)
				}
				hxeClient.Service.Print(progs)
				return nil
			},
		},
		{
			Name:        "restart",
			Usage:       "Restart a service",
			Description: `Restart a service by name.`,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				hxeClient := client.NewClient(serverURL, username, password)
				if hxeClient, err = hxeClient.Login(); err != nil {
					return err
				}

				progs, err := hxeClient.Service.MultiRestart(cmd.Args().Slice()...)
				if err != nil {
					return fmt.Errorf("failed to run command: %w", err)
				}
				hxeClient.Service.Print(progs)
				return nil
			},
		},
		{
			Name:        "status",
			Usage:       "Show service status",
			Description: `Show the status of a specific service or all services.`,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				hxeClient := client.NewClient(serverURL, username, password)
				if hxeClient, err = hxeClient.Login(); err != nil {
					return err
				}

				progs, err := hxeClient.Service.MultiStatus(cmd.Args().Slice()...)
				if err != nil {
					return fmt.Errorf("failed to run command: %w", err)
				}
				hxeClient.Service.Print(progs)
				return nil
			},
		},
		{
			Name:        "reload",
			Usage:       "Reload configuration",
			Description: `Reload the configuration without restarting services.`,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				hxeClient := client.NewClient(serverURL, username, password)
				if hxeClient, err = hxeClient.Login(); err != nil {
					return err
				}

				prog, err := hxeClient.Service.Reload(cmd.Args().First())
				if err != nil {
					return fmt.Errorf("failed to run command: %w", err)
				}

				hxeClient.Service.PrintDetail(prog)
				return nil
			},
		},
		{
			Name:        "enable",
			Usage:       "Enable a service",
			Description: `Enable a service to start automatically.`,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				hxeClient := client.NewClient(serverURL, username, password)
				if hxeClient, err = hxeClient.Login(); err != nil {
					return err
				}

				prog, err := hxeClient.Service.Enable(cmd.Args().First())
				if err != nil {
					return fmt.Errorf("failed to run command: %w", err)
				}

				hxeClient.Service.PrintDetail(prog)

				return nil
			},
		},
		{
			Name:        "disable",
			Usage:       "Disable a service",
			Description: `Disable a service from starting automatically.`,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				hxeClient := client.NewClient(serverURL, username, password)
				if hxeClient, err = hxeClient.Login(); err != nil {
					return err
				}

				prog, err := hxeClient.Service.Disable(cmd.Args().First())
				if err != nil {
					return fmt.Errorf("failed to run command: %w", err)
				}

				hxeClient.Service.PrintDetail(prog)
				return nil
			},
		},
		{
			Name:        "delete",
			Usage:       "Delete registered service by ID",
			Description: `Delete a registered service by ID.`,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				hxeClient := client.NewClient(serverURL, username, password)
				if hxeClient, err = hxeClient.Login(); err != nil {
					return err
				}

				progs, err := hxeClient.Service.MultiDelete(cmd.Args().Slice()...)
				if err != nil {
					return fmt.Errorf("failed to run command: %w", err)
				}

				hxeClient.Service.Print(progs)
				return nil
			},
		},
		{
			Name:        "log",
			Usage:       "Show service logs",
			Description: `Show the logs of a service.`,
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
				hxeClient := client.NewClient(serverURL, username, password)
				if hxeClient, err = hxeClient.Login(); err != nil {
					return err
				}

				prog, err := hxeClient.Service.Log(cmd.Args().First())
				if err != nil {
					return fmt.Errorf("failed to run command: %w", err)
				}

				hxeClient.Service.PrintDetail(prog)
				return nil
			},
		},
		{
			Name:        "shell",
			Usage:       "Open shell for a service",
			Description: `Open an interactive shell in the context of a service.`,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				hxeClient := client.NewClient(serverURL, username, password)
				if hxeClient, err = hxeClient.Login(); err != nil {
					return err
				}

				prog, err := hxeClient.Service.Shell(cmd.Args().First())
				if err != nil {
					return fmt.Errorf("failed to run command: %w", err)
				}

				hxeClient.Service.PrintDetail(prog)
				return nil
			},
		},
	},
}
