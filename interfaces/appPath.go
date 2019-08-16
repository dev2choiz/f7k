package interfaces

type AppPath interface {
	F7kGitImportPath() 			string
	SetF7kGitImportPath(string)
	F7kPath()          			string
	SetF7kPath(string)
	AppRoot()                	string
	SetAppRoot(string)
	GopathSrc()              	string
	SetGopathSrc(string)
	Gopath()                 	string
	SetGopath(string)

}
