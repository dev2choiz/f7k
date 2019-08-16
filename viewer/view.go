package viewer

type ConfViewYaml struct {
	ViewDir 		  string              `yaml:"viewDirectory"`
	DefaultLayout     string              `yaml:"defaultLayout"`
	DefaultLayoutFile string              `yaml:"defaultLayoutFile"`
	FilesToParse      *[]string  `yaml:"filesToParse"`
}

type ConfView struct {
	*ConfViewYaml
}

func (v *ConfView) ViewDir() string {
	return v.ConfViewYaml.ViewDir
}
func (v *ConfView) SetViewDir(s string) {
	v.ConfViewYaml.ViewDir = s
}

func (v *ConfView) DefaultLayout() string {
	return v.ConfViewYaml.DefaultLayout
}
func (v *ConfView) SetDefaultLayout(s string) {
	v.ConfViewYaml.DefaultLayout = s
}

func (v *ConfView) DefaultLayoutFile() string {
	return v.ConfViewYaml.DefaultLayoutFile
}
func (v *ConfView) SetDefaultLayoutFile(s string) {
	v.ConfViewYaml.DefaultLayoutFile = s
}

func (v *ConfView) FilesToParse() *[]string {
	return v.ConfViewYaml.FilesToParse
}
func (v *ConfView) SetFilesToParse(s *[]string) {
	v.ConfViewYaml.FilesToParse = s
}

func (c *ConfView) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if nil == c.ConfViewYaml {
		c.ConfViewYaml = &ConfViewYaml{}
	}

	err := unmarshal(c.ConfViewYaml)
	if err != nil {
		return err
	}

	return nil
}


