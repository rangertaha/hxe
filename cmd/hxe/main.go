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
	"github.com/rangertaha/hxe/internal/log"
	"github.com/rs/zerolog"
	"github.com/urfave/cli/v3"
)

// var (
// 	CfgOption *config.Config
// 	HxeConfig *config.Config
// 	newAgent  *agent.Agent
// 	hxeClient *client.Client
// 	serverURL string
// 	username  string
// 	password  string
// 	err       error
// )

func init() {
	log.SetGlobalLevel(zerolog.ErrorLevel)
}
	
func main() {
	app := &cli.Command{
		Name:                  "hxe",
		Usage:                 "Hxe Host-based Process Execution Manager",
		Description:           `Hxe Host-based Process Execution Manager`,
		Version:               internal.VERSION,
		Authors:               []any{"Rangertaha"},
		Copyright:             "Rangertaha",
		EnableShellCompletion: true,
		Suggest:               true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:   "debug",
				Value:  false,
				Hidden: false,
				Usage:  "Log debug messages for development",
				Action: nil,
			},
		},
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			if cmd.Bool("debug") {
				log.SetGlobalLevel(zerolog.DebugLevel)
			}

			return ctx, nil
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			cli.ShowAppHelpAndExit(cmd, 1)
			return nil
		},
		Commands: []*cli.Command{
			serverCmd,
			programCmd,
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
