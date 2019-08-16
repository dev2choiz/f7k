package appLoader

import (
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/model/events"
)

type HttpServerLoader struct {
	CliServerLoader
}

var httpServerLoaderInstance *HttpServerLoader

func GetHttpServerLoader(conf interfaces.ConfigInterface) *HttpServerLoader {
	if nil == httpServerLoaderInstance {
		httpServerLoaderInstance = &HttpServerLoader{*DefaultCliServerLoader(conf)}
	}
	return httpServerLoaderInstance
}

func (l *HttpServerLoader) Load() interfaces.AppLoader {
	l.CliServerLoader.LoadApp()
	f7k.Router.Init()

	return l
}

func (l *HttpServerLoader) PostAppLoad() interfaces.AppLoader {
	// internal listeners of PostAppLoad event
	// ...
	// dispatch of PostAppLoad event
	e := &events.Event{}
	f7k.Dispatcher.Dispatch(events.OnPostLoadEvent, e)

	return l
}