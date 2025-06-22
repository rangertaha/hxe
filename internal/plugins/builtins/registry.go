package builtins

import (
	"fmt"

	"github.com/rangertaha/hxe/internal/interfaces"
)

type Creator func() interfaces.Plugin

var Plugins = map[string]Creator{}

func Add(name string, creator Creator) {
	Plugins[name] = creator
}

func Get(name string) (Creator, error) {
	if plugin, ok := Plugins[name]; ok {
		return plugin, nil
	}

	return nil, fmt.Errorf("unable to locate plugin/%s", name)
}

func List() []string {
	names := make([]string, 0, len(Plugins))
	for name := range Plugins {
		names = append(names, name)
	}

	return names
}
