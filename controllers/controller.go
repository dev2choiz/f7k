package controllers

import (
	"bytes"
	"github.com/dev2choiz/f7k"
	"github.com/dev2choiz/f7k/interfaces"
	"github.com/dev2choiz/f7k/model"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

type Controller struct {
	Request        *http.Request
	ResponseWriter *http.ResponseWriter
	Response       interfaces.ResponseInterface
	UrlParams      map[string]interfaces.RouteParam
	appConfig      interfaces.ConfigInterface

	View interfaces.ConfViewer
}

func (c *Controller) GetRequest() *http.Request {
	return c.Request
}

func (c *Controller) SetRequest(r *http.Request) {
	c.Request = r
}

func (c *Controller) GetResponse() interfaces.ResponseInterface {
	return c.Response
}

func (c *Controller) SetResponse(r interfaces.ResponseInterface) {
	c.Response = r
}

func (c *Controller) GetResponseWriter() *http.ResponseWriter {
	return c.ResponseWriter
}

func (c *Controller) SetResponseWriter(r *http.ResponseWriter) {
	c.ResponseWriter = r
}


func (k *Controller) GetAppConfig () interfaces.ConfigInterface {
	return k.appConfig
}
func (k *Controller) SetAppConfig (c interfaces.ConfigInterface) {
	k.appConfig = c
}

func (c *Controller) NewJsonResponse() *model.JsonResponse {
	r := &model.JsonResponse{}
	r.Data = make(map[string]interface{})
	r.Success = true
	r.SetStatus(http.StatusOK)

	return r
}

func (c *Controller) NewHtmlResponse() *model.HtmlResponse {
	r := &model.HtmlResponse{}
	r.SetStatus(http.StatusOK)

	return r
}

func (c *Controller) GetUrlParams() map[string]interfaces.RouteParam {
	return c.UrlParams
}

func (c *Controller) SetUrlParams(u map[string]interfaces.RouteParam) {
	c.UrlParams = u
}

func (c *Controller) Render(data interface{}, tmpls ...string) string {
	var toParse []string
	viewConf := f7k.ViewConfig
	currWD, _ := os.Getwd()
	defaultLayout := filepath.Join(currWD, viewConf.ViewDir(), viewConf.DefaultLayoutFile())
	toParse = append(toParse, defaultLayout)
	toParse = append(toParse, c.getFilesToParse(tmpls)...)
	toParse = append(toParse, c.getFilesToParse(*viewConf.FilesToParse())...)

	tmpl := template.Must(template.ParseFiles(toParse...))
	buf := bytes.NewBuffer(nil)

	tmpl.ExecuteTemplate(buf, viewConf.DefaultLayout(), data)

	return buf.String()
}

func (c *Controller) getFilesToParse(tmpls []string) []string {
	var toParse []string
	viewConf := f7k.ViewConfig
	currWD, _ := os.Getwd()

	for _, viewPath := range tmpls {
		path := filepath.Join("/", currWD, viewConf.ViewDir())
		path = filepath.Join(path, viewPath)

		toParse = append(toParse, path)
	}
	return toParse
}
