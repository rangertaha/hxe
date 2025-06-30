// package main

// import (
// 	"fmt"
// 	"time"

// 	"github.com/nats-io/nats-server/v2/server"
// 	"github.com/nats-io/nats.go"
// 	"github.com/rangertaha/hxe/internal/log"
// )

// func main() {
// 	// Create server
// 	var err error
// 	opts := &server.Options{}
// 	natsServer, err := server.NewServer(opts)
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("failed to create messaging server")
// 	}

// 	// Configure server logger
// 	// natsServer.SetLogger(log.MsgLogger(), true, true)

// 	// Start server in background
// 	go natsServer.Start()

// 	// Wait for server to be ready
// 	if !natsServer.ReadyForConnections(5 * time.Second) {
// 		log.Fatal().Msg("server not ready for connections")
// 	}

// 	// Configure client options
// 	clientOpts := []nats.Option{
// 		nats.Name("hxe-server"),
// 		nats.InProcessServer(natsServer),
// 		nats.FlusherTimeout(5 * time.Second),
// 		// nats.PendingLimits(-1, -1), // Unlimited pending messages
// 	}

// 	// Connect to our own NATS server
// 	natsClient, err := nats.Connect(natsServer.ClientURL(), clientOpts...)
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("failed to connect to messaging server")
// 	}
// 	subject := "my-subject"

// 	// Subscribe to the subject
// 	natsClient.Subscribe(subject, func(msg *nats.Msg) {
// 		// Print message data
// 		data := string(msg.Data)
// 		fmt.Println(data)

// 		// Shutdown the server (optional)
// 		natsServer.Shutdown()
// 	})

// 	// Publish data to the subject
// 	natsClient.Publish(subject, []byte("Hello embedded NATS!"))

// 	// Wait for server shutdown
// 	natsServer.WaitForShutdown()

// }

package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/rangertaha/hxe/internal/log"
)

func main() {
	nc, ns, err := EmbedNats()

	if err != nil {
		log.Fatal().Err(err).Msg("failed to create messaging server")
	}
	// Subscribe to the subject
	nc.Subscribe("test", func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))

		// Shutdown the server (optional)
		// ns.Shutdown()
	})
	for i := 0; i < 10; i++ {
		nc.Publish("test", []byte("Hello, World!"))
		time.Sleep(1 * time.Second)
	}

	ns.Shutdown()
}

func EmbedNats() (nc *nats.Conn, ns *server.Server, err error) {
	opts := &server.Options{Port: 4242}

	// Initialize new server with options
	ns, err = server.NewServer(opts)

	if err != nil {
		log.Fatal().Err(err).Msg("failed to create messaging server")
	}

	// Start the server via goroutine
	go ns.Start()

	// Wait for server to be ready for connections
	if !ns.ReadyForConnections(5 * time.Second) {
		log.Fatal().Msg("server not ready for connections")
	}

	// Configure client options
	clientOpts := []nats.Option{
		nats.Name("hxe-server"),
		nats.InProcessServer(ns),
		nats.FlusherTimeout(5 * time.Second),
		// nats.PendingLimits(-1, -1), // Unlimited pending messages
	}

	// Connect to server
	nc, err = nats.Connect(ns.ClientURL(), clientOpts...)

	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to messaging server")
	}

	return nc, ns, nil

	// subject := "my-subject"

	// // Subscribe to the subject
	// nc.Subscribe(subject, func(msg *nats.Msg) {
	//     // Print message data
	//     data := string(msg.Data)
	//     fmt.Println(data)

	//     // Shutdown the server (optional)
	//     ns.Shutdown()
	// })

	// // Publish data to the subject
	// nc.Publish(subject, []byte("Hello embedded NATS!"))

	// // Wait for server shutdown
	// ns.WaitForShutdown()
}
