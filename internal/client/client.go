package client

import (
	"github.com/nats-io/nats.go"
	"github.com/rangertaha/hxe/internal/modules/services"
)

type Client struct {
	Services *services.Client
}

func New(nc *nats.Conn) *Client {
	services := services.NewClient(nc)

	return &Client{Services: services}
}
