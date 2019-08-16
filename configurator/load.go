package configurator

import (
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/model/events"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Load(f string, config interfaces.ConfigInterface) interfaces.ConfigInterface {
	loadConfig(f, config)

	e := &events.ConfigEvent{}
	e.SetAppConfig(config)
	f7k.Dispatcher.Dispatch(events.OnConfigEvent, e)

	return config
}

func loadConfig(f string, config interfaces.ConfigInterface) {
	p, err := os.Getwd()
	s, err := ioutil.ReadFile(filepath.Join(p, f))
	if nil != err {
		panic(err)
	}

	m := config.Data()
	err = yaml.Unmarshal(s, m)
	if nil != err {
		panic(err)
	}
}
