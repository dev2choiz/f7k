package interfaces

type RouteParam interface {
	Value() interface{}
	SetValue(interface{})
	Name() string
	SetName(v string)
	Order() uint8
	SetOrder(v uint8)
	AddRequirement(ParamRequirement)
	GetRequirement(n string) ParamRequirement
	Requirements() map[string]ParamRequirement
	SetRequirements(map[string]ParamRequirement)
}
