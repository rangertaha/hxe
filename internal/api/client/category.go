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
	"github.com/nats-io/nats.go"
	models "github.com/rangertaha/hxe/internal/api/models"
)

type Category struct {
	nc *nats.Conn
}

func NewCategory(nc *nats.Conn) *Category {
	return &Category{
		nc: nc,
	}
}

// Load a service
func (e *Category) Load(cid uint) (err error) {
	return
}

// List all categories
func (e *Category) List() (cats []*models.Category, err error) {
	return
}

// Get a service by ID
func (e *Category) Get(cid uint) (ca *models.Category, err error) {
	return
}

// Create a new category
func (e *Category) Create(cat *models.Category) (c *models.Category, err error) {
	return
}

// Update a category
func (e *Category) Update(cat *models.Category) (c *models.Category, err error) {
	return
}

// Delete a category
func (e *Category) Delete(cid uint) (c *models.Category, err error) {
	return
}

// Start a category
func (e *Category) Start(cid uint) (c *models.Category, err error) {
	return
}

// Stop a category
func (e *Category) Stop(cid uint) (c *models.Category, err error) {
	return
}

// Restart a category
func (e *Category) Restart(cid uint) (c *models.Category, err error) {
	return
}
