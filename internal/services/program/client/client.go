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

package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rangertaha/hxe/internal/log"
	"github.com/rangertaha/hxe/internal/services/program/models"
	"github.com/rs/zerolog"
)

// Client is a client used internally to manage services
type Client struct {
	nc  *nats.Conn
	log zerolog.Logger
}

type Request struct {
	Program *models.Program `json:"service"`
}

type Response struct {
	Status   error             `json:"status"`
	Programs []*models.Program `json:"programs"`
}

func New(nc *nats.Conn) *Client {
	return &Client{nc: nc, log: log.With().Logger()}
}

// List all services
func (c *Client) List() (resp *Response, err error) {

	msg, err := c.nc.Request("program.list", []byte("{}"), 5*time.Second)
	if err != nil {
		errMsg := fmt.Sprintf("%s: failed to request program.list", c.nc.ConnectedUrl())
		c.log.Error().Err(err).Str("url", c.nc.ConnectedUrl()).Msg(errMsg)
		return nil, errors.New(errMsg)
	}
	c.log.Debug().Msgf("programs.list response: %s", string(msg.Data))

	resp = &Response{}
	err = json.Unmarshal(msg.Data, resp)
	if err != nil {
		errMsg := "failed to unmarshal programs.program.list response"
		c.log.Error().Err(err).Msg(errMsg)
		return nil, errors.New(errMsg)
	}
	return resp, nil
}

// Print formats and prints the list of programs in table format
func (s *Response) Print() {
	for _, program := range s.Programs {
		fmt.Printf("Program: %s\n", program.Name)
	}
}
