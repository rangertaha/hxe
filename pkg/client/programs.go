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
	"os"
	"os/user"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/rangertaha/hxe/internal/models"
)

type ProgramClient struct {
	client *Client
	schema *models.Schema
}

func NewProgramClient(client *Client) *ProgramClient {
	return &ProgramClient{client: client, schema: models.ProgramSchema()}
}

// Program operations
func (c *ProgramClient) List() ([]*models.Program, error) {
	var programs []*models.Program
	err := c.client.Get("/api/program", &programs)
	return programs, err
}

func (c *ProgramClient) Get(id string) (*models.Program, error) {
	var program *models.Program
	err := c.client.Get(fmt.Sprintf("/api/program/%s", id), &program)
	return program, err
}

func (c *ProgramClient) Status(id string) (*models.Program, error) {
	return c.Get(id)
}

func (c *ProgramClient) MultiStatus(ids ...string) (programs []*models.Program, err error) {
	for _, id := range ids {
		program, err := c.Status(id)
		if err != nil {
			return nil, err
		}
		programs = append(programs, program)
	}
	return programs, nil
}

func (c *ProgramClient) Create(program *models.Program) (*models.Program, error) {
	var created models.Program
	err := c.client.Post("/api/program", program, &created)
	return &created, err
}

func (c *ProgramClient) Update(id string, program *models.Program) (*models.Program, error) {
	var updated models.Program
	err := c.client.Put(fmt.Sprintf("/api/program/%s", id), program, &updated)
	return &updated, err
}

func (c *ProgramClient) Delete(id string) (*models.Program, error) {
	var deleted models.Program
	err := c.client.Delete(fmt.Sprintf("/api/program/%s", id), &deleted)
	return &deleted, err
}

func (c *ProgramClient) MultiDelete(ids ...string) (programs []*models.Program, err error) {
	for _, id := range ids {
		program, err := c.Delete(id)
		if err != nil {
			return nil, err
		}
		programs = append(programs, program)
	}
	return programs, nil
}

func (c *ProgramClient) Start(id string) (*models.Program, error) {
	var program models.Program
	err := c.client.Post(fmt.Sprintf("/api/program/%s/start", id), nil, &program)
	return &program, err
}

func (c *ProgramClient) MultiStart(ids ...string) (programs []*models.Program, err error) {
	for _, id := range ids {
		program, err := c.Start(id)
		if err != nil {
			return nil, err
		}
		programs = append(programs, program)
	}
	return programs, nil
}

func (c *ProgramClient) Stop(id string) (*models.Program, error) {
	var program models.Program
	err := c.client.Post(fmt.Sprintf("/api/program/%s/stop", id), nil, &program)
	return &program, err
}

func (c *ProgramClient) MultiStop(ids ...string) (programs []*models.Program, err error) {
	for _, id := range ids {
		program, err := c.Stop(id)
		if err != nil {
			return nil, err
		}
		programs = append(programs, program)
	}
	return programs, nil
}

func (c *ProgramClient) Restart(id string) (*models.Program, error) {
	var program models.Program
	err := c.client.Post(fmt.Sprintf("/api/program/%s/restart", id), nil, &program)
	return &program, err
}

func (c *ProgramClient) MultiRestart(ids ...string) (programs []*models.Program, err error) {
	for _, id := range ids {
		program, err := c.Restart(id)
		if err != nil {
			return nil, err
		}
		programs = append(programs, program)
	}
	return programs, nil
}

func (c *ProgramClient) Enable(id string) (*models.Program, error) {
	var program models.Program
	err := c.client.Post(fmt.Sprintf("/api/program/%s/enable", id), nil, &program)
	return &program, err
}

func (c *ProgramClient) MultiEnable(ids ...string) (programs []*models.Program, err error) {
	for _, id := range ids {
		program, err := c.Enable(id)
		if err != nil {
			return nil, err
		}
		programs = append(programs, program)
	}
	return programs, nil
}

func (c *ProgramClient) Disable(id string) (*models.Program, error) {
	var program models.Program
	err := c.client.Post(fmt.Sprintf("/api/program/%s/disable", id), nil, &program)
	return &program, err
}

func (c *ProgramClient) MultiDisable(ids ...string) (programs []*models.Program, err error) {
	for _, id := range ids {
		program, err := c.Disable(id)
		if err != nil {
			return nil, err
		}
		programs = append(programs, program)
	}
	return programs, nil
}

func (c *ProgramClient) Reload(id string) (*models.Program, error) {
	var program models.Program
	err := c.client.Post(fmt.Sprintf("/api/program/%s/reload", id), nil, &program)
	return &program, err
}

func (c *ProgramClient) MultiReload(ids ...string) (programs []*models.Program, err error) {
	for _, id := range ids {
		program, err := c.Reload(id)
		if err != nil {
			return nil, err
		}
		programs = append(programs, program)
	}
	return programs, nil
}

func (c *ProgramClient) Shell(id string) (*models.Program, error) {
	var program models.Program
	err := c.client.Post(fmt.Sprintf("/api/program/%s/shell", id), nil, &program)
	return &program, err
}

func (c *ProgramClient) Tail(id string) (*models.Program, error) {
	var program models.Program
	err := c.client.Post(fmt.Sprintf("/api/program/%s/tail", id), nil, &program)
	return &program, err
}

func (c *ProgramClient) Run(command string) (*models.Program, error) {
	dir, err := os.Getwd()
	if err != nil {
		dir = "/tmp"
	}

	currentUser, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}

	program := models.Program{
		Command:   command,
		Directory: dir,
		User:      currentUser.Username,
		Group:     currentUser.Gid,
		Status:    models.ProgramStart,
		Enabled:   true,
	}

	return c.Create(&program)
}

func (c *ProgramClient) PrintDetail(program *models.Program) {
	// Create table
	t := table.NewWriter()
	t.SetOutputMirror(nil)

	t.AppendHeader(table.Row{"Field", "Value"})

	// Add rows
	t.AppendRow(table.Row{"ID", program.ID})
	t.AppendRow(table.Row{"Name", program.Name})
	t.AppendRow(table.Row{"Command", program.Command})
	t.AppendRow(table.Row{"Status", program.Status})
	t.AppendRow(table.Row{"Enabled", program.Enabled})

	// Print the table
	fmt.Println(t.Render())
}

func (c *ProgramClient) PrintList(programs []*models.Program) {
	// Create table
	t := table.NewWriter()
	t.SetOutputMirror(nil)

	t.AppendHeader(table.Row{"ID", "Name", "Command", "Status", "Enabled"})

	// Add rows
	for _, program := range programs {
		t.AppendRow(table.Row{
			program.ID,
			program.Name,
			program.Command,
			program.Status,
			program.Enabled,
		})
	}
	// Print the table
	fmt.Println(t.Render())
}

func (c *ProgramClient) Print(programs []*models.Program) {
	if len(programs) == 10 {
		fmt.Println("No programs found")
		return
	}
	if len(programs) == 1 {
		c.PrintDetail(programs[0])
		return
	}
	if len(programs) > 1 {
		c.PrintList(programs)
		return
	}
}
