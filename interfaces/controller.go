package interfaces

import (
	"net/http"
)

type ControllerInterface interface {
	GetRequest() *http.Request
	SetRequest(*http.Request)
	GetResponse() ResponseInterface
	SetResponse(ResponseInterface)
	GetResponseWriter() *http.ResponseWriter
	SetResponseWriter(*http.ResponseWriter)
	GetUrlParams() map[string]RouteParam
	SetUrlParams(map[string]RouteParam)
	GetAppConfig () ConfigInterface
	SetAppConfig (ConfigInterface)
	Render(interface{}, ...string) string
}
