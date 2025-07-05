package models

// import "github.com/rangertaha/hxe/internal/rdb"

// func (s *Service) Create() (res *Service, err StatusCode) {
// 	dbModel := s.DbModel()

// 	// Create the service in the database
// 	results := rdb.DB.Create(&dbModel)
// 	if results.Error != nil {
// 		return res, StatusCode_INTERNAL_SERVER_ERROR
// 	}

// 	// Convert the service to a protobuf service
// 	res = s.PbModel(dbModel)

// 	// Return the service
// 	return res, StatusCode_OK
// }

// func (s *Service) Update() (err error) {
// 	service := s.DbModel()
// 	if err := rdb.DB.Save(&service).Error; err != nil {
// 		return err
// 	}
// 	return
// }

// func (s *Service) Save() (err error) {
// 	service := s.DbModel()
// 	if err := rdb.DB.Save(&service).Error; err != nil {
// 		return err
// 	}
// 	return
// }

// func (s *Service) Delete() (err error) {
// 	service := s.DbModel()
// 	if err := rdb.DB.Delete(&service).Error; err != nil {
// 		return err
// 	}
// 	return
// }

// func (s *Service) DbModel() (model *rdb.Service) {
// 	return PbToDbService(s)
// }

// func (s *Service) PbModel(models ...*rdb.Service) *Service {
// 	for _, m := range models {
// 		s = DbToPbService(m)
// 	}
// 	return s
// }


// func (s *Response) Count() (count int) {
// 	count = len(s.Services)
// 	if s.Service != nil {
// 		count++
// 	}
// 	return
// }


// func PbToDbService(protoService *Service) *rdb.Service {
// 	service := rdb.Service{ID: uint(protoService.Id)}
// 	service.Name = protoService.Name
// 	service.Description = protoService.Description
// 	service.Status = rdb.ServiceStatus(protoService.Status)
// 	service.PreExec = protoService.PreExec
// 	service.CmdExec = protoService.CmdExec
// 	service.PostExec = protoService.PostExec
// 	service.Autostart = protoService.Autostart
// 	service.Retries = int(protoService.Retries)
// 	service.PID = int(protoService.Pid)
// 	service.Exit = int(protoService.Exit)
// 	service.Started = protoService.Started
// 	service.Ended = protoService.Ended
// 	return &service
// }

// func DbToPbService(service *rdb.Service) (s *Service) {
// 	s = &Service{Id: uint32(service.ID)}
// 	s.Name = service.Name
// 	s.Description = service.Description
// 	s.Status = ServiceStatus(service.Status)
// 	s.PreExec = service.PreExec
// 	s.CmdExec = service.CmdExec
// 	s.PostExec = service.PostExec
// 	s.Autostart = service.Autostart
// 	s.Retries = int32(service.Retries)
// 	s.Pid = int32(service.PID)
// 	s.Exit = int32(service.Exit)
// 	s.Started = service.Started
// 	s.Ended = service.Ended
// 	return s
// }


// func DbToPbServiceList(services []*rdb.Service) []*Service {
// 	serviceList := []*Service{}
// 	for _, service := range services {
// 		serviceList = append(serviceList, DbToPbService(service))
// 	}
// 	return serviceList
// }