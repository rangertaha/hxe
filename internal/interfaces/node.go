package interfaces

import "time"


type Noder interface {
	Fill(Filler) error
	Train(Trainer) error
	Test(Tester) error
	Exec(Executer) error
}

type Filler interface {
	Init() error
	Exec(time.Time, time.Time) error
}

type Trainer interface {
	Init() error
	Exec(time.Time, time.Time) error
}

type Tester interface {
	Init() error
	Exec(time.Time, time.Time) error
}

type Executer interface {
	Init() error
	Exec(time.Time, time.Time) error
}
