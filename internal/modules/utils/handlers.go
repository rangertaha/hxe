package utils

import (
	"errors"

	"github.com/nats-io/nats.go/micro"
)

type (
	// ProtoHandlerFunc is a function implementing [Handler].
	// It allows using a function as a request handler, without having to implement Handle
	// on a separate type.
	HandlerFunc func(micro.Request) micro.HandlerFunc
)

var (
	ErrRespond         = errors.New("NATS error when sending response")
	ErrMarshalResponse = errors.New("marshaling response")
	ErrArgRequired     = errors.New("argument required")
)

func ProtoHandlerFunc(fn HandlerFunc) micro.HandlerFunc {
	return micro.HandlerFunc(func(req micro.Request) {
		fn(req)
	})
}
