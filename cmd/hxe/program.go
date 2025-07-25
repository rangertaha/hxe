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
	"time"

	"github.com/rangertaha/hxe/internal"
	"github.com/rangertaha/hxe/internal/client"
	"github.com/rangertaha/hxe/internal/config"
	"github.com/urfave/cli/v3"
)

var (
	clientConfig *config.Client
	hxeClient    *client.Client
	serverURL    string
	username     string
	password     string
	profile      string
)

var programCmd *cli.Command = &cli.Command{
	Name:                  "program",
	Aliases:               []string{"p", "prog"},
	Usage:                 "Process management",
	Description:           `Process management`,
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
		&cli.StringFlag{
			Name:        "profile",
			Aliases:     []string{"p"},
			Hidden:      false,
			Value:       "default",
			Usage:       "Client profile to use",
			Destination: &profile,
		},
		&cli.StringFlag{
			Name:        "username",
			Aliases:     []string{"u"},
			Value:       "",
			Hidden:      false,
			Usage:       "Username for authentication",
			Destination: &username,
		},
		&cli.StringFlag{
			Name:        "password",
			Aliases:     []string{"p"},
			Value:       "",
			Hidden:      false,
			Usage:       "Password for authentication",
			Destination: &password,
		},
		&cli.StringFlag{
			Name:        "url",
			Value:       "",
			Hidden:      false,
			Usage:       "Hxe API server URL",
			Destination: &serverURL,
		},
		&cli.DurationFlag{
			Name:   "timeout",
			Value:  10 * time.Second,
			Hidden: false,
			Usage:  "Timeout for API requests",
		},
	},
	Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
		// Create new client config
		if clientConfig, err = config.NewClientConfig(
			config.ClientDefaultOptions(),
			config.ClientCliOpts(ctx, cmd),
			config.ClientProfileOpts(profile),
		); err != nil {
			return ctx, err
		}

		// Create new client
		hxeClient, err = client.New(clientConfig)
		if err != nil {
			return ctx, err
		}

		return ctx, nil
	},
	Action: func(ctx context.Context, cmd *cli.Command) error {
		cli.ShowSubcommandHelpAndExit(cmd, 1)
		return nil
	},
	Commands: []*cli.Command{

		// {
		// 	Name:  "run",
		// 	Usage: "Run a service",
		// 	Description: `Run a service. If no service name is provided,
		// 		executes the command directly.`,
		// 	Action: func(ctx context.Context, cmd *cli.Command) error {
		// 		// hxeClient := client.NewClient(serverURL, username, password)
		// 		// if hxeClient, err = hxeClient.Login(); err != nil {
		// 		// 	return err
		// 		// }

		// 		command := strings.Join(cmd.Args().Slice(), " ")
		// 		prog, err := hxeClient.Service.Run(command)
		// 		if err != nil {
		// 			return fmt.Errorf("failed to run command: %w", err)
		// 		}

		// 		hxeClient.Service.Print(prog)
		// 		return nil
		// 	},
		// },
		// {
		// 	Name:        "reload",
		// 	Usage:       "Reload service configurations",
		// 	Description: `Reload service configurations.`,
		// 	Action: func(ctx context.Context, cmd *cli.Command) error {
		// 		services := []*services.Service{}
		// 		for _, serviceId := range cmd.Args().Slice() {
		// 			service, err := hxeClient.Services.Reload(serviceId)
		// 			if err != nil {
		// 				return fmt.Errorf("failed to reload services: %w", err)
		// 			}
		// 			services = append(services, service)
		// 		}
		// 		hxeClient.Services.PrintList(services)
		// 		return nil
		// 	},
		// },
		{
			Name:        "list",
			Usage:       "List all programs",
			Description: `List all programs and their status.`,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				programs, err := hxeClient.Programs.List()
				if err != nil {
					return fmt.Errorf("failed to list programs: %w", err)
				}

				// hxeClient.Programs.Print(programs)
				programs.Print()
				return nil
			},
		},
		// {
		// 	Name:        "start",
		// 	Usage:       "Start a service",
		// 	Description: `Start a service by name.`,
		// 	Action: func(ctx context.Context, cmd *cli.Command) error {
		// 		// hxeClient := client.NewClient(serverURL, username, password)
		// 		// if hxeClient, err = hxeClient.Login(); err != nil {
		// 		// 	return err
		// 		// }
		// 		res, err := hxeClient.Service.Start(cmd.Args().First())
		// 		if err != nil {
		// 			return fmt.Errorf("failed to run command: %w", err)
		// 		}
		// 		hxeClient.Service.Print(res)

		// 		return nil
		// 	},
		// },
		// {
		// 	Name:        "stop",
		// 	Usage:       "Stop a service",
		// 	Description: `Stop a service by name.`,
		// 	Action: func(ctx context.Context, cmd *cli.Command) error {
		// 		// hxeClient := client.NewClient(serverURL, username, password)
		// 		// if hxeClient, err = hxeClient.Login(); err != nil {
		// 		// 	return err
		// 		// }

		// 		progs, err := hxeClient.Service.Stop(cmd.Args().First())
		// 		if err != nil {
		// 			return fmt.Errorf("failed to run command: %w", err)
		// 		}
		// 		hxeClient.Service.Print(progs)
		// 		return nil
		// 	},
		// },
		// {
		// 	Name:        "restart",
		// 	Usage:       "Restart a service",
		// 	Description: `Restart a service by name.`,
		// 	Action: func(ctx context.Context, cmd *cli.Command) error {
		// 		// hxeClient := client.NewClient(serverURL, username, password)
		// 		// if hxeClient, err = hxeClient.Login(); err != nil {
		// 		// 	return err
		// 		// }

		// 		progs, err := hxeClient.Service.Restart(cmd.Args().First())
		// 		if err != nil {
		// 			return fmt.Errorf("failed to run command: %w", err)
		// 		}
		// 		hxeClient.Service.Print(progs)
		// 		return nil
		// 	},
		// },
		// {
		// 	Name:        "get",
		// 	Usage:       "Get service details by ID",
		// 	Description: `Get a service details by ID.`,
		// 	Action: func(ctx context.Context, cmd *cli.Command) error {
		// 		// hxeClient := client.NewClient(serverURL, username, password)
		// 		// if hxeClient, err = hxeClient.Login(); err != nil {
		// 		// 	return err
		// 		// }

		// 		progs, err := hxeClient.Service.Get(cmd.Args().First())
		// 		if err != nil {
		// 			return fmt.Errorf("failed to run command: %w", err)
		// 		}
		// 		hxeClient.Service.Print(progs)
		// 		return nil
		// 	},
		// },
		// {
		// 	Name:        "reload",
		// 	Usage:       "Reload configuration",
		// 	Description: `Reload the configuration without restarting services.`,
		// 	Action: func(ctx context.Context, cmd *cli.Command) error {
		// 		// hxeClient := client.NewClient(serverURL, username, password)
		// 		// if hxeClient, err = hxeClient.Login(); err != nil {
		// 		// 	return err
		// 		// }

		// 		prog, err := hxeClient.Service.Reload(cmd.Args().First())
		// 		if err != nil {
		// 			return fmt.Errorf("failed to run command: %w", err)
		// 		}

		// 		hxeClient.Service.Print(prog)
		// 		return nil
		// 	},
		// },
		// {
		// 	Name:        "enable",
		// 	Usage:       "Enable a service",
		// 	Description: `Enable a service to start automatically.`,
		// 	Action: func(ctx context.Context, cmd *cli.Command) error {
		// 		hxeClient := client.NewClient(serverURL, username, password)
		// 		if hxeClient, err = hxeClient.Login(); err != nil {
		// 			return err
		// 		}

		// 		prog, err := hxeClient.Service.Enable(cmd.Args().First())
		// 		if err != nil {
		// 			return fmt.Errorf("failed to run command: %w", err)
		// 		}

		// 		hxeClient.Service.Print(prog)

		// 		return nil
		// 	},
		// },
		// {
		// 	Name:        "disable",
		// 	Usage:       "Disable a service",
		// 	Description: `Disable a service from starting automatically.`,
		// 	Action: func(ctx context.Context, cmd *cli.Command) error {
		// 		// hxeClient := client.NewClient(serverURL, username, password)
		// 		// if hxeClient, err = hxeClient.Login(); err != nil {
		// 		// 	return err
		// 		// }

		// 		prog, err := hxeClient.Service.Disable(cmd.Args().First())
		// 		if err != nil {
		// 			return fmt.Errorf("failed to run command: %w", err)
		// 		}

		// 		hxeClient.Service.Print(prog)
		// 		return nil
		// 	},
		// },
		// {
		// 	Name:        "delete",
		// 	Usage:       "Delete registered service by ID",
		// 	Description: `Delete a registered service by ID.`,
		// 	Action: func(ctx context.Context, cmd *cli.Command) error {
		// 		// hxeClient := client.NewClient(serverURL, username, password)
		// 		// if hxeClient, err = hxeClient.Login(); err != nil {
		// 		// 	return err
		// 		// }

		// 		progs, err := hxeClient.Service.Delete(cmd.Args().First())
		// 		if err != nil {
		// 			return fmt.Errorf("failed to run command: %w", err)
		// 		}

		// 		hxeClient.Service.Print(progs)
		// 		return nil
		// 	},
		// },
		// {
		// 	Name:        "log",
		// 	Usage:       "Show service logs",
		// 	Description: `Show the logs of a service.`,
		// 	Flags: []cli.Flag{
		// 		&cli.IntFlag{
		// 			Name:    "lines",
		// 			Aliases: []string{"n"},
		// 			Usage:   "Number of lines to show",
		// 			Value:   50,
		// 		},
		// 		&cli.BoolFlag{
		// 			Name:    "follow",
		// 			Aliases: []string{"f"},
		// 			Usage:   "Follow log output",
		// 			Value:   true,
		// 		},
		// 	},
		// 	Action: func(ctx context.Context, cmd *cli.Command) error {
		// 		// hxeClient := client.NewClient(serverURL, username, password)
		// 		// if hxeClient, err = hxeClient.Login(); err != nil {
		// 		// 	return err
		// 		// }

		// 		// prog, err := hxeClient.Service.Log(cmd.Args().Slice())...)
		// 		// if err != nil {
		// 		// 	return fmt.Errorf("failed to run command: %w", err)
		// 		// }

		// 		// hxeClient.Service.PrintDetail(prog)
		// 		return nil
		// 	},
		// },
		// {
		// 	Name:        "shell",
		// 	Usage:       "Open shell for a service",
		// 	Description: `Open an interactive shell in the context of a service.`,
		// 	Action: func(ctx context.Context, cmd *cli.Command) error {
		// 		// hxeClient := client.NewClient(serverURL, username, password)
		// 		// if hxeClient, err = hxeClient.Login(); err != nil {
		// 		// 	return err
		// 		// }

		// 		// prog, err := hxeClient.Service.Shell(cmd.Args().Slice())...)
		// 		// if err != nil {
		// 		// 	return fmt.Errorf("failed to run command: %w", err)
		// 		// }

		// 		// hxeClient.Service.PrintDetail(prog)
		// 		return nil
		// 	},
		// },
	},
}
