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

package executors

import (
	"fmt"

	"github.com/rangertaha/hxe/internal"
)

type Creator func() internal.Executor

var Plugins = map[string]Creator{}

func Add(name string, creator Creator) {
	Plugins[name] = creator
}

func Get(name string) (Creator, error) {
	if plugin, ok := Plugins[name]; ok {
		return plugin, nil
	}

	return nil, fmt.Errorf("unable to locate plugin/%s", name)
}

func List() []string {
	names := make([]string, 0, len(Plugins))
	for name := range Plugins {
		names = append(names, name)
	}

	return names
}
