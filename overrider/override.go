package overrider

type Override struct {
	ImportPath   string `yaml:"importPath"`
	PackageName  string `yaml:"packageName"`
	FunctionName string `yaml:"functionName"`
}

func (c *Override) GetTargetImportPath() string {
	return c.ImportPath
}

func (c *Override) GetTargetPackageName() string {
	return c.PackageName
}

func (c *Override) GetFunctionName() string {
	return c.FunctionName
}
