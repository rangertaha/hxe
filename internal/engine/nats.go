package engine

import (
	"fmt"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/rangertaha/hxe/internal/config"
	"github.com/rangertaha/hxe/internal/log"
)

func NewMessaging(cfg *config.Mq) (ns *server.Server, nc *nats.Conn, err error) {
	log.Info().Msg("initializing messaging service")
	opts := &server.Options{Port: 3143, Host: "0.0.0.0"}
	ns, err = server.NewServer(opts)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create messaging server: %w", err)
	}
	// Start server in background
	go ns.Start()

	// Wait for server to be ready
	if !ns.ReadyForConnections(5 * time.Second) {
		return nil, nil, fmt.Errorf("server not ready for connections")
	}

	clientOpts := []nats.Option{
		nats.Name("hxe-server"),
		nats.InProcessServer(ns),
		nats.FlusherTimeout(5 * time.Second),
	}
	nc, err = nats.Connect(ns.ClientURL(), clientOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to messaging server: %w", err)
	}

	log.Info().Msgf("connected to messaging server: %s", ns.ClientURL())

	return ns, nc, nil

}

// // New creates a new configuration
// func New(options ...func(*Server) error) (srv *Server, err error) {
// 	srv = &Server{
// 		log: log.With().Logger(),
// 		// Client: &nats.Conn{},
// 		// Server: &server.Server{},
// 		done: make(chan bool),
// 	}

// 	// Apply config options
// 	for _, opt := range options {
// 		err := opt(srv)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	err = srv.Init()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return srv, nil
// }

// func ServerOpts(c config.Mq) func(*Server) (err error) {
// 	return func(s *Server) (err error) {
// 		s.svcOpts = &server.Options{
// 			Debug: c.Debug,
// 			Host:  c.Host,
// 			Port:  c.Port,
// 		}

// 		return nil
// 	}
// }

// // func StartEmbeddedServer() (ns *server.Server, nc *nats.Conn, err error) {

// // 	opts := &server.Options{Port: 3443, Host: "0.0.0.0"}
// // 	ns, err = server.NewServer(opts)
// // 	if err != nil {
// // 		return nil, nil, fmt.Errorf("failed to create messaging server: %w", err)
// // 	}

// // 	// Start server in background
// // 	go ns.Start()

// // 	// Wait for server to be ready
// // 	if !ns.ReadyForConnections(5 * time.Second) {
// // 		return nil, nil, fmt.Errorf("server not ready for connections")
// // 	}

// // 	clientOpts := []nats.Option{
// // 		nats.Name("hxe-server"),
// // 		nats.InProcessServer(ns),
// // 		nats.FlusherTimeout(5 * time.Second),
// // 		// nats.PendingLimits(-1, -1), // Unlimited pending messages
// // 	}
// // 	log.Info().Msgf("connecting to messaging server: %s", ns.ClientURL())
// // 	nc, err = nats.Connect(ns.ClientURL(), clientOpts...)
// // 	if err != nil {
// // 		return nil, nil, fmt.Errorf("failed to connect to messaging server: %w", err)
// // 	}

// // 	return ns, nc, nil
// // }

// func (s *Server) Init() (err error) {
// 	s.log.Info().Msg("initializing messaging service")
// 	opts := &server.Options{Port: 3143, Host: "0.0.0.0"}
// 	s.Server, err = server.NewServer(opts)
// 	if err != nil {
// 		return fmt.Errorf("failed to create messaging server: %w", err)
// 	}
// 	// Start server in background
// 	go s.Server.Start()

// 	// Wait for server to be ready
// 	if !s.Server.ReadyForConnections(5 * time.Second) {
// 		return fmt.Errorf("server not ready for connections")
// 	}

// 	clientOpts := []nats.Option{
// 		nats.Name("hxe-server"),
// 		nats.InProcessServer(s.Server),
// 		nats.FlusherTimeout(5 * time.Second),
// 	}
// 	s.Client, err = nats.Connect(s.Server.ClientURL(), clientOpts...)
// 	if err != nil {
// 		return fmt.Errorf("failed to connect to messaging server: %w", err)
// 	}

// 	s.log.Info().Msgf("connected to messaging server: %s", s.Server.ClientURL())

// 	return nil
// }

// func (s *Server) GetClient() (client *nats.Conn, err error) {
// 	s.log.Info().Msg("initializing messaging service")

