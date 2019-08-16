package router

import (
	"errors"
	"github.com/dev2choiz/f7k/interfaces"
	"regexp"
	"strings"
)

type routerMetadata struct {
	Routes map[string]*Route `yaml:"routes"`
}

type router struct {
	*routerMetadata
	currentRoute interfaces.Route
}

var routerInstance *router

func Instance() *router {
	if nil == routerInstance {
		routes := make(map[string]*Route)
		i := &router {&routerMetadata{routes}, nil}
		routerInstance = i
	}

	return routerInstance
}

func (r *router) Init() interfaces.Router {
	/*
	Routes can be declared with a yaml file or with annotations.
	As like it is, a route declared by annotation will override
	another in the yaml file if they have the same name.
	 */
	r.PopulateWithControllersAnnotations()

	for _, route := range r.routerMetadata.Routes {
		route.SetParams(make(map[string]interfaces.RouteParam))
		r.ParseParameters(route)
		r.ParseRequirements(route)
	}

	return r
}

func (r *router) IsCurrentRoute(uri string, route interfaces.Route) bool {
	uri = strings.Split(uri, "?")[0]

	var re = regexp.MustCompile(route.CheckPattern())
	m := re.FindAllStringSubmatch(uri, -1)
	if 0 == len(m) {
		return false
	}

	for _, p := range route.Params() {
		p.SetValue(m[0][p.Order()])
	}

	return true
}

func (r *router) SearchRoute(uri string) (interfaces.Route, error) {
	for name := range r.routerMetadata.Routes {
		route := r.routerMetadata.Routes[name]
		route.SetName(name)
		if ! r.IsCurrentRoute(uri, route) {
			continue
		}
		return route, nil
	}

	return nil, errors.New(uri + " does not match")
}

func (r *router) RouteNames() []string {
	var l []string
	for name := range r.routerMetadata.Routes {
		l = append(l, name)
	}
	return l
}

func (r *router) GetRoute(name string) interfaces.Route {
	return r.routerMetadata.Routes[name]
}

func (r *router) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if nil == r.routerMetadata {
		r.routerMetadata = &routerMetadata{}
	}
	err := unmarshal(r.routerMetadata)
	if err != nil {
		return err
	}
	return nil
}

func (r *router) CurrentRoute() interfaces.Route {
	return r.currentRoute
}

func (r *router) SetCurrentRoute(route interfaces.Route) {
	r.currentRoute = route
}

