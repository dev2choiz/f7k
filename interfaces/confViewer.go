package interfaces

type ConfViewer interface {
	ViewDir() string
	SetViewDir(string)
	DefaultLayout() string
	SetDefaultLayout(string)
	DefaultLayoutFile() string
	SetDefaultLayoutFile(string)
	FilesToParse() *[]string
	SetFilesToParse(*[]string)
}
