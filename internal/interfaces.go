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
