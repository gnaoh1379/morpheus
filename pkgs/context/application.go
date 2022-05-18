package context

import (
	"errors"
	"github.com/spf13/viper"
	"sync"
)

type Application struct {
	mu        sync.RWMutex
	plugins   []Plugin
	pluginMap map[any]int
	cfg       *viper.Viper
}

func New(plugins []Plugin) (*Application, error) {
	var a = &Application{
		plugins:   make([]Plugin, len(plugins)),
		pluginMap: make(map[any]int),
		cfg:       viper.New(),
	}
	for index, plugin := range plugins {
		if _, existed := a.pluginMap[plugin.Signature()]; existed {
			return nil, errors.New("plugin signature duplicated")
		}
		a.plugins[index] = plugin
		a.pluginMap[plugin.Signature()] = index
	}
	return a, nil
}

func (a *Application) Run() {

}
