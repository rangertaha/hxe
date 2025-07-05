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

package client

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/rangertaha/hxe/internal/log"
	"github.com/rangertaha/hxe/internal/modules/services"
	"github.com/rangertaha/hxe/internal/rdb"
	"github.com/rs/zerolog"
)

type ServiceClient struct {
	client *Client
	schema *rdb.Schema
	log    zerolog.Logger
}

func NewServiceClient(client *Client) *ServiceClient {
	return &ServiceClient{client: client, schema: rdb.ServiceSchema(), log: log.With().Logger()}
}

// Service operations
func (c *ServiceClient) List() (res *services.Response, err error) {
	res = &services.Response{}
	err = c.client.Get("/api/service", res)
	c.Print(res)
	return
}

// func (c *ServiceClient) Get(id string) (res *models.Response, err error) {
// 	res = &models.Response{}
// 	err = c.client.Get(fmt.Sprintf("/api/service/%s", id), &res)
// 	return
// }

// // func (c *ServiceClient) Status(ids ...string) (res *models.Response, err error) {
// // 	res = &models.Response{}
// // 	for _, id := range ids {
// // 		resp, err := c.Get(id)

// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 		res.Services = append(res.Services, resp.Services...)
// // 		res.Services = append(res.Services, resp.Service)
// // 	}
// // 	return
// // }

// func (c *ServiceClient) Create(req *models.Request) (res *models.Response, err error) {
// 	res = &models.Response{}
// 	err = c.client.Post("/api/service", req, &res)
// 	return
// }

// // func (c *ServiceClient) Creates(reqs ...*models.Request) (res *models.Response, err error) {
// // 	res = &models.Response{}
// // 	for _, req := range reqs {
// // 		resp, err := c.Create(req)
// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 		res.Services = append(res.Services, resp.Service)
// // 		res.Services = append(res.Services, resp.Services...)
// // 	}
// // 	return
// // }

// func (c *ServiceClient) Update(req *models.Request) (res *models.Response, err error) {
// 	res = &models.Response{}
// 	err = c.client.Put(fmt.Sprintf("/api/service/%d", req.Service.Id), req, &res)
// 	return
// }

// // func (c *ServiceClient) Updates(reqs ...*models.Request) (res *models.Response, err error) {
// // 	res = &models.Response{}
// // 	for _, req := range reqs {
// // 		resp, err := c.Update(req)
// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 		res.Services = append(res.Services, resp.Service)
// // 		res.Services = append(res.Services, resp.Services...)
// // 	}
// // 	return
// // }

// func (c *ServiceClient) Delete(id string) (res *models.Response, err error) {
// 	res = &models.Response{}
// 	err = c.client.Delete(fmt.Sprintf("/api/service/%s", id), &res)
// 	return
// }

// // func (c *ServiceClient) Deletes(ids ...string) (res *models.Response, err error) {
// // 	res = &models.Response{}
// // 	for _, id := range ids {
// // 		resp, err := c.Delete(id)
// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 		res.Services = append(res.Services, resp.Service)
// // 		res.Services = append(res.Services, resp.Services...)
// // 	}
// // 	return
// // }

// func (c *ServiceClient) Start(id string) (res *models.Response, err error) {
// 	res = &models.Response{}
// 	err = c.client.Post(fmt.Sprintf("/api/service/%s/start", id), &models.Request{Service: &models.Service{Id: 1}}, &res)
// 	// fmt.Println(res)
// 	return
// }

// // func (c *ServiceClient) Starts(ids ...string) (res *models.Response, err error) {
// // 	res = &models.Response{}
// // 	for _, id := range ids {
// // 		resp, err := c.Start(id)
// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 		res.Services = append(res.Services, resp.Service)
// // 		res.Services = append(res.Services, resp.Services...)
// // 	}
// // 	return
// // }

// func (c *ServiceClient) Stop(id string) (res *models.Response, err error) {
// 	res = &models.Response{}
// 	err = c.client.Post(fmt.Sprintf("/api/service/%s/stop", id), nil, &res)
// 	return
// }

// // func (c *ServiceClient) Stops(ids ...string) (res *models.Response, err error) {
// // 	res = &models.Response{}
// // 	for _, id := range ids {
// // 		resp, err := c.Stop(id)
// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 		res.Services = append(res.Services, resp.Service)
// // 		res.Services = append(res.Services, resp.Services...)
// // 	}
// // 	return
// // }

// func (c *ServiceClient) Restart(id string) (res *models.Response, err error) {
// 	res = &models.Response{}
// 	err = c.client.Post(fmt.Sprintf("/api/service/%s/restart", id), nil, &res)
// 	return
// }

