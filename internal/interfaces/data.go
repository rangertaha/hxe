/*
Copyright © 2025 Rangertaha <rangertaha@gmail.com>

Licensed under the MIT License (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://opensource.org/licenses/MIT

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
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
