package appLoader

import (
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/fileFiller"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/internal/createProject/installer"
	"github.com/dev2choiz/f7k/internal/grpcBuilder"
	"github.com/dev2choiz/f7k/internal/sampler/sampler"
)

type AppLoader struct {
	config                 interfaces.ConfigInterface
	confFile, viewConfFile string
	LoadExtListeners       bool
	LoadCacheGenerators    bool
}

const DefaultConfFile string = "conf/conf.yaml"
const DefaultConfViewFile string = "conf/view.yaml"

var instance *AppLoader

func Instance() interfaces.AppLoader {
	return instance
}

func DefaultAppLoader(conf interfaces.ConfigInterface) *AppLoader {
	if nil == instance {
		instance = &AppLoader{
			conf,
			DefaultConfFile,
			DefaultConfViewFile,
			true,
			true,
		}
		f7k.AppLoader = instance
	}

	return instance
}

func NewAppLoader(conf interfaces.ConfigInterface, confFile, confViewFile string, loadExtListeners, loadCacheGenerators bool) *AppLoader {
	if nil == instance {
		instance = &AppLoader{
			conf,
			confFile,
			confViewFile,
			loadExtListeners,
			loadCacheGenerators,
		}
	}

	return instance
}

func (l *AppLoader) Load() interfaces.AppLoader {
	f7k.Version = "0.0.1"

	return l
}

func (l *AppLoader) ConfFile() string {
	return l.confFile
}

func (l *AppLoader) ViewConfFile() string {
	return l.viewConfFile
}

func (l *AppLoader) PostAppLoad() interfaces.AppLoader {
	return l
}


// Reference for ensure import if it is not used
var _ = grpcBuilder.New
var _ = fileFiller.New
var _ = sampler.New
var _ = installer.New
