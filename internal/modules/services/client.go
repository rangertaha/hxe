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

package services

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rangertaha/hxe/internal/log"
	"github.com/rangertaha/hxe/internal/rdb"
	"github.com/rs/zerolog"
)

// Client is a client used internally to manage services
type Client struct {
	nc  *nats.Conn
	log zerolog.Logger
}

type Request struct {
	Service *rdb.Service `json:"service"`
}

type Response struct {
	Status   error          `json:"status"`
	Service  *rdb.Service   `json:"service"`
	Services []*rdb.Service `json:"services"`
}

func NewClient(nc *nats.Conn) *Client {
	return &Client{nc: nc, log: log.With().Logger()}
}

// // Load a service
// func (c *Client) Load(id uint) (resp *Response, err error) {
// 	req := &Request{Service: &rdb.Service{ID: uint(id)}}
// 	data, err := json.Marshal(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	msg, err := c.nc.Request("service.load", data, 5*time.Second)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp = &Response{}
// 	err = proto.Unmarshal(msg.Data, resp)
// 	return resp, err
// }

// // List all services
// func (c *Client) List() (resp *Response, err error) {
// 	req := &Request{}
// 	data, err := proto.Marshal(req)
// 	c.log.Debug().Msgf("service.list request: %s", string(data))
// 	if err != nil {
// 		errMsg := "failed to marshal service list request"
// 		c.log.Error().Err(err).Msg(errMsg)
// 		return nil, errors.New(errMsg)
// 	}

// 	msg, err := c.nc.Request("service.list", data, 5*time.Second)
// 	if err != nil {
// 		errMsg := "failed to request service.list"
// 		c.log.Error().Err(err).Msg(errMsg)
// 		return nil, errors.New(errMsg)
// 	}
// 	c.log.Debug().Msgf("service.list response: %s", string(msg.Data))

// 	resp = &models.Response{}
// 	err = proto.Unmarshal(msg.Data, resp)
// 	if err != nil {
// 		errMsg := "failed to unmarshal service.list response"
// 		c.log.Error().Err(err).Msg(errMsg)
// 		return nil, errors.New(errMsg)
// 	}
// 	return resp, nil
// }

// func (c *Client) Get(id uint) (resp *models.Response, err error) {
// 	req := &models.Request{Service: &models.Service{Id: uint32(id)}}
// 	data, err := proto.Marshal(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	msg, err := c.nc.Request("service.get", data, 5*time.Second)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp = &models.Response{}
// 	err = proto.Unmarshal(msg.Data, resp)
// 	return resp, err
// }

// // Create a new service
// func (c *Client) Create(req *models.Service) (resp *models.Response, err error) {
// 	data, err := proto.Marshal(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	msg, err := c.nc.Request("service.create", data, 5*time.Second)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp = &models.Response{}
// 	err = proto.Unmarshal(msg.Data, resp)
// 	return resp, err
// }

// // Update a service
// func (c *Client) Update(req *models.Service) (resp *models.Response, err error) {
// 	data, err := proto.Marshal(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	msg, err := c.nc.Request("service.update", data, 5*time.Second)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp = &models.Response{}
// 	err = proto.Unmarshal(msg.Data, resp)
// 	return resp, err
// }

// // Delete a service
// func (c *Client) Delete(id uint) (resp *models.Response, err error) {
// 	req := &models.Request{Service: &models.Service{Id: uint32(id)}}
// 	data, err := proto.Marshal(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	msg, err := c.nc.Request("service.delete", data, 5*time.Second)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp = &models.Response{}
// 	err = proto.Unmarshal(msg.Data, resp)
// 	return resp, err
// }

// // Start a service
// func (c *Client) Start(id uint) (resp *models.Response, err error) {
// 	req := &models.Request{Service: &models.Service{Id: uint32(id)}}
// 	data, err := proto.Marshal(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	msg, err := c.nc.Request("service.start", data, 5*time.Second)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp = &models.Response{}
// 	err = proto.Unmarshal(msg.Data, resp)
// 	return resp, err
// }

// // Stop a service
// func (c *Client) Stop(id uint) (resp *models.Response, err error) {
// 	req := &models.Request{Service: &models.Service{Id: uint32(id)}}
// 	data, err := proto.Marshal(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	msg, err := c.nc.Request("service.stop", data, 5*time.Second)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp = &models.Response{}
// 	err = proto.Unmarshal(msg.Data, resp)
// 	return resp, err
// }

// // Restart a service
// func (c *Client) Restart(id uint) (resp *models.Response, err error) {
// 	req := &models.Request{Service: &models.Service{Id: uint32(id)}}
// 	data, err := proto.Marshal(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	msg, err := c.nc.Request("service.restart", data, 5*time.Second)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp = &models.Response{}
// 	err = proto.Unmarshal(msg.Data, resp)
// 	return resp, err
// }

// // Status returns service status
// func (c *Client) Status(id uint) (resp *models.Response, err error) {
// 	req := &models.Request{Service: &models.Service{Id: uint32(id)}}
// 	data, err := proto.Marshal(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	msg, err := c.nc.Request("service.status", data, 5*time.Second)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp = &models.Response{}
// 	err = proto.Unmarshal(msg.Data, resp)
// 	return resp, err
// }

// // Reload reloads a service
// func (c *Client) Reload(id uint) (resp *models.Response, err error) {
// 	req := &models.Request{Service: &models.Service{Id: uint32(id)}}
// 	data, err := proto.Marshal(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	msg, err := c.nc.Request("service.reload", data, 5*time.Second)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp = &models.Response{}
// 	err = proto.Unmarshal(msg.Data, resp)
// 	return resp, err
// }

// // Enable enables a service
// func (c *Client) Enable(id uint) (resp *models.Response, err error) {
// 	req := &models.Request{Service: &models.Service{Id: uint32(id)}}
// 	data, err := proto.Marshal(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	msg, err := c.nc.Request("service.enable", data, 5*time.Second)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp = &models.Response{}
// 	err = proto.Unmarshal(msg.Data, resp)
// 	return resp, err
// }

// // Disable disables a service
// func (c *Client) Disable(id uint) (resp *models.Response, err error) {
// 	req := &models.Request{Service: &models.Service{Id: uint32(id)}}
// 	data, err := proto.Marshal(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	msg, err := c.nc.Request("service.disable", data, 5*time.Second)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp = &models.Response{}
// 	err = proto.Unmarshal(msg.Data, resp)
// 	return resp, err
// }

// Disable disables a service
func (c *Client) Send(topic string, req []byte) (res []byte, err error) {
	msg, err := c.nc.Request(topic, req, 5*time.Second)
	if err != nil {
		return nil, err
	}
	res = msg.Data

	return
}

// // Shell opens a shell for a service
// func (c *Client) Shell(id uint) (resp *models.Response, err error) {
// 	req := &models.Request{Service: &models.Service{Id: uint32(id)}}
// 	data, err := proto.Marshal(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	msg, err := c.nc.Request("service.shell", data, 5*time.Second)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp = &models.Response{}
// 	err = proto.Unmarshal(msg.Data, resp)
// 	return resp, err
// }

// // Log returns service logs
// func (c *Client) Log(id uint, stream string) (resp chan []byte, err error) {
// 	resp = make(chan []byte, 100)
// 	close(resp)
// 	return resp, nil
// }
