package interfaces

type AppLoader interface {
	Load() AppLoader
	PostAppLoad() AppLoader
	ConfFile() string
	ViewConfFile() string
}
