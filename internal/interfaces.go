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

package internal

import (
	"time"

	"github.com/nats-io/nats.go"
)

type Broker interface {
	Init() error
	Subscribe(subject string, handler nats.MsgHandler) error
	Publish(subject string, data []byte) error
	Request(subject string, data []byte, timeout time.Duration) (*nats.Msg, error)
}

type Middleware interface {
	Produce(in string, indata []byte) (out string, outdata []byte)
	Consume(in string, indata []byte) (out string, outdata []byte)
}

type Producer interface {
	Produce(subject string, data []byte) (n int, err error)
}

type Consumer interface {
	Consume(subject string, data []byte) (n int, err error)
}

type Executor interface {
	Stdin() error
	Stdcmd() error
	Stdout() error
	Stderr() error
}
