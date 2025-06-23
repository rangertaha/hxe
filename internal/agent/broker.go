package agent

import (
	"fmt"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/rangertaha/hxe/internal"
	"github.com/rangertaha/hxe/internal/log"
	"github.com/rs/zerolog"
)

// Messaging handles NATS server and client operations
type Broker struct {
	Server *server.Server
	Client *nats.Conn
	log    zerolog.Logger
}

// NewMessaging creates a new Messaging instance
func NewBroker(a *Agent) (*Broker, error) {
	// Create server
	natsServer, err := server.NewServer(a.Config.Broker.Options())
	if err != nil {
		return nil, fmt.Errorf("failed to create messaging server: %w", err)
	}

	// Configure server logger
	// natsServer.SetLogger(log.MsgLogger(), true, true)

	// Start server in background
	go natsServer.Start()

	// Wait for server to be ready
	if !natsServer.ReadyForConnections(5 * time.Second) {
		return nil, fmt.Errorf("server not ready for connections")
	}

	// Configure client options
	clientOpts := []nats.Option{
		nats.Name(internal.NAME),
		nats.InProcessServer(natsServer),
		nats.FlusherTimeout(5 * time.Second),
		// nats.PendingLimits(-1, -1), // Unlimited pending messages
	}

	// Connect to our own NATS server
	natsClient, err := nats.Connect(natsServer.ClientURL(), clientOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to messaging server: %w", err)
	}

	return &Broker{
		Server: natsServer,
		Client: natsClient,
		log:    log.With().Logger(),
	}, nil
}

func (b *Broker) Init() error {
	b.log.Info().Msg("initializing messaging broker")

	return nil
}

// Start starts listening for messages
func (b *Broker) Start() error {
	b.log.Info().Msg("starting messaging broker")

	// // Subscribe to ares subject
	// if err := m.Subscribe("ares", m.handleMessage); err != nil {
	// 	return fmt.Errorf("failed to subscribe to ares: %w", err)
	// }
	return nil
}

// Stop stops the messaging system
func (b *Broker) Stop() error {
	b.log.Info().Msg("stopping messaging broker")

	// First close the client to stop new messages
	if b.Client != nil {
		// Drain subscriptions before closing
		if err := b.Client.Drain(); err != nil {
			b.log.Warn().Err(err).Msg("Error draining client subscriptions")
		}
		b.Client.Close()
	}

	// Then shutdown the server
	if b.Server != nil {
		// Wait for server to finish processing messages
		b.Server.Shutdown()
	}

	b.log.Debug().Msg("stopped messaging broker")
	return nil
}

// Subscribe subscribes to a NATS subject
func (b *Broker) Subscribe(subject string, handler nats.MsgHandler) error {
	if b.Client == nil {
		return fmt.Errorf("messaging client not initialized")
	}
	_, err := b.Client.Subscribe(subject, handler)
	return err
}

// Publish publishes a message to a NATS subject
func (b *Broker) Publish(subject string, data []byte) error {
	if b.Client == nil {
		return fmt.Errorf("messaging client not initialized")
	}
	return b.Client.Publish(subject, data)
}

// Request sends a request to a NATS subject and waits for a response
func (b *Broker) Request(subject string, data []byte, timeout time.Duration) (*nats.Msg, error) {
	if b.Client == nil {
		return nil, fmt.Errorf("messaging client not initialized")
	}
	return b.Client.Request(subject, data, timeout)
}

// func (b *Broker) Req(msg *nats.Msg) (req *client.Request, err error) {
// 	b.log.Debug().Msgf("server unmarshaling: %s", string(msg.Data))
// 	req = &client.Request{}

// 	// Parse the request
// 	if err = proto.Unmarshal(msg.Data, req); err != nil {
// 		b.Err(msg, err)
// 		return
// 	}
// 	return
// }

// func (b *Broker) Err(msg *nats.Msg, err error) {
// 	resData, marshalErr := proto.Marshal(&client.Response{Status: []*client.Status{
// 		{
// 			State:   client.STATE_ERRORING,
// 			Message: err.Error(),
// 		},
// 	}})
// 	if marshalErr != nil {
// 		b.log.Error().Err(marshalErr).Msg("failed to marshal error response")
// 		return
// 	}
// 	msg.Respond(resData)
// }

// func (b *Broker) Res(msg *nats.Msg, resources []*client.Resource) {
// 	b.log.Debug().Msgf("server responding with %d bots", len(resources))

// 	// Marshal and send response
// 	resData, err := proto.Marshal(&client.Response{Resources: resources})
// 	if err != nil {
// 		resData, err := proto.Marshal(&client.Response{Status: []*client.Status{
// 			{
// 				State:   client.STATE_ERRORING,
// 				Message: "server failed to marshal response",
// 			},
// 		}})
// 		b.log.Error().Err(err).Msg("server failed to marshal response")
// 		msg.Respond(resData)
// 		return
// 	}

// 	msg.Respond(resData)
// }
