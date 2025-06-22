package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rangertaha/hxe/internal"
	"github.com/rangertaha/hxe/internal/agent"
	"github.com/rangertaha/hxe/internal/config"
	"github.com/urfave/cli/v3"
)

var (
	err       error
	cfgOption *config.Config
	newAgent  *agent.Agent
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
		},
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			// Load configuration file
			cfgFile := cmd.String("config")
			defaultOpts := config.DefaultOptions()
			cliOptions := config.CliOptions(ctx, cmd)
			fileOptions := config.FileOption(cfgFile)

			// Create new config
			if cfgOption, err = config.New(defaultOpts, fileOptions, cliOptions); err != nil {
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
				Name: "server",
				// Aliases:     []string{"s"},
				Usage:       "",
				Description: ``,
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
				Name: "run",
				// Aliases: []string{"r"},
				Usage: "Run a program",
				Description: `Run a program. If no program name is provided, 
				executes the command directly.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if err := newAgent.Execute(cmd.Args().Slice()...); err != nil {
						fmt.Printf("\n\tError: %v\n\n", err)
						cli.ShowAppHelpAndExit(cmd, 1)
						return err
					}
					return nil
				},
			},
			{
				Name: "list",
				// Aliases:     []string{"ls", "l"},
				Usage:       "List all services",
				Description: `List all configured services with their current status.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					cli.ShowAppHelpAndExit(cmd, 1)
					return nil
				},
			},
			{
				Name: "start",
				// Aliases:     []string{"s"},
				Usage:       "Start a service",
				Description: `Start a service by name.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					cli.ShowAppHelpAndExit(cmd, 1)
					return nil
				},
			},
			{
				Name: "stop",
				// Aliases:     []string{"st"},
				Usage:       "Stop a service",
				Description: `Stop a service by name.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					cli.ShowAppHelpAndExit(cmd, 1)
					return nil
				},
			},
			{
				Name: "restart",
				// Aliases:     []string{"rs"},
				Usage:       "Restart a service",
				Description: `Restart a service by name.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					cli.ShowAppHelpAndExit(cmd, 1)
					return nil
				},
			},
			{
				Name: "status",
				// Aliases:     []string{"st"},
				Usage:       "Show service status",
				Description: `Show the status of a specific service or all services.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					cli.ShowAppHelpAndExit(cmd, 1)
					return nil
				},
			},
			{
				Name: "tail",
				// Aliases:     []string{"t"},
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
					return nil
				},
			},
			{
				Name: "reload",
				// Aliases:     []string{"rl"},
				Usage:       "Reload configuration",
				Description: `Reload the configuration without restarting services.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					cli.ShowAppHelpAndExit(cmd, 1)
					return nil
				},
			},
			{
				Name: "enable",
				// Aliases:     []string{"e"},
				Usage:       "Enable a service",
				Description: `Enable a service to start automatically.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					cli.ShowAppHelpAndExit(cmd, 1)

					return nil
				},
			},
			{
				Name: "disable",
				// Aliases:     []string{"d"},
				Usage:       "Disable a service",
				Description: `Disable a service from starting automatically.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					cli.ShowAppHelpAndExit(cmd, 1)
					return nil
				},
			},
			{
				Name: "shell",
				// Aliases:     []string{"sh"},
				Usage:       "Open shell for a service",
				Description: `Open an interactive shell in the context of a service.`,
				Action: func(ctx context.Context, cmd *cli.Command) error {
					cli.ShowAppHelpAndExit(cmd, 1)
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
