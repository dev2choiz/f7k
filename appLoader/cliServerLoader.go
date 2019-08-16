package appLoader

import (
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/appPath"
	"github.com/dev2choiz/f7k/cacheGen"
	"github.com/dev2choiz/f7k/command"
	"github.com/dev2choiz/f7k/configurator"
	"github.com/dev2choiz/f7k/controllers"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/overrider"
	"github.com/dev2choiz/f7k/router"
	"github.com/dev2choiz/f7k/viewer"
)

type CliServerLoader struct {
	AppLoader
}

func DefaultCliServerLoader(conf interfaces.ConfigInterface) *CliServerLoader {
	return &CliServerLoader{*DefaultAppLoader(conf)}
}

func GetCliServerLoader(conf interfaces.ConfigInterface, confFile, confViewFile string, loadExtListeners, loadCacheGenerators bool) *CliServerLoader {
	return &CliServerLoader{*NewAppLoader(conf, confFile, confViewFile, loadExtListeners, loadCacheGenerators)}
}

func (i *CliServerLoader) LoadApp() {
	i.Load()
	f7k.AppPath = appPath.Instance()
	if "" != i.confFile {
		f7k.AppConfig = configurator.Load(i.confFile, i.config)
	}
	if "" != i.viewConfFile {
		f7k.ViewConfig = viewer.Load(i.viewConfFile)
	}
	f7k.Router = overrider.LoadRouterInstance()
	f7k.Kernel = overrider.LoadKernelInstance()
	if i.LoadCacheGenerators {
		cacheGen.LoadCacheGenerators()
	}
	if i.LoadExtListeners && nil != f7k.AppConfig {
		f7k.AppConfig.PostConfig()
	}
}

func (i *CliServerLoader) Load() interfaces.AppLoader {
	// f7k cacheGen listeners
	cacheGen.WaitingForListen = append(cacheGen.WaitingForListen, command.OnCacheGenCommand)
	cacheGen.WaitingForListen = append(cacheGen.WaitingForListen, overrider.OnCacheGenOverloader)
	cacheGen.WaitingForListen = append(cacheGen.WaitingForListen, controllers.OnCacheGenAnnotations)
	cacheGen.WaitingForListen = append(cacheGen.WaitingForListen, controllers.OnCacheGenInstantiate)
	cacheGen.WaitingForListen = append(cacheGen.WaitingForListen, router.OnCacheGenYaml)

	i.AppLoader.Load()
	f7k.Dispatcher = overrider.LoadEventDispatcherInstance()

	return i
}
