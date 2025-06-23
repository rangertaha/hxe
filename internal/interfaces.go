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
