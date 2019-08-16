package annotation

import (
	"errors"
	"fmt"
	"github.com/dev2choiz/f7k/pkg/goParser"
	"go/ast"
)

type Reader struct {
}

func New() *Reader {
	return &Reader{}
}

func (c *Reader) GetPackageAnnotations(dir string) map[string]map[string]IAnnotation {
	annotations := make(map[string]map[string]IAnnotation, 0)

	gp := goParser.Instance()
	decls, err := gp.ExtractDecls(dir, "*ast.FuncDecl")
	if nil != err {
		panic(err)
	}

	pack, err := gp.GetPackage(dir)
	if nil != err {
		panic(err)
	}

	for _, d := range decls {
		decl := d.(*ast.FuncDecl)
		var st = "func"
		if nil != decl.Recv {
			st = fmt.Sprintf("%v", decl.Recv.List[0].Type)
		}

		someAnnotations, err := c.ExtractAnnotations(decl)
		if nil != err {
			continue
		}

		k := fmt.Sprintf("%s.%s.%s", pack.Name, st, decl.Name.Name)
		if 0 == len(annotations[k]) && 0 < len(someAnnotations) {
			annotations[k] = make(map[string]IAnnotation, 0)
		}
		for _, a := range someAnnotations {
			annotations[k][a.Name()] = a
		}
	}

	return annotations
}

func (c *Reader) ExtractAnnotations(decl *ast.FuncDecl) ([]IAnnotation, error) {
	annotations := make([]IAnnotation, 0)
	if nil == decl.Doc {
		return annotations, errors.New("there is not annotation")
	}

	for _, c := range decl.Doc.List {
		com := c.Text
		annsMeth := parseMethodAnnotations(com)
		annotations = append(annotations, annsMeth...)
	}

	return annotations, nil
}

