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

package program

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"github.com/rangertaha/hxe/internal"
	"github.com/rangertaha/hxe/internal/db"
	"github.com/rangertaha/hxe/internal/log"
	pc "github.com/rangertaha/hxe/internal/services/program/client"
	"github.com/rangertaha/hxe/internal/services/program/models"
	"github.com/rs/zerolog"
)

type Crud interface {
	List(*models.Program) ([]*models.Program, error)
}

type Microservice struct {
	service micro.Service
	log     zerolog.Logger
}

func NewMicroservice(nc *nats.Conn) *Microservice {

	svc, err := micro.AddService(nc, micro.Config{
		Name:        "programs",
		Version:     internal.VERSION,
		Description: "Program process manager",
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to add service")
		return nil
	}

	return &Microservice{
		service: svc,
		log:     log.With().Str("service", "program").Logger(),
	}
}

func (s *Microservice) Init() (err error) {
	svc := s.service.AddGroup("program")
	// svc.AddEndpoint("load", JSONHandler(e.Load))
	svc.AddEndpoint("list", JSONHandler(s.List))
	svc.AddEndpoint("get", JSONHandler(s.Get))
	svc.AddEndpoint("create", JSONHandler(s.Create))
	svc.AddEndpoint("update", JSONHandler(s.Update))
	svc.AddEndpoint("delete", JSONHandler(s.Delete))
	// svc.AddEndpoint("start", JSONHandler(e.Start))
	// svc.AddEndpoint("stop", JSONHandler(e.Stop))
	// svc.AddEndpoint("restart", JSONHandler(e.Restart))
	// svc.AddEndpoint("status", JSONHandler(e.Status))
	// svc.AddEndpoint("log", JSONHandler(e.Log))
	// svc.AddEndpoint("shell", JSONHandler(e.Shell))

	return models.AutoMigrate()
}

// Load a service (no request needed)
// func (s *Service) Load(req *Request) (res *Response) {
// 	// s.Runner.Load(req.Service)
// 	return
// }

// List all services
func (s *Microservice) List(req *pc.Request) (res *pc.Response) {
	progs := []*models.Program{}
	db.DB.Find(&progs)
	return &pc.Response{Programs: progs}
}

// Get a service by ID
func (s *Microservice) Get(req *pc.Request) (res *pc.Response) {
	progs := []*models.Program{}
	db.DB.Find(&progs, "id = ?", req.Program.ID)
	return &pc.Response{Programs: progs}
}

// Create a new service
func (s *Microservice) Create(req *pc.Request) (res *pc.Response) {
	db.DB.Create(req.Program)
	return &pc.Response{Programs: []*models.Program{req.Program}}
}

// Update a service
func (s *Microservice) Update(req *pc.Request) (res *pc.Response) {
	db.DB.Save(req.Program)
	return &pc.Response{Programs: []*models.Program{req.Program}}
}

// Delete a service
func (s *Microservice) Delete(req *pc.Request) (res *pc.Response) {
	db.DB.Delete(req.Program)
	return &pc.Response{Programs: []*models.Program{req.Program}}
}

// // Start a service
// func (s *Service) Start(req *models.Request) (res *models.Response) {
// 	res = &models.Response{}
// 	res.Service, res.Status = s.Storage.Get(req.Service)
// 	s.Runner.Start(res.Service)
// 	res.Service.Save()
// 	return res
// }

// // Stop a service
// func (s *Service) Stop(req *models.Request) (res *models.Response) {
// 	res = &models.Response{}
// 	res.Service, res.Status = s.Storage.Get(req.Service)
// 	s.Runner.Stop(res.Service)
// 	return
// }

// // Restart a service
// func (s *Service) Restart(req *models.Request) (res *models.Response) {
// 	res = &models.Response{}
// 	res.Service, res.Status = s.Storage.Get(req.Service)
// 	s.Runner.Restart(res.Service)
// 	res.Service.Save()
// 	return
// }

// // Status of a service
// func (s *Service) Status(req *models.Request) (res *models.Response) {
// 	return s.Get(req)
// }

// // Log a service
// func (s *Service) Log(req *models.Request) (res *models.Response) {

// 	return
// }

// // Shell a service
// func (s *Service) Shell(req *models.Request) (res *models.Response) {

// 	return
// }

// JSONHandler wraps a handler function with automatic marshaling/unmarshaling
func JSONHandler(handler func(*pc.Request) *pc.Response) micro.HandlerFunc {
	return func(msg micro.Request) {
		var req pc.Request

		// Unmarshal request
		if err := json.Unmarshal(msg.Data(), &req); err != nil {
			msg.Error("500", "Unmarshal error", nil)
			return
		}

		// Call the handler
		res := handler(&req)

		// Marshal and send response
		data, err := json.Marshal(res)
		if err != nil {
			msg.Error("500", "Marshal error", nil)
			return
		}
		msg.Respond(data)
	}
}
