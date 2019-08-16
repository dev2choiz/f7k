package router

import (
	"github.com/dev2choiz/f7k/interfaces"
)

type routeMetadata struct {
	Path        string `yaml:"path"`
	Ctrl        string `yaml:"controller"`
	PackageName string `yaml:"package"`
	Action      string `yaml:"action,omitempty"`
	Methods     []string `yaml:"methods,omitempty"`
}

type Route struct {
	*routeMetadata

	name         string
	checkPattern string
	params       map[string]interfaces.RouteParam
}

func NewRoute() *Route {
	return &Route{ routeMetadata: &routeMetadata{} }
}

func (r *Route) Name() string {
	return r.name
}
func (r *Route) SetName(v string) interfaces.Route {
	r.name = v

	return r
}
func (r *Route) Path() string {
	return r.routeMetadata.Path
}
func (r *Route) SetPath(v string) interfaces.Route {
	r.routeMetadata.Path = v

	return r
}
func (r *Route) Controller() string {
	return r.routeMetadata.Ctrl
}
func (r *Route) SetController(v string) interfaces.Route {
	r.routeMetadata.Ctrl = v

	return r
}
func (r *Route) Action() string {
	return r.routeMetadata.Action
}
func (r *Route) SetAction(v string) interfaces.Route {
	r.routeMetadata.Action = v

	return r
}
func (r *Route) Package() string {
	return r.routeMetadata.PackageName
}
func (r *Route) SetPackage(v string) interfaces.Route {
	r.routeMetadata.PackageName = v

	return r
}
func (r *Route) Methods() []string {
	return r.routeMetadata.Methods
}
func (r *Route) SetMethods(v []string) interfaces.Route {
	r.routeMetadata.Methods = v

	return r
}
func (r *Route) CheckPattern() string {
	return r.checkPattern
}
func (r *Route) SetCheckPattern(v string)interfaces.Route {
	r.checkPattern = v

	return r
}

func (r *Route) AddParam(p interfaces.RouteParam) {
	if nil == r.params {
		r.params = make(map[string]interfaces.RouteParam)
	}
	r.params[p.Name()] = p
}

func (r *Route) Param(n string) interfaces.RouteParam {
	if nil == r.params {
		r.params = map[string]interfaces.RouteParam{}
		return nil
	}

	p, ok := r.params[n]
	if !ok {
		return nil
	}

	return p
}

func (r *Route) Params() map[string]interfaces.RouteParam {
	if nil == r.params {
		r.params = map[string]interfaces.RouteParam{}
		return nil
	}

	return r.params
}

func (r *Route) SetParams(p map[string]interfaces.RouteParam) interfaces.Route {
	r.params = p

	return r
}

func (r *Route) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if nil == r.routeMetadata {
		r.routeMetadata = &routeMetadata{}
	}
	err := unmarshal(r.routeMetadata)
	if err != nil {
		return err
	}
	return nil
}
