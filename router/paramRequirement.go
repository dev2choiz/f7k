package router

type ParamRequirement struct {
	Name         string
}

func (r *ParamRequirement) GetName() string {
	return r.Name
}
func (r *ParamRequirement) SetName(v string) {
	r.Name = v
}
