package installer

import (
	"github.com/dev2choiz/f7k/fileFiller"
	"path"
)

var placeholders map[string]string

func (i *installer) populatePlaceholders() *installer {
	placeholders = make(map[string]string)
	placeholders["{{APP_ABS_PATH}}"] = path.Join(i.CurrentDir, i.ProjectName)
	placeholders["{{APP_IMPORT_PATH}}"] = i.AppImportPath
	placeholders["{{F7K_IMPORT_PATH}}"] = i.F7kImportPath

	return i
}

func (i *installer) replacePlaceholders() *installer {
	ff := fileFiller.New()
	ff.Placeholders = placeholders
	ff.FillFiles(i.getAbsProjectDir())

	return i
}
