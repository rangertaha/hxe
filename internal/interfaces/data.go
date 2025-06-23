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

package interfaces

import "time"

type Provider interface {
	Init() error
	Update() error
	Get() error
	Save() error
	Stop()
}

type Metadata interface {
	Markets() []string
	Symbols() []string
	Exchanges() []string
	Currencies() []string
	Assets() []string
	Types() []string
}

type Dataset interface {
	Source() string
	Interval() time.Duration
	Start() time.Time
	Stop() time.Time

	// Market returns the market name
	Market() string

	// Symbol returns the symbol name
	Symbol() string

	// Exchange returns the exchange name
	Exchange() string

	// Currency returns the currency name
	Currency() string

	// Asset returns the asset name
	Asset() string

	// Type returns the asset type name
	Type() string
}

type Backfiller interface {
	Init(Dataset) error
	Ticker() time.Ticker
	Get() error
	Save() error
	Stop()
}
