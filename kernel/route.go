package kernel

import (
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/model"
	"github.com/dev2choiz/f7k/model/events"
	"github.com/dev2choiz/f7k/router"
	"net/http"
)

func (k *Kernel) RetrieveRoute() (interfaces.Route, interfaces.KernelEvent, *error) {
	f7k.Router.SetCurrentRoute(nil)
	dispatcher := f7k.Dispatcher
	url := Request.URL.Path
	e := &events.RequestEvent{}
	e.ResponseWriter = ResponseWrite
	e.Request = Request

	route, err := router.Instance().SearchRoute(url)
	if nil != err {
		dispatcher.Dispatch(events.OnRouteNotFoundEvent, e)
		if nil != e.Response {
			return nil, e, nil
		} else {
			r := &model.JsonResponse{}
			r.SetStatus(http.StatusNotFound)
			r.Success = false
			r.Message = url + " not found."
			e.Response = r
			return nil, e, nil
		}
	}

	f7k.Router.SetCurrentRoute(route)
	dispatcher.Dispatch(events.OnRequestEvent, e)

	return route, e, nil
}
