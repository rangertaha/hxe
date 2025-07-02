package models

import "github.com/rangertaha/hxe/internal/rdb"

func (s *Service) Create() (err error) {
	service := FromProtoService(s)
	if err := rdb.DB.Create(&service).Error; err != nil {
		return err
	}
	return
}

func (s *Service) Update() (err error) {
	service := FromProtoService(s)
	if err := rdb.DB.Save(&service).Error; err != nil {
		return err
	}
	return
}

func (s *Service) Save() (err error) {
	service := FromProtoService(s)
	if err := rdb.DB.Save(&service).Error; err != nil {
		return err
	}
	return
}

func (s *Service) Delete() (err error) {
	service := FromProtoService(s)
	if err := rdb.DB.Delete(&service).Error; err != nil {
		return err
	}
	return
}

func (s *Response) Count() (count  int) {
	count = len(s.Services)
	if s.Service != nil {
		count++
	}
	return
}

