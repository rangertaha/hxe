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

package services

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"github.com/rangertaha/hxe/internal"
	"github.com/rangertaha/hxe/internal/log"
	"github.com/rangertaha/hxe/internal/modules/services/models"
	"github.com/rangertaha/hxe/internal/rdb"
	"github.com/rs/zerolog"
)

type Service struct {
	Runner  *Runner
	Service micro.Service
	Storage *models.Service
	log     zerolog.Logger
}

func NewService(nc *nats.Conn) (s *Service, err error) {
	s = &Service{
		Storage: models.NewService(rdb.DB),
		Runner:  NewRunner(),
		log:     log.With().Logger(),
	}
	config := micro.Config{
		Name:        "Services",
		Version:     internal.VERSION,
		Description: "Process manager",
	}

	s.Service, err = micro.AddService(nc, config)
	if err != nil {
		s.log.Fatal().Err(err).Msg("failed to add service")
	}

	if err := s.Init(); err != nil {
		s.log.Fatal().Err(err).Msg("failed to init service")
		return nil, err
	}

	return
}

func (e *Service) Init() (err error) {
	svc := e.Service.AddGroup("service")
	// svc.AddEndpoint("load", JSONHandler(e.Load))
	svc.AddEndpoint("list", JSONHandler(e.List))
	// svc.AddEndpoint("get", JSONHandler(e.Get))
	// svc.AddEndpoint("create", JSONHandler(e.Create))
	// svc.AddEndpoint("update", JSONHandler(e.Update))
	// svc.AddEndpoint("delete", JSONHandler(e.Delete))
	// svc.AddEndpoint("start", JSONHandler(e.Start))
	// svc.AddEndpoint("stop", JSONHandler(e.Stop))
	// svc.AddEndpoint("restart", JSONHandler(e.Restart))
	// svc.AddEndpoint("status", JSONHandler(e.Status))
	// svc.AddEndpoint("log", JSONHandler(e.Log))
	// svc.AddEndpoint("shell", JSONHandler(e.Shell))

	return
}

// Load a service (no request needed)
// func (s *Service) Load(req *Request) (res *Response) {
// 	// s.Runner.Load(req.Service)
// 	return
// }

// List all services
func (s *Service) List(req *Request) (res *Response) {
	res = &Response{Service: &rdb.Service{ID: 0}}
	res.Services, res.Status = s.Storage.List(req.Service)
	s.log.Info().Msgf("Services: %d", len(res.Services))
	for _, service := range res.Services {
		s.log.Info().Msgf("Service: %s", service.Name)
	}
	return
}

// // Get a service by ID
// func (s *Service) Get(req *models.Request) (res *models.Response) {
// 	res = &models.Response{}
// 	res.Service, res.Status = s.Storage.Get(req.Service)
// 	return
// }

// // Create a new service
// func (s *Service) Create(req *models.Request) (res *models.Response) {
// 	res = &models.Response{}
// 	fmt.Println("pre Create service", req)
// 	res.Service, res.Status = s.Storage.Create(req.Service)
// 	fmt.Println("post Create service")
// 	return
// }

// // Update a service
// func (s *Service) Update(req *models.Request) (res *models.Response) {
// 	res = &models.Response{}
// 	res.Service, res.Status = s.Storage.Update(req.Service)
// 	return
// }

// // Delete a service
// func (s *Service) Delete(req *models.Request) (res *models.Response) {
// 	res = &models.Response{}
// 	res.Service, res.Status = s.Storage.Delete(req.Service)
// 	return
// }

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

// Proto wraps a handler function with automatic marshaling/unmarshaling
func JSONHandler(handler func(*Request) *Response) micro.HandlerFunc {
	return func(msg micro.Request) {
		var req Request

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