// 	// // Wait for server to be ready
// 	// if !s.Server.ReadyForConnections(5 * time.Second) {
// 	// 	return nil, fmt.Errorf("server not ready for connections")
// 	// }

// 	clientOpts := []nats.Option{
// 		nats.Name("hxe-server"),
// 		nats.InProcessServer(s.Server),
// 		nats.FlusherTimeout(5 * time.Second),
// 		// nats.PendingLimits(-1, -1), // Unlimited pending messages
// 	}
// 	log.Info().Msgf("connecting to messaging server: %s", s.Server.ClientURL())
// 	s.Client, err = nats.Connect(s.Server.ClientURL(), clientOpts...)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect to messaging server: %w", err)
// 	}
// 	return s.Client, nil
// }

// func (s *Server) Load() error {
// 	s.log.Info().Msg("configuring messaging service")
// 	return nil
// }

// func (s *Server) Start() (err error) {
// 	s.log.Info().Msg("starting messaging service")

// 	// for {
// 	// 	select {
// 	// 	case <-s.done:
// 	// 		s.log.Info().Msg("stopping messaging service")
// 	// 		// s.Client.Close()
// 	// 		// s.Server.Shutdown()

// 	// 		return nil
// 	// 	default:
// 	// 		time.Sleep(1 * time.Second)
// 	// 	}
// 	// }
// 	return nil
// }

// func (s *Server) Stop() error {
// 	// close(s.done)
// 	return nil
// }

// func (s *Server) Restart() error {
// 	s.log.Info().Msg("restarting messaging service")
// 	return nil
// }

// func (s *Server) Status() error {
// 	s.log.Info().Msg("messaging service status")
// 	return nil
// }

// import (
// 	"fmt"
// 	"time"

// 	"github.com/nats-io/nats-server/v2/server"
// 	"github.com/nats-io/nats.go"
// 	"github.com/rangertaha/hxe/internal"
// 	"github.com/rangertaha/hxe/internal/log"
// 	"github.com/rs/zerolog"
// )

// // Messaging handles NATS server and client operations
// type Broker struct {
// 	Server *server.Server
// 	Client *nats.Conn
// 	log    zerolog.Logger
// }

// // NewMessaging creates a new Messaging instance
// func NewBroker(a *Agent) (*Broker, error) {
// 	// Create server
// 	natsServer, err := server.NewServer(a.Config.Broker.Options())
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create messaging server: %w", err)
// 	}

// 	// Configure server logger
// 	// natsServer.SetLogger(log.MsgLogger(), true, true)

// 	// Start server in background
// 	go natsServer.Start()

// 	// Wait for server to be ready
// 	if !natsServer.ReadyForConnections(5 * time.Second) {
// 		return nil, fmt.Errorf("server not ready for connections")
// 	}

// 	// Configure client options
// 	clientOpts := []nats.Option{
// 		nats.Name(internal.NAME),
// 		nats.InProcessServer(natsServer),
// 		nats.FlusherTimeout(5 * time.Second),
// 		// nats.PendingLimits(-1, -1), // Unlimited pending messages
// 	}

// 	// Connect to our own NATS server
// 	natsClient, err := nats.Connect(natsServer.ClientURL(), clientOpts...)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect to messaging server: %w", err)
// 	}

// 	return &Broker{
// 		Server: natsServer,
// 		Client: natsClient,
// 		log:    log.With().Logger(),
// 	}, nil
// }

// func (b *Broker) Init() error {
// 	b.log.Info().Msg("initializing messaging broker")

// 	return nil
// }

// // Start starts listening for messages
// func (b *Broker) Start() error {
// 	b.log.Info().Msg("starting messaging broker")

// 	// // Subscribe to ares subject
// 	// if err := b.Subscribe("ares", b.handleMessage); err != nil {
// 	// 	return fmt.Errorf("failed to subscribe to ares: %w", err)
// 	// }
// 	return nil
// }

// // Stop stops the messaging system
// func (b *Broker) Stop() error {
// 	b.log.Info().Msg("stopping messaging broker")

// 	// First close the client to stop new messages
// 	if b.Client != nil {
// 		// Drain subscriptions before closing
// 		if err := b.Client.Drain(); err != nil {
// 			b.log.Warn().Err(err).Msg("Error draining client subscriptions")
// 		}
// 		b.Client.Close()
// 	}

