package goParser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
)

type GoParser struct {
}

var instance *GoParser

func Instance() *GoParser {
	if nil == instance {
		instance = &GoParser{}
	}
	return instance
}

func (gp *GoParser) fetchControllersAnnotations() *GoParser {
	//@todo
	return gp
}

func (gp *GoParser) GetPackage(dir string) (*ast.Package, error) {
	fileset := token.NewFileSet()
	packs, err := parser.ParseDir(fileset, dir, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	// first package
	var pack *ast.Package
	for _, pack = range packs {
		break
	}

	return pack, nil
}

func (gp *GoParser) GetAstFile(filename string) (*ast.File, error) {
	fileset := token.NewFileSet()
	file, err := parser.ParseFile(fileset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (gp *GoParser) ExtractDecls(dir, declType string) ([]ast.Decl, error) {
	decls := make([]ast.Decl, 0)
	pack, err := gp.GetPackage(dir)
	if err != nil {
		return nil, err
	}

	for _, astFile := range pack.Files {
		for _, d := range astFile.Decls {
			if reflect.TypeOf(d).String() != declType {
				continue
			}

			decls = append(decls, d)
		}
	}

	return decls, nil
}

func (gp *GoParser) ExtractStructsSpecs(dir string) ([]*ast.TypeSpec, error) {
	specs := make([]*ast.TypeSpec, 0)
	pack, err := gp.GetPackage(dir)
	if err != nil {
		return nil, err
	}

	for _, astFile := range pack.Files {
		for _, d := range astFile.Decls {
			if reflect.TypeOf(d).String() != "*ast.GenDecl" {
				continue
			}

			decl := d.(*ast.GenDecl)
			if token.TYPE != decl.Tok {
				continue
			}

			for _, s := range decl.Specs {
				if reflect.TypeOf(s).String() != "*ast.TypeSpec" {
					continue
				}
				spec := s.(*ast.TypeSpec)
				if reflect.TypeOf(spec.Type).String() != "*ast.StructType" {
					continue
				}

				specs = append(specs, spec)
			}
		}
	}

	return specs, nil
}
