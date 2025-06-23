package models

type ProgramStatus string

const (
	ProgramStop    ProgramStatus = "stop"
	ProgramStart   ProgramStatus = "start"
	ProgramReload  ProgramStatus = "reload"
	ProgramRestart ProgramStatus = "restart"
)

type Group struct {
	Base
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`

	Programs []Program `json:"programs" gorm:"foreignKey:GID"`
}

type Action struct {
	Name  string     `json:"name"`
	Desc  string     `json:"desc"`
	Props []Property `json:"props"`
}

type Program struct {
	Base
	AID         uint          `json:"aid" gorm:"column:aid"`
	GID         uint          `json:"gid" gorm:"column:gid"`
	SID         uint          `json:"sid" gorm:"column:sid"`
	Name        string        `json:"name" gorm:"column:name"`
	Description string        `json:"desc" gorm:"column:desc"`
	Command     string        `json:"command" gorm:"column:command"`
	Args        string        `json:"args" gorm:"column:args"`
	Directory   string        `json:"cwd"`
	User        string        `json:"user"`
	Group       string        `json:"group"`
	Status      ProgramStatus `json:"status"`
	PID         int           `json:"pid" gorm:"column:pid"`
	ExitCode    int           `json:"exitCode" gorm:"column:exit"`
	StartTime   int64         `json:"startTime" gorm:"column:start_time"`
	EndTime     int64         `json:"endTime" gorm:"column:end_time"`
	Autostart   bool          `json:"autostart"`
	Enabled     bool          `json:"enabled"`
	Retries     int           `json:"retries"`
	MaxRetries  int           `json:"maxRetries"`

	// Metrics map[string]float64 `json:"metrics" gorm:"serialize:json"`
	// Actions []Action           `json:"actions" gorm:"serialize:json"`
}

// func (p *Program) AfterSave(tx *gorm.DB) (err error) {
// 	// Implement your post-save logic here
// 	log.Info().Str("program", p.Name).Str("status", string(p.Status)).Str("command", p.Command).Msgf("Program successfully saved!")
// 	return nil
// }