// 	// Then shutdown the server
// 	if b.Server != nil {
// 		// Wait for server to finish processing messages
// 		b.Server.Shutdown()
// 	}

// 	b.log.Debug().Msg("stopped messaging broker")
// 	return nil
// }

// // Subscribe subscribes to a NATS subject
// func (b *Broker) Subscribe(subject string, handler nats.MsgHandler) error {
// 	if b.Client == nil {
// 		return fmt.Errorf("messaging client not initialized")
// 	}
// 	_, err := b.Client.Subscribe(subject, handler)
// 	return err
// }

// // Publish publishes a message to a NATS subject
// func (b *Broker) Publish(subject string, data []byte) error {
// 	if b.Client == nil {
// 		return fmt.Errorf("messaging client not initialized")
// 	}
// 	return b.Client.Publish(subject, data)
// }

// // Request sends a request to a NATS subject and waits for a response
// func (b *Broker) Request(subject string, data []byte, timeout time.Duration) (*nats.Msg, error) {
// 	if b.Client == nil {
// 		return nil, fmt.Errorf("messaging client not initialized")
// 	}
// 	return b.Client.Request(subject, data, timeout)
// }

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

// func New(cfg *config.Messaging) (srv *Server, nc *nats.Conn, err error) {
// 	srv = &Server{
// 		config: cfg,
// 		log:    log.With().Logger(),
// 		done:   make(chan bool),
// 	}
// 	if cfg.Embed {
// 		srv.svcOpts = cfg.Options()
// 		if srv.svcOpts.Port == 0 && srv.svcOpts.Host == "" {
// 			srv.isIPC = true
// 		} else {
// 			return nil, nil, fmt.Errorf("invalid server host/port options: %v", srv.svcOpts)
// 		}
// 	}

// 	if srv.svcOpts.Port == 0 && srv.svcOpts.Host == "" {

// 	}

// 	// Create server
// 	if cfg.Embed {
// 		opts := &server.Options{}
// 		if cfg.Port != 0 && cfg.Host != "" {
// 			opts := &server.Options{Port: cfg.Port, Host: cfg.Host}
// 		}

// 		opts := &server.Options{Port: cfg.Port}
// 		srv.ns, err = server.NewServer(opts)
// 		if err != nil {
// 			return nil, nil, fmt.Errorf("failed to create messaging server: %w", err)
// 		}

// 		// Configure server logger
// 		// natsServer.SetLogger(log.MsgLogger(), true, true)

// 		// Start server in background
// 		go srv.ns.Start()

// 		// Wait for server to be ready
// 		if !srv.ns.ReadyForConnections(5 * time.Second) {
// 			return nil, nil, fmt.Errorf("server not ready for connections")
// 		}

// 		// Configure client options
// 		clientOpts := []nats.Option{
// 			nats.Name("hxe-server"),
// 			nats.InProcessServer(srv.ns),
// 			nats.FlusherTimeout(5 * time.Second),
// 			// nats.PendingLimits(-1, -1), // Unlimited pending messages
// 		}

// 		// Connect to our own NATS server
// 		nc, err = nats.Connect(srv.ns.ClientURL(), clientOpts...)
// 		if err != nil {
// 			return nil, nil, fmt.Errorf("failed to connect to messaging server: %w", err)
// 		}
// 	} else {
// 		// Connect to external NATS server
// 		url := nats.DefaultURL
// 		if cfg.Host != "" && cfg.Port != 0 {
// 			url = fmt.Sprintf("nats://%s:%d", cfg.Host, cfg.Port)
// 		}

// 		nc, err = nats.Connect(url)
// 		if err != nil {
// 			return nil, nil, fmt.Errorf("failed to connect to messaging server: %w", err)
// 		}
// 	}

// 	// subject := "my-subject"

// 	// // Subscribe to the subject
// 	// natsClient.Subscribe(subject, func(msg *nats.Msg) {
// 	// 	// Print message data
// 	// 	data := string(msg.Data)
// 	// 	fmt.Println(data)

// 	// 	// Shutdown the server (optional)
// 	// 	natsServer.Shutdown()
// 	// })

// 	// Publish data to the subject
// 	// nc.Publish(subject, []byte("Hello embedded NATS!"))

// 	// // Wait for server shutdown
// 	// natsServer.WaitForShutdown()

// 	if err := srv.Init(); err != nil {
// 		return nil, nil, err
// 	}

// 	return
// }
