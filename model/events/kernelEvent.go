package events

import (
	"github.com/dev2choiz/f7k/interfaces"
	"net/http"
)

const OnPostLoadEvent = "kernel.on_post_load_event"
const OnRequestEvent = "kernel.on_request_event"
const OnControllerEvent = "kernel.on_controller_event"
const OnResponseEvent = "kernel.on_response_event"
const OnRouteNotFoundEvent = "kernel.on_route_not_found_event"
const OnConfigEvent = "kernel.on_config_event"

type KernelEvent struct {
	Event
	Request        *http.Request
	ResponseWriter *http.ResponseWriter
	Controller     interfaces.ControllerInterface
	Response       interfaces.ResponseInterface
	appConfig      interfaces.ConfigInterface
}

func (k *KernelEvent) GetEventName() string {
	return "KernelRequestEvent"
}

func (k *KernelEvent) GetRequest () *http.Request {
	return k.Request
}
func (k *KernelEvent) SetRequest (r *http.Request) {
	k.Request = r
}
func (k *KernelEvent) GetResponseWriter () *http.ResponseWriter {
	return k.ResponseWriter
}
func (k *KernelEvent) SetResponseWriter (r *http.ResponseWriter) {
	k.ResponseWriter = r
}
func (k *KernelEvent) GetController () interfaces.ControllerInterface {
	return k.Controller
}
func (k *KernelEvent) SetController (c interfaces.ControllerInterface) {
	k.Controller = c
}
func (k *KernelEvent) GetResponse () interfaces.ResponseInterface {
	return k.Response
}
func (k *KernelEvent) SetResponse (r interfaces.ResponseInterface) {
	k.Response = r
}
func (k *KernelEvent) AppConfig () interfaces.ConfigInterface {
	return k.appConfig
}
func (k *KernelEvent) SetAppConfig (c interfaces.ConfigInterface) {
	k.appConfig = c
}
