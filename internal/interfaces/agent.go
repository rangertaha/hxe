package interfaces

type Agent interface {
	Init() error
	Stop() error
	Start() error
	Reload() error
}
