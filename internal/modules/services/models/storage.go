package models

import (
	"github.com/rangertaha/hxe/internal/rdb"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) List(req *Service) (res []*Service, err StatusCode) {
	services := []rdb.Service{}
	if err := s.db.Find(&services).Error; err != nil {
		return nil, StatusCode_INTERNAL_SERVER_ERROR
	}
	if len(services) == 0 {
		return nil, StatusCode_NOT_FOUND
	}
	return ToProtoServices(services), StatusCode_OK
}

func (s *Storage) Get(req *Service) (res *Service, err StatusCode) {
	service := rdb.Service{}
	if err := s.db.First(&service, uint(req.Id)).Error; err != nil {
		return nil, StatusCode_INTERNAL_SERVER_ERROR
	}
	if service.ID == 0 {
		return nil, StatusCode_NOT_FOUND
	}

	return ToProtoService(service), StatusCode_OK
}

func (s *Storage) Create(req *Service) (res *Service, err StatusCode) {
	service := FromProtoService(req)
	if err := s.db.Create(&service).Error; err != nil {
		return nil, StatusCode_INTERNAL_SERVER_ERROR
	}

	return ToProtoService(service), StatusCode_OK
}

func (s *Storage) Update(req *Service) (res *Service, err StatusCode) {
	service := FromProtoService(req)
	if err := s.db.Model(&service).Updates(service).Error; err != nil {
		return nil, StatusCode_INTERNAL_SERVER_ERROR
	}
	return ToProtoService(service), StatusCode_OK
}

func (s *Storage) Delete(req *Service) (res *Service, err StatusCode) {
	service := FromProtoService(req)
	if err := s.db.Delete(&service).Error; err != nil {
		return nil, StatusCode_INTERNAL_SERVER_ERROR
	}
	return ToProtoService(service), StatusCode_OK
}
