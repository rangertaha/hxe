package services

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/rangertaha/hxe/internal/interfaces"
)


type Creator func(nc *nats.Conn) interfaces.Service

var Services = map[string]Creator{}

func Add(name string, creator Creator) {
	Services[name] = creator
}

func Get(name string) (Creator, error) {
	if plugin, ok := Services[name]; ok {
		return plugin, nil
	}

	return nil, fmt.Errorf("unable to locate service: %s", name)
}
