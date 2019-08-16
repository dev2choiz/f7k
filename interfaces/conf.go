package interfaces

type ConfigInterface interface {
	GetPort() int
	GetImportPath() string
	GetControllerDir() string
	PostConfig()

	// return struct used to unmarshal the yaml config
	Data() ConfigInterface
}
