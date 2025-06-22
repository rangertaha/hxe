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
