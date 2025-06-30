package models

import (
	"github.com/rangertaha/hxe/internal/rdb"
)

func ToProtoServices(services []rdb.Service) *Services {
	protoServices := &Services{}
	for _, service := range services {
		protoServices.Services = append(protoServices.Services, ToProtoService(service))
	}
	return protoServices
}

func ToProtoService(service rdb.Service) (s *Service) {
	s = &Service{Id: uint32(service.ID)}
	s.Name = service.Name
	s.Description = service.Description
	s.Status = ServiceStatus(service.Status)
	s.PreExec = service.PreExec
	s.CmdExec = service.CmdExec
	s.PostExec = service.PostExec
	s.Autostart = service.Autostart
	s.Retries = int32(service.Retries)
	s.Pid = int32(service.PID)
	s.Exit = int32(service.ExitCode)
	s.Started = service.StartTime
	s.Ended = service.EndTime
	return s
}

func FromProtoService(protoService *Service) rdb.Service {
	service := rdb.Service{ID: uint(protoService.Id)}
	service.Name = protoService.Name
	service.Description = protoService.Description
	service.Status = rdb.ServiceStatus(protoService.Status)
	service.PreExec = protoService.PreExec
	service.CmdExec = protoService.CmdExec
	service.PostExec = protoService.PostExec
	service.Autostart = protoService.Autostart
	service.Retries = int(protoService.Retries)
	service.PID = int(protoService.Pid)
	service.ExitCode = int(protoService.Exit)
	service.StartTime = protoService.Started
	service.EndTime = protoService.Ended
	return service
}
