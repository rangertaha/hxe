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
	"github.com/nats-io/nats.go"
	"github.com/rangertaha/hxe/internal/interfaces"
	"github.com/rangertaha/hxe/internal/log"
	"github.com/rangertaha/hxe/internal/services"
	"github.com/rs/zerolog"
)

type Service struct {
	micro *Microservice
	log   zerolog.Logger
}

func (s *Service) Init() (err error) {
	return s.micro.Init()
}

func (s *Service) Start() (err error) {
	return
}

func (s *Service) Stop() (err error) {
	return
}

func (s *Service) Reload() (err error) {
	return
}

func (s *Service) Restart() (err error) {
	return
}

func (s *Service) Status() (status string) {
	return
}

// Register the service
func init() {
	services.Add("programs", func(nc *nats.Conn) interfaces.Service {
		return &Service{
			log:   log.With().Logger(),
			micro: NewMicroservice(nc),
		}
	})
}
