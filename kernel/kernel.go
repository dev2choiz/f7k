package kernel

import (
	"errors"
	"fmt"
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/controllers"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/model"
	"github.com/dev2choiz/f7k/model/events"
	"github.com/dev2choiz/f7k/pkg/prompt"
	"net/http"
	"reflect"
)

var kernelInstance interfaces.Kernel
var ResponseWrite *http.ResponseWriter
var Request *http.Request

type Kernel struct{}

func Instance() interfaces.Kernel {
	if nil == kernelInstance {
		kernelInstance = &Kernel{}
	}
	return kernelInstance
}

func (k *Kernel) Handle(w http.ResponseWriter, r *http.Request) {
	ResponseWrite = &w
	Request = r
	if f7k.Verbose {
		k.logStartRequest(r)
	}

	route, eReq, err := k.RetrieveRoute()
	if nil != err {
		panic(err)
	}

	if nil != eReq.GetResponse() {
		k.Finish(eReq.GetResponse())
		return
	}

	ctrl, eCon, err := k.GetController(route)
	if err != nil {
		panic(err)
	}
	if nil != eCon.GetResponse() {
		k.Finish(eCon.GetResponse())
		return
	}

	eRes, err := k.CallController(ctrl, route)
	if nil != err {
		panic(err)
	}

	k.Finish(eRes.GetResponse())
}

func (k *Kernel) GetController(r interfaces.Route) (interfaces.ControllerInterface, interfaces.KernelEvent, *error) {
	name := r.Package() + "." + r.Controller()
	reg := controllers.RegistryInstance()
	ctrl, ok := reg.Controllers[name]
	if !ok {
		panic(fmt.Errorf("controller '%s' not found", name))
	}
	ctrl.SetRequest(Request)
	ctrl.SetResponseWriter(ResponseWrite)
	ctrl.SetResponse(&model.JsonResponse{})
	ctrl.SetUrlParams(r.Params())
	ctrl.SetAppConfig(f7k.AppConfig)

	e := &events.ControllerEvent{}
	e.ResponseWriter = ResponseWrite
	e.Request = Request
	e.Controller = ctrl
	e.SetAppConfig(f7k.AppConfig)
	f7k.Dispatcher.Dispatch(events.OnControllerEvent, e)

	return ctrl, e, nil
}

func (k *Kernel) CallController(ctrl interfaces.ControllerInterface, r interfaces.Route) (interfaces.KernelEvent, *error) {
	value := reflect.ValueOf(ctrl)
	act := r.Action()
	action := value.MethodByName(act)
	v := action.Call([]reflect.Value{})
	if 0 == len(v) {
		panic(errors.New("controller should return a ResponseInterface"))
	}
	res := v[0].Interface().(interfaces.ResponseInterface)

	e := &events.ResponseEvent{}
	e.ResponseWriter = ResponseWrite
	e.Request = Request
	e.Response = res

	f7k.Dispatcher.Dispatch(events.OnResponseEvent, e)

	return e, nil
}

func (k *Kernel) Finish(r interfaces.ResponseInterface) {
	if nil == r {
		return
	}

	if f7k.Verbose {
		k.logEndRequest(r)
	}
	_ = r.Send(*ResponseWrite)

}

func (k *Kernel) logStartRequest(req *http.Request) {
	prompt.New("info").Printfln(
		"%s %s %s\n%s-->%s  User-Agent(s): %s",
		req.Method,
		req.RequestURI,
		req.Proto,
		req.RemoteAddr,
		req.Host,
		req.Header.Get("User-Agent"),
	)
}

func (k *Kernel) logEndRequest(r interfaces.ResponseInterface) {
	s := r.Status()
	p := "info"
	if 100 <= s && 200 > s  {
		p = "info"
	} else if 200 <= s && 300 > s  {
		p = "info"
	} else if 300 <= s && 400 > s  {
		p = "info"
	} else if 400 <= s && 500 > s  {
		p = "danger"
	} else if 500 <= s && 600 > s  {
		p = "danger"
	}
	prompt.New(p).Printfln("Status %d   %s", int(s), http.StatusText(int(s)))
}
