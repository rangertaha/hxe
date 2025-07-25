package client

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rangertaha/hxe/internal/config"
	prog "github.com/rangertaha/hxe/internal/services/program/client"
)

const (
	DefaultBaseURL  = "nats://0.0.0.0:8080"
	DefaultUsername = "admin"
	DefaultPassword = "admin"
	DefaultTimeout  = 30 * time.Second
)

type (
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	Client struct {
		Config   *config.Client
		Programs *prog.Client
		conn     *nats.Conn
	}
)

// New creates a new client
func New(c *config.Client) (clt *Client, err error) {

	clt = &Client{
		Config: c,
	}

	// Initialize the client
	if err = clt.Init(); err != nil {
		return nil, err
	}

	clt.Programs = prog.New(clt.conn)

	return clt, nil
}

func (c *Client) Init() (err error) {
	// clientOpts := []nats.Option{
	// 	nats.Name("hxe-client"),
	// 	nats.FlusherTimeout(5 * time.Second),
	// }

	c.conn, err = nats.Connect(c.Config.Url, c.Config.Options())
	if err != nil {
		return fmt.Errorf("failed to connect to messaging server: %w", err)
	}

	return nil
}

func (c *Client) Close() (err error) {
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
	return nil
}

func (c *Client) IsConnected() (connected bool) {
	return c.conn != nil
}

func (c *Client) Login() (err error) {
	return nil
}

func (c *Client) Logout() (err error) {
	return nil
}
