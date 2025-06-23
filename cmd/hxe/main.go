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
				Usage:       "List all services",
				Description: `List all configured services with their current status.`,
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
				Usage:       "Start a service",
				Description: `Start a service by name.`,
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
				Usage:       "Stop a service",
				Description: `Stop a service by name.`,
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
				Usage:       "Restart a service",
				Description: `Restart a service by name.`,
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
				Usage:       "Show service status",
				Description: `Show the status of a specific service or all services.`,
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
				Usage:       "Enable a service",
				Description: `Enable a service to start automatically.`,
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
				Usage:       "Disable a service",
				Description: `Disable a service from starting automatically.`,
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
