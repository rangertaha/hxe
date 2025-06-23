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

package interfaces

import "time"

type Status interface {
	State() int32
	Uptime() time.Duration
	Message() string
	Progress() int32
}

type Runner interface {
	Id() string
	Init() error
	Start() error
	Stop() error
	Fill() error
	Test() error
	Train() error
	Status() Status
}

type Plugin interface {
	Init() error
	Stdin() error
	Stdcmd() error
	Stdout() error
	Stderr() error
}
