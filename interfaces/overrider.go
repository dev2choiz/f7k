package interfaces

type Overrider interface {
	GetTargetImportPath()	string
	GetTargetPackageName()	string
	GetFunctionName()		string
}
