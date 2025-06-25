package services

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"github.com/rangertaha/hxe/internal/api/handlers"
)

type API struct {
	srv micro.Service
	svc *handlers.Service
	tsk *handlers.Task
}

func NewAPI(nc *nats.Conn) (e *API, err error) {
	srv, err := micro.AddService(nc, micro.Config{
		Name:        "engine",
		Version:     "0.0.1",
		Description: "Hxe engine exposing API and managing plugins",
	})
	if err != nil {
		return nil, err
	}
	e = &API{
		srv: srv,
		svc: handlers.NewService(srv.AddGroup("svc")),
		tsk: handlers.NewTask(srv.AddGroup("tsk")),
	}
	err = e.Init()
	return
}

func (e *API) Init() (err error) {
	// Initialize the service
	if err = e.svc.Init(); err != nil {
		return err
	}
	if err = e.tsk.Init(); err != nil {
		return err
	}
	return
}

func (e *API) Start() (err error) {
	return
}

func (e *API) Stop() (err error) {
	return
}
