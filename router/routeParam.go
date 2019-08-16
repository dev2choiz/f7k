package router

import "github.com/dev2choiz/f7k/interfaces"

type RouteParam struct {
	value        interface{}
	name         string
	order        uint8

	requirements map[string]interfaces.ParamRequirement
}

func (r *RouteParam) Value() interface{} {
	return r.value
}

func (r *RouteParam) SetValue(v interface{}) {
	r.value = v
}

func (r *RouteParam) Name() string {
	return r.name
}

func (r *RouteParam) SetName(v string) {
	r.name = v
}

func (r *RouteParam) Order() uint8 {
	return r.order
}

func (r *RouteParam) SetOrder(v uint8) {
	r.order = v
}

func (r *RouteParam) AddRequirement(p interfaces.ParamRequirement) {
	if nil == r.requirements {
		r.requirements = make(map[string]interfaces.ParamRequirement)
	}
	r.requirements[p.GetName()] = p
}

func (r *RouteParam) GetRequirement(n string) interfaces.ParamRequirement {
	if nil == r.requirements {
		r.requirements = map[string]interfaces.ParamRequirement{}
		return nil
	}

	p, ok := r.requirements[n]
	if !ok {
		return nil
	}

	return p
}

func (r *RouteParam) Requirements() map[string]interfaces.ParamRequirement {
	if nil == r.requirements {
		r.requirements = map[string]interfaces.ParamRequirement{}
		return nil
	}

	return r.requirements
}

func (r *RouteParam) SetRequirements(reqs map[string]interfaces.ParamRequirement)  {
	r.requirements = reqs
}