// // func (c *ServiceClient) Restarts(ids ...string) (res *models.Response, err error) {
// // 	res = &models.Response{}
// // 	for _, id := range ids {
// // 		resp, err := c.Restart(id)
// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 		res.Services = append(res.Services, resp.Service)
// // 		res.Services = append(res.Services, resp.Services...)
// // 	}
// // 	return
// // }

// func (c *ServiceClient) Enable(id string) (res *models.Response, err error) {
// 	res = &models.Response{}
// 	err = c.client.Post(fmt.Sprintf("/api/service/%s/enable", id), nil, &res)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

// // func (c *ServiceClient) Enables(ids ...string) (res *models.Response, err error) {
// // 	res = &models.Response{}
// // 	for _, id := range ids {
// // 		resp, err := c.Enable(id)
// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 		res.Services = append(res.Services, resp.Service)
// // 		res.Services = append(res.Services, resp.Services...)
// // 	}
// // 	return
// // }

// func (c *ServiceClient) Disable(id string) (res *models.Response, err error) {
// 	res = &models.Response{}
// 	err = c.client.Post(fmt.Sprintf("/api/service/%s/disable", id), nil, &res)
// 	return
// }

// // func (c *ServiceClient) Disables(ids ...string) (res *models.Response, err error) {
// // 	res = &models.Response{}
// // 	for _, id := range ids {
// // 		resp, err := c.Disable(id)
// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 		res.Services = append(res.Services, resp.Service)
// // 		res.Services = append(res.Services, resp.Services...)
// // 	}
// // 	return
// // }

// func (c *ServiceClient) Reload(id string) (res *models.Response, err error) {
// 	res = &models.Response{}
// 	err = c.client.Post(fmt.Sprintf("/api/service/%s/reload", id), nil, &res)
// 	return
// }

// // func (c *ServiceClient) Reloads(ids ...string) (res *models.Response, err error) {
// // 	res = &models.Response{}
// // 	for _, id := range ids {
// // 		resp, err := c.Reload(id)
// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 		res.Services = append(res.Services, resp.Services...)
// // 	}
// // 	return
// // }

// func (c *ServiceClient) Shell(id string) (res *models.Response, err error) {
// 	res = &models.Response{}
// 	err = c.client.Post(fmt.Sprintf("/api/service/%s/shell", id), nil, &res)
// 	return
// }

// func (c *ServiceClient) Log(id string) (res *models.Response, err error) {
// 	res = &models.Response{}
// 	err = c.client.Post(fmt.Sprintf("/api/service/%s/log", id), nil, &res)
// 	return
// }

// func (c *ServiceClient) Run(command string) (res *models.Response, err error) {
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		dir = "/tmp"
// 	}

// 	service := models.Service{
// 		CmdExec:   command,
// 		Directory: dir,
// 		User:      "root",
// 		Group:     "root",
// 		Status:    models.ServiceStatus_STARTING,
// 		Enabled:   true,
// 	}
// 	req := &models.Request{
// 		Service: &service,
// 	}

// 	return c.Create(req)
// }

func (c *ServiceClient) PrintDetail(res rdb.Service) {
	// Create table
	t := table.NewWriter()
	t.SetOutputMirror(nil)

	t.AppendHeader(table.Row{"Field", "Value"})

	// Add rowss
	t.AppendRow(table.Row{"ID", res.ID})
	t.AppendRow(table.Row{"Name", res.Name})
	t.AppendRow(table.Row{"Command", res.CmdExec})
	t.AppendRow(table.Row{"Status", res.Status})
	t.AppendRow(table.Row{"Enabled", res.Enabled})

	// Print the table
	fmt.Println(t.Render())
}

func (c *ServiceClient) PrintList(res *services.Response) {
	// Create table
	t := table.NewWriter()
	t.SetOutputMirror(nil)

	t.AppendHeader(table.Row{"ID", "Name", "Command", "Status", "Enabled", "Category", "App"})

	// Add rows
	for _, service := range res.Services {
		t.AppendRow(table.Row{
			service.ID,
			service.Name,
			service.CmdExec,
			service.Status,
			service.Enabled,
			service.Category,
			service.App,
		})
	}
	// Print the table
	fmt.Println(t.Render())
}

func (c *ServiceClient) Print(res *services.Response) {

	if len(res.Services) > 1 {
		c.PrintList(res)
	}

	// if res.Service != nil && res.Service.ID != 0 {
	// 	c.PrintDetail(res.Service)
	// }
}
