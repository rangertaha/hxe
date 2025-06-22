/*
Copyright © 2023 Rangertaha  <rangertaha@gmail.com>
*/
package config

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/rangertaha/hxe/internal"

	_ "embed"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

const (
	CONFIG_DIR     = "hxe"
	CONFIG_FILE    = "config.hcl"
	PROGRAM_FILE   = "example.hcl"
	DefaultSubject = "hxe.cmd"
)

var (
	//go:embed config.hcl
	DefaultConfig []byte

	//go:embed program.hcl
	DefaultProgram []byte
)

type (
	Config struct {
		ID      string        `hcl:"id,optional"`
		Debug   bool          `hcl:"debug,optional"`
		Level   zerolog.Level `hcl:"level,optional"`
		Version string        `hcl:"version,optional"`
		Banner  bool          `hcl:"banner,optional"`
		ProgDir string        `hcl:"programs,optional"`

		Programs []Program
		Broker   Broker `hcl:"broker,block"`
		API      API    `hcl:"api,block"`

		ConfigFile string
	}
	Program struct {
		Name    string   `hcl:"name,optional"`
		Command string   `hcl:"command,optional"`
		Args    []string `hcl:"args,optional"`
	}
	API struct {
		Host     string `hcl:"addr,optional"`
		Port     int    `hcl:"port,optional"`
		Username string `hcl:"username,optional"`
		Password string `hcl:"password,optional"`
	}
)

func init() {
	// Outputs to nowhere
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(io.Discard)

}

// New creates a new configuration
func New(options ...func(*Config) error) (*Config, error) {
	s := &Config{
		Level:  zerolog.TraceLevel,
		Banner: true,
	} // Default values

	// Apply config options
	for _, opt := range options {
		err := opt(s)
		if err != nil {
			return nil, err
		}
	}

	if s.Debug {
		zerolog.SetGlobalLevel(zerolog.Level(s.Level))
		log.Logger = log.Output(os.Stdout)
	}

	if s.Banner {
		internal.PrintBanner()
	}

	return s, nil
}

func (c *Config) LoadConfig(path string) (err error) {
	if err = hclsimple.DecodeFile(path, CtxFunctions, c); err != nil {
		return fmt.Errorf("error parsing config file: %w", err)
	}
	return
}

func (c *Config) LoadProgram(ppath string) (err error) {
	log.Info().Str("path", ppath).Msg("loading programs from directory")
	err = filepath.WalkDir(ppath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Error().Err(err).Str("path", path).Msg("error accessing path")
			return err
		}

		if !d.IsDir() {
			log.Info().Str("file", path).Msg("loading program")
			var program Program
			if strings.HasSuffix(path, ".hcl") {
				if err := hclsimple.DecodeFile(path, CtxFunctions, &program); err != nil {
					log.Warn().Err(err).Str("file", path).Msg("failed to parse program file")
				} else {
					c.Programs = append(c.Programs, program)
				}
			}
		}
		return nil
	})

	// Errors that occurred during the walk.
	if err != nil {
		log.Error().Err(err).Str("path", c.ProgDir).Msg("error walking the path")
		os.Exit(1)
	}

	// Update the config with loaded programs
	log.Info().Int("count", len(c.Programs)).Msg("finished loading programs")
	return
}

func CliOptions(ctx context.Context, cmd *cli.Command) func(c *Config) error {

	return func(c *Config) error {
		c.Debug = cmd.Bool("debug")
		c.Level = zerolog.Level(cmd.Int("level"))
		return nil
	}
}

func FileOption(path string) func(*Config) error {
	return func(c *Config) error {
		if path == "" {
			return nil
		}

		c.ConfigFile = path

		if err := c.LoadConfig(c.ConfigFile); err != nil {
			return fmt.Errorf("error parsing config file: %w", err)
		}

		if err := c.LoadProgram(c.ProgDir); err != nil {
			return fmt.Errorf("error parsing program file: %w", err)
		}

		return nil
	}
}
func DefaultOptions() func(*Config) error {
	return func(c *Config) error {

		userConfigDir, err := os.UserConfigDir()
		if err != nil {
			return fmt.Errorf("error getting user config directory: %w", err)
		}
		c.ConfigFile = filepath.Join(userConfigDir, CONFIG_DIR, CONFIG_FILE)
		if err := createFileIfNotExists(c.ConfigFile, DefaultConfig); err != nil {
			return fmt.Errorf("error creating config file: %w", err)
		}

		c.ProgDir = filepath.Join(userConfigDir, CONFIG_DIR, "programs")
		exampleFile := filepath.Join(c.ProgDir, PROGRAM_FILE)
		if err := createFileIfNotExists(exampleFile, DefaultProgram); err != nil {
			return fmt.Errorf("error creating default program config: %w", err)
		}

		if err := c.LoadConfig(c.ConfigFile); err != nil {
			return fmt.Errorf("error parsing config file: %w", err)
		}

		if err := c.LoadProgram(c.ProgDir); err != nil {
			return fmt.Errorf("error parsing program file: %w", err)
		}

		return nil
	}
}

func createFileIfNotExists(filename string, contents []byte) error {
	// Check if file exists
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// Create directory if it doesn't exist
		dir := filepath.Dir(filename)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create config directory: %w", err)
		}

		// Write default config to file
		if err := os.WriteFile(filename, contents, 0644); err != nil {
			return fmt.Errorf("failed to create default config file: %w", err)
		}

		fmt.Printf("Created default configuration file: %s\n", filename)

	} else if err != nil {
		return fmt.Errorf("error checking config file: %w", err)
	}

	return nil
}
