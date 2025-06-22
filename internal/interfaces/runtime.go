package interfaces

import "time"

type Status interface {
	State() int32
	Uptime() time.Duration
	Message() string
	Progress() int32
}

type Runner interface {
	Id() string
	Init() error
	Start() error
	Stop() error
	Fill() error
	Test() error
	Train() error
	Status() Status
}

type Plugin interface {
	Init() error
	Stdin() error
	Stdcmd() error
	Stdout() error
	Stderr() error
}
