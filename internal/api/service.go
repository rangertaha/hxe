package services

import (
	"fmt"

	"github.com/nats-io/nats.go/micro"
)

type Service struct {
	micro.Group
}

func NewService(group micro.Group) *Service {
	return &Service{group}
}

func (e *Service) Init() (err error) {
	e.AddEndpoint("load", micro.HandlerFunc(e.Load))
	e.AddEndpoint("list", micro.HandlerFunc(e.List))
	e.AddEndpoint("list", micro.HandlerFunc(e.Get))
	e.AddEndpoint("create", micro.HandlerFunc(e.Create))
	e.AddEndpoint("update", micro.HandlerFunc(e.Update))
	e.AddEndpoint("delete", micro.HandlerFunc(e.Delete))
	e.AddEndpoint("start", micro.HandlerFunc(e.Start))
	e.AddEndpoint("stop", micro.HandlerFunc(e.Stop))
	e.AddEndpoint("restart", micro.HandlerFunc(e.Restart))
	e.AddEndpoint("status", micro.HandlerFunc(e.Status))
	e.AddEndpoint("log", micro.HandlerFunc(e.Log))
	e.AddEndpoint("shell", micro.HandlerFunc(e.Shell))

	return
}

// Runtime CRUD operations in a database

// Load a service
func (s *Service) Load(msg micro.Request) {
	fmt.Println("load:service")
}

// List all services
func (s *Service) List(msg micro.Request) {
	fmt.Println("list:services")
}

// Get a service by ID
func (s *Service) Get(msg micro.Request) {
	fmt.Println("get:service")
}

// Create a new service
func (s *Service) Create(msg micro.Request) {
	fmt.Println("create:service")
}

// Update a service
func (s *Service) Update(msg micro.Request) {
	fmt.Println("update:service")
}

// Delete a service
func (s *Service) Delete(msg micro.Request) {
	fmt.Println("delete:service")
}

// Runtime operations sending commands to a service executor plugins

// Start a service
func (s *Service) Start(msg micro.Request) {
	fmt.Println("start:service")
}

// Stop a service
func (s *Service) Stop(msg micro.Request) {
	fmt.Println("stop:service")
}

// Restart a service
func (s *Service) Restart(msg micro.Request) {
	fmt.Println("restart:service")
}

// Status of a service
func (s *Service) Status(msg micro.Request) {
	fmt.Println("status:service")
}

// Log a service
func (s *Service) Log(msg micro.Request) {
	fmt.Println("log:service")
}

func (s *Service) Shell(msg micro.Request) {
	fmt.Println("shell:service")
}
