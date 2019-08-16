package interfaces

type Router interface {
	Init() Router
	PopulateWithControllersAnnotations()
	IsCurrentRoute(uri string, route Route) bool
	SetCurrentRoute(route Route)
	CurrentRoute() Route
	RouteNames() []string
	GetRoute(string) Route
	SearchRoute(uri string) (Route, error)
}
