package utils

var instance *Utils

type Utils struct {
	Slice *slice
}

func Instance() *Utils {
	if nil == instance {
		instance = &Utils{
			&slice{},
		}
	}

	return instance
}
