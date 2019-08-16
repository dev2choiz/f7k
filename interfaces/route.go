package interfaces

type Route interface {
	Name() string
	SetName(string) Route
	Path() string
	SetPath(string) Route
	Controller() string
	SetController(string) Route
	Package() string
	SetPackage(string) Route
	Action() string
	SetAction(string) Route
	Methods() []string
	SetMethods([]string) Route
	CheckPattern() string
	SetCheckPattern(string) Route
	AddParam(p RouteParam)
	Param(n string) RouteParam
	Params() map[string]RouteParam
	SetParams(map[string]RouteParam) Route
}
