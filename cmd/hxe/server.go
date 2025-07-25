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

	"github.com/rangertaha/hxe/internal/agent"
	"github.com/rangertaha/hxe/internal/config"
	"github.com/urfave/cli/v3"
)

var (
	agentConfig *config.AgentConfig
	newAgent    *agent.Agent
	err         error
)

var serverCmd *cli.Command = &cli.Command{
	Name:        "server",
	Usage:       "Start the Hxe server",
	Description: `Start the Hxe server.`,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Hidden:  false,
			Usage:   "Configuration from `FILE`",
		},
	},
	Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
		// Create new config
		if agentConfig, err = config.NewAgentConfig(
			config.AgentDefaultOpts(),
			config.AgentCliOpts(ctx, cmd),
		); err != nil {
			return ctx, err
		}

		return ctx, nil
	},
	Action: func(ctx context.Context, cmd *cli.Command) error {
		// Create new agent
		if newAgent, err = agent.New(agentConfig); err != nil {
			return err
		}

		// Start the agent
		return newAgent.Start()
	},
}
