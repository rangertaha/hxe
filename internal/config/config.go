/*
 * HXE - Host-based Process Execution Engine
 * Copyright (C) 2025 rangertaha@gmail.com
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package config

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/rangertaha/hxe/internal"
	"github.com/rangertaha/hxe/internal/log"
	"github.com/rangertaha/hxe/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "embed"

	"github.com/rs/zerolog"
	"github.com/urfave/cli/v3"
)

const (
	CONFIG_DIR      = "hxe"
	CONFIG_FILE     = "config.hcl"
	PROGRAM_FILE    = "example.hcl"
	DATABASE_FILE   = "hxe.db"
	DEFAULT_SUBJECT = "hxe"
)

var (
	//go:embed config.hcl
	DefaultConfig []byte

	//go:embed program.hcl
	DefaultProgram []byte
)

type (
	Config struct {
		ID      string `hcl:"id,optional"`
		Debug   bool   `hcl:"debug,optional"`
		Version string `hcl:"version,optional"`
		Banner  bool   `hcl:"banner,optional"`
		ProgDir string `hcl:"programs,optional"`

		Programs []models.Program
		Database Database `hcl:"database,block"`
		Broker   Broker   `hcl:"broker,block"`
		API      API      `hcl:"api,block"`

		ConfigFile string
	}
	API struct {
		Host     string `hcl:"addr,optional"`
		Port     int    `hcl:"port,optional"`
		Username string `hcl:"username,optional"`
		Password string `hcl:"password,optional"`
		Token    string `hcl:"token,optional"`
		Client   Client `hcl:"client,block"`
	}
	Client struct {
		Host     string `hcl:"addr,optional"`
		Port     int    `hcl:"port,optional"`
		URL      string `hcl:"url,optional"`
		Token    string `hcl:"token,optional"`
		Username string `hcl:"username,optional"`
		Password string `hcl:"password,optional"`
	}
	Database struct {
		Type     string `hcl:"type,optional"`
		Host     string `hcl:"host,optional"`
		Port     int    `hcl:"port,optional"`
		Username string `hcl:"username,optional"`
		Password string `hcl:"password,optional"`
		Migrate  bool   `hcl:"migrate,optional"`
	}
)

func init() {
	log.SetGlobalLevel(zerolog.ErrorLevel)
}

// New creates a new configuration
func New(options ...func(*Config) error) (*Config, error) {
	s := &Config{
		Banner:  true,
		Debug:   false,
		Version: internal.VERSION,
	}

	// Default values
	DefaultOptions()(s)

	// Apply config options
	for _, opt := range options {
		err := opt(s)
		if err != nil {
			return nil, err
		}
	}

	if s.Debug {
		log.SetGlobalLevel(zerolog.DebugLevel)
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
			var program models.Program
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

func (c *Config) LoadDatabase(db Database) (err error) {
	log.Info().Msg("loading programs from directory")

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("error getting user config directory: %w", err)
	}

	// Use SQLite database
	if db.Type == "sqlite" {
		dbFile := filepath.Join(userConfigDir, CONFIG_DIR, DATABASE_FILE)
		log.Info().Str("file", dbFile).Msg("using existing SQLite database")
		models.DB, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
		if err != nil {
			return err
		}
	}
	// Use in-memory SQLite database

	// Auto migrate the schema
	if db.Migrate {
		models.AutoMigrate(models.DB)
		SeedPrograms(models.DB)
	}

	if c.Debug {
		models.DB.Logger = models.DB.Logger.LogMode(logger.Info)
	}

	return
}

func CliOptions(ctx context.Context, cmd *cli.Command) func(c *Config) error {

	return func(c *Config) error {
		c.Debug = cmd.Bool("debug")
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

		if err := c.LoadDatabase(c.Database); err != nil {
			return fmt.Errorf("error parsing database file: %w", err)
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
