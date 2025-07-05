package models

import (
	"fmt"

	"github.com/rangertaha/hxe/internal/log"
	"github.com/rangertaha/hxe/internal/rdb"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Service struct {
	db  *gorm.DB
	log zerolog.Logger
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db, log: log.With().Logger()}
}

func (s *Service) List(req *rdb.Service) (services []*rdb.Service, err error) {
	// Find all services
	services = []*rdb.Service{}
	results := s.db.Find(&services)
	if results.Error != nil {
		return nil, fmt.Errorf("failed:List() to list services: %w", results.Error)
	}

	for _, service := range services {
		s.log.Info().Msgf("Service: %s", service.Name)
	}

	return services, results.Error
}

func (s *Service) Get(req *rdb.Service) (res *rdb.Service, err error) {
	// service := &rdb.Service{}
	if results := s.db.First(res, req.ID); results.Error != nil {
		return nil, results.Error
	}

	return res, nil
}

func (s *Service) Create(service *rdb.Service) (*rdb.Service, error) {
	// Create te service in the database
	results := s.db.Create(&service)
	if results.Error != nil {
		return nil, results.Error
	}

	// Return the service
	return service, nil
}

func (s *Service) Update(service *rdb.Service) (*rdb.Service, error) {
	// Create te service in the database
	results := s.db.Create(&service)
	if results.Error != nil {
		return nil, results.Error
	}

	// Return the service
	return service, nil
}

func (s *Service) Delete(service *rdb.Service) (*rdb.Service, error) {
	if results := s.db.Delete(&service); results.Error != nil {
		return nil, results.Error
	}
	return service, nil
}
