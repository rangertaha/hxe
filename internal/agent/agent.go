/*
Copyright © 2023 Rangertaha  <rangertaha@gmail.com>
*/
package agent

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
	"github.com/rangertaha/hxe/internal/config"
	"github.com/rangertaha/hxe/internal/log"

	"github.com/rs/zerolog"
)

// Agent is the main struct for the hxe backend
type Agent struct {
	Config     *config.Config // Configuration
	Broker     *Broker        // Message broker
	Supervisor *Supervisor    // Process manager
	API        *APIServer     // API server
	log        zerolog.Logger
	sig        chan os.Signal
	conf       chan config.Config
	done       chan struct{}
}
type APIServer struct {
	log    zerolog.Logger
	router *echo.Echo
}

// New creates a new Agent instance
func New(cfg *config.Config) (agent *Agent, err error) {
	agent = &Agent{
		Config: cfg,
		log:    log.With().Logger(),
		sig:    make(chan os.Signal, 1),
		done:   make(chan struct{}),
		conf:   make(chan config.Config),
	}

	if err = agent.Init(); err != nil {
		return nil, err
	}

	return
}

// Init initializes the agent
func (a *Agent) Init() (err error) {
	a.log.Info().Msg("initializing agent")

	// Create messaging with server options
	a.Broker, err = NewBroker(a)
	if err != nil {
		return err
	}

	// Create task manager
	a.Supervisor, err = NewSupervisor(a)
	if err != nil {
		return err
	}

	// Create api server
	a.API, err = NewAPIServer(a)
	if err != nil {
		return err
	}

	return nil
}

// Stop the agent
func (a *Agent) Stop() {
	a.log.Info().Msg("stopping agent")

	// Stop api server
	a.API.Stop()

	// Stop message broker
	a.Broker.Stop()

	// close(a.done)
	a.log.Info().Msg("stopped agent")
}

// Start the agent
func (a *Agent) Start() error {

	// Watch for program changes in goroutine
	go func() {
		if err := a.Watch(a.Config.ConfigFile, a.Config.LoadConfig); err != nil {
			a.log.Error().Err(err).Msg("failed to watch config file")
		}
	}()

	// Watch for program changes in goroutine
	go func() {
		if err := a.Watch(a.Config.ProgDir, a.Config.LoadProgram); err != nil {
			a.log.Error().Err(err).Msg("failed to watch programs directory")
		}
	}()

	// Start broker in goroutine
	go func() {
		if err := a.Broker.Start(); err != nil {
			a.log.Error().Err(err).Msg("failed to start broker")
		}
	}()

	// Start supervisor in goroutine
	go func() {
		if err := a.Supervisor.StartSupervisor(); err != nil {
			a.log.Error().Err(err).Msg("failed to start supervisor")
		}
	}()

	// Start api server in goroutine
	go func() {
		if err := a.API.Start(); err != nil {
			a.log.Error().Err(err).Msg("failed to start api server")
		}
	}()

	// Setup signal handling
	signal.Notify(a.sig, syscall.SIGINT, syscall.SIGTERM)

	// Wait for shutdown signal
	select {
	case sig := <-a.sig:
		a.log.Debug().Msgf("received system signal %v, initiating shutdown", sig)
		a.Stop()
	case <-a.done:
		a.log.Debug().Msg("received done signal, initiating shutdown")
		a.Stop()
	}
	return nil
}

// // Load all programs from the programs directory
// func (a *Agent) LoadConfig() (err error) {
// 	if err = hclsimple.DecodeFile(a.Config.ConfigFile, config.CtxFunctions, a.Config); err != nil {
// 		a.log.Warn().Err(err).Str("file", a.Config.ConfigFile).Msg("failed to parse config file")
// 	}
// 	return
// }

// // Load all programs from the programs directory
// func (a *Agent) Load(paths ...string) error {
// 	// for _, p := range paths {
// 	// 	a.log.Info().Str("file", p).Msg("loading program")
// 	// }

// 	// var extract = func(path string) (string, error) {
// 	// 	dir, file := filepath.Split(path)
// 	// 	return filepath.Join(dir, strings.TrimSuffix(file, ".hcl")), nil
// 	// }

// 	a.log.Info().Str("dir", a.Config.ProgDir).Msg("loading programs from directory")
// 	var programs []config.Program
// 	err := filepath.WalkDir(a.Config.ProgDir, func(path string, d fs.DirEntry, err error) error {
// 		// If there was an error accessing a file or directory,
// 		// the function will be called with the error. We should handle it.
// 		if err != nil {
// 			fmt.Printf("Error accessing path %q: %v\n", path, err)
// 			return err
// 		}

// 		// We can check if the entry is a directory or a file.
// 		if d.IsDir() {
// 			fmt.Printf("Directory found: %s\n", path)
// 		} else {
// 			fmt.Printf("File found: %s\n", path)
// 			var program config.Program
// 			if strings.HasSuffix(path, ".hcl") {
// 				if err := hclsimple.DecodeFile(path, config.CtxFunctions, &program); err != nil {
// 					a.log.Warn().Err(err).Str("file", path).Msg("failed to parse program file")
// 				}
// 				programs = append(programs, program)
// 			}

// 		}

// 		// Return nil to continue walking the tree.
// 		// Returning a non-nil error will stop the walk immediately.
// 		return nil
// 	})

// 	// Handle any error that occurred during the walk itself.
// 	if err != nil {
// 		fmt.Printf("Error walking the path %q: %v\n", a.Config.ProgDir, err)
// 		os.Exit(1)
// 	}

// 	// Update the config with loaded programs
// 	a.Config.Program = &programs

// 	a.log.Info().Int("count", len(programs)).Msg("finished loading programs")
// 	return nil
// }

func (a *Agent) Watch(path string, loader func(path string) error) error {
	// Create a new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	// Start listening for events in a goroutine.
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				a.log.Info().Str("event", event.String()).Msg("event")
				if event.Has(fsnotify.Write) {
					loader(event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				a.log.Error().Err(err).Msg("error")
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		return err
	}

	// Keep the main goroutine alive until a signal is received (e.g., Ctrl+C).
	<-done
	return nil
}

// Execute the agent
func (a *Agent) Execute(args ...string) error {
	return nil
}
